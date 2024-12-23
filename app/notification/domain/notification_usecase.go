package domain

import (
	"github.com/gin-gonic/gin"
)

type NotificationUseCase interface {
	Pagination(ctx *gin.Context) ([]NotificationEntity, int)
	HasUnread(ctx *gin.Context) bool
	MarkAsRead(ctx *gin.Context, id string)
}
