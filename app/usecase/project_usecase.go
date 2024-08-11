package usecase

import (
	"context"

	projectdomain "github.com/codespace-id/codespace-x/app/domain/project"
)

type projectUsecase struct {
	projectRepo projectdomain.Repository
}

func NewProjectUsecase(projectRepo projectdomain.Repository) projectdomain.Usecase {
	return &projectUsecase{
		projectRepo: projectRepo,
	}
}

// CreateNewInquiry implements projectdomain.Usecase.
func (p *projectUsecase) CreateNewInquiry(ctx context.Context, dto projectdomain.Entity) error {
	panic("unimplemented")
}

// ListProject implements projectdomain.Usecase.
func (p *projectUsecase) ListProject(ctx context.Context, userID int64, page int, perPage int) (res []projectdomain.Entity, err error) {
	panic("unimplemented")
}

// ProjectDetail implements projectdomain.Usecase.
func (p *projectUsecase) ProjectDetail(ctx context.Context, ID int64) (res projectdomain.Entity, err error) {
	panic("unimplemented")
}

// UpdateDetailProject implements projectdomain.Usecase.
func (p *projectUsecase) UpdateDetailProject(ctx context.Context, dto projectdomain.Entity) error {
	panic("unimplemented")
}
