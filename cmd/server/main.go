package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"kv-store/pkg/config"
	"kv-store/pkg/handlers"
	"kv-store/pkg/middlewares"
)

func main() {

	config.LoadAppConfig()

	// db, err := sqlite.GetDBInstance()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// recordRepo := repository.NewGormRecordRepository(db)
	// tenantRepo := repository.NewGormTenantRepository(db)

	healthHandler := handlers.NewHealthHandler()

	router := mux.NewRouter()

	router.Use(middlewares.TraceRequest)
	router.Use(middlewares.LogRequest)
	router.Use(middlewares.LogResponse)
	// router.Use(middlewares.NotFound)

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")

	port := fmt.Sprint(config.AppConfig.Port)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("KV Store Api started on port %s", port)

	log.Fatal(srv.ListenAndServe())

}
