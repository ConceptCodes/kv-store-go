package models

import (
	"gorm.io/gorm"
)

type RecordModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:text;unique_index"`
	TenantId string `gorm:"primaryKey;type:varchar(36);foreignkey:TenantRefer"`
	Value    string `gorm:"type:text"`
	TTL      uint64 `gorm:"type:uint"`
	Tenant   TenantModel
}

type TenantModel struct {
	gorm.Model
	ID     string `gorm:"primaryKey;type:varchar(36);unique_index"`
	Name   string `gorm:"type:varchar(100);unique_index"`
	Secret string `gorm:"type:varchar(36)"`
}
