package entity

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

type Member struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uuid.UUID `json:"id" bson:"id"`
	GroupID          uuid.UUID `json:"group_id" bson:"group_id"`
	Name             string    `json:"name" bson:"name"`
}
