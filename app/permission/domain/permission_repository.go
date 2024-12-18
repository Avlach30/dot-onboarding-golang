package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Pagination(ctx *gin.Context) ([]Permission, int)
	Create(ctx *gin.Context, payload *Permission) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*Permission, error)
	FindByKey(ctx *gin.Context, key string, trashed bool) (*Permission, error)
	FindByNameAndKey(ctx *gin.Context, name string, key string) (*Permission, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *Permission) error
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
	IsKeyExist(ctx *gin.Context, key string) bool
	IsKeyExistExceptPermissionId(ctx *gin.Context, key string, id uuid.UUID) bool
}
