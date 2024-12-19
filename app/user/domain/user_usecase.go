package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Pagination(ctx *gin.Context) ([]UserEntity, int)
	Create(ctx *gin.Context, payload *UserEntity) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *UserEntity)
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
}
