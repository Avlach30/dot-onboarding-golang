package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"` // UUID primary key
	Name      string         `gorm:"size:255;not null" json:"name"`
	Key       string         `gorm:"size:255;not null" json:"key"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`

	// Relations
	// Roles []domain.RoleEntity `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles,omitempty"`
}
