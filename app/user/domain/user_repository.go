package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRepository interface {
	Pagination(ctx *gin.Context) ([]UserEntity, int)
	Create(ctx *gin.Context, payload *UserEntity) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *UserEntity) error
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
	IsEmailExist(ctx *gin.Context, email string) bool
	IsEmailExistExceptUserId(ctx *gin.Context, email string, id uuid.UUID) bool
}
