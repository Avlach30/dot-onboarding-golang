package domain

import (
	"context"

	"github.com/google/uuid"
)

type PermissionRepository interface {
	Create(context *context.Context, payload *PermissionEntity) error
	FindById(context *context.Context, id uuid.UUID, trashed bool) (*PermissionEntity, error)
	FindByKey(context *context.Context, key string, trashed bool) (*PermissionEntity, error)
	FindByNameAndKey(context *context.Context, name string, key string) (*PermissionEntity, error)
	Update(context *context.Context, id uuid.UUID, payload *PermissionEntity) error
	Delete(context *context.Context, id uuid.UUID)
	ForceDelete(context *context.Context, id uuid.UUID)
	IsKeyExist(context *context.Context, key string) bool
	IsKeyExistExceptPermissionId(context *context.Context, key string, id uuid.UUID) bool
}
