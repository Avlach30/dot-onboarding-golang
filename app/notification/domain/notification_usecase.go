package domain

import "context"

type Usecase interface {
	ListNotif(ctx context.Context, userID int64, page, perPage int) (res []Entity, err error)
	CountUnreadNotif(ctx context.Context, userID int64) (res int, err error)
	SetReadNotif(ctx context.Context, ID, userID int64) error
}
