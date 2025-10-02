package entity

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	ID               uuid.UUID `json:"id" bson:"id,omitempty"`
	Name             string    `json:"name" bson:"name"`
}
