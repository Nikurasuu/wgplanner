package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	Logger struct {
		Level string
	}
	Server struct {
		Port int
	}
	Mongo struct {
		Host     string
		Port     int
		DataBase string
	}
}

// NewConfig loads configuration from environment variables and .env file.
func NewConfig() (*Config, error) {
	// Load .env file if present
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.Logger.Level = os.Getenv("LOGGER_LEVEL")

	if portStr := os.Getenv("SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Server.Port = port
		}
	}

	cfg.Mongo.Host = os.Getenv("MONGO_HOST")
	if mongoPortStr := os.Getenv("MONGO_PORT"); mongoPortStr != "" {
		if mongoPort, err := strconv.Atoi(mongoPortStr); err == nil {
			cfg.Mongo.Port = mongoPort
		}
	}
	cfg.Mongo.DataBase = os.Getenv("MONGO_DATABASE")

	return cfg, nil
}
