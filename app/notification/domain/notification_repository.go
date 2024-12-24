package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationRepository interface {
	Pagination(httpContext *gin.Context, userId uuid.UUID) ([]NotificationEntity, int)
	HasUnread(httpContext *gin.Context, userId uuid.UUID) bool
	MarkAsRead(httpContext *gin.Context, id string, userId uuid.UUID)
}
