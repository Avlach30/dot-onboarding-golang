package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"` // UUID primary key
	TaskName  string         `gorm:"size:255;not null"`
	Payload   string         `gorm:"type:text;not null"`
	Booked    bool           `gorm:"type:bool;not null;default:false"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index"`
}
