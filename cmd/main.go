package main

import (
	"fmt"

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

	dsn := "host=" + cfg.Database.Host +
		" user=" + cfg.Database.User +
		" password=" + cfg.Database.Password +
		" port=" + fmt.Sprint(cfg.Database.Port) +
		" sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	server := server.NewServer(cfg, logger, db)
	server.Run()
}
