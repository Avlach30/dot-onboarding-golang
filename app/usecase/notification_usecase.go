package usecase

import (
	"context"

	notificationdomain "github.com/codespace-id/codespace-x/app/domain/notification"
)

type notificationUsecase struct {
	notificationRepo notificationdomain.Repository
}

func NewNotificationUsecase(notificationRepo notificationdomain.Repository) notificationdomain.Usecase {
	return &notificationUsecase{
		notificationRepo: notificationRepo,
	}
}

// CountUnreadNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) CountUnreadNotif(ctx context.Context, userID int64) (res int, err error) {
	panic("unimplemented")
}

// ListNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) ListNotif(ctx context.Context, userID int64, page int, perPage int) (res []notificationdomain.Entity, err error) {
	panic("unimplemented")
}

// SetReadNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) SetReadNotif(ctx context.Context, ID int64, userID int64) error {
	panic("unimplemented")
}
