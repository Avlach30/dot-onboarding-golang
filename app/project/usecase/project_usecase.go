package usecase

import (
	"context"
	"github.com/codespace-id/codespace-x/app/project/domain"
)

type projectUsecase struct {
	projectRepo domain.Repository
}

func NewProjectUsecase(projectRepo domain.Repository) domain.Usecase {
	return &projectUsecase{
		projectRepo: projectRepo,
	}
}

// CreateNewInquiry implements projectdomain.Usecase.
func (p *projectUsecase) CreateNewInquiry(ctx context.Context, dto domain.Entity) error {
	panic("unimplemented")
}

// ListProject implements projectdomain.Usecase.
func (p *projectUsecase) ListProject(ctx context.Context, userID int64, page int, perPage int) (res []domain.Entity, err error) {
	panic("unimplemented")
}

// ProjectDetail implements projectdomain.Usecase.
func (p *projectUsecase) ProjectDetail(ctx context.Context, ID int64) (res domain.Entity, err error) {
	panic("unimplemented")
}

// UpdateDetailProject implements projectdomain.Usecase.
func (p *projectUsecase) UpdateDetailProject(ctx context.Context, dto domain.Entity) error {
	panic("unimplemented")
}
