package domain

import (
	"github.com/google/uuid"
)

type PermissionUsecase interface {
	Create(schema *PermissionEntity) error
	FindById(id uuid.UUID) (*PermissionEntity, error)
	FindByKey(key string) (*PermissionEntity, error)
	Update(id uuid.UUID, schema *PermissionEntity)
	Delete(id uuid.UUID)
}
