package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EncounterExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EncounterExecutionRepository) FindByUserId(userID int) (model.EncounterExecution, error) {
	encounter := model.EncounterExecution{}
	dbResult := repo.DatabaseConnection.Where("user_id = ?", userID).Find(&encounter)
	if dbResult.Error != nil {
		return encounter, dbResult.Error
	}
	return encounter, nil
}

func (repo *EncounterExecutionRepository) Update(encounter *model.EncounterExecution) error {
	dbResult := repo.DatabaseConnection.Save(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncounterExecutionRepository) Create(encounter *model.EncounterExecution) error {
	dbResult := repo.DatabaseConnection.Create(encounter)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncounterExecutionRepository) Delete(encounterExecId int) error {
	dbResult := repo.DatabaseConnection.Exec("DELETE FROM encounter_executions WHERE id = ?", encounterExecId)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	return nil
}

func (repo *EncounterExecutionRepository) GetAll() ([]*model.EncounterExecution, error) {
	var encounters []*model.EncounterExecution
	dbResult := repo.DatabaseConnection.Find(&encounters)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return encounters, nil
}
