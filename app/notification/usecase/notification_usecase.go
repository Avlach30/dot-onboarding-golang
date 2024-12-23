package usecase

import (
	"github.com/gin-gonic/gin"
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

func (notification *NotificationUseCase) Pagination(ctx *gin.Context) ([]domain.NotificationEntity, int) {
	return notification.notificationRepo.Pagination(ctx)
}

func (notification *NotificationUseCase) HasUnread(ctx *gin.Context) bool {
	return notification.notificationRepo.HasUnread(ctx)
}

func (notification *NotificationUseCase) MarkAsRead(ctx *gin.Context, id string) {
	notification.notificationRepo.MarkAsRead(ctx, id)
}
