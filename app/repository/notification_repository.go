package repository

import (
	"context"
	"database/sql"

	notificationdomain "github.com/codespace-id/codespace-x/app/domain/notification"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) notificationdomain.Repository {
	return &NotificationRepository{
		db: db,
	}
}

// CountUnreadByUserID implements notificationdomain.Repository.
func (n *NotificationRepository) CountUnreadByUserID(ctx context.Context, userID int64) (res int, err error) {
	panic("unimplemented")
}

// GetByUserID implements notificationdomain.Repository.
func (n *NotificationRepository) GetByUserID(ctx context.Context, userID int64, skip string, limit string) (res []notificationdomain.Entity, err error) {
	panic("unimplemented")
}

// UpdateStatus implements notificationdomain.Repository.
func (n *NotificationRepository) UpdateStatus(ctx context.Context, ID int64, userID int64, status enum.NotificationStatusType) error {
	panic("unimplemented")
}

// Get implements notificationdomain.Repository.
func (n *NotificationRepository) Get(ctx context.Context, skip string, limit string) (res []notificationdomain.Entity, err error) {
	panic("unimplemented")
}
