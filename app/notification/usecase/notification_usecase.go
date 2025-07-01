package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type NotificationUseCase struct {
	notificationRepo domain.NotificationRepository
}

func NewNotificationUseCase(notificationRepo domain.NotificationRepository) domain.NotificationUseCase {
	return &NotificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (notification *NotificationUseCase) Pagination(httpContext *gin.Context, userId uuid.UUID, queryDto *querydto.QueryDto) ([]entities.NotificationEntity, int) {
	return notification.notificationRepo.Pagination(httpContext, userId, queryDto)
}

func (notification *NotificationUseCase) HasUnread(httpContext *gin.Context, userId uuid.UUID) bool {
	return notification.notificationRepo.HasUnread(httpContext, userId)
}

func (notification *NotificationUseCase) MarkAsRead(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) {
	notification.notificationRepo.MarkAsRead(httpContext, id, userId)
}

func (notification *NotificationUseCase) Detail(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) entities.NotificationEntity {
	return notification.notificationRepo.FindOneById(httpContext, id, userId)
}
