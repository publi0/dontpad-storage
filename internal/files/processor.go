package files

import (
	"crypto/md5"
	"dontpad-storage/internal/dontpad"
	"dontpad-storage/internal/user"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type ProcessorAPI interface {
	UploadFile(file []byte, fileName string) error
	GetFileContentAndName(id string) ([]byte, string, error)
	ListFiles() ([]user.FileData, error)
	DeleteFile(id string) error
}

type Processor struct {
	encrypter EncrypterAPI
	dontpad   dontpad.ClientAPI
	user      user.DataAPI
}

func NewProcessor(encrypter EncrypterAPI, dontpad dontpad.ClientAPI, user user.DataAPI) *Processor {
	return &Processor{
		encrypter: encrypter,
		dontpad:   dontpad,
		user:      user,
	}
}

func (p *Processor) UploadFile(file []byte, fileName string) error {
	id := uuid.New().String()

	encrypt, err := p.encrypter.encrypt(file, getEncryptKey())
	if err != nil {
		return err
	}

	base64 := base64.StdEncoding.EncodeToString(encrypt)

	hash := md5.Sum([]byte(base64))
	hashString := hex.EncodeToString(hash[:])

	err = p.dontpad.UploadFile(base64, id)
	if err != nil {
		return err
	}

	err = p.user.SaveFile(user.FileData{
		ID:        id,
		Name:      fileName,
		MD5:       hashString,
		CreatedAt: time.Now().String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) GetFileContentAndName(id string) ([]byte, string, error) {
	getFile, err := p.user.GetFile(id)
	if err != nil {
		return nil, "", err
	}

	log.Println("File name: "+getFile.Name, " File MD5: "+getFile.MD5)

	file, err := p.dontpad.DownloadFile(id)
	if err != nil {
		return nil, "", err
	}

	checkMD5 := p.encrypter.checkMD5(file, getFile.MD5)
	if !checkMD5 {
		return nil, "", errors.New("md5 check failed")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return nil, "", err
	}

	decrypt, err := p.encrypter.decrypt(decodedBytes, getEncryptKey())
	if err != nil {
		return nil, "", err
	}

	return decrypt, getFile.Name, nil
}

func (p *Processor) ListFiles() ([]user.FileData, error) {
	return p.user.GetFiles()
}

func (p *Processor) DeleteFile(id string) error {
	err := p.dontpad.UploadFile("deleted", id)
	if err != nil {
		return err
	}

	err = p.user.DeleteFile(id)
	if err != nil {
		return err
	}

	return nil
}

func getEncryptKey() string {
	file, err := os.ReadFile("resources/key.txt")
	if err != nil {
		log.Println("Error reading key from file")
		log.Println("Generating new key")
		s := uuid.New().String()
		err := os.WriteFile("resources/key.txt", []byte(s), 0644)
		if err != nil {
			log.Println("Error writing key to file")
			panic(err)
		}
		return s
	}
	return string(file)
}
