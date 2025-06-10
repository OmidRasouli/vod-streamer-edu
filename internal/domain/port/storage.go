package port

import (
	"io"

	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

type VideoStorage interface {
	Save(reader io.Reader, path ...string) (entity.Path, error) // Save video data to storage
	Open(path ...string) (io.ReadCloser, error)                 // Open video file for reading
	GetPath(path ...string) (entity.Path, error)                // Get absolute path of stored video file
}
