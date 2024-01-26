package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")

	port := strconv.Itoa(config.AppConfig.Port)

	log.Println(fmt.Printf("KV Store Api started on port %s", port))

	log.Fatal(http.ListenAndServe(":"+port, router))

}
