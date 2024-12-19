package domain

import (
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"` // Password is excluded from JSON
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true index" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`

	// Relations
	Roles []domain.RoleEntity `gorm:"many2many:user_roles;foreignKey:ID;joinForeignKey:user_id;References:ID;joinReferences:role_id" json:"roles"`
}

func (UserEntity) TableName() string {
	return "users"
}
