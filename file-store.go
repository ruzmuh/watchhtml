package main

import (
	"fmt"
	"log"
	"os"
)

type FileStore struct {
	storeDir string
}

func NewFileStore(dir string) *FileStore {
	return &FileStore{
		storeDir: dir,
	}
}

func (fs *FileStore) PutValue(key, value string) (err error) {
	filePath := fmt.Sprintf("%s/%s", fs.storeDir, key)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(value))
	if err != nil {
		return
	}
	log.Printf("successfully put a value file=%s, value=%s", filePath, value)
	return
}

func (fs *FileStore) GetValue(key string) (result string, err error) {
	filePath := fmt.Sprintf("%s/%s", fs.storeDir, key)
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	result = string(buf)
	log.Printf("got a value from the file=%s, value=%s", filePath, result)
	return
}
