package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_cron "github.com/robfig/cron/v3"

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
	// router.NotFoundHandler = middlewares.NotFound(nil)

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")

	router.HandleFunc("/api/tenant/onboard", tenantHandler.OnboardTenantHandler).Methods("GET")

	router.HandleFunc("/api/records/{id}", recordHandler.GetRecordHandler).Methods("GET")
	router.HandleFunc("/api/records", recordHandler.SaveRecordHandler).Methods("POST")

	port := fmt.Sprint(config.AppConfig.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: time.Duration(config.AppConfig.Timeout) * time.Second,
		ReadTimeout:  time.Duration(config.AppConfig.Timeout) * time.Second,
	}

	log.Printf("KV Store Api started on port %s", port)

	c := _cron.New()

	helpers.RecordDeletionCronJob(c, recordRepo)

	log.Fatal(srv.ListenAndServe())

	select {}

}
