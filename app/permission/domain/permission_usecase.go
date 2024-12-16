package domain

import (
	"context"

	"github.com/google/uuid"
)

type PermissionUsecase interface {
	Create(context *context.Context, schema *PermissionEntity) error
	FindById(context *context.Context, id uuid.UUID) (*PermissionEntity, error)
	FindByKey(context *context.Context, key string) (*PermissionEntity, error)
	Update(context *context.Context, id uuid.UUID, schema *PermissionEntity)
	Delete(context *context.Context, id uuid.UUID)
}
