package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel"
)

type EncounterHandler struct {
	EncounterService *service.EncounterService
}

func (handler *EncounterHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	log.Printf("INFO: Request to get encounter with id: %s", id)
	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterHandler) Create(writer http.ResponseWriter, req *http.Request) {
	log.Println("INFO: Entered Create Encounter handler")
	ctx := req.Context()
	tracer := otel.Tracer("encounters-service")
	_, span := tracer.Start(ctx, "CreateEncounter")
	defer span.End()

	var encounter model.Encounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdEncounter, err := handler.EncounterService.Create(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to create encounter: %v", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	log.Printf("INFO: Created encounter: %v", createdEncounter)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"id":               createdEncounter.ID.Hex(),
		"name":             createdEncounter.Name,
		"description":      createdEncounter.Description,
		"xpPoints":         createdEncounter.XpPoints,
		"status":           createdEncounter.Status,
		"type":             createdEncounter.Type,
		"latitude":         createdEncounter.Latitude,
		"longitude":        createdEncounter.Longitude,
		"shouldBeApproved": createdEncounter.ShouldBeApproved,
	}
	log.Printf("INFO: Response data: %v", response)
	json.NewEncoder(writer).Encode(response)
}

func (handler *EncounterHandler) CreateSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	log.Println("INFO: Entered Create Social Encounter handler")
	var encounter model.SocialEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EncounterService.CreateSocialEncounter(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to create social encounter: %v", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	log.Printf("INFO: Created social encounter: %v", encounter)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) CreateHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	log.Println("INFO: Entered Create Hidden Location Encounter handler")
	var encounter model.HiddenLocationEncounter
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("INFO: Parsed encounter data: %v", encounter)
	err = handler.EncounterService.CreateHiddenLocationEncounter(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to create hidden location encounter: %v", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	log.Printf("INFO: Created hidden location encounter: %v", encounter)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (h *EncounterHandler) GetAllEncounters(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO: Entered Get All Encounters handler")
	ctx := r.Context()
	tracer := otel.Tracer("encounters-service")
	_, span := tracer.Start(ctx, "GetAllEncounters")
	defer span.End()
	encounters, err := h.EncounterService.GetAllEncounters()
	if err != nil {
		log.Printf("ERROR: Failed to get encounters: %v", err)
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Retrieved encounters: %v", encounters)

	modifiedJSON := modifyEncountersJSON(encounters)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(modifiedJSON))
}

func modifyEncountersJSON(encounters []*model.Encounter) string {

	var modifiedJSON strings.Builder
	modifiedJSON.WriteString("[")
	for i, encounter := range encounters {
		encounterJSON, err := json.Marshal(encounter)
		if err != nil {
			log.Printf("Error marshaling encounter %d: %s\n", i, err.Error())
			continue
		}

		encounterJSONString := string(encounterJSON)
		encounterJSONString = strings.Replace(encounterJSONString, "\"_id\"", "\"id\"", 1)

		modifiedJSON.WriteString(encounterJSONString)

		if i < len(encounters)-1 {
			modifiedJSON.WriteString(",")
		}
	}
	modifiedJSON.WriteString("]")
	return modifiedJSON.String()
}

func (h *EncounterHandler) GetAllSocialEncounters(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO: Entered Get All Social Encounters handler")
	encounters, err := h.EncounterService.GetAllSocialEncounters()
	if err != nil {
		log.Printf("ERROR: Failed to get social encounters: %v", err)
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Retrieved social encounters: %v", encounters)

	modifiedJSON := modifyEncountersJSONsoc(encounters)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(modifiedJSON))
}

func modifyEncountersJSONsoc(encounters []*model.SocialEncounter) string {

	var modifiedJSON strings.Builder
	modifiedJSON.WriteString("[")
	for i, encounter := range encounters {
		encounterJSON, err := json.Marshal(encounter)
		if err != nil {
			log.Printf("ERROR: Error marshaling encounter %d: %s\n", i, err.Error())
			continue
		}

		encounterJSONString := string(encounterJSON)
		encounterJSONString = strings.Replace(encounterJSONString, "\"_id\"", "\"id\"", 1)

		modifiedJSON.WriteString(encounterJSONString)

		if i < len(encounters)-1 {
			modifiedJSON.WriteString(",")
		}
	}
	modifiedJSON.WriteString("]")
	return modifiedJSON.String()
}

