package domain

import (
	"time"

	"github.com/google/uuid"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gorm.io/gorm"
)

type UserRoleEntity struct {
	ID        uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	RoleID    uuid.UUID             `gorm:"type:uuid;index:idx_composite_user_role" json:"role_id"`
	UserID    uuid.UUID             `gorm:"type:uuid;index:idx_composite_user_role" json:"user_id"`
	Role      roleDomain.RoleEntity `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE" json:"role"`
	User      userDomain.UserEntity `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	DeletedAt gorm.DeletedAt        `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt time.Time             `gorm:"autoUpdateTime:true;index" json:"updated_at"`
	CreatedAt time.Time             `gorm:"autoCreateTime:true;index" json:"created_at"`
}

func (UserRoleEntity) TableName() string {
	return "user_roles"
}
