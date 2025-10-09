package entity

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `json:"name"`
	GroupID   uuid.UUID `gorm:"type:uuid;index" json:"groupId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
