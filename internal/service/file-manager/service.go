package file_manager

import (
	"encoding/base64"
	"os"
)

type FileManager struct {
}

func (f FileManager) SaveBase64File(filepath string, encodedContent string) error {
	data, err := base64.StdEncoding.DecodeString(encodedContent)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f FileManager) DeleteFile(filepath string) error {
	return os.Remove(filepath)
}
