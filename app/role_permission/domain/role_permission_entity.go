package domain

import (
	"time"

	"github.com/google/uuid"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type RolePermission struct {
	ID           uuid.UUID                   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleID       uuid.UUID                   `gorm:"type:uuid;index:idx_composite_role_permission"`
	PermissionID uuid.UUID                   `gorm:"type:uuid;index:idx_composite_role_permission"`
	Role         roleDomain.Role             `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permisison   permissionDomain.Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
	DeletedAt    gorm.DeletedAt              `gorm:"index"`
	UpdatedAt    time.Time                   `gorm:"column:some_data;autoUpdateTime:true;index"`
	CreatedAt    time.Time                   `gorm:"autoCreateTime:true;index"`
}
