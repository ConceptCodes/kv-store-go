package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"kv-store/pkg/config"
	"kv-store/pkg/handlers"
	"kv-store/pkg/helpers"
	"kv-store/pkg/middlewares"
	"kv-store/pkg/models"
	repository "kv-store/pkg/repositories"
	"kv-store/pkg/storage/sqlite"
)

func main() {

	config.LoadAppConfig()

	db, err := sqlite.GetDBInstance()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.TenantModel{}, &models.RecordModel{})

	recordRepo := repository.NewGormRecordRepository(db)
	tenantRepo := repository.NewGormTenantRepository(db)

	healthHandler := handlers.NewHealthHandler()
	tenantHandler := handlers.NewTenantHandler(tenantRepo)
	recordHandler := handlers.NewRecordHandler(recordRepo)

	router := mux.NewRouter()

	router.Use(middlewares.TraceRequest)
	router.Use(middlewares.LogRequest)
	router.Use(middlewares.LogResponse)
	router.NotFoundHandler = middlewares.NotFound(nil)
	// router.Use(middlewares.NotFound)

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")

	router.HandleFunc("/api/tenant/onboard", tenantHandler.OnboardTenantHandler).Methods("GET")

	router.HandleFunc("/api/records/{id:^[a-zA-Z0-9]*$}", recordHandler.GetRecordHandler).Methods("GET")
	router.HandleFunc("/api/records", recordHandler.SaveRecordHandler).Methods("POST")

	port := fmt.Sprint(config.AppConfig.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("KV Store Api started on port %s", port)

	log.Fatal(srv.ListenAndServe())

	helpers.RecordDeletionCronJob(recordRepo)

}
