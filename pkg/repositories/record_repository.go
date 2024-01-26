package repository

import (
	"gorm.io/gorm"

	"kv-store/pkg/constants"
	"kv-store/pkg/models"
)

type RecordRepository interface {
	FindById(tenant_id string, id string) (*models.RecordModel, error)
	Save(record *models.RecordModel) error
	Delete(tenant_id string, id string) error
}

type GormRecordRepository struct {
	db *gorm.DB
}

func (r *GormRecordRepository) FindById(tenant_id string, id string) (*models.RecordModel, error) {
	var record models.RecordModel
	if err := r.db.Where(constants.FindByTenantIdAndKey, tenant_id, id).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *GormRecordRepository) Save(record *models.RecordModel) error {
	return r.db.Save(record).Error
}

func (r *GormRecordRepository) Delete(tenant_id string, id string) error {
	return r.db.Delete(&models.RecordModel{}, constants.FindByTenantIdAndKey, tenant_id, id).Error
}

func NewGormRecordRepository(db *gorm.DB) RecordRepository {
	return &GormRecordRepository{db}
}
