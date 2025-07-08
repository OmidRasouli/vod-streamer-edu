// config.go
// Provides configuration loading and access for the VOD Streamer educational project.
// This file defines the structure of the configuration, loads it from a YAML file,
// and exposes helper methods to access different configuration sections.

package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the main configuration struct that holds all app settings.
// It is populated from the YAML config file.
type Config struct {
	Server  ServerConfig  `yaml:"server"`  // Server-related settings (e.g., port)
	FFmpeg  FFmpegConfig  `yaml:"ffmpeg"`  // FFmpeg processing settings
	Storage StorageConfig `yaml:"storage"` // Storage paths/settings
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Port int `yaml:"port"` // Port on which the HTTP server will listen
}

// FFmpegConfig holds settings for the FFmpeg video processing service.
type FFmpegConfig struct {
	CpuCoreLimit   float32 `yaml:"cpu_core_limit"`   // Max CPU cores FFmpeg can use
	CpuCoreRequest float32 `yaml:"cpu_core_request"` // Requested CPU cores for FFmpeg
}

// StorageConfig holds paths for storing video files.
type StorageConfig struct {
	RawVideoPath string `yaml:"raw_video_path"` // Directory for raw (uploaded) videos
}

// loadConfig reads and parses the YAML configuration file at the given path.
// Returns a Config struct or an error if loading fails.
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// MustLoad loads the configuration and panics if there is any error.
// Useful for ensuring the app never starts with invalid or missing config.
func MustLoad(path string) *Config {
	cfg, err := loadConfig(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

// GetFFmpegConfig returns the FFmpeg section of the config.
func (c *Config) GetFFmpegConfig() FFmpegConfig {
	return c.FFmpeg
}

// GetServerConfig returns the server section of the config.
func (c *Config) GetServerConfig() ServerConfig {
	return c.Server
}

// GetStorageConfig returns the storage section of the config.
func (c *Config) GetStorageConfig() StorageConfig {
	return c.Storage
}
