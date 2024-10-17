package usecase

import (
	"context"

	userdto "github.com/codespace-id/codespace-x/app/user/dto"
	"github.com/codespace-id/codespace-x/app/user/userdomain"
	"github.com/codespace-id/codespace-x/pkg/common/generator"

	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/pkg/errors"
)

type userUsecase struct {
	userRepo userdomain.Repository
}

func NewUserUsecase(userRepo userdomain.Repository) userdomain.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Create(ctx context.Context, dto userdto.RegisterRequest) error {
	traceTag := "UserUsecase.CreateTx"

	userData, err := u.userRepo.Find(ctx, dto.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, traceTag)
	}
	if userData.Fullname != "" {
		newErr := errors.New("DuplicatePhone")
		return errors.Wrap(newErr, traceTag)
	}

	if err := u.userRepo.Create(ctx, userdomain.Entity{
		Fullname:       dto.Fullname,
		IdentityNumber: dto.Email,
		PhoneNumber:    dto.PhoneNumber,
		Gender:         enum.UNKNOWN.Value(),
		Password:       "x",
		ImageURL:       generator.GenerateRandomPhotoProfile(),
		Email:          dto.Email,
	}); err != nil {
		return errors.WithMessage(err, traceTag)
	}

	return nil
}

func (u *userUsecase) Profile(ctx context.Context, phoneNumber string) (res userdomain.Entity, err error) {

	userData, err := u.userRepo.Find(ctx, phoneNumber)
	if err != nil {
		return res, errors.WithMessage(err, "UserUsecase.Profile")
	}

	return userdomain.Entity{
		Fullname:       userData.Fullname,
		IdentityNumber: userData.IdentityNumber,
		PhoneNumber:    userData.PhoneNumber,
		Gender:         userData.Gender,
		Email:          userData.Email,
		ImageURL:       userData.ImageURL,
		Roles:          userData.Roles,
	}, nil
}

func (u *userUsecase) Delete(ctx context.Context, phoneNumber string) error {
	traceTag:= "UserUsecase.DeleteTx"

	if err := u.userRepo.Delete(ctx, phoneNumber); err != nil {
		return errors.WithMessage(err, traceTag)
	}

	return nil
}
