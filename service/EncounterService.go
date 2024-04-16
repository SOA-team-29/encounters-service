package service

import (
	"database-example/model"
	"database-example/repo"
	"errors"
)

var ErrEncounterNotFound = errors.New("encounter not found")

type EncounterService struct {
	EncounterRepo *repo.EncounterRepository
}

func (service *EncounterService) Create(encounter *model.Encounter) (*model.Encounter, error) {
	createdEncounter, err := service.EncounterRepo.CreateEncounter(encounter)
	if err != nil {
		return nil, err
	}

	return createdEncounter, nil
}

func (service *EncounterService) CreateSocialEncounter(encounter *model.SocialEncounter) error {
	err := service.EncounterRepo.CreateSocialEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterService) CreateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {
	err := service.EncounterRepo.CreateHiddenLocationEncounter(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (s *EncounterService) GetAllEncounters() ([]*model.Encounter, error) {
	// Poziv baze podataka ili nekog drugog skladišta podataka da dobijemo sve susrete
	encounters, err := s.EncounterRepo.GetAllEncounters()
	if err != nil {
		// Ukoliko dođe do greške, vraćamo praznu listu i grešku
		return nil, err
	}

	return encounters, nil
}

func (s *EncounterService) GetAllHiddenLocationEncounters() ([]*model.HiddenLocationEncounter, error) {
	encounters, err := s.EncounterRepo.GetAllHiddenLocationEncounters()
	if err != nil {
		return nil, err
	}

	return encounters, nil
}

func (s *EncounterService) GetAllSocialEncounters() ([]*model.SocialEncounter, error) {
	encounters, err := s.EncounterRepo.GetAllSocialEncounters()
	if err != nil {
		return nil, err
	}

	return encounters, nil
}

func (s *EncounterService) Update(encounter *model.Encounter) error {
	// Ažuriranje susreta u repozitorijumu
	err := s.EncounterRepo.Update(encounter)
	if err != nil {
		// Provera da li je susret pronađen
		if errors.Is(err, ErrEncounterNotFound) {
			return ErrEncounterNotFound
		}
		// Vraćanje drugih grešaka ako se nešto drugo dogodi
		return err
	}

	return nil
}

func (s *EncounterService) UpdateHiddenLocationEncounter(encounter *model.HiddenLocationEncounter) error {

	err := s.EncounterRepo.UpdateHiddenLocationEncounter(encounter)
	if err != nil {
		if errors.Is(err, ErrEncounterNotFound) {
			return ErrEncounterNotFound
		}
		return err
	}

	return nil
}

func (s *EncounterService) UpdateSocialEncounter(encounter *model.SocialEncounter) error {

	err := s.EncounterRepo.UpdateSocialEncounter(encounter)
	if err != nil {
		if errors.Is(err, ErrEncounterNotFound) {
			return ErrEncounterNotFound
		}
		return err
	}

	return nil
}

func (s *EncounterService) DeleteEncounter(baseEncounterID string) error {

	err := s.EncounterRepo.DeleteEncounter(baseEncounterID)
	if err != nil {
		return err
	}
	return nil
}

/*


func (s *EncounterService) GetHiddenLocationEncounterId(baseEncounterID int) (int, error) {
	// Pozivamo odgovarajuću funkciju u repozitorijumu za pronalaženje ID-a skrivenog lokacijskog susreta
	hiddenLocationEncounterID, err := s.EncounterRepo.GetHiddenLocationEncounterId(baseEncounterID)
	if err != nil {
		// Provera da li je susret pronađen
		if errors.Is(err, ErrEncounterNotFound) {
			return -1, ErrEncounterNotFound
		}
		// Vraćanje drugih grešaka ako se nešto drugo dogodi
		return -1, err
	}

	return hiddenLocationEncounterID, nil
}

// GetSocialEncounterID vraća ID socijalnog susreta na osnovu ID-a osnovnog susreta
func (s *EncounterService) GetSocialEncounterId(baseEncounterID int) (int, error) {
	// Pozivamo odgovarajuću funkciju u repozitorijumu za pronalaženje ID-a socijalnog susreta
	socialEncounterID, err := s.EncounterRepo.GetSocialEncounterId(baseEncounterID)
	if err != nil {
		// Provera da li je susret pronađen
		if errors.Is(err, ErrEncounterNotFound) {
			return -1, ErrEncounterNotFound
		}
		// Vraćanje drugih grešaka ako se nešto drugo dogodi
		return -1, err
	}

	return socialEncounterID, nil
}

func (s *EncounterService) DeleteSocialEncounter(socialEncounterID int) error {
	// Poziv funkcije za brisanje socijalnog susreta iz repozitorijuma
	err := s.EncounterRepo.DeleteSocialEncounter(socialEncounterID)
	if err != nil {
		// Ukoliko dođe do greške prilikom brisanja, prosledi je nadole
		return err
	}
	return nil
}

func (s *EncounterService) DeleteHiddenLocationEncounter(hiddenLocationEncounterID int) error {
	// Poziv funkcije za brisanje skrivenog susreta iz repozitorijuma
	err := s.EncounterRepo.DeleteHiddenLocationEncounter(hiddenLocationEncounterID)
	if err != nil {
		// Ukoliko dođe do greške prilikom brisanja, prosledi je nadole
		return err
	}
	return nil
}



func (s *EncounterService) GetHiddenLocationEncounterByEncounterId(encounterId int) (*model.HiddenLocationEncounter, error) {
	// Call the repository method to fetch the hidden location encounter by encounter ID
	hiddenLocationEncounter, err := s.EncounterRepo.GetHiddenLocationEncounterByEncounterId(encounterId)
	if err != nil {
		// If there's an error, return nil and the error
		return nil, err
	}

	return hiddenLocationEncounter, nil
}

// (s *EncounterService) - prijemnik funkcije, na kom tipu (objektu) je funkcija definisana
// (encounterId int) - parametri funkcije
// (*model.Encounter, error) - povratni tipovi funkcije
func (s *EncounterService) GetEncounterById(encounterId int) (*model.Encounter, error) {
	encounter, err := s.EncounterRepo.GetEncounterById(encounterId)
	if err != nil {
		return nil, err
	}

	return encounter, nil
}*/
