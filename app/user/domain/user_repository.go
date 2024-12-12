package domain

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(payload *UserEntity) error
	FindById(id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(id uuid.UUID, payload *UserEntity)
	Delete(id uuid.UUID)
	ForceDelete(id uuid.UUID)
}
