package handler

import (
	"time"

	"wgplanner/internal/entity"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
		return nil, fuego.BadRequestError{Title: "Invalid request body", Err: err}
	}

	group.Name = body.Name
	group.ID = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	h.logger.Debugf("Creating group with ID %s", group.ID)
	err = h.mongoCollection.Create(&group)
	if err != nil {
		h.logger.Errorf("Error creating group: %v", err)
		return nil, fuego.InternalServerError{Title: "Error creating group", Err: err}
	}

	h.logger.Debugf("Created group with ID %s", group.ID)
	return &group, nil
}

func (h *GroupHandler) GetGroupFromID(c fuego.ContextNoBody) (*entity.Group, error) {
	var group entity.Group

	id := c.PathParam("groupId")
	uuid, err := uuid.Parse(id)
	if err != nil {
		h.logger.Errorf("Error parsing group ID %s: %v", id, err)
		return nil, fuego.BadRequestError{Title: "Invalid group ID", Err: err}
	}

	h.logger.Debugf("Fetching group with ID %s", id)
	filter := bson.M{"id": uuid}
	err = h.mongoCollection.First(filter, &group)
	if err != nil {
		h.logger.Errorf("Error fetching group with ID %s: %v", id, err)
		return nil, fuego.NotFoundError{Title: "Group not found", Err: err}
	}

	h.logger.Debugf("Fetched group with ID %s", id)
	return &group, nil
}

type AddMemberRequest struct {
	Name string `json:"name" validate:"required"`
}

func (h *GroupHandler) AddMemberToGroup(c fuego.ContextWithBody[AddMemberRequest]) (*entity.Member, error) {
	var group entity.Group

	id := c.PathParam("groupId")
	groupId, err := uuid.Parse(id)
	if err != nil {
		h.logger.Errorf("Error parsing group ID %s: %v", id, err)
		return nil, fuego.BadRequestError{Title: "Invalid group ID", Err: err}
	}

	h.logger.Debugf("Fetching group with ID %s", id)
	filter := bson.M{"id": groupId}
	err = h.mongoCollection.First(filter, &group)
	if err != nil {
		h.logger.Errorf("Error fetching group with ID %s: %v", id, err)
		return nil, fuego.NotFoundError{Title: "Group not found", Err: err}
	}

	body, err := c.Body()
	if err != nil {
		h.logger.Errorf("Error parsing request body: %v", err)
		return nil, fuego.BadRequestError{Title: "Invalid request body", Err: err}
	}

	var member entity.Member
	member.ID = uuid.New()
	member.GroupID = group.ID
	member.Name = body.Name
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()

	group.Members = append(group.Members, member)
	group.UpdatedAt = time.Now()

	h.logger.Debugf("Updating group with ID %s", id)
	err = h.mongoCollection.Update(&group)
	if err != nil {
		h.logger.Errorf("Error updating group with ID %s: %v", id, err)
		return nil, fuego.InternalServerError{Title: "Error updating group", Err: err}
	}

	h.logger.Debugf("Added member with ID %s to group with ID %s", member.ID, id)
	return &member, nil
}
