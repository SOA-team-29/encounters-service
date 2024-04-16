package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EncounterStatus int

const (
	Draft EncounterStatus = iota
	Archived
	Active
)

type EncounterType int

const (
	Social EncounterType = iota
	Location
	Misc
)

type Encounter struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	XpPoints         int                `json:"xpPoints"`
	Status           string             `json:"status"`
	Type             string             `json:"type"`
	Latitude         float64            `json:"latitude"`
	Longitude        float64            `json:"longitude"`
	ShouldBeApproved bool               `json:"shouldBeApproved"`
}
