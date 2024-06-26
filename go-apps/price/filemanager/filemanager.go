package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
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

func WriteJsonFile(path string, data interface{}, c chan bool, ec chan error) error {
	time.Sleep(3 * time.Second)
	file, err := os.Create(path)

	// test error
	// ec <- errors.New("Something wrong")

	if err != nil {
		ec <- err
		return err
	}
	// after every run, always close file
	defer file.Close()
	defer fmt.Printf("File %f closed!!!\n", path)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		ec <- err
		return errors.New("Failed to convert data to JSON")
	}
	c <- true
	return nil
}

func NewFileManager(path string) *FileManager {
	return &FileManager{
		path: path,
	}
}
