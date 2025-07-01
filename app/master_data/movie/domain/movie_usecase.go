package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieEntity, int)
	Create(httpContext *gin.Context, payload *entities.MovieEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}