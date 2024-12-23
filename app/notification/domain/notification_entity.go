package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"size:255;not null" json:"content"`
	IsRead    bool           `gorm:"default:false" json:"is_read"`
	Href      string         `gorm:"size:255;not null" json:"href"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`

	// Relations
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
}

// Specify the table name for the Notification struct
func (NotificationEntity) TableName() string {
	return "notifications"
}
