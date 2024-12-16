package domain

import (
	"context"

	"github.com/google/uuid"
)

type RoleRepository interface {
	Create(context *context.Context, payload *RoleEntity) error
	FindById(context *context.Context, id uuid.UUID, trashed bool) (*RoleEntity, error)
	FindByKey(context *context.Context, key string, trashed bool) (*RoleEntity, error)
	FindByNameAndKey(context *context.Context, name string, key string) (*RoleEntity, error)
	Update(context *context.Context, id uuid.UUID, payload *RoleEntity) error
	Delete(context *context.Context, id uuid.UUID)
	ForceDelete(context *context.Context, id uuid.UUID)
	IsKeyExist(context *context.Context, key string) bool
	IsKeyExistExceptRoleId(context *context.Context, key string, id uuid.UUID) bool
}
