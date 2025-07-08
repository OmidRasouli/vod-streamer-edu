// main.go
// Entry point for the VOD Streamer educational project.
// This file sets up the server, loads configuration, and wires up the main dependencies.
// The project demonstrates a clean architecture approach for a video-on-demand streaming backend.

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
	// Load configuration from YAML file.
	// MustLoad will panic if the config file is missing or invalid,
	// ensuring the application never runs with incomplete settings.
	cfg := configs.MustLoad("configs/config.yaml")
	runServer(cfg)
}

// runServer initializes all dependencies and starts the HTTP server.
// It demonstrates dependency injection and separation of concerns.
func runServer(cfg *configs.Config) {
	// Initialize local storage for raw video files.
	// This abstracts file operations and can be replaced with other storage backends.
	localStorage := storage.NewLocalStorage(cfg.Storage.RawVideoPath)

	// Load FFmpeg configuration and initialize the FFmpeg service.
	// This service handles video processing tasks using FFmpeg.
	ffmpegConf := cfg.GetFFmpegConfig()
	ffmpegService := service.NewFFmpegService(ffmpegConf.CpuCoreRequest, ffmpegConf.CpuCoreLimit)

	// Create the video use case, which contains business logic for video operations.
	// It receives storage and FFmpeg services as dependencies.
	videoUseCase := usecase.NewVideoUsecase(localStorage, ffmpegService)

	// Set up the HTTP controller, which handles incoming API requests.
	videoController := http.NewVideoController(videoUseCase)

	// Create the HTTP router and register routes.
	router := http.NewRouter(videoController)

	// Prepare the server address using the configured port.
	port := cfg.GetServerConfig().Port
	addr := ":" + strconv.Itoa(port)

	log.Printf("Server is starting on port %d", port)
	// Start the HTTP server and log any fatal errors.
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
