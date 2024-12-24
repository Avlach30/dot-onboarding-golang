package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	notificationModel *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &NotificationRepository{
		notificationModel: db.Model(&domain.NotificationEntity{}),
	}
}

// Pagination get notification data with pagination
func (notification *NotificationRepository) Pagination(httpContext *gin.Context, userId uuid.UUID) ([]domain.NotificationEntity, int) {
	notification.notificationModel = notification.notificationModel.WithContext(httpContext)
	var notifications []domain.NotificationEntity
	var total int64

	// Query filter
	notification.queryFilter(httpContext)
	// Query sort
	notification.querySort(httpContext)

	notification.notificationModel.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).
		Where("user_id = ?", userId).
		Find(&notifications).
		Count(&total)

	return notifications, int(total)
}

// func filter for pagination
func (notification *NotificationRepository) queryFilter(httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		notification.notificationModel = notification.notificationModel.
			Where("title LIKE ?", search+"%")
	}

	return notification.notificationModel
}

// func query sort for pagination
func (notification *NotificationRepository) querySort(httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"created_at"}

	if sort := httpContext.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			notification.notificationModel = notification.notificationModel.Order(sort + " " + httpContext.Query("order"))
		}
	}

	return notification.notificationModel
}

// HasUnread check if user has unread notification
func (notification *NotificationRepository) HasUnread(httpContext *gin.Context, userId uuid.UUID) bool {
	notification.notificationModel = notification.notificationModel.WithContext(httpContext)
	var total int64

	err := notification.notificationModel.
		Where("is_read = ?", false).
		Where("user_id = ?", userId).
		Count(&total).Error

	if err != nil {
		panic(*exception.ServerErrorException("Failed to get unread notification"))
	}

	return total > 0
}

// MarkAsRead mark notification as read
func (notification *NotificationRepository) MarkAsRead(httpContext *gin.Context, id string, userId uuid.UUID) {
	notification.notificationModel = notification.notificationModel.WithContext(httpContext)

	err := notification.notificationModel.
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Update("is_read", true).Error

	if err != nil {
		panic(*exception.ServerErrorException("Failed to mark notification as read"))
	}
}
