package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type NotificationUseCase interface {
	Pagination(httpContext *gin.Context, userId uuid.UUID) ([]entities.NotificationEntity, int)
	HasUnread(httpContext *gin.Context, userId uuid.UUID) bool
	MarkAsRead(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID)
	Detail(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) entities.NotificationEntity
}
