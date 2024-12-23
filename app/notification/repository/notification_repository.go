package repository

import (
	"github.com/gin-gonic/gin"
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
func (notification *NotificationRepository) Pagination(ctx *gin.Context) ([]domain.NotificationEntity, int) {
	notification.notificationModel = notification.notificationModel.WithContext(ctx)
	var notifications []domain.NotificationEntity
	var total int64

	// Query filter
	notification.queryFilter(ctx)
	// Query sort
	notification.querySort(ctx)

	notification.notificationModel.Session(&gorm.Session{}).
		Scopes(utils.Paginate(ctx)).
		Find(&notifications).
		Count(&total)

	return notifications, int(total)
}

// func filter for pagination
func (notification *NotificationRepository) queryFilter(ctx *gin.Context) *gorm.DB {
	if search := ctx.Query("search"); search != "" {
		notification.notificationModel = notification.notificationModel.
			Where("title LIKE ?", search+"%")
	}

	return notification.notificationModel
}

// func query sort for pagination
func (notification *NotificationRepository) querySort(ctx *gin.Context) *gorm.DB {
	sortableColumns := []string{"created_at"}

	if sort := ctx.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			notification.notificationModel = notification.notificationModel.Order(sort + " " + ctx.Query("order"))
		}
	}

	return notification.notificationModel
}

// HasUnread check if user has unread notification
func (notification *NotificationRepository) HasUnread(ctx *gin.Context) bool {
	notification.notificationModel = notification.notificationModel.WithContext(ctx)
	var total int64

	err := notification.notificationModel.
		Where("is_read = ?", false).
		Count(&total).Error

	if err != nil {
		panic(*exception.ServerErrorException("Failed to get unread notification"))
	}

	return total > 0
}

// MarkAsRead mark notification as read
func (notification *NotificationRepository) MarkAsRead(ctx *gin.Context, id string) {
	notification.notificationModel = notification.notificationModel.WithContext(ctx)

	err := notification.notificationModel.
		Where("id = ?", id).
		Update("is_read", true).Error

	if err != nil {
		panic(*exception.ServerErrorException("Failed to mark notification as read"))
	}
}
