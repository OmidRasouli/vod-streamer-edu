package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

type LocalStorage struct {
	BasePath entity.Path
}

func NewLocalStorage(basePath string) port.VideoStorage {
	return &LocalStorage{
		BasePath: entity.NewPath(basePath),
	}
}

func (s *LocalStorage) Save(reader io.Reader, path ...string) (entity.Path, error) {
	fullPath := s.BasePath.Join(path...).String()

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return entity.Path{}, fmt.Errorf("failed to create directories: %w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return entity.Path{}, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return entity.Path{}, fmt.Errorf("failed to save video data: %w", err)
	}

	return entity.StringPathToPath(fullPath), nil
}

func (s *LocalStorage) Open(path ...string) (io.ReadCloser, error) {
	fullPath := s.BasePath.Join(path...).String()
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open video file: %w", err)
	}
	return file, nil
}

func (s *LocalStorage) GetPath(path ...string) (entity.Path, error) {
	fullPath := s.BasePath.Join(path...).String()

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return entity.Path{}, fmt.Errorf("file does not exist: %s", fullPath)
	}

	return entity.StringPathToPath(fullPath), nil
}
