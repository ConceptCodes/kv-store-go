package helpers

import (
	"encoding/base64"
	"errors"
	"kv-store/internal/models"
	repository "kv-store/internal/repositories"
	"kv-store/pkg/logger"
	"kv-store/pkg/storage/sqlite"
	"strings"
)

func ValidateToken(token string) (*models.UserModel, error) {
	db, err := sqlite.GetDBInstance()

	log := logger.GetLogger()

	if err != nil {
		log.Error().Err(err).Msg("Error getting db instance")
	}

	log.Debug().Msgf("Validating token: %s", token)

	tenantRepo := repository.NewGormTenantRepository(db)

	data, err := base64.StdEncoding.DecodeString(token)

	if err != nil {
		log.Error().Err(err).Msg("Error decoding token")
		return nil, err
	}

	parts := strings.Split(string(data), ":")
	if len(parts) < 2 {
		return nil, errors.New("invalid token format")
	}

	tenantId := parts[0]
	tenantSecret := parts[1]

	tenant, err := tenantRepo.FindById(tenantId)

	if err != nil {
		return nil, err
	}

	if tenant.Secret == tenantSecret {
		return &models.UserModel{
			ID: tenant.ID,
		}, nil
	}

	return nil, err
}

func GenerateToken(tenantId string, tenantSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(tenantId + ":" + tenantSecret))
}
