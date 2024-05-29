package main

import (
	"context"
	"database-example/handler"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

/*
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
*/
var tp *trace.TracerProvider

func initDB() *mongo.Database {

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("SOAencounters")
	log.Println("Successfully connected to MongoDB")
	return database
}

func startServer(handlerEnc *handler.EncounterHandler) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/encounters/create", handlerEnc.Create).Methods("POST")
	router.HandleFunc("/encounters/createSocialEncounter", handlerEnc.CreateSocialEncounter).Methods("POST")
	router.HandleFunc("/encounters/createHiddenLocationEncounter", handlerEnc.CreateHiddenLocationEncounter).Methods("POST")

	router.HandleFunc("/encounters", handlerEnc.GetAllEncounters).Methods("GET")
	router.HandleFunc("/hiddenLocationEncounters", handlerEnc.GetAllHiddenLocationEncounters).Methods("GET")
	router.HandleFunc("/socialEncounters", handlerEnc.GetAllSocialEncounters).Methods("GET")

	router.HandleFunc("/encounters/update", handlerEnc.Update).Methods("PUT")
	router.HandleFunc("/encounters/updateHiddenLocationEncounter", handlerEnc.UpdateHiddenLocationEncounter).Methods("PUT")
	router.HandleFunc("/encounters/updateSocialEncounter", handlerEnc.UpdateSocialEncounter).Methods("PUT")

	router.HandleFunc("/encounters/deleteEncounter/{baseEncounterId}", handlerEnc.DeleteEncounter).Methods("DELETE")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	println("Server starting")
	log.Fatal(http.ListenAndServe(":4000", router))

}

func main() {

	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	client := database.Client()
	encounterRepo := &repo.EncounterRepository{DatabaseConnection: client}
	//encounterRepo := &repo.EncounterRepository{DatabaseConnection: database}
	encounterService := &service.EncounterService{EncounterRepo: encounterRepo}
	encounterHandler := &handler.EncounterHandler{EncounterService: encounterService}
	/*
		encounterExecutionRepo := &repo.EncounterExecutionRepository{DatabaseConnection: database}
		encounterExecutionService := &service.EncounterExecutionService{EncounterExecutionRepo: encounterExecutionRepo}
		encounterExecutionHandler := &handler.EncounterExecutionHandler{EncounterExecutionService: encounterExecutionService}
	*/
	startServer(encounterHandler)
}

func initTracer() (*trace.TracerProvider, error) {
	url := "http://jaeger:14268/api/traces"
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("encounters-service"),
		)),
	), nil
}

/*

	router.HandleFunc("/encounters/getEncounterById/{encounterId}", handlerEnc.GetEncounterById).Methods("GET")

	router.HandleFunc("/encounters/getSocialEncounterId/{baseEncounterId}", handlerEnc.GetSocialEncounterId).Methods("GET")
	router.HandleFunc("/encounters/getHiddenLocationEncounterId/{baseEncounterId}", handlerEnc.GetHiddenLocationEncounterId).Methods("GET")
	router.HandleFunc("/encounters/getHiddenLocationEncounter/{encounterId}", handlerEnc.GetHiddenLocationEncounterByEncounterId).Methods("GET")

	router.HandleFunc("/encounters/deleteEncounter/{baseEncounterId}", handlerEnc.DeleteEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteSocialEncounter/{socialEncounterId}", handlerEnc.DeleteSocialEncounter).Methods("DELETE")
	router.HandleFunc("/encounters/deleteHiddenLocationEncounter/{hiddenLocationEncounterId}", handlerEnc.DeleteHiddenLocationEncounter).Methods("DELETE")

*/
