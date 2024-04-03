package model

type HiddenLocationEncounter struct {
	ID               int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ImageURL         string  `json:"imageURL"`
	ImageLatitude    float64 `json:"imageLatitude"`
	ImageLongitude   float64 `json:"imageLongitude"`
	DistanceTreshold float64 `json:"distanceTreshold"`
	EncounterID      int     `json:"encounterId" gorm:"foreignKey:EncounterID"`
}
