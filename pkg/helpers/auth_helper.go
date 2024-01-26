package helpers

import (
	"encoding/base64"
	"errors"
	"kv-store/pkg/models"
	repository "kv-store/pkg/repositories"
	"kv-store/pkg/storage/sqlite"
	"log"
	"strings"
)

func ValidateToken(token string) (*models.UserModel, error) {
	db, err := sqlite.GetDBInstance()

	if err != nil {
		log.Print(err)
	}

	tenantRepo := repository.NewGormTenantRepository(db)

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
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
