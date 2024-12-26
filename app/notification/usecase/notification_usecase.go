package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
)

type NotificationUseCase struct {
	notificationRepo domain.NotificationRepository
}

func NewNotificationUseCase(notificationRepo domain.NotificationRepository) domain.NotificationUseCase {
	return &NotificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (notification *NotificationUseCase) Pagination(httpContext *gin.Context, userId uuid.UUID) ([]domain.NotificationEntity, int) {
	return notification.notificationRepo.Pagination(httpContext, userId)
}

func (notification *NotificationUseCase) HasUnread(httpContext *gin.Context, userId uuid.UUID) bool {
	return notification.notificationRepo.HasUnread(httpContext, userId)
}

func (notification *NotificationUseCase) MarkAsRead(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) {
	notification.notificationRepo.MarkAsRead(httpContext, id, userId)
}

func (notification *NotificationUseCase) Detail(httpContext *gin.Context, id uuid.UUID) domain.NotificationEntity {
	return notification.notificationRepo.FindOneById(httpContext, id)
}
