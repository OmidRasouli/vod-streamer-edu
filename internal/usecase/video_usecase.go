package usecase

import (
	"io"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
)

type VideoUsecase struct {
	storage port.VideoStorage
}

func NewVideoUsecase(storage port.VideoStorage) *VideoUsecase {
	return &VideoUsecase{storage: storage}
}

// Save saves video data using the storage interface
func (uc *VideoUsecase) Save(filename string, reader io.Reader) error {
	return uc.storage.Save(reader, filename)
}

// Open returns a ReadCloser for the video file
func (uc *VideoUsecase) Open(filename string) (io.ReadCloser, error) {
	return uc.storage.Open(filename)
}

// GetPath returns the absolute path of the video file
func (uc *VideoUsecase) GetPath(filename string) (string, error) {
	return uc.storage.GetPath(filename)
}
