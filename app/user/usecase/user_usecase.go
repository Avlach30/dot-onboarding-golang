package usecase

import (
	"context"
	"github.com/codespace-id/codespace-x/app/user/domain"
	userdto "github.com/codespace-id/codespace-x/app/user/dto"

	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/pkg/errors"
)

type userUsecase struct {
	userRepo domain.Repository
}

func NewUserUsecase(userRepo domain.Repository) domain.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Create(ctx context.Context, dto userdto.RegisterRequest) error {
	traceTag := "UserUsecase.Create"

	userData, err := u.userRepo.Find(ctx, dto.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, traceTag)
	}
	if userData.Fullname != "" {
		newErr := errors.New("DuplicatePhone")
		return errors.Wrap(newErr, traceTag)
	}

	if err := u.userRepo.Create(ctx, domain.Entity{
		Fullname:       dto.Fullname,
		IdentityNumber: dto.Email,
		PhoneNumber:    dto.PhoneNumber,
		Gender:         enum.UNKNOWN.Value(),
		Password:       "x",
	}); err != nil {
		return errors.WithMessage(err, traceTag)
	}

	return nil
}

func (u *userUsecase) Profile(ctx context.Context, phoneNumber string) (res domain.Entity, err error) {

	userData, err := u.userRepo.Find(ctx, phoneNumber)
	if err != nil {
		return res, errors.WithMessage(err, "UserUsecase.Profile")
	}

	return domain.Entity{
		Fullname:       userData.Fullname,
		IdentityNumber: userData.IdentityNumber,
		PhoneNumber:    userData.PhoneNumber,
		Gender:         userData.Gender,
	}, nil
}
