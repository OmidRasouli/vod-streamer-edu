package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OmidRasouli/vod-streamer-edu/internal/entity"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

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

	if err := s.generateMasterPlaylist(input.Parent().String()); err != nil {
		return fmt.Errorf("failed to generate master playlist: %w", err)
	}

	return nil
}
