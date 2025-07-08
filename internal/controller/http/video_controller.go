package http

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
	"github.com/gin-gonic/gin"
)

// VideoController handles HTTP requests related to video operations.
// It acts as a bridge between HTTP layer and business logic (usecase).
type VideoController struct {
	VideoUsecase *usecase.VideoUsecase
}

// NewVideoController creates a new VideoController with the given usecase dependency.
func NewVideoController(videoUsecase *usecase.VideoUsecase) *VideoController {
	return &VideoController{VideoUsecase: videoUsecase}
}

// UploadVideo handles video upload requests.
// It expects a multipart form with a "video" file field.
// The uploaded file is passed to the usecase for processing and saving.
// Responds with success or error message.
func (v *VideoController) UploadVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
		return
	}
	defer file.Close()

	filename := header.Filename

	if err := v.VideoUsecase.ProcessAndSave(filename, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully", "filename": filename})
}

// GetMaster serves the master playlist (master.m3u8) for a given video ID.
// This is used by HLS players to get available video qualities.
func (v *VideoController) GetMaster(c *gin.Context) {
	fileID := c.Param("id")

	video := v.VideoUsecase.GetVideoMaster(fileID)
	c.File(video)
}

// Stream serves individual video segments or playlists for a given video ID and quality.
// Validates the file extension to allow only .m3u8 and .ts files for security.
func (v *VideoController) Stream(c *gin.Context) {
	fileID := c.Param("id")
	quality := c.Param("quality")
	file := c.Param("file")

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(file))
	if ext != ".m3u8" && ext != ".ts" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension. Only .m3u8 and .ts are allowed."})
		return
	}

	video := v.VideoUsecase.GetVideoSegment(fileID, quality, file)
	c.File(video)
}
