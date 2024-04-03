package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EncounterExecutionHandler struct {
	EncounterExecutionService *service.EncounterExecutionService
}

func (handler *EncounterExecutionHandler) GetExecutionByUser(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr, ok := vars["userId"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encounter, err := handler.EncounterExecutionService.GetExecutionByUser(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func (handler *EncounterExecutionHandler) CompleteEncounter(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIDStr, ok := vars["userId"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encounter, err := handler.EncounterExecutionService.CompleteEncounter(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResponse)
}

func (handler *EncounterExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var encounter model.EncounterExecution
	err := json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterExecutionService.CreateEncounter(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterExecutionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	encIdStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encId, err := strconv.Atoi(encIdStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var encounter model.EncounterExecution
	err = json.NewDecoder(req.Body).Decode(&encounter)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterExecutionService.UpdateEncounter(encId, &encounter)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(encounter)
}

func (handler *EncounterExecutionHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	encIdStr, ok := vars["id"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encId, err := strconv.Atoi(encIdStr)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.EncounterExecutionService.DeleteEncounter(encId)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (handler *EncounterExecutionHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	encounters, err := handler.EncounterExecutionService.GetAllEncounters()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	encountersJson, err := json.Marshal(encounters)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(encountersJson)
}
