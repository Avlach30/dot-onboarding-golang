package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieScheduleUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieScheduleEntity, int)
	Create(httpContext *gin.Context, payload *entities.MovieScheduleEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieScheduleEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieScheduleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}