package usecase

import (
	"fmt"
	"io"

	"github.com/OmidRasouli/vod-streamer-edu/internal/domain/port"
	"github.com/google/uuid"
)

// VideoUsecase contains the business logic for video operations.
// It coordinates storage and video processing (FFmpeg) through interfaces (ports).
type VideoUsecase struct {
	storage port.VideoStorage // Interface for saving and retrieving video files
	ffmpeg  port.FFmpeg       // Interface for video processing (transcoding, probing)
}

// NewVideoUsecase constructs a new VideoUsecase with the given storage and ffmpeg dependencies.
func NewVideoUsecase(storage port.VideoStorage, ffmpeg port.FFmpeg) *VideoUsecase {
	return &VideoUsecase{storage: storage, ffmpeg: ffmpeg}
}

// ProcessAndSave saves video data using the storage interface, probes its details, and transcodes it.
// 1. Generates a unique ID for the video.
// 2. Saves the uploaded file under /videos/{id}/{filename}.
// 3. Extracts video metadata using FFmpeg (FFprobe).
// 4. Checks orientation (portrait/landscape).
// 5. Transcodes the video for adaptive streaming.
// Returns an error if any step fails.
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

// GetVideoSegment returns the absolute path to a video segment or playlist for a given video ID and quality.
// Used by the HTTP controller to serve HLS segments or playlists.
func (uc *VideoUsecase) GetVideoSegment(fileID string, quality string, file string) string {
	videoPath, _ := uc.storage.GetPath(fileID, "normal_hls", quality, file)
	return videoPath.String()
}

// GetVideoMaster returns the absolute path to the master playlist (master.m3u8) for a given video ID.
// Used by the HTTP controller to serve the HLS master playlist.
func (uc *VideoUsecase) GetVideoMaster(fileID string) string {
	playlistPath, _ := uc.storage.GetPath(fileID, "normal_hls", "master.m3u8")
	return playlistPath.String()
}
