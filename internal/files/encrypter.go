package files

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"log"
)

type EncrypterAPI interface {
	encrypt(file []byte, key string) ([]byte, error)
	decrypt(file []byte, key string) ([]byte, error)
	checkMD5(input, expectedHash string) bool
}

type Encrypter struct {
}

func NewEncrypter() *Encrypter {
	return &Encrypter{}
}

func (e Encrypter) encrypt(plaintext []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func (e Encrypter) decrypt(ciphertext []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println("Error creating cipher block:", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error creating GCM:", err)
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (e Encrypter) checkMD5(input, expectedHash string) bool {
	hasher := md5.New()
	hasher.Write([]byte(input))
	calculatedHash := hex.EncodeToString(hasher.Sum(nil))
	return calculatedHash == expectedHash
}
