package port

import (
	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/model"
	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
)

type FFmpeg interface {
	Transcode(input entity.Path, isPortrait bool) error         // Transcode video data
	GetVideoDetails(path entity.Path) (*model.VideoData, error) // Get video details using FFprobe
}
