package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=SOA host=database1 port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Encounter{}, &model.SocialEncounter{}, &model.HiddenLocationEncounter{}, &model.EncounterExecution{})
	return database
}

func startServer(handlerEnc *handler.EncounterHandler, handlerExec *handler.EncounterExecutionHandler) {
	router := mux.NewRouter().StrictSlash(true) //za rukovanje http zahtevima i definisanje ruta

	//za zahteve iz c# proj ka ovamo
	router.HandleFunc("/encounters/create", handlerEnc.Create).Methods("POST")
	router.HandleFunc("/encounters/createSocialEncounter", handlerEnc.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/createHiddenLocationEncounter", handlerEnc.CreateHiddenLocationEncounter).Methods("POST")

	router.HandleFunc("/encounters/update", handlerEnc.Update).Methods("PUT")
	router.HandleFunc("/encounters/updateHiddenLocationEncounter", handlerEnc.UpdateHiddenLocationEncounter).Methods("PUT")
	router.HandleFunc("/encounters/updateSocialEncounter", handlerEnc.UpdateSocialEncounter).Methods("PUT")

	router.HandleFunc("/encounters", handlerEnc.GetAllEncounters).Methods("GET")
	router.HandleFunc("/hiddenLocationEncounters", handlerEnc.GetAllHiddenLocationEncounters).Methods("GET")
	router.HandleFunc("/socialEncounters", handlerEnc.GetAllSocialEncounters).Methods("GET")

	router.HandleFunc("/encounters/getEncounterById/{encounterId}", handlerEnc.GetEncounterById).Methods("GET")

	router.HandleFunc("/encounters/getSocialEncounterId/{baseEncounterId}", handlerEnc.GetSocialEncounterId).Methods("GET")
	router.HandleFunc("/encounters/getHiddenLocationEncounterId/{baseEncounterId}", handlerEnc.GetHiddenLocationEncounterId).Methods("GET")
	router.HandleFunc("/encounters/getHiddenLocationEncounter/{encounterId}", handlerEnc.GetHiddenLocationEncounterByEncounterId).Methods("GET")

	router.HandleFunc("/encounters/deleteEncounter/{baseEncounterId}", handlerEnc.DeleteEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteSocialEncounter/{socialEncounterId}", handlerEnc.DeleteSocialEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteHiddenLocationEncounter/{hiddenLocationEncounterId}", handlerEnc.DeleteHiddenLocationEncounter).Methods("DELETE")

	// Encounter Execution
	router.HandleFunc("/encounterExecution", handlerExec.GetAll).Methods("GET")
	router.HandleFunc("/encounterExecution/getActive/{userId}", handlerExec.GetExecutionByUser).Methods("GET")
	router.HandleFunc("/encounterExecution/completeExecution/{userId}", handlerExec.CompleteEncounter).Methods("GET")
	router.HandleFunc("/encounterExecution/create", handlerExec.Create).Methods("POST")
	router.HandleFunc("/encounterExecution/update/{id}", handlerExec.Update).Methods("PUT")
	router.HandleFunc("/encounterExecution/delete/{id}", handlerExec.Delete).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	encounterRepo := &repo.EncounterRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{EncounterRepo: encounterRepo}
	encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}

	encounterExecutionRepo := &repo.EncounterExecutionRepository{DatabaseConnection: database}
	encounterExecutionService := &service.EncounterExecutionService{EncounterExecutionRepo: encounterExecutionRepo}
	encounterExecutionHandler := &handler.EncounterExecutionHandler{EncounterExecutionService: encounterExecutionService}

	startServer(encounterHandler, encounterExecutionHandler)
}
