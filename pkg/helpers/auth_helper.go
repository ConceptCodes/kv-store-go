package helpers

import (
	"encoding/base64"
	"kv-store/pkg/models"
	repository "kv-store/pkg/repositories"
	"kv-store/pkg/storage/sqlite"
	"log"
)

func ValidateToken(token string) (*models.UserModel, error) {
	db, err := sqlite.GetDBInstance()
	if err != nil {
		log.Fatal(err)
	}

	tenantRepo := repository.NewGormTenantRepository(db)

	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Fatal("error:", err)
	}

	tenantId := string(data[:][0])
	tenantSecret := string(data[:][1])

	tenant, err := tenantRepo.FindById(tenantId)

	if err != nil {
		log.Fatal("error:", err)
	}

	if tenant.Secret == tenantSecret {
		return &models.UserModel{
			ID: tenant.ID,
		}, nil
	}
	// if the user is not found, i will return an error
	return nil, err
}
