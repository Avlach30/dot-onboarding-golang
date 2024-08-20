package domain

import (
	"context"

	"github.com/codespace-id/codespace-x/pkg/common/enum"
)

type Repository interface {
	GetByUserID(ctx context.Context, userID int64, skip, limit string) (res []Entity, err error)
	CountUnreadByUserID(ctx context.Context, userID int64) (res int, err error)
	UpdateStatus(ctx context.Context, ID, userID int64, status enum.NotificationStatusType) error
}
