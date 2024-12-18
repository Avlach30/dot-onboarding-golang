package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionUsecase interface {
	Pagination(ctx *gin.Context) ([]Permission, int)
	Create(ctx *gin.Context, schema *Permission) error
	FindById(ctx *gin.Context, id uuid.UUID) (*Permission, error)
	FindByKey(ctx *gin.Context, key string) (*Permission, error)
	Update(ctx *gin.Context, id uuid.UUID, schema *Permission)
	Delete(ctx *gin.Context, id uuid.UUID)
}
