package http

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
	"github.com/gin-gonic/gin"
)

type VideoController struct {
	VideoUsecase *usecase.VideoUsecase
}

func NewVideoController(videoUsecase *usecase.VideoUsecase) *VideoController {
	return &VideoController{VideoUsecase: videoUsecase}
}

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

func (v *VideoController) GetMaster(c *gin.Context) {
	fileID := c.Param("id")

	video := v.VideoUsecase.GetVideoMaster(fileID)
	c.File(video)
}

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
