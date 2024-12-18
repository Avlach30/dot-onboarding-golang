package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"` // Password is excluded from JSON
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`
}
