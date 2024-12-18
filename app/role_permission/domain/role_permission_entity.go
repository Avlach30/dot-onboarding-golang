package domain

import (
	"time"

	"github.com/google/uuid"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type RolePermissionEntity struct {
	ID           uuid.UUID                         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	RoleID       uuid.UUID                         `gorm:"type:uuid;index:idx_composite_role_permission" json:"role_id"`
	PermissionID uuid.UUID                         `gorm:"type:uuid;index:idx_composite_role_permission" json:"permission_id"`
	Role         roleDomain.RoleEntity             `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE" json:"role"`
	Permisison   permissionDomain.PermissionEntity `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE" json:"permission"`
	DeletedAt    gorm.DeletedAt                    `gorm:"index" json:"deleted_at,omitempty"`
	UpdatedAt    time.Time                         `gorm:"column:some_data;autoUpdateTime:true;index" json:"updated_at"`
	CreatedAt    time.Time                         `gorm:"autoCreateTime:true;index" json:"created_at"`
}

func (RolePermissionEntity) TableName() string {
	return "role_permissions"
}
