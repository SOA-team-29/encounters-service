package repo

import (
	"database-example/model"
	"errors"

	"gorm.io/gorm"
)

type EncounterRepository struct {
	DatabaseConnection *gorm.DB //za konekciju sa bazom podataka
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (r *EncounterRepository) GetAllEncounters() ([]*model.Encounter, error) {
	var encounters []*model.Encounter
	if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
		return nil, err
	}

	return encounters, nil
}

func (r *EncounterRepository) GetAllHiddenLocationEncounters() ([]*model.HiddenLocationEncounter, error) {
	var encounters []*model.HiddenLocationEncounter
	if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
		return nil, err
	}

	return encounters, nil
}

func (r *EncounterRepository) GetAllSocialEncounters() ([]*model.SocialEncounter, error) {
	var encounters []*model.SocialEncounter
	if err := r.DatabaseConnection.Find(&encounters).Error; err != nil {
		return nil, err
	}

	return encounters, nil
}

func (repo *EncounterRepository) Update(encounter *model.Encounter) error {
	dbResult := repo.DatabaseConnection.Model(&model.Encounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
		"name":               encounter.Name,
		"description":        encounter.Description,
		"xp_points":          encounter.XpPoints,
		"status":             encounter.Status,
		"type":               encounter.Type,
		"longitude":          encounter.Longitude,
		"latitude":           encounter.Latitude,
		"should_be_approved": encounter.ShouldBeApproved,
	})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("encounter not found")
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) UpdateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	dbResult := repo.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
		"image_url":         encounter.ImageURL,
		"image_latitude":    encounter.ImageLatitude,
		"image_longitude":   encounter.ImageLongitude,
		"distance_treshold": encounter.DistanceTreshold,
		"encounter_id":      encounter.EncounterID,
	})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("HiddenLocationEncounter not found")
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EncounterRepository) UpdateSocialEncounter(encounter *model.SocialEncounter) error {
	dbResult := repo.DatabaseConnection.Model(&model.SocialEncounter{}).Where("id = ?", encounter.ID).Updates(map[string]interface{}{
		"tourists_required_for_completion": encounter.TouristsRequiredForCompletion,
		"distance_treshold":                encounter.DistanceTreshold,
		"tourist_ids":                      encounter.TouristIDs,
	})
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult.RowsAffected == 0 {
		return errors.New("SocialEncounter not found")
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (r *EncounterRepository) GetSocialEncounterId(baseEncounterID int) (int, error) {
	var socialEncounterID int

	result := r.DatabaseConnection.Model(&model.SocialEncounter{}).Select("id").Where("encounter_id = ?", baseEncounterID).First(&socialEncounterID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1, nil
		}
		return 0, result.Error
	}
	return socialEncounterID, nil
}

func (r *EncounterRepository) GetHiddenLocationEncounterId(baseEncounterID int) (int, error) {
	var hiddenLocationEncounterID int

	result := r.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Select("id").Where("encounter_id = ?", baseEncounterID).First(&hiddenLocationEncounterID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return -1, nil
		}
		return 0, result.Error
	}

	// Ako nema greške, vraćamo ID društvenog susreta
	return hiddenLocationEncounterID, nil
}

func (r *EncounterRepository) DeleteSocialEncounter(socialEncounterID int) error {
	result := r.DatabaseConnection.Exec("DELETE FROM social_encounters WHERE id = ?", socialEncounterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *EncounterRepository) DeleteHiddenLocationEncounter(hiddenLocationEncounterID int) error {

	result := r.DatabaseConnection.Exec("DELETE FROM hidden_location_encounters WHERE id = ?", hiddenLocationEncounterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *EncounterRepository) DeleteEncounter(baseEncounterID int) error {

	result := r.DatabaseConnection.Where("id = ?", baseEncounterID).Delete(&model.Encounter{})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}

func (r *EncounterRepository) GetHiddenLocationEncounterByEncounterId(baseEncounterID int) (*model.HiddenLocationEncounter, error) {
	var hiddenLocationEncounter model.HiddenLocationEncounter

	result := r.DatabaseConnection.Model(&model.HiddenLocationEncounter{}).Where("encounter_id = ?", baseEncounterID).First(&hiddenLocationEncounter)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &hiddenLocationEncounter, nil
}

func (r *EncounterRepository) GetEncounterById(encounterId int) (*model.Encounter, error) {
	var encounter model.Encounter

	result := r.DatabaseConnection.Model(&model.Encounter{}).Where("id = ?", encounterId).First(&encounter)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &encounter, nil
}
