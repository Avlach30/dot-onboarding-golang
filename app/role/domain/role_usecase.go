package domain

import (
	"context"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	Create(context *context.Context, schema *RoleEntity) error
	FindById(context *context.Context, id uuid.UUID) (*RoleEntity, error)
	FindByKey(context *context.Context, key string) (*RoleEntity, error)
	Update(context *context.Context, id uuid.UUID, dto *RoleEntity)
	Delete(context *context.Context, id uuid.UUID)
}
