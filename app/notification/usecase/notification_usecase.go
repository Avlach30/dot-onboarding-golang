package usecase

import (
	"context"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
)

type notificationUsecase struct {
	notificationRepo domain.Repository
}

func NewNotificationUsecase(notificationRepo domain.Repository) domain.Usecase {
	return &notificationUsecase{
		notificationRepo: notificationRepo,
	}
}

// CountUnreadNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) CountUnreadNotif(ctx context.Context, userID int64) (res int, err error) {
	panic("unimplemented")
}

// ListNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) ListNotif(ctx context.Context, userID int64, page int, perPage int) (res []domain.Entity, err error) {
	panic("unimplemented")
}

// SetReadNotif implements notificationdomain.Usecase.
func (n *notificationUsecase) SetReadNotif(ctx context.Context, ID int64, userID int64) error {
	panic("unimplemented")
}
