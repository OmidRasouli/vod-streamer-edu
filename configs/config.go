package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	FFmpeg  FFmpegConfig  `yaml:"ffmpeg"`
	Storage StorageConfig `yaml:"storage"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type FFmpegConfig struct {
	CpuCoreLimit   float32 `yaml:"cpu_core_limit"`
	CpuCoreRequest float32 `yaml:"cpu_core_request"`
}

type StorageConfig struct {
	RawVideoPath string `yaml:"raw_video_path"`
}

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

func MustLoad(path string) *Config {
	cfg, err := loadConfig(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func (c *Config) GetFFmpegConfig() FFmpegConfig {
	return c.FFmpeg
}

func (c *Config) GetServerConfig() ServerConfig {
	return c.Server
}

func (c *Config) GetStorageConfig() StorageConfig {
	return c.Storage
}
