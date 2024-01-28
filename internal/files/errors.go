package files

import "fmt"

type FileNotFoundErr struct {
	FileName string
}

func (e *FileNotFoundErr) Error() string {
	return fmt.Sprintf("file [%s] not found", e.FileName)
}
