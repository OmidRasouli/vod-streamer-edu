package main

import (
	"log"
	"strconv"

	"github.com/OmidRasouli/vod-streamer-edu/configs"
	"github.com/OmidRasouli/vod-streamer-edu/internal/controller/http"
	service "github.com/OmidRasouli/vod-streamer-edu/internal/infrastructure/ffmpeg"
	"github.com/OmidRasouli/vod-streamer-edu/internal/infrastructure/storage"
	"github.com/OmidRasouli/vod-streamer-edu/internal/usecase"
)

func main() {
	// Load configuration (panic if not found or invalid)
	cfg := configs.MustLoad("configs/config.yaml")
	runServer(cfg)
}

func runServer(cfg *configs.Config) {
	// Initialize infrastructure
	localStorage := storage.NewLocalStorage(cfg.Storage.RawVideoPath)
	ffmpegConf := cfg.GetFFmpegConfig()
	ffmpegService := service.NewFFmpegService(ffmpegConf.CpuCoreRequest, ffmpegConf.CpuCoreLimit)

	// Initialize use case and controller
	videoUseCase := usecase.NewVideoUsecase(localStorage, ffmpegService)
	videoController := http.NewVideoController(videoUseCase)
	router := http.NewRouter(videoController)

	port := cfg.GetServerConfig().Port
	addr := ":" + strconv.Itoa(port)

	log.Printf("Server is starting on port %d", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
