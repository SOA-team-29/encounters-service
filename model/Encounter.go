package model

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
	ID               int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name             string  `json:"name" gorm:"not null;type:string"`
	Description      string  `json:"description"`
	XpPoints         int     `json:"xpPoints"`
	Status           string  `json:"status"`
	Type             string  `json:"type"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	ShouldBeApproved bool    `json:"shouldBeApproved"`
}
