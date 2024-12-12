package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUID primary key
	Name      string         `gorm:"size:255;not null"`
	Email     string         `gorm:"size:255;not null"`
	Password  string         `gorm:"size:255;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index"`
}
