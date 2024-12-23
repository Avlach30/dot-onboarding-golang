package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationUseCase interface {
	Pagination(ctx *gin.Context, userId uuid.UUID) ([]NotificationEntity, int)
	HasUnread(ctx *gin.Context, userId uuid.UUID) bool
	MarkAsRead(ctx *gin.Context, id string, userId uuid.UUID)
}
