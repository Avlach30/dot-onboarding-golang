package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
)

type NotificationUseCase struct {
	notificationRepo domain.NotificationRepository
}

func NewNotificationUseCase(notificationRepo domain.NotificationUseCase) domain.NotificationUseCase {
	return &NotificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (notification *NotificationUseCase) Pagination(ctx *gin.Context, userId uuid.UUID) ([]domain.NotificationEntity, int) {
	return notification.notificationRepo.Pagination(ctx, userId)
}

func (notification *NotificationUseCase) HasUnread(ctx *gin.Context, userId uuid.UUID) bool {
	return notification.notificationRepo.HasUnread(ctx, userId)
}

func (notification *NotificationUseCase) MarkAsRead(ctx *gin.Context, id string, userId uuid.UUID) {
	notification.notificationRepo.MarkAsRead(ctx, id, userId)
}
