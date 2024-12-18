package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionUsecase interface {
	Pagination(ctx *gin.Context) ([]PermissionEntity, int)
	Create(ctx *gin.Context, schema *PermissionEntity) error
	FindById(ctx *gin.Context, id uuid.UUID) (*PermissionEntity, error)
	FindByKey(ctx *gin.Context, key string) (*PermissionEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, schema *PermissionEntity)
	Delete(ctx *gin.Context, id uuid.UUID)
}
