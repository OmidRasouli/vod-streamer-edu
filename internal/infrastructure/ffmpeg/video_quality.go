package service

import "fmt"

type VideoQuality struct {
	Name    string
	Width   int
	Height  int
	Bitrate string
	Maxrate string
	Bufsize string
}

func (vq VideoQuality) ScaleHorizontally() string {
	return fmt.Sprintf("scale=w=%d:h=%d:force_original_aspect_ratio=decrease", vq.Width, vq.Height)
}

func (vq VideoQuality) LandScape() string {
	return fmt.Sprintf("%dx%d", vq.Width, vq.Height)
}

func (vq VideoQuality) ScaleVertically() string {
	return fmt.Sprintf("scale='min(%d,iw*%d/ih)':-1", vq.Width, vq.Height)
}
