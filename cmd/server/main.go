package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"kv-store/config"
	"kv-store/internal/handlers"
	"kv-store/internal/middlewares"
	"kv-store/internal/models"
	repository "kv-store/internal/repositories"
	"kv-store/pkg/logger"
	"kv-store/pkg/storage/redis"
	"kv-store/pkg/storage/sqlite"
)

func Run() {

	config.LoadAppConfig()

	db, err := sqlite.GetDBInstance()
	log := logger.GetLogger()

	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to database")
	}

	db.AutoMigrate(&models.TenantModel{})

	_redis := redis.New()

	tenantRepo := repository.NewGormTenantRepository(db)

	healthHandler := handlers.NewHealthHandler()
	tenantHandler := handlers.NewTenantHandler(tenantRepo)
	recordHandler := handlers.NewRecordHandler(_redis)

	router := mux.NewRouter()

	router.Use(middlewares.TraceRequest)
	router.Use(middlewares.RequestLogger)
	router.Use(middlewares.ContentTypeJSON)
	router.NotFoundHandler = middlewares.NotFound(nil)

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")
	router.HandleFunc("/api/tenant/onboard", tenantHandler.OnboardTenantHandler).Methods("GET")
	router.HandleFunc("/api/records/{key}", recordHandler.GetRecordHandler).Methods("GET")
	router.HandleFunc("/api/records", recordHandler.SaveRecordHandler).Methods("POST")

	port := fmt.Sprint(config.AppConfig.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: time.Duration(config.AppConfig.Timeout) * time.Second,
		ReadTimeout:  time.Duration(config.AppConfig.Timeout) * time.Second,
	}

	log.Debug().Msgf("KV Store Api started on port %s", port)

	log.
		Fatal().
		Err(srv.ListenAndServe()).
		Msg("Error while starting server")

	select {}

}
