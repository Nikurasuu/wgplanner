package entity

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `json:"name"`
	Members   []Member  `gorm:"foreignKey:GroupID" json:"members"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
