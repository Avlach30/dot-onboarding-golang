package usecase

import (
	"context"
	commonrepo "github.com/codespace-id/codespace-x/app/common/repository"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/app/project/dto"
	userdto "github.com/codespace-id/codespace-x/app/user/dto"
	"github.com/codespace-id/codespace-x/app/user/userdomain"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/pkg/errors"
	"time"
)

type projectPublicUsecase struct {
	projectRepo     domain.Repository
	sqlTxRepo       commonrepo.SqlTx
	userProjectRepo domain.UserProjectRepository
	userRepo        userdomain.Repository
}

func NewProjectPublicUsecase(projectRepo domain.Repository, sqlTxRepo commonrepo.SqlTx, userProjectRepo domain.UserProjectRepository, userRepo userdomain.Repository) domain.PublicUsecase {
	return &projectPublicUsecase{
		projectRepo:     projectRepo,
		sqlTxRepo:       sqlTxRepo,
		userProjectRepo: userProjectRepo,
		userRepo:        userRepo,
	}
}

func (uc *projectPublicUsecase) ListProject(ctx context.Context, phoneNumber, roles string, page, perPage int) (res []dto.ListProjectResponse, err error) {
	var data []domain.Entity

	data, err = uc.projectRepo.GetByStatus(ctx, page, perPage, enum.FINISHED.Value())
	if err != nil {
		return nil, errors.WithMessage(err, "projectUsecase.ListProject")
	}

	for _, val := range data {
		res = append(res, dto.ListProjectResponse{
			UUID:              val.UUID,
			Name:              val.Name,
			Description:       val.Description,
			ThumbnailImageURL: val.ThumbnailImageURL,
			ServiceType:       enum.GetTransformServiceType(val.ServiceType),
			Status:            val.Status,
			CreatedAt:         val.CreatedAt.Format(time.RFC3339),
			Astrodevs:         make([]userdto.GetProfileTalentResponse, 0),
		})
	}

	return res, nil
}
