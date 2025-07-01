package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type TicketRepository interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.TicketEntity, int)
	Create(httpContext *gin.Context, payload *entities.TicketEntity)
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.TicketEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.TicketEntity
}