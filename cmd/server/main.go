package main

import (
	"log"
	"strconv"

	"github.com/OmidRasouli/vod-streamer-edu/configs"
	"github.com/OmidRasouli/vod-streamer-edu/internal/controller/http"
	"github.com/OmidRasouli/vod-streamer-edu/internal/infrastructure/storage"
	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
)

func main() {
	cfg := configs.MustLoad("configs/config.yaml")

	runRouter(cfg)
}

func runRouter(cfg *configs.Config) {
	storagePath := cfg.Storage.RawVideoPath
	localStorage := storage.NewLocalStorage(storagePath)
	videoUseCase := usecase.NewVideoUsecase(localStorage)
	uploadHandler := http.NewUploadHandler(videoUseCase)
	router := http.NewRouter(uploadHandler)
	port := cfg.GetServerConfig().Port

	if err := router.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	log.Printf("Server is running on port %d", port)
}
