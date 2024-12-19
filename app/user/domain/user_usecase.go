package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
)

type UserUsecase interface {
	Pagination(ctx *gin.Context) ([]UserEntity, int)
	Create(ctx *gin.Context, payload *dto.UserCreateRequest) error
	FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*UserEntity, error)
	Update(ctx *gin.Context, id uuid.UUID, payload *UserEntity)
	Delete(ctx *gin.Context, id uuid.UUID)
	ForceDelete(ctx *gin.Context, id uuid.UUID)
}
