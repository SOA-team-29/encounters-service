package repo

import (
	"context"
	"database-example/model"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EncounterRepository struct {
	//DatabaseConnection *gorm.DB //za konekciju sa bazom podataka
	DatabaseConnection *mongo.Client
}

func (repo *EncounterRepository) CreateEncounter(encounter *model.Encounter) (*model.Encounter, error) {
	collection := repo.DatabaseConnection.Database("SOAencounters").Collection("encounters")

	ctx := context.TODO()
	encounter.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, encounter)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: encounter.ID}}

	var createdEncounter model.Encounter
	err = collection.FindOne(ctx, filter).Decode(&createdEncounter)
	if err != nil {
		return nil, err
	}

	return &createdEncounter, nil
}

func (repo *EncounterRepository) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	collection := repo.DatabaseConnection.Database("SOAencounters").Collection("socialEncounters")

	ctx := context.TODO()

	_, err := collection.InsertOne(ctx, encounter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *EncounterRepository) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	collection := repo.DatabaseConnection.Database("SOAencounters").Collection("hiddenLocationEncounters")

	ctx := context.TODO()

	_, err := collection.InsertOne(ctx, encounter)
	if err != nil {
		return err
	}
	return nil
}

func (r *EncounterRepository) GetAllEncounters() ([]*model.Encounter, error) {

	filter := bson.D{}

	cursor, err := r.DatabaseConnection.Database("SOAencounters").Collection("encounters").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var encounters []*model.Encounter
	for cursor.Next(context.Background()) {
		var encounter model.Encounter
		if err := cursor.Decode(&encounter); err != nil {
			return nil, err
		}

		encounters = append(encounters, &encounter)
	}

	return encounters, nil
}

func (r *EncounterRepository) GetAllHiddenLocationEncounters() ([]*model.HiddenLocationEncounter, error) {
	filter := bson.D{}

	cursor, err := r.DatabaseConnection.Database("SOAencounters").Collection("hiddenLocationEncounters").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var encounters []*model.HiddenLocationEncounter
	for cursor.Next(context.Background()) {
		var encounter model.HiddenLocationEncounter
		if err := cursor.Decode(&encounter); err != nil {
			return nil, err
		}

		encounters = append(encounters, &encounter)
	}

	return encounters, nil
}

func (r *EncounterRepository) GetAllSocialEncounters() ([]*model.SocialEncounter, error) {
	filter := bson.D{}

	cursor, err := r.DatabaseConnection.Database("SOAencounters").Collection("socialEncounters").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var encounters []*model.SocialEncounter
	for cursor.Next(context.Background()) {
		var encounter model.SocialEncounter
		if err := cursor.Decode(&encounter); err != nil {
			return nil, err
		}

		encounters = append(encounters, &encounter)
	}

	return encounters, nil
}

func (repo *EncounterRepository) Update(encounter *model.Encounter) error {
	filter := bson.M{"_id": encounter.ID}

	update := bson.M{
		"$set": bson.M{
			"name":             encounter.Name,
			"description":      encounter.Description,
			"xppoints":         encounter.XpPoints,
			"status":           encounter.Status,
			"type":             encounter.Type,
			"longitude":        encounter.Longitude,
			"latitude":         encounter.Latitude,
			"shouldbeapproved": encounter.ShouldBeApproved,
		},
	}

	_, err := repo.DatabaseConnection.Database("SOAencounters").Collection("encounters").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *EncounterRepository) UpdateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {

	filter := bson.M{"_id": encounter.ID}

	update := bson.M{
		"$set": bson.M{
			"imageurl":         encounter.ImageURL,
			"imagelatitude":    encounter.ImageLatitude,
			"imagelongitude":   encounter.ImageLongitude,
			"distancetreshold": encounter.DistanceTreshold,
			"encounterid":      encounter.EncounterID,
		},
	}

	_, err := repo.DatabaseConnection.Database("SOAencounters").Collection("hiddenLocationEncounters").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *EncounterRepository) UpdateSocialEncounter(encounter *model.SocialEncounter) error {

	filter := bson.M{"_id": encounter.ID}

	update := bson.M{
		"$set": bson.M{
			"encounterid":                   encounter.EncounterID,
			"touristsrequiredforcompletion": encounter.TouristsRequiredForCompletion,
			"distancetreshold":              encounter.DistanceTreshold,
			"touristids":                    encounter.TouristIDs,
		},
	}

	_, err := repo.DatabaseConnection.Database("SOAencounters").Collection("socialEncounters").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *EncounterRepository) DeleteEncounter(baseEncounterID string) error {

	objectID, err := primitive.ObjectIDFromHex(baseEncounterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := r.DatabaseConnection.Database("SOAencounters").Collection("encounters").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("encounter not found")
	}

	hiddenLocationFilter := bson.M{"encounterid": baseEncounterID}
	_, err = r.DatabaseConnection.Database("SOAencounters").Collection("hiddenLocationEncounters").DeleteMany(context.TODO(), hiddenLocationFilter)
	if err != nil {
		return err
	}

	socialFilter := bson.M{"encounterid": baseEncounterID}
	_, err = r.DatabaseConnection.Database("SOAencounters").Collection("socialEncounters").DeleteMany(context.TODO(), socialFilter)
	if err != nil {
		return err
	}

	return nil
}

/* ono od pre jer vise ne treba jer sam uradila brisanje povezanih social i location na laksi nacin odmah u brisanju
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
*/
