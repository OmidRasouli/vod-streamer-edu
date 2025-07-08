package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

// LocalStorage implements the VideoStorage interface for local filesystem storage.
// It saves, opens, and retrieves video files under a configurable base directory.
type LocalStorage struct {
	BasePath entity.Path // Root directory for all stored files
}

// NewLocalStorage creates a new LocalStorage instance with the given base path.
// This allows the application to store all files under a specific directory.
func NewLocalStorage(basePath string) port.VideoStorage {
	return &LocalStorage{
		BasePath: entity.NewPath(basePath),
	}
}

// Save writes the contents from the provided reader to a file at the given path (relative to BasePath).
// It creates any necessary directories and returns the resulting Path.
// Returns an error if writing fails.
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

// Open opens a file for reading at the given path (relative to BasePath).
// Returns a ReadCloser for the file, or an error if the file cannot be opened.
func (s *LocalStorage) Open(path ...string) (io.ReadCloser, error) {
	fullPath := s.BasePath.Join(path...).String()
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open video file: %w", err)
	}
	return file, nil
}

// GetPath returns the absolute Path for the given relative path, if the file exists.
// Returns an error if the file does not exist.
func (s *LocalStorage) GetPath(path ...string) (entity.Path, error) {
	fullPath := s.BasePath.Join(path...).String()

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return entity.Path{}, fmt.Errorf("file does not exist: %s", fullPath)
	}

	return entity.StringPathToPath(fullPath), nil
}
