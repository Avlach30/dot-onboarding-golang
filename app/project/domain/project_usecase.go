package domain

import (
	"context"
)

type Usecase interface {
	CreateNewInquiry(ctx context.Context, dto Entity) error
	ProjectDetail(ctx context.Context, ID int64) (res Entity, err error)
	ListProject(ctx context.Context, userID int64, page, perPage int) (res []Entity, err error)
	UpdateDetailProject(ctx context.Context, dto Entity) error
}
