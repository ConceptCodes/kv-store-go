package models

import (
    "time"

    "gorm.io/gorm"
)

type RecordModel struct {
    gorm.Model
    ID        string    `gorm:"primaryKey;type:text;unique_index"`
    TenantId  string    `gorm:"primaryKey;type:varchar(36);"`
    Value     string    `gorm:"type:text"`
    ExpiresAt time.Time `gorm:"type:timestamp"`
    Tenant    TenantModel `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE;"`
}

type TenantModel struct {
    gorm.Model
    ID     string `gorm:"primaryKey;type:varchar(36);unique_index"`
    Name   string `gorm:"type:varchar(100);unique_index"`
    Secret string `gorm:"type:varchar(36)"`
}