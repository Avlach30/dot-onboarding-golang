package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	notificationModel *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &NotificationRepository{
		notificationModel: db.Model(&entities.NotificationEntity{}),
	}
}

// Pagination get notification data with pagination
func (notification *NotificationRepository) Pagination(httpContext *gin.Context, userId uuid.UUID, queryDto *querydto.QueryDto) ([]entities.NotificationEntity, int) {
	query := notification.notificationModel.WithContext(httpContext)
	var notifications []entities.NotificationEntity
	var total int64

	// Query filter
	query = notification.queryFilter(query, queryDto, userId)

	// Query sort
	query = notification.querySort(query, queryDto)

	// Count all column first before paginate the query
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count user", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&notifications).Error

	if err != nil {
		log.Println("Error pagination notification", err)
		panic(*exception.ServerErrorException(err))
	}

	return notifications, int(total)
}

// func filter for pagination
func (notification *NotificationRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto, userId uuid.UUID) *gorm.DB {
	query = query.Where("user_id = ?", userId)

	if search := queryDto.Search; search != "" {
		query = query.Where("title ILIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (notification *NotificationRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"title", "created_at", "updated_at"}

	if sort := queryDto.SortBy; sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		// Handle order query
		if order := queryDto.Order; order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}
			query = query.Order(sort + " " + order)
		}
	} else {
		query = query.Order("created_at desc")
	}

	return query
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
		panic(*exception.ServerErrorException(err))
	}

	return total > 0
}

// MarkAsRead mark notification as read
func (notification *NotificationRepository) MarkAsRead(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) {
	notification.notificationModel = notification.notificationModel.WithContext(httpContext)

	err := notification.notificationModel.
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Update("is_read", true).Error

	if err != nil {
		panic(*exception.ServerErrorException(err))
	}
}

// FindOneById get notification detail by id
func (notification *NotificationRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, userId uuid.UUID) entities.NotificationEntity {
	notification.notificationModel = notification.notificationModel.WithContext(httpContext)
	notificationEntity := entities.NotificationEntity{}

	err := notification.notificationModel.
		Joins("User").
		Where("notifications.id = ?", id).
		Where("notifications.user_id = ?", userId).
		First(&notificationEntity).Error

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Notification not found"))
	} else if err != nil {
		log.Println("err detail notification", err)
		panic(*exception.ServerErrorException(err))
	}

	return notificationEntity
}
