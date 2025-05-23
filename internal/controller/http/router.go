package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(uploadHandler *UploadHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/upload", uploadHandler.UploadVideo)

	return router
}
