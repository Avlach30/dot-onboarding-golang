package domain

import (
	"context"
	"github.com/codespace-id/codespace-x/app/project/dto"
)

type PublicUsecase interface {
	ListProject(ctx context.Context, phoneNumber string, page, perPage int) (res []dto.ListProjectResponse, err error)
}

type Usecase interface {
	CreateNewInquiry(ctx context.Context, phoneNumber string, project Entity) (res Entity, err error)
	ProjectDetail(ctx context.Context, ID int64) (res Entity, err error)
	ListProject(ctx context.Context, phoneNumber string, page, perPage int) (res []dto.ListProjectResponse, err error)
	UpdateDetailProject(ctx context.Context, dto Entity) error
	PublicUsecase
}
