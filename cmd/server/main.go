package main

import (
	"log"
	"strconv"

	"github.com/OmidRasouli/vod-streamer-edu/configs"
	"github.com/OmidRasouli/vod-streamer-edu/internal/controller/http"
)

func main() {
	cfg := configs.MustLoad("configs/config.yaml")

	runRouter(cfg)
}

func runRouter(cfg *configs.Config) {
	// Set up HTTP server
	router := http.NewRouter()
	port := cfg.GetServerConfig().Port

	if err := router.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	log.Printf("Server is running on port %d", port)
}
