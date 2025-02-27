package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	RoleID    uuid.UUID      `gorm:"type:uuid;index:idx_composite_user_role" json:"role_id"`
	UserID    uuid.UUID      `gorm:"type:uuid;index:idx_composite_user_role" json:"user_id"`
	Role      RoleEntity     `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE" json:"role"`
	User      UserEntity     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:true;index" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime:true;index" json:"created_at"`
}

func (UserRoleEntity) TableName() string {
	return "user_roles"
}
