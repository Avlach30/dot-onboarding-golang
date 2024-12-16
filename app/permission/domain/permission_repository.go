package domain

import (
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Create(payload *PermissionEntity) error
	FindById(id uuid.UUID, trashed bool) (*PermissionEntity, error)
	FindByKey(key string, trashed bool) (*PermissionEntity, error)
	FindByNameAndKey(name string, key string) (*PermissionEntity, error)
	Update(id uuid.UUID, payload *PermissionEntity) error
	Delete(id uuid.UUID)
	ForceDelete(id uuid.UUID)
	IsKeyExist(key string) bool
	IsKeyExistExceptPermissionId(key string, id uuid.UUID) bool
}
