package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	path string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	scanError := scanner.Err()
	if scanError != nil {
		file.Close()
		return nil, scanError
	}
	return lines, nil
}

func WriteJsonFile(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data to JSON")
	}
	file.Close()
	return nil
}

func NewFileManager(path string) *FileManager {
	return &FileManager{
		path: path,
	}
}
