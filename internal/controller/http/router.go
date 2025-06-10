package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(uploadHandler *UploadHandler) *gin.Engine {
	router := gin.Default()

	// Add a simple health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/upload", uploadHandler.UploadVideo)

	return router
}
