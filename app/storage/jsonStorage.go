package storage

import (
	"os"
)

type JsonStorage struct {
	*Storage
}

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{
		&Storage{filename: filename},
	}
}
func (s *JsonStorage) Save(data []byte) error {
	return os.WriteFile(s.GetFilename(), data, 0644)
}
func (s *JsonStorage) Load() ([]byte, error) {
	data, err := os.ReadFile(s.GetFilename())
	return data, err
}
