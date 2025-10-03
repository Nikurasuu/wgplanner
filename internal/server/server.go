package server

import (
	"strconv"

	"wgplanner/internal/config"
	"wgplanner/internal/entity"
	"wgplanner/internal/handler"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg    *config.Config
	logger *logrus.Logger
}

func NewServer(cfg *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Run() error {
	s.logger.Infof("Starting server on port %d", s.cfg.Server.Port)

	srv := fuego.NewServer(
		fuego.WithAddr("localhost:" + strconv.Itoa(s.cfg.Server.Port)),
	)

	api := fuego.Group(srv, "/api",
		option.Tags("API"),
	)

	groupCollection := mgm.Coll(&entity.Group{})
	groupHandler := handler.NewGroupHandler(s.logger, groupCollection)
	addGroupRoutes(api, groupHandler)

	return srv.Run()
}
