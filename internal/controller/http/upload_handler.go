package http

import (
	"net/http"

	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	VideoUsecase *usecase.VideoUsecase
}

func NewUploadHandler(videoUsecase *usecase.VideoUsecase) *UploadHandler {
	return &UploadHandler{VideoUsecase: videoUsecase}
}

func (h *UploadHandler) UploadVideo(c *gin.Context) {
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
