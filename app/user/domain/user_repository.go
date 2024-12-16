package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(context *context.Context, payload *UserEntity) error
	FindById(context *context.Context, id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(context *context.Context, id uuid.UUID, payload *UserEntity) error
	Delete(context *context.Context, id uuid.UUID)
	ForceDelete(context *context.Context, id uuid.UUID)
	IsEmailExist(context *context.Context, email string) bool
	IsEmailExistExceptUserId(context *context.Context, email string, id uuid.UUID) bool
}
