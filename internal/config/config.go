package config

import (
	"log"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Env         string     `yaml:"env" env:"ENV" env-default:"local"`
    StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
    HTTPServer  HTTPServer `yaml:"http_server"`
}


type HTTPServer struct {
    Address     string        `yaml:"address" env:"ADDRESS" env-default:"localhost:8080"`
    TimeOut     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`
    IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}


func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is required")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	return &cfg
}