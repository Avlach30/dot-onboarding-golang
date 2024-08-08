package usecase

import (
	"context"
	domainUser "github.com/codespace-id/codespace-x/app/domain/user"
	"github.com/codespace-id/codespace-x/app/dto"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/pkg/errors"
)

type userUsecase struct {
	userRepo domainUser.Repository
}

func NewUserUsecase(userRepo domainUser.Repository) domainUser.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Create(ctx context.Context, dto dto.RegisterRequest) error {

	if err := u.userRepo.Create(ctx, domainUser.Entity{
		Fullname:       dto.Fullname,
		IdentityNumber: dto.Email,
		PhoneNumber:    dto.PhoneNumber,
		Gender:         enum.UNKNOWN.Value(),
		Password:       "x",
	}); err != nil {
		return errors.WithMessage(err, "UserUsecase.Create")
	}

	return nil
}
