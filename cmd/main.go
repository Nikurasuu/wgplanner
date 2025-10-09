package main

import (
	"wgplanner/internal/config"
	"wgplanner/internal/server"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	if cfg.Logger.Level == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	server := server.NewServer(cfg, logger, db)
	server.Run()
}
