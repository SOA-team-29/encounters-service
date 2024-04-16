package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type HiddenLocationEncounter struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ImageURL         string             `json:"imageURL"`
	ImageLatitude    float64            `json:"imageLatitude"`
	ImageLongitude   float64            `json:"imageLongitude"`
	DistanceTreshold float64            `json:"distanceTreshold"`
	EncounterID      string             `json:"encounterId"`
}
