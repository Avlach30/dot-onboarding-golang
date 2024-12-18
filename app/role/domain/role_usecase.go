package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleUsecase interface {
	Pagination(ctx *gin.Context) ([]Role, int)
	Create(ctx *gin.Context, schema *Role) error
	FindById(ctx *gin.Context, id uuid.UUID) (*Role, error)
	FindByKey(ctx *gin.Context, key string) (*Role, error)
	Update(ctx *gin.Context, id uuid.UUID, dto *Role)
	Delete(ctx *gin.Context, id uuid.UUID)
}
