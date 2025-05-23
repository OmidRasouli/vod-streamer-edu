package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) port.VideoStorage {
	return &LocalStorage{
		BasePath: basePath,
	}
}

func (s *LocalStorage) Save(reader io.Reader, path ...string) error {
	fullPath := filepath.Join(append([]string{s.BasePath}, path...)...)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("failed to save video data: %w", err)
	}

	return nil
}

func (s *LocalStorage) Open(filename string) (io.ReadCloser, error) {
	fullPath := filepath.Join(s.BasePath, filename)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open video file: %w", err)
	}
	return file, nil
}

func (s *LocalStorage) GetPath(path ...string) (string, error) {
	fullPath := filepath.Join(append([]string{s.BasePath}, path...)...)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", fullPath)
	}

	return fullPath, nil
}
