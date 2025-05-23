package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/model"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func (s *FFmpegService) GetVideoDetails(path string) (*model.VideoData, error) {
	videoDetailsJSON, err := ffmpeg_go.Probe(path)
	if err != nil {
		log.Printf("FFprobe error: %v", err)
		return nil, fmt.Errorf("failed to probe video: %w", err)
	}

	var videoData model.VideoData
	if err := json.Unmarshal([]byte(videoDetailsJSON), &videoData); err != nil {
		return nil, fmt.Errorf("failed to parse video data: %w", err)
	}

	if len(videoData.Streams) == 0 {
		return nil, fmt.Errorf("no video stream found")
	}

	return &videoData, nil
}
