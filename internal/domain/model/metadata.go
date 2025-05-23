package model

type Stream struct {
	Index              int    `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Profile            string `json:"profile"`
	CodecType          string `json:"codec_type"`
	CodecTagString     string `json:"codec_tag_string"`
	CodecTag           string `json:"codec_tag"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	CodedWidth         int    `json:"coded_width"`
	CodedHeight        int    `json:"coded_height"`
	ClosedCaptions     int    `json:"closed_captions"`
	HasBFrames         int    `json:"has_b_frames"`
	SampleAspectRatio  string `json:"sample_aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	PixFmt             string `json:"pix_fmt"`
	Level              int    `json:"level"`
	ChromaLocation     string `json:"chroma_location"`
	Refs               int    `json:"refs"`
	IsAVC              string `json:"is_avc"`
	NalLengthSize      string `json:"nal_length_size"`
	RFrameRate         string `json:"r_frame_rate"`
	AvgFrameRate       string `json:"avg_frame_rate"`
	TimeBase           string `json:"time_base"`
	StartPTS           int    `json:"start_pts"`
	StartTime          string `json:"start_time"`
	DurationTS         int    `json:"duration_ts"`
	Duration           string `json:"duration"`
	BitRate            string `json:"bit_rate"`
	BitsPerRawSample   string `json:"bits_per_raw_sample"`
	NbFrames           string `json:"nb_frames"`
	Disposition        struct {
		Default         int `json:"default"`
		Dub             int `json:"dub"`
		Original        int `json:"original"`
		Comment         int `json:"comment"`
		Lyrics          int `json:"lyrics"`
		Karaoke         int `json:"karaoke"`
		Forced          int `json:"forced"`
		HearingImpaired int `json:"hearing_impaired"`
		VisualImpaired  int `json:"visual_impaired"`
		CleanEffects    int `json:"clean_effects"`
		AttachedPic     int `json:"attached_pic"`
		TimedThumbnails int `json:"timed_thumbnails"`
	} `json:"disposition"`
	Tags struct {
		CreationTime string `json:"creation_time"`
		Language     string `json:"language"`
		HandlerName  string `json:"handler_name"`
		VendorID     string `json:"vendor_id"`
	} `json:"tags"`
}

func (s Stream) IsPortrait() bool {
	return s.Height > s.Width
}

type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
	Tags           struct {
		MajorBrand       string `json:"major_brand"`
		MinorVersion     string `json:"minor_version"`
		CompatibleBrands string `json:"compatible_brands"`
		CreationTime     string `json:"creation_time"`
		Encoder          string `json:"encoder"`
	} `json:"tags"`
}

type VideoData struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}
