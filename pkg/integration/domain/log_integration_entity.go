package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogIntegrationEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	URL       string         `json:"url" gorm:"type:text;not null"`
	Request   string         `json:"request" gorm:"type:text;not null"`
	Response  string         `json:"response" gorm:"type:text;not null"`
	Status    string         `json:"status" gorm:"type:varchar(255);not null"`
	Scheme    string         `json:"scheme" gorm:"type:varchar(255);not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index"`
}

func (LogIntegrationEntity) TableName() string {
	return "log_integrations"
}
