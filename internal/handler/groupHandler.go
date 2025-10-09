package handler

import (
	"time"

	"wgplanner/internal/entity"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GroupHandler struct {
	logger *logrus.Logger
	db     *gorm.DB
}

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

func NewGroupHandler(logger *logrus.Logger, db *gorm.DB) *GroupHandler {
	db.AutoMigrate(&entity.Group{}, &entity.Member{})
	return &GroupHandler{
		logger: logger,
		db:     db,
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
	if err := h.db.Create(&group).Error; err != nil {
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
	if err := h.db.Preload("Members").First(&group, "id = ?", uuid).Error; err != nil {
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

	body, err := c.Body()
	if err != nil {
		h.logger.Errorf("Error parsing request body: %v", err)
		return nil, fuego.BadRequestError{Title: "Invalid request body", Err: err}
	}

	if err := h.db.First(&group, "id = ?", groupId).Error; err != nil {
		h.logger.Errorf("Error fetching group with ID %s: %v", id, err)
		return nil, fuego.NotFoundError{Title: "Group not found", Err: err}
	}

	newMember := entity.Member{
		ID:        uuid.New(),
		Name:      body.Name,
		GroupID:   groupId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.db.Create(&newMember).Error; err != nil {
		h.logger.Errorf("Error creating member for group %s: %v", id, err)
		return nil, fuego.InternalServerError{Title: "Error adding member", Err: err}
	}

	h.logger.Debugf("Added member with ID %s to group with ID %s", newMember.ID, id)
	return &newMember, nil
}

type RemoveMemberRequest struct {
	MemberID string `json:"member_id" validate:"dive,uuid"`
}

type RemoveMemberResponse struct {
	Message string `json:"message"`
}

func (h *GroupHandler) RemoveMembersFromGroup(c fuego.ContextNoBody) (*RemoveMemberResponse, error) {
	id := c.PathParam("groupId")
	groupId, err := uuid.Parse(id)
	if err != nil {
		h.logger.Errorf("Error parsing group ID %s: %v", id, err)
		return nil, fuego.BadRequestError{Title: "Invalid group ID", Err: err}
	}

	if err := h.db.First(&entity.Group{}, "id = ?", groupId).Error; err != nil {
		h.logger.Errorf("Error fetching group with ID %s: %v", id, err)
		return nil, fuego.NotFoundError{Title: "Group not found", Err: err}
	}

	memberIDParam := c.PathParam("memberId")
	memberID, err := uuid.Parse(memberIDParam)
	if err != nil {
		h.logger.Errorf("Error parsing member ID %s: %v", memberIDParam, err)
		return nil, fuego.BadRequestError{Title: "Invalid member ID", Err: err}
	}

	var member entity.Member
	if err := h.db.First(&member, "id = ? AND group_id = ?", memberID, groupId).Error; err != nil {
		h.logger.Errorf("Error fetching member with ID %s in group %s: %v", memberID, id, err)
		return nil, fuego.NotFoundError{Title: "Member not found in group", Err: err}
	}

	if err := h.db.Delete(&member).Error; err != nil {
		h.logger.Errorf("Error removing member with ID %s from group %s: %v", memberID, id, err)
		return nil, fuego.InternalServerError{Title: "Error removing member", Err: err}
	}

	response := &RemoveMemberResponse{
		Message: "Member removed successfully",
	}
	h.logger.Debugf("Removed member with ID %s from group with ID %s", memberID, id)
	return response, nil
}
