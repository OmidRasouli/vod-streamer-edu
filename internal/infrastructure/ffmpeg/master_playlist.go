package service

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// generateMasterPlaylist creates an HLS master playlist (master.m3u8) in the given output directory.
// The master playlist references all available video qualities for adaptive streaming.
// Each quality is described with its bandwidth and resolution, and points to its own variant playlist.
// This is essential for HLS players to select the best stream based on network conditions.
func (s *FFmpegService) generateMasterPlaylist(outputDir string) error {
	masterFilePath := filepath.Join(outputDir, "master.m3u8")

	masterFile, err := os.Create(masterFilePath)
	if err != nil {
		return err
	}
	defer masterFile.Close()

	writer := bufio.NewWriter(masterFile)
	defer writer.Flush()

	// Write the HLS playlist header.
	if _, err := writer.WriteString("#EXTM3U\n"); err != nil {
		return err
	}

	// Add an entry for each video quality.
	for _, q := range s.videoQualities {
		bandwidth := extractBandwidth(q.Bitrate)
		line := fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", bandwidth, q.LandScape(), q.Name)
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

// extractBandwidth converts a bitrate string (e.g., "4500k") to an integer in bits per second.
// This is required for the BANDWIDTH attribute in the HLS playlist.
func extractBandwidth(bitrate string) int {
	if strings.HasSuffix(bitrate, "k") {
		bitrate = strings.TrimSuffix(bitrate, "k")
	}
	kbps, err := strconv.Atoi(bitrate)
	if err != nil {
		return 0
	}
	return kbps
}
