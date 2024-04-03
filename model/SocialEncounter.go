package model

type SocialEncounter struct {
	ID                            int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	EncounterID                   int     `json:"encounterId" gorm:"foreignKey:EncounterID"`
	TouristsRequiredForCompletion int     `json:"touristsRequiredForCompletion"`
	DistanceTreshold              float64 `json:"distanceTreshold"`
	TouristIDs                    []int   `json:"touristIDs" gorm:"type:integer[]"`
}
