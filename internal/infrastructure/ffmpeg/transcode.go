package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

// Transcode processes the input video file into multiple HLS renditions for adaptive streaming.
// For each supported video quality, it creates an output directory, generates TS segments and a playlist.
// It applies scaling filters based on the video's orientation (portrait or landscape).
// After transcoding all qualities, it generates a master playlist referencing all variants.
// Returns an error if any step fails.
func (s *FFmpegService) Transcode(input entity.Path, isPortrait bool) error {
	for _, q := range s.videoQualities {
		outputPath := input.Parent().String()
		qualityDir := filepath.Join(outputPath, "normal_hls", q.Name)
		if err := os.MkdirAll(qualityDir, 0755); err != nil {
			return fmt.Errorf("failed to create output dir %s: %w", qualityDir, err)
		}

		segmentPath := filepath.Join(qualityDir, "%03d.ts")
		playlistPath := filepath.Join(qualityDir, "index.m3u8")
		scaleFilter := q.ScaleHorizontally()
		if isPortrait {
			scaleFilter = q.ScaleVertically()
		}

		// Build and run the FFmpeg command for this quality.
		cmd := ffmpeg_go.Input(input.String()).
			WithCpuCoreRequest(s.cpuCoreRequest).
			WithCpuCoreLimit(s.cpuCoreLimit).
			Output(playlistPath, s.getFFmpegArgs(q, segmentPath, []string{scaleFilter, q.LandScape()}))

		err := cmd.
			OverWriteOutput().
			WithOutput(nil, os.Stdout).
			Run()
		if err != nil {
			return fmt.Errorf("ffmpeg failed for quality %s: %w", q.Name, err)
		}
	}

	// Generate the HLS master playlist referencing all created renditions.
	if err := s.generateMasterPlaylist(input.Parent().String()); err != nil {
		return fmt.Errorf("failed to generate master playlist: %w", err)
	}

	return nil
}
