package server

import (
	"strconv"

	"wgplanner/internal/config"
	"wgplanner/internal/handler"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	logger *logrus.Logger
	gormDB *gorm.DB
}

func NewServer(cfg *config.Config, logger *logrus.Logger, gormDB *gorm.DB) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		gormDB: gormDB,
	}
}

func (s *Server) Run() error {
	s.logger.Infof("Starting server on port %d", s.cfg.Server.Port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://wgplanner.niklas-malkusch.de", "https://wgplanner.niklas-malkusch.de/", "http://localhost:5173", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "content-type", "Authorization"},
		AllowCredentials: true,
	})

	srv := fuego.NewServer(
		fuego.WithGlobalMiddlewares(c.Handler),
		fuego.WithAddr("127.0.0.1:"+strconv.Itoa(s.cfg.Server.Port)),
	)

	api := fuego.Group(srv, "/api",
		option.Tags("API"),
	)

	groupHandler := handler.NewGroupHandler(s.logger, s.gormDB)
	addGroupRoutes(api, groupHandler)

	return srv.Run()
}
