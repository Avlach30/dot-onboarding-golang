package domain

import (
	"github.com/google/uuid"
)

type RoleUsecase interface {
	Create(schema *RoleEntity) error
	FindById(id uuid.UUID) (*RoleEntity, error)
	FindByKey(key string) (*RoleEntity, error)
	Update(id uuid.UUID, dto *RoleEntity)
	Delete(id uuid.UUID)
}
