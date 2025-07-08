package port

import (
	"io"

	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

// VideoStorage defines the interface (port) for video file storage operations.
// This abstraction allows the usecase/business logic to interact with storage
// without depending on a specific implementation (e.g., local disk, cloud storage).
// This is useful for testing and for swapping out storage backends if needed.
//
// Save: Saves video data to storage and returns the path.
// Open: Opens a video file for reading.
// GetPath: Gets the absolute path of a stored video file.
type VideoStorage interface {
	Save(reader io.Reader, path ...string) (entity.Path, error) // Save video data to storage
	Open(path ...string) (io.ReadCloser, error)                 // Open video file for reading
	GetPath(path ...string) (entity.Path, error)                // Get absolute path of stored video file
}
