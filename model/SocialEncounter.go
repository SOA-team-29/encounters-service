package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SocialEncounter struct {
	ID                            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EncounterID                   string             `json:"encounterId"`
	TouristsRequiredForCompletion int                `json:"touristsRequiredForCompletion"`
	DistanceTreshold              float64            `json:"distanceTreshold"`
	TouristIDs                    []int              `json:"touristIDs"`
}
