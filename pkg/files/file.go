package files

import (
	"io"
	"os"
)

func CreateDir(name string) error {
	
	if err := os.MkdirAll(name,0755); err != nil {
		return nil
	}

	return nil
}

func CopyToNewFile(source io.Reader, destination string, mode int64) error {
	// create or open file
	file, err := os.OpenFile(destination,os.O_CREATE|os.O_RDWR, os.FileMode(mode))
	if err != nil {
		return err
	}
	defer file.Close()
	// copy contents from source to destination
	if _, err = io.Copy(file,source); err != nil {
		return err
	}

	return nil
}

func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	return err
}