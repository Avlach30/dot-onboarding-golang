package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleUsecase interface {
	Create(ctx *gin.Context, schema *RoleEntity) error
	FindById(ctx *gin.Context, id uuid.UUID) (*RoleEntity, error)
	FindByKey(ctx *gin.Context, key string) (*RoleEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, dto *RoleEntity)
	Delete(ctx *gin.Context, id uuid.UUID)
}
