package service

import "fmt"

// VideoQuality represents a specific video resolution and bitrate profile for transcoding.
// Each quality includes a name, dimensions, and bitrate settings used for adaptive streaming.
type VideoQuality struct {
	Name    string // e.g., "1080p"
	Width   int    // Target width in pixels
	Height  int    // Target height in pixels
	Bitrate string // Target video bitrate (e.g., "4500k")
	Maxrate string // Maximum allowed bitrate for this quality
	Bufsize string // Buffer size for rate control
}

// ScaleHorizontally returns a scaling filter string for FFmpeg to resize video to this quality,
// maintaining aspect ratio and fitting within the target width and height.
func (vq VideoQuality) ScaleHorizontally() string {
	return fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", vq.Width, vq.Height)
}

// LandScape returns the resolution string in "WIDTHxHEIGHT" format (e.g., "1920x1080").
func (vq VideoQuality) LandScape() string {
	return fmt.Sprintf("%dx%d", vq.Width, vq.Height)
}

// ScaleVertically returns a scaling filter string for FFmpeg to resize portrait videos to this quality,
// ensuring the minimum dimension fits while maintaining aspect ratio.
func (vq VideoQuality) ScaleVertically() string {
	return fmt.Sprintf("scale='min(%d,iw*%d/ih)':-1", vq.Width, vq.Height)
}
