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
		Host string
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}

// NewConfig loads configuration from environment variables and .env file.
func NewConfig() (*Config, error) {
	// Load .env file if present
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.Logger.Level = os.Getenv("LOGGER_LEVEL")

	cfg.Server.Host = os.Getenv("SERVER_HOST")

	if portStr := os.Getenv("SERVER_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Server.Port = port
		}
	}

	cfg.Database.Host = os.Getenv("POSTGRES_HOST")
	cfg.Database.User = os.Getenv("POSTGRES_USER")
	cfg.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	if portStr := os.Getenv("POSTGRES_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			cfg.Database.Port = port
		}
	}
	cfg.Database.Database = os.Getenv("POSTGRES_DB")

	return cfg, nil
}