func (h *EncounterHandler) GetAllHiddenLocationEncounters(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO: Entered Get All Hidden Location Encounters handler")
	encounters, err := h.EncounterService.GetAllHiddenLocationEncounters()
	if err != nil {
		log.Printf("ERROR: Failed to get hidden location encounters: %v", err)
		http.Error(w, "Error getting encounters", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Retrieved hidden location encounters: %v", encounters)

	modifiedJSON := modifyEncountersJSONhidd(encounters)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(modifiedJSON))
}

func modifyEncountersJSONhidd(encounters []*model.HiddenLocationEncounter) string {

	var modifiedJSON strings.Builder
	modifiedJSON.WriteString("[")
	for i, encounter := range encounters {
		encounterJSON, err := json.Marshal(encounter)
		if err != nil {
			log.Printf("ERROR: Error marshaling encounter %d: %s\n", i, err.Error())
			continue
		}

		encounterJSONString := string(encounterJSON)
		encounterJSONString = strings.Replace(encounterJSONString, "\"_id\"", "\"id\"", 1)

		modifiedJSON.WriteString(encounterJSONString)

		if i < len(encounters)-1 {
			modifiedJSON.WriteString(",")
		}
	}
	modifiedJSON.WriteString("]")
	return modifiedJSON.String()
}

func (handler *EncounterHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var encounter model.Encounter
	log.Println("INFO: Entered Update Encounter handler")
	// Dekodiranje JSON-a u mapu kao intermedijernu strukturu
	var encounterMap map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&encounterMap)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("INFO: Parsed encounter data: %v", encounterMap)
	// Konverzija ID-a iz stringa u primitive.ObjectID
	if id, ok := encounterMap["Id"].(string); ok {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Printf("ERROR: Failed to convert ID to ObjectID: %v", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		encounterMap["Id"] = objID
		log.Printf("INFO: Converted ID to ObjectID: %v", objID)
	}

	// Konvertovanje mape u strukturu
	err = mapstructure.Decode(encounterMap, &encounter)
	if err != nil {
		log.Printf("ERROR: Failed to decode JSON to struct: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.Update(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to update encounter: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Updated encounter: %v", encounter)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) UpdateHiddenLocationEncounter(writer http.ResponseWriter, req *http.Request) {
	var encounter model.HiddenLocationEncounter
	log.Println("INFO: Entered Update Hidden Location Encounter handler")
	// Dekodiranje JSON-a u mapu kao intermedijernu strukturu
	var encounterMap map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&encounterMap)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("INFO: Parsed hidden location encounter data: %v", encounterMap)
	// Konverzija ID-a iz stringa u primitive.ObjectID
	if id, ok := encounterMap["Id"].(string); ok {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Printf("ERROR: Failed to convert ID to ObjectID: %v", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		encounterMap["Id"] = objID
		log.Printf("INFO: Converted ID to ObjectID: %v", objID)
	}

	// Konvertovanje mape u strukturu
	err = mapstructure.Decode(encounterMap, &encounter)
	if err != nil {
		log.Printf("ERROR: Failed to decode JSON to struct: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.UpdateHiddenLocationEncounter(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to update hidden location encounter: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Updated hidden location encounter: %v", encounter)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) UpdateSocialEncounter(writer http.ResponseWriter, req *http.Request) {
	var encounter model.SocialEncounter
	log.Println("INFO: Entered Update Social Encounter handler")
	var encounterMap map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&encounterMap)
	if err != nil {
		log.Printf("ERROR: Failed to parse JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("INFO: Parsed social encounter data: %v", encounterMap)
	if id, ok := encounterMap["Id"].(string); ok {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Printf("ERROR: Failed to convert ID to ObjectID: %v", err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		encounterMap["Id"] = objID
		log.Printf("INFO: Converted ID to ObjectID: %v", objID)
	}

	err = mapstructure.Decode(encounterMap, &encounter)
	if err != nil {
		log.Printf("ERROR: Failed to decode JSON to struct: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterService.UpdateSocialEncounter(&encounter)
	if err != nil {
		log.Printf("ERROR: Failed to update social encounter: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Updated social encounter: %v", encounter)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterHandler) DeleteEncounter(writer http.ResponseWriter, req *http.Request) {
	log.Println("INFO: Entered Delete Encounter handler")
	vars := mux.Vars(req)
	baseEncounterID := vars["baseEncounterId"]
	log.Printf("INFO: Base Encounter ID to delete: %s", baseEncounterID)

	err := handler.EncounterService.DeleteEncounter(baseEncounterID)
	if err != nil {
		log.Printf("ERROR: Error deleting encounter with ID %s: %v", baseEncounterID, err)
		http.Error(writer, "Error deleting encounter", http.StatusInternalServerError)
		return
	}
	log.Printf("INFO: Successfully deleted encounter with ID %s", baseEncounterID)

	// Ako je brisanje uspešno, vraćamo status 204 No Content
	writer.WriteHeader(http.StatusNoContent)
}

/*

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
}*/
