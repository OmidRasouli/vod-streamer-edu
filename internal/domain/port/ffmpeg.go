package port

import (
	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/model"
	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

// FFmpeg defines the interface (port) for video processing operations.
// This abstraction allows the usecase/business logic to interact with FFmpeg functionality
// without depending on a specific implementation. This is useful for testing and for
// swapping out the underlying FFmpeg logic if needed.
//
// Transcode: Transcodes the input video file, possibly adjusting for portrait/landscape orientation.
// GetVideoDetails: Uses FFprobe to extract metadata and details from a video file.
type FFmpeg interface {
	Transcode(input entity.Path, isPortrait bool) error         // Transcode video data
	GetVideoDetails(path entity.Path) (*model.VideoData, error) // Get video details
}
