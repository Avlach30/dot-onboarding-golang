package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	ticketDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/ticket/dto"
)

type TicketUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.TicketEntity, int)
	Create(httpContext *gin.Context, payload *ticketDto.TicketCreateRequest)
	Update(httpContext *gin.Context, id uuid.UUID, payload *ticketDto.TicketUpdateRequest)
	Delete(httpContext *gin.Context, id uuid.UUID)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.TicketEntity
}