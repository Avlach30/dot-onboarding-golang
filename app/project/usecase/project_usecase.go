package usecase

import (
	"context"
	"encoding/json"
	commonrepo "github.com/codespace-id/codespace-x/app/common/repository"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/app/project/dto"
	userdto "github.com/codespace-id/codespace-x/app/user/dto"
	"github.com/codespace-id/codespace-x/app/user/userdomain"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/Integrations/notifications"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/codespace-id/codespace-x/pkg/common/generator"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type projectUsecase struct {
	projectRepo        domain.Repository
	sqlTxRepo          commonrepo.SqlTx
	userProjectRepo    domain.UserProjectRepository
	userRepo           userdomain.Repository
	projectImagesRepo  domain.ProjectImagesRepository
	projectHistoryRepo domain.ProjectHistoryRepository
	discordNotif       notifications.NotificationProxy
}

func NewProjectUsecase(
	projectRepo domain.Repository,
	sqlTxRepo commonrepo.SqlTx,
	userProjectRepo domain.UserProjectRepository,
	userRepo userdomain.Repository,
	projectImagesRepo domain.ProjectImagesRepository,
	projectHistoryRepo domain.ProjectHistoryRepository,
	discordNotif notifications.NotificationProxy,
) domain.Usecase {
	return &projectUsecase{
		projectRepo:        projectRepo,
		sqlTxRepo:          sqlTxRepo,
		userProjectRepo:    userProjectRepo,
		userRepo:           userRepo,
		projectImagesRepo:  projectImagesRepo,
		projectHistoryRepo: projectHistoryRepo,
		discordNotif:       discordNotif,
	}
}

// CreateNewInquiry implements projectdomain.Usecase.
func (uc *projectUsecase) CreateNewInquiry(ctx context.Context, phoneNumber string, entity domain.Entity) (res domain.Entity, err error) {

	user, err := uc.userRepo.Find(ctx, phoneNumber)
	if err != nil {
		return res, errors.WithMessage(err, "ProjectUsecase.CreateNewInquiry")
	}

	if entity.Name == "" {
		entity.Name = generator.GenerateRandomProjectName()
	}

	dbTx, err := uc.sqlTxRepo.Begin(ctx)
	if err != nil {
		return res, errors.WithMessage(err, "ProjectUsecase.CreateNewInquiry")
	}

	project, err := uc.projectRepo.CreateTx(ctx, dbTx, entity)
	if err != nil {
		dbTx.Rollback()
		return res, errors.WithMessage(err, "ProjectUsecase.CreateNewInquiry")
	}
	if project.UUID == "" {
		dbTx.Rollback()
		return res, errors.WithMessage(errors.New("Create Project Failed"), "ProjectUsecase.CreateNewInquiry")
	}

	if err := uc.projectImagesRepo.CreateTx(ctx, dbTx, generator.GenerateRandomProjectImage(), project.ID); err != nil {
		dbTx.Rollback()
		return res, errors.WithMessage(err, "ProjectUsecase.CreateNewInquiry")
	}

	if err := uc.userProjectRepo.CreateTx(ctx, dbTx, user.ID, project.ID); err != nil {
		dbTx.Rollback()
		return res, errors.WithMessage(err, "ProjectUsecase.CreateNewInquiry")
	}

	err = dbTx.Commit()

	uc.discordNotif.Send(config.WebhookNewInquiry, "**New Inquiry From CodespaceX App** 🥳\nCC: Sam <@422102506198401054> \n\nNama: "+user.Fullname+"\nPhone: "+phoneNumber+"\nProject Name: "+entity.Name+"\n\nDeskripsi: \n"+entity.Description)

	res = domain.Entity{
		UUID: project.UUID,
	}

	return res, nil
}

// ListProject implements projectdomain.Usecase.
func (uc *projectUsecase) ListProject(ctx context.Context, phoneNumber, roles string, page int, perPage int) (res []dto.ListProjectResponse, err error) {

	var projectData []domain.Entity

	projectData, err = uc.projectRepo.GetByPhoneNumber(ctx, phoneNumber, page, perPage)
	if err != nil {
		return nil, errors.WithMessage(err, "projectUsecase.ListProject")
	}

	if strings.Contains(roles, "admin") {
		projectData, err = uc.projectRepo.Get(ctx, page, perPage)
		if err != nil {
			return nil, errors.WithMessage(err, "projectUsecase.ListProject")
		}
	}

	for _, val := range projectData {

		var astroDev []userdto.GetProfileTalentResponse
		json.Unmarshal([]byte(val.Astrodevs), &astroDev)

		res = append(res, dto.ListProjectResponse{
			UUID:              val.UUID,
			Name:              val.Name,
			Description:       val.Description,
			ThumbnailImageURL: val.ThumbnailImageURL,
			ServiceType:       enum.GetTransformServiceType(val.ServiceType),
			Status:            val.Status,
			CreatedAt:         val.CreatedAt.Format(time.RFC3339),
			Astrodevs:         astroDev,
		})
	}

	return res, nil
}

// ProjectDetail implements projectdomain.Usecase.
func (uc *projectUsecase) ProjectDetail(ctx context.Context, UUID string) (res dto.ProjectDetailResponse, err error) {

	data, err := uc.projectRepo.Find(ctx, UUID)
	if err != nil {
		return res, errors.WithMessage(err, "projectUsecase.ProjectDetail")
	}

	var astroDev []userdto.GetProfileTalentResponse
	json.Unmarshal([]byte(data.Astrodevs), &astroDev)

	deadline := "Not Started Yet"
	if data.Status == enum.ON_DEVELOPMENT.Value() {
		deadline = enum.GetProjectTimeType(data.TargetTime)
	} else if data.Status == enum.BAST_AND_GUARANTEE.Value() {
		deadline = ""
	} else if data.Status == enum.FINISHED.Value() {
		deadline = ""
	}

	return dto.ProjectDetailResponse{
		UUID:              data.UUID,
		Name:              data.Name,
		Description:       data.Description,
		ThumbnailImageURL: data.ThumbnailImageURL,
		ServiceType:       enum.GetTransformServiceType(data.ServiceType),
		Status:            data.Status,
		CreatedAt:         data.CreatedAt.Format("2006-01-02 15:04:05"),
		Astrodevs:         astroDev,
		Deadline:          deadline,
	}, nil

}

// UpdateDetailProject implements projectdomain.Usecase.
func (uc *projectUsecase) UpdateDetailProject(ctx context.Context, dto domain.Entity) error {
	panic("unimplemented")
}

func (uc *projectUsecase) ListProjectHistory(ctx context.Context, projectUUID string, page int, perPage int) (res []dto.ProjectHistoryRes, err error) {

	projectData, err := uc.projectHistoryRepo.Get(ctx, projectUUID, page, perPage)
	if err != nil {
		return nil, errors.WithMessage(err, "projectUsecase.ListProjectHistory")
	}

	for _, val := range projectData {
		res = append(res, dto.ProjectHistoryRes{
			HistoryType:     val.HistoryType,
			Title:           val.Title,
			Description:     val.Description,
			AttachmentUrl:   val.AttachmentUrl,
			AttachmentTitle: val.AttachmentTitle,
			CreatedAt:       val.CreatedAt.Format(time.RFC3339),
		})
	}

	if len(res) <= 0 {
		return make([]dto.ProjectHistoryRes, 0), nil
	}

	return res, nil
}
