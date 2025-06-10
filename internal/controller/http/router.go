package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(videoController *VideoController) *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Add a simple health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/upload", videoController.UploadVideo)
	router.GET("/stream/:id/master.m3u8", videoController.GetMaster)
	router.GET("/stream/:id/:quality/:file", videoController.Stream)

	return router
}
