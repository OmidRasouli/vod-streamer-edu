package service

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *FFmpegService) generateMasterPlaylist(outputDir string) error {
	masterFilePath := filepath.Join(outputDir, "master.m3u8")

	masterFile, err := os.Create(masterFilePath)
	if err != nil {
		return err
	}
	defer masterFile.Close()

	writer := bufio.NewWriter(masterFile)
	defer writer.Flush()

	if _, err := writer.WriteString("#EXTM3U\n"); err != nil {
		return err
	}

	for _, q := range s.videoQualities {
		bandwidth := extractBandwidth(q.Bitrate)
		line := fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", bandwidth, q.LandScape(), q.Name)
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func extractBandwidth(bitrate string) int {
	if strings.HasSuffix(bitrate, "k") {
		bitrate = strings.TrimSuffix(bitrate, "k")
	}
	kbps, err := strconv.Atoi(bitrate)
	if err != nil {
		return 0
	}
	return kbps * 1000
}
