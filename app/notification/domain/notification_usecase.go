package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type NotificationUseCase interface {
	Pagination(ctx *gin.Context, userId uuid.UUID) ([]entities.NotificationEntity, int)
	HasUnread(ctx *gin.Context, userId uuid.UUID) bool
	MarkAsRead(ctx *gin.Context, id uuid.UUID, userId uuid.UUID)
	Detail(ctx *gin.Context, id uuid.UUID, userId uuid.UUID) entities.NotificationEntity
}
