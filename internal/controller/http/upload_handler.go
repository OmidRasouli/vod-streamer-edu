package http

import (
	"net/http"

	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
	"github.com/gin-gonic/gin"
)

type VideoController struct {
	VideoUsecase *usecase.VideoUsecase
}

func NewVideoController(videoUsecase *usecase.VideoUsecase) *VideoController {
	return &VideoController{VideoUsecase: videoUsecase}
}

func (h *VideoController) UploadVideo(c *gin.Context) {
	file, header, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video file is required"})
		return
	}
	defer file.Close()

	filename := header.Filename

	if err := h.VideoUsecase.Save(filename, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully", "filename": filename})
}
