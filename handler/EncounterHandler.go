package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func (handler *EncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Encounter sa id-em %s", id)
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(req.Body).Decode(&encounter) //dekodiranje json zahteva
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.Create(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) CreateSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	//ResponseWriter - pisanje odgovora
	//Request - dolazni zahtev
	var encounter model.SocialEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.CreateSocialEncounter(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) CreateHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {

	var encounter model.HiddenLocationEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter) //dekodiranje json zahteva
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.CreateHiddenLocationEncounter(&encounter)
	if err != nil {
		println("Error while creating a new encounter")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (h *EncounterHandler) GetAllEncounters(w http.ResponseWriter, r *http.Request) {
	encounters, err := h.EncounterService.GetAllEncounters()
	if err != nil {
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}

	// Konvertujemo susrete u JSON format
	encountersJSON, err := json.Marshal(encounters)
	if err != nil {
		http.Error(w, "Error converting encounters to JSON", http.StatusInternalServerError)
		return
	}

	// Postavljamo Content-Type zaglavlje na application/json
	w.Header().Set("Content-Type", "application/json")

	// Šaljemo odgovor sa susretima u JSON formatu
	w.WriteHeader(http.StatusOK)
	w.Write(encountersJSON)
}

func (h *EncounterHandler) GetAllSocialEncounters(w http.ResponseWriter, r *http.Request) {
	// Ovde bi trebalo da dobijemo sve susrete iz baze podataka
	encounters, err := h.EncounterService.GetAllSocialEncounters()
	if err != nil {
		// Ukoliko dođe do greške prilikom dobijanja susreta, vraćamo odgovarajući status i poruku o grešci
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}

	// Konvertujemo susrete u JSON format
	encountersJSON, err := json.Marshal(encounters)
	if err != nil {
		// Ukoliko dođe do greške prilikom konvertovanja u JSON, vraćamo odgovarajući status i poruku o grešci
		http.Error(w, "Error converting encounters to JSON", http.StatusInternalServerError)
		return
	}

	// Postavljamo Content-Type zaglavlje na application/json
	w.Header().Set("Content-Type", "application/json")

	// Šaljemo odgovor sa susretima u JSON formatu
	w.WriteHeader(http.StatusOK)
	w.Write(encountersJSON)
}

func (h *EncounterHandler) GetAllHiddenLocationEncounters(w http.ResponseWriter, r *http.Request) {
	// Ovde bi trebalo da dobijemo sve susrete iz baze podataka
	encounters, err := h.EncounterService.GetAllHiddenLocationEncounters()
	if err != nil {
		// Ukoliko dođe do greške prilikom dobijanja susreta, vraćamo odgovarajući status i poruku o grešci
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}

	// Konvertujemo susrete u JSON format
	encountersJSON, err := json.Marshal(encounters)
	if err != nil {
		// Ukoliko dođe do greške prilikom konvertovanja u JSON, vraćamo odgovarajući status i poruku o grešci
		http.Error(w, "Error converting encounters to JSON", http.StatusInternalServerError)
		return
	}

	// Postavljamo Content-Type zaglavlje na application/json
	w.Header().Set("Content-Type", "application/json")

	// Šaljemo odgovor sa susretima u JSON formatu
	w.WriteHeader(http.StatusOK)
	w.Write(encountersJSON)
}

func (handler *EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.Update(&encounter)
	if err != nil {
		println("Error while updating the encounter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) UpdateHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	var encounter model.HiddenLocationEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.UpdateHiddenLocationEncounter(&encounter)
	if err != nil {
		println("Error while updating the encounter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) UpdateSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	var encounter model.SocialEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		println("Error while parsing json")
		println("Greska:", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.UpdateSocialEncounter(&encounter)
	if err != nil {
		println("Error while updating the encounter")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) GetSocialEncounterId(writer http.ResponseWriter, req *http.Request) {
	baseEncounterId, err := strconv.Atoi(mux.Vars(req)["baseEncounterId"])
	if err != nil {
		log.Println("Error converting baseEncounterId to int:", err)
		http.Error(writer, "Invalid baseEncounterId", http.StatusBadRequest)
		return
	}

	socialEncounterId, err := handler.EncounterService.GetSocialEncounterId(baseEncounterId)
	if err != nil {
		log.Println("Error getting social encounter ID:", err)
		http.Error(writer, "Error getting social encounter ID", http.StatusInternalServerError)
		return
	}

	response := struct {
		SocialEncounterId int `json:"socialEncounterId"`
	}{
		SocialEncounterId: socialEncounterId,
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func (handler *EncounterHandler) GetHiddenLocationEncounterId(writer http.ResponseWriter, req *http.Request) {
	baseEncounterId, err := strconv.Atoi(mux.Vars(req)["baseEncounterId"])
	if err != nil {
		log.Println("Error converting baseEncounterId to int:", err)
		http.Error(writer, "Invalid baseEncounterId", http.StatusBadRequest)
		return
	}

	hiddenLocationEncounterId, err := handler.EncounterService.GetHiddenLocationEncounterId(baseEncounterId)
	if err != nil {
		log.Println("Error getting hidden location encounter ID:", err)
		http.Error(writer, "Error getting hidden location encounter ID", http.StatusInternalServerError)
		return
	}

	response := struct {
		HiddenLocationEncounterId int `json:"hiddenLocationEncounterId"`
	}{
		HiddenLocationEncounterId: hiddenLocationEncounterId,
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func (handler *EncounterHandler) DeleteSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	// Dobijanje ID-a socijalnog susreta iz URL putanje
	vars := mux.Vars(req)
	socialEncounterID, err := strconv.Atoi(vars["socialEncounterId"])
	if err != nil {
		log.Println("Error converting socialEncounterId to integer:", err)
		http.Error(writer, "Invalid socialEncounterId", http.StatusBadRequest)
		return
	}

	// Poziv metode u servisu za brisanje socijalnog susreta
	err = handler.EncounterService.DeleteSocialEncounter(socialEncounterID)
	if err != nil {
		log.Println("Error while deleting the social encounter:", err)
		http.Error(writer, "Error while deleting the social encounter", http.StatusInternalServerError)
		return
	}

	// Uspesan odgovor
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) DeleteHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	// Dobijanje ID-a skrivenog susreta iz URL putanje
	vars := mux.Vars(req)
	hiddenLocationEncounterID, err := strconv.Atoi(vars["hiddenLocationEncounterId"])
	if err != nil {
		log.Println("Error converting hiddenLocationEncounterId to integer:", err)
		http.Error(writer, "Invalid hiddenLocationEncounterId", http.StatusBadRequest)
		return
	}

	// Poziv metode u servisu za brisanje skrivenog susreta
	err = handler.EncounterService.DeleteHiddenLocationEncounter(hiddenLocationEncounterID)
	if err != nil {
		log.Println("Error while deleting the hidden location encounter:", err)
		http.Error(writer, "Error while deleting the hidden location encounter", http.StatusInternalServerError)
		return
	}

	// Uspesan odgovor
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) DeleteEncounter(writer http.ResponseWriter, req *http.Request) {
	// Uzimanje ID-ja susreta iz putanje zahteva
	vars := mux.Vars(req)
	baseEncounterID, err := strconv.Atoi(vars["baseEncounterId"])
	if err != nil {
		log.Println("Error converting baseEncounterId to int:", err)
		http.Error(writer, "Invalid baseEncounterId", http.StatusBadRequest)
		return
	}

	// Pozivanje odgovarajuće funkcije za brisanje susreta iz servisa
	err = handler.EncounterService.DeleteEncounter(baseEncounterID)
	if err != nil {
		log.Println("Error deleting encounter:", err)
		http.Error(writer, "Error deleting encounter", http.StatusInternalServerError)
		return
	}

	// Ako je brisanje uspešno, vraćamo status 204 No Content
	writer.WriteHeader(http.StatusNoContent)
}

// GetHiddenLocationEncounterByEncounterId handles the GET request for getting a hidden location encounter by encounter ID
func (handler *EncounterHandler) GetHiddenLocationEncounterByEncounterId(w http.ResponseWriter, r *http.Request) {
	// Extract the encounterId from the URL parameters
	vars := mux.Vars(r)
	encounterIdStr := vars["encounterId"]

	// Convert encounterIdStr to int
	encounterId, err := strconv.Atoi(encounterIdStr)
	if err != nil {
		http.Error(w, "Invalid encounterId", http.StatusBadRequest)
		return
	}

	// Call the method from the service package to retrieve the hidden location encounter by encounter ID
	hiddenLocationEncounter, err := handler.EncounterService.GetHiddenLocationEncounterByEncounterId(encounterId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(hiddenLocationEncounter)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(responseJSON)
}

// ODAVDE IDI DALJE U SERVISE I REPO
func (handler *EncounterHandler) GetEncounterById(w http.ResponseWriter, r *http.Request) {
	//Ekstrahovanje parametara iz URL-a ili tela zahteva, ako je potrebno
	//Pozivanje odgovarajuće funkcionalnosti iz servisnog sloja ili repozitorijuma kako bi se dobio traženi susret
	//Pretvaranje dobijenih podataka u odgovarajući format (npr. JSON) kako bi se poslali nazad klijentu
	//Slanje odgovora nazad klijentu putem http.ResponseWriter

	//mux: izvlacenje varijabli iz url parametra
	//encounterId - parametar putanje
	vars := mux.Vars(r)
	encounterIdStr := vars["encounterId"] //izvucena vrednost se cuva kao string

	//konvertovanje stringa u int
	encounterId, err := strconv.Atoi(encounterIdStr)
	if err != nil {
		http.Error(w, "Invalid encounterId", http.StatusBadRequest)
		return
	}

	//poziv metode servisa da se dobavi encounter na osnovu encounterId
	encounter, err := handler.EncounterService.GetEncounterById(encounterId)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//konvertovanje odgovora (encounter) u json - marshal
	responseJSON, err := json.Marshal(encounter)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	//postavlja se Content-Type zaglavlje HTTP odgovora na application/json, što označava da je odgovor JSON
	w.Header().Set("Content-Type", "application/json")

	//json odgovor se pise u http.ResponseWriter sto ce se proslediti kao odgovor
	w.Write(responseJSON)
}
