package storage

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func NewZipStorage(filename string) *ZipStorage {
	return &ZipStorage{
		&Storage{filename: filename},
	}
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.GetFilename())
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, err := zw.Create("data")
	_, err = w.Write(data)
	return err
}

func (z *ZipStorage) Load() ([]byte, error) {
	r, err := zip.OpenReader(z.GetFilename())
	if err != nil {
		return nil, errors.New("ошибка при чтении из архива")
	}
	defer r.Close()

	if len(r.File) == 0 {
		return nil, errors.New("архив пуст")
	}

	file := r.File[0]
	rc, err := file.Open()
	if err != nil {
		return nil, errors.New("ошибка при открытии файла")
	}
	defer rc.Close()

	return io.ReadAll(rc)
}
