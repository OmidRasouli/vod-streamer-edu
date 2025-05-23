package port

import "io"

type VideoStorage interface {
	Save(reader io.Reader, filename string) error // Save video data to storage
	Open(filename string) (io.ReadCloser, error)  // Open video file for reading
	GetPath(filename string) (string, error)      // Get absolute path of stored video file
}
