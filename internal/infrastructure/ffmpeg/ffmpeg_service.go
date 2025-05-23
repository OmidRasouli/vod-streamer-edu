package service

import ffmpeg_go "github.com/u2takey/ffmpeg-go"

type FFmpegService struct {
	VideoQualities []VideoQuality
}

func NewFFmpegService() *FFmpegService {
	return &FFmpegService{
		VideoQualities: []VideoQuality{
			{"1080p", 1920, 1080, "4500k", "4700k", "6000k"},
			{"720p", 1280, 720, "2500k", "2675k", "3750k"},
			{"480p", 854, 480, "1000k", "1075k", "1500k"},
			{"360p", 640, 360, "600k", "650k", "900k"},
			{"240p", 426, 240, "400k", "450k", "600k"},
			{"144p", 256, 144, "250k", "275k", "400k"},
		},
	}
}

func (s *FFmpegService) getFFmpegArgs(q VideoQuality, segmentPath string, filters []string) ffmpeg_go.KwArgs {
	return ffmpeg_go.KwArgs{
		"c:v":                  "h264",                          // Use H.264 video codec
		"profile:v":            "main",                          // Set video encoding profile to "main" for broad compatibility
		"crf":                  "20",                            // Constant Rate Factor - balances quality and compression (lower = better quality)
		"sc_threshold":         "0",                             // Disable scene change detection for keyframes (forces regular keyframes)
		"g":                    "48",                            // GOP size: one keyframe every 48 frames (assuming ~2s GOP for 24fps)
		"keyint_min":           "48",                            // Minimum interval between keyframes (same as GOP)
		"b:v":                  q.Bitrate,                       // Target video bitrate for this quality level
		"maxrate":              q.Maxrate,                       // Maximum allowed video bitrate
		"bufsize":              q.Bufsize,                       // Buffer size for rate control
		"c:a":                  "aac",                           // Use AAC audio codec
		"ar":                   "48000",                         // Audio sampling rate (48kHz)
		"b:a":                  "128k",                          // Audio bitrate
		"hls_list_size":        "0",                             // Ensure the entire playlist is written (not a sliding window)
		"hls_time":             "6",                             // Duration of each segment in seconds
		"hls_playlist_type":    "vod",                           // Indicate this is a video-on-demand playlist
		"start_number":         "1",                             // Start segment numbering from 1
		"hls_segment_filename": segmentPath,                     // Pattern for naming the TS segment files
		"hls_flags":            "round_durations+split_by_time", // Round segment durations and split strictly by time
		"hls_allow_cache":      "1",                             // Allow caching of HLS segments
		"vf":                   filters[0],                      // Video filter (e.g., scaling)
		"s":                    filters[1],                      // Output resolution (explicit)
	}
}
