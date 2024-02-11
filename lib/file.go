package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(fileName, extension string, content []byte) {
	if exists, err := isFileExists(fileName + extension); exists || err != nil {
		fmt.Printf("Error: File %s%s already exists\n", fileName, extension)
		return
	}

	dirPath := filepath.Dir(fileName)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	file, err := os.Create(fileName + extension)
	if err != nil {
		fmt.Println("Create file error", err)
		return
	}
	defer file.Close()

	if _, err = file.Write(content); err != nil {
		fmt.Println("Write file content error", err)
	}
}

func RemoveFile(fileName string) {
	if err := os.RemoveAll(fileName); err != nil {
		fmt.Printf("Remove file %s error %e", fileName, err)
	}
}

func isFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
