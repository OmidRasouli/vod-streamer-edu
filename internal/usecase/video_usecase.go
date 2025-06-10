package usecase

import (
	"fmt"
	"io"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
	"github.com/google/uuid"
)

type VideoUsecase struct {
	storage port.VideoStorage
	ffmpeg  port.FFmpeg
}

func NewVideoUsecase(storage port.VideoStorage, ffmpeg port.FFmpeg) *VideoUsecase {
	return &VideoUsecase{storage: storage, ffmpeg: ffmpeg}
}

// ProcessAndSave saves video data using the storage interface
func (uc *VideoUsecase) ProcessAndSave(filename string, reader io.Reader) error {
	// Generate a new UUID for the filename
	id := uuid.New().String()

	// Save the video file using the storage interface
	// The video will be saved in the storage in a directory structure like: /videos/{id}/{filename}
	savedDetails, err := uc.storage.Save(reader, id, filename)
	if err != nil {
		return err
	}
	fmt.Println("Video saved with details:", savedDetails)

	// Get video details using the ffmpeg service
	// TODO: We will save it later in the database
	videoDetails, err := uc.ffmpeg.GetVideoDetails(savedDetails)
	if err != nil {
		return err
	}

	// Check if the video is portrait
	if videoDetails == nil {
		return fmt.Errorf("no video details available") // No video details available, nothing to process
	}
	isPortrait := videoDetails.IsPortrait()

	// Transcode the video using the ffmpeg service
	if err := uc.ffmpeg.Transcode(savedDetails, isPortrait); err != nil {
		return err
	}

	return nil
}

func (uc *VideoUsecase) GetVideoSegment(fileID string, quality string, file string) string {
	videoPath, _ := uc.storage.GetPath(fileID, "normal_hls", quality, file)
	return videoPath.String()
}

func (uc *VideoUsecase) GetVideoMaster(fileID string) string {
	playlistPath, _ := uc.storage.GetPath(fileID, "normal_hls", "master.m3u8")
	return playlistPath.String()
}
