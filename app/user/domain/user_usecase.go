package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
)

type UserUsecase interface {
	Pagination(httpContext *gin.Context) ([]UserEntity, int)
	Create(httpContext *gin.Context, payload *dto.UserCreateRequest)
	FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) *UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *dto.UserUpdateRequest)
	Delete(httpContext *gin.Context, id uuid.UUID)
	ForceDelete(httpContext *gin.Context, id uuid.UUID)
}
