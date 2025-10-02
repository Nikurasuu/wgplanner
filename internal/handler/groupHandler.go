package handler

import (
	"time"

	"wgplanner/internal/entity"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
)

type GroupHandler struct {
	logger          *logrus.Logger
	mongoCollection *mgm.Collection
}

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

func NewGroupHandler(logger *logrus.Logger, mongoCollection *mgm.Collection) *GroupHandler {
	return &GroupHandler{
		logger:          logger,
		mongoCollection: mongoCollection,
	}
}

func (h *GroupHandler) CreateGroup(c fuego.ContextWithBody[CreateGroupRequest]) (*entity.Group, error) {
	var group entity.Group
	body, err := c.Body()
	if err != nil {
		h.logger.Errorf("Error parsing request body: %v", err)
		return &group, fuego.BadRequestError{Title: "Invalid request body", Err: err}
	}

	group.Name = body.Name
	group.ID = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	h.logger.Debugf("Creating group with ID %s", group.ID)
	err = h.mongoCollection.Create(&group)
	if err != nil {
		h.logger.Errorf("Error creating group: %v", err)
		return &group, fuego.InternalServerError{Title: "Error creating group", Err: err}
	}

	h.logger.Debugf("Created group with ID %s", group.ID)
	return &group, nil
}
