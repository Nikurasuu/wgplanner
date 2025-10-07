package main

import (
	"wgplanner/internal/config"
	"wgplanner/internal/server"

	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	mgmErr := mgm.SetDefaultConfig(nil, cfg.Mongo.Database, options.Client().ApplyURI(cfg.Mongo.ConnectionString))
	if mgmErr != nil {
		logger.Fatalf("Error setting up mgm: %v", mgmErr)
	}

	server := server.NewServer(cfg, logger)
	server.Run()
}
