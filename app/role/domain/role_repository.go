package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Create(ctx *gin.Context, payload *RoleEntity) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*RoleEntity, error)
	FindByKey(ctx *gin.Context, key string, trashed bool) (*RoleEntity, error)
	FindByNameAndKey(ctx *gin.Context, name string, key string) (*RoleEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *RoleEntity) error
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
	IsKeyExist(ctx *gin.Context, key string) bool
	IsKeyExistExceptRoleId(ctx *gin.Context, key string, id uuid.UUID) bool
}
