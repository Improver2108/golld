package file

import (
	"fmt"
	"os"
)

type FileStorage struct {
}

func Init() *FileStorage {
	return &FileStorage{}
}

func (f *FileStorage) Save(data string) {
	if err := os.WriteFile("./documenteditor/good/document.txt", []byte(data), 0644); err != nil {
		fmt.Println("Error: Unable to open file for writing.")
		return
	}
	fmt.Println("Document saved to document.txt")
}
