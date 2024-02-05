package config

import (
	"kv-store/pkg/logger"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.
			GetLogger().
			Error().
			Err(err).
			Msg("Error while loading env vars")
	}
}
