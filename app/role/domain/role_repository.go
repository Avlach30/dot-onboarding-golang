package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Pagination(ctx *gin.Context) ([]Role, int)
	Create(ctx *gin.Context, payload *Role) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*Role, error)
	FindByKey(ctx *gin.Context, key string, trashed bool) (*Role, error)
	FindByNameAndKey(ctx *gin.Context, name string, key string) (*Role, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *Role) error
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
	IsKeyExist(ctx *gin.Context, key string) bool
	IsKeyExistExceptRoleId(ctx *gin.Context, key string, id uuid.UUID) bool
}
