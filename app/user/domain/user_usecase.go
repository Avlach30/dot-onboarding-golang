package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type UserUsecase interface {
	Pagination(httpContext *gin.Context) ([]entities.UserEntity, int)
	Create(httpContext *gin.Context, payload *entities.UserEntity, roleIds []uuid.UUID)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.UserEntity, roleIds []uuid.UUID)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
