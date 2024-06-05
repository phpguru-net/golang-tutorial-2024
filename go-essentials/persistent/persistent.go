package persistent

import "os"

const WRITE_FILE_MODE = 0644

type FileStorage struct {
	filePath string
}

func New(filePath string) *FileStorage {
	return &FileStorage{
		filePath: filePath,
	}
}

func (fs *FileStorage) SaveData(data []byte) error {
	return os.WriteFile(fs.filePath, data, WRITE_FILE_MODE)
}

func (fs *FileStorage) GetData() ([]byte, error) {
	return os.ReadFile(fs.filePath)
}
