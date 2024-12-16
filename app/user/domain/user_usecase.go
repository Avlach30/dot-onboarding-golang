package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Create(context *context.Context, payload *UserEntity) error
	FindById(context *context.Context, id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(context *context.Context, id uuid.UUID, payload *UserEntity)
	Delete(context *context.Context, id uuid.UUID)
	ForceDelete(context *context.Context, id uuid.UUID)
}
