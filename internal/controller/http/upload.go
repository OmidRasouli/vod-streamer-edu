package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadVideoHandler(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video file is required"})
		return
	}

	fmt.Println("Received file:", file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"message": "video uploaded successfully",
	})
}
