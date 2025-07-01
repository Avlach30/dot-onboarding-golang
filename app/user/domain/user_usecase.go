package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type UserUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.UserEntity, int)
	Create(httpContext *gin.Context, payload *entities.UserEntity, roleIds []uuid.UUID)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.UserEntity, roleIds []uuid.UUID)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
