package service

import (
	"fmt"
	"io"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func (s *FFmpegService) Transcode(input io.Reader, outputDir string, isPortrait bool) error {
	for _, q := range s.VideoQualities {
		segmentPath := fmt.Sprintf("%s/%s/%%03d.ts", outputDir, q.Name)
		playlistPath := fmt.Sprintf("%s/%s/index.m3u8", outputDir, q.Name)
		scaleFilter := q.ScaleHorizontally()

		if isPortrait {
			scaleFilter = q.ScaleVertically()
		}

		cmd := ffmpeg_go.Input("pipe:").
			Output(
				playlistPath,
				s.getFFmpegArgs(q, segmentPath, []string{scaleFilter, q.LandScape()}),
			).
			OverWriteOutput().
			WithInput(input)

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
