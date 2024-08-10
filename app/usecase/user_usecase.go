package usecase

import (
	"context"

	domainUser "github.com/codespace-id/codespace-x/app/domain/user"
	userdto "github.com/codespace-id/codespace-x/app/dto/user"
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

	if err := u.userRepo.Create(ctx, domainUser.Entity{
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

func (u *userUsecase) Profile(ctx context.Context, phoneNumber string) (res domainUser.Entity, err error) {

	userData, err := u.userRepo.Find(ctx, phoneNumber)
	if err != nil {
		return res, errors.WithMessage(err, "UserUsecase.Profile")
	}

	return domainUser.Entity{
		Fullname:       userData.Fullname,
		IdentityNumber: userData.IdentityNumber,
		PhoneNumber:    userData.PhoneNumber,
		Gender:         userData.Gender,
	}, nil
}
