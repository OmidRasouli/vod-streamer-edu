// router.go
// Defines the HTTP routing for the VOD Streamer educational project.
// This file sets up the Gin router, applies middleware (like CORS), and registers API endpoints.

package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creates and configures a new Gin HTTP router.
// It registers all API endpoints and applies necessary middleware.
// The router is the main entry point for HTTP requests.
func NewRouter(videoController *VideoController) *gin.Engine {
	router := gin.Default()

	// Add CORS middleware to allow cross-origin requests (useful for frontend integration).
	router.Use(cors.Default())

	// Add a simple health check endpoint.
	// Useful for monitoring and verifying the server is running.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// Register endpoint for uploading videos.
	router.POST("/upload", videoController.UploadVideo)

	// Register endpoint for serving the master playlist (HLS streaming).
	router.GET("/stream/:id/master.m3u8", videoController.GetMaster)

	// Register endpoint for streaming video segments at different qualities.
	router.GET("/stream/:id/:quality/:file", videoController.Stream)

	return router
}
