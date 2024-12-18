package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Create(ctx *gin.Context, payload *PermissionEntity) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*PermissionEntity, error)
	FindByKey(ctx *gin.Context, key string, trashed bool) (*PermissionEntity, error)
	FindByNameAndKey(ctx *gin.Context, name string, key string) (*PermissionEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *PermissionEntity) error
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
	IsKeyExist(ctx *gin.Context, key string) bool
	IsKeyExistExceptPermissionId(ctx *gin.Context, key string, id uuid.UUID) bool
}
