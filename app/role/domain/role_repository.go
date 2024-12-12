package domain

import (
	"github.com/google/uuid"
)

type RoleRepository interface {
	Create(payload *RoleEntity) error
	FindById(id uuid.UUID, trashed bool) (*RoleEntity, error)
	FindByKey(key string, trashed bool) (*RoleEntity, error)
	FindByNameAndKey(name string, key string) (*RoleEntity, error)
	Update(id uuid.UUID, payload *RoleEntity)
	Delete(id uuid.UUID)
	ForceDelete(id uuid.UUID)
}
