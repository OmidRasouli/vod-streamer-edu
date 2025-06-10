package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(videoController *VideoController) *gin.Engine {
	router := gin.Default()

	// Add a simple health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/upload", videoController.UploadVideo)

	return router
}
