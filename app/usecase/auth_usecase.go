package usecase

import (
	"context"
	"database/sql"
	userdomain "github.com/codespace-id/codespace-x/app/domain/user"
	"strconv"
	"time"

	authdomain "github.com/codespace-id/codespace-x/app/domain/auth"
	authdto "github.com/codespace-id/codespace-x/app/dto/auth"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/Integrations/otp"
	"github.com/pkg/errors"
)

type authUsecase struct {
	otpService otp.OtpProxy
	otpRepo    authdomain.OtpRepository
	userRepo   userdomain.Repository
}

func NewAuthUsecase(otp otp.OtpProxy, otpRepo authdomain.OtpRepository, userRepo userdomain.Repository) authdomain.Usecase {
	return &authUsecase{
		otpService: otp,
		otpRepo:    otpRepo,
		userRepo:   userRepo,
	}
}

// OtpRequest implements authdomain.Usecase.
func (uc *authUsecase) OtpRequest(ctx context.Context, phoneNumber string) error {
	errTrace := "authUsecase.OtpRequest"

	newOtp, err := otp.GenerateOTP(4)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}

	expiredInMins, err := strconv.ParseInt(config.OtpExpiredInMins, 10, 64)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}
	expiredAt := time.Now().UTC().Add(time.Minute * time.Duration(expiredInMins))

	if err := uc.otpRepo.Upsert(ctx, authdomain.OtpEntity{
		Code:       newOtp,
		Identifier: phoneNumber,
		Trial:      0,
		IsValid:    0,
		ExpiredAt:  sql.NullTime{Time: expiredAt, Valid: true},
	}); err != nil {
		return errors.WithMessage(err, errTrace)
	}

	uc.otpService.SendOTP(phoneNumber, newOtp)

	return nil
}

// OtpValidate implements authdomain.Usecase.
func (uc *authUsecase) OtpValidate(ctx context.Context, payload authdto.OtpValidateRequest) error {
	errTrace := "authUsecase.OtpValidate"

	otpData, err := uc.otpRepo.FindByIdentifier(ctx, payload.PhoneNumber)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}

	if payload.Otp != otpData.Code {
		return errors.New("wrong otp, please try again")
	}

	if time.Now().After(otpData.ExpiredAt.Time) {
		return errors.New("otp expired")
	}

	otpData.IsValid = 1
	otpData.Trial = 0
	uc.otpRepo.UpdateByIdentifier(ctx, payload.PhoneNumber, otpData)

	return nil
}

// OtpResend implements authdomain.Usecase.
func (uc *authUsecase) OtpResend(ctx context.Context, phoneNumber string) error {
	errTrace := "authUsecase.OtpResend"

	otpData, err := uc.otpRepo.FindByIdentifier(ctx, phoneNumber)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}

	if otpData.Trial >= 5 {
		return errors.New("max retry otp reached")
	}

	newOtp, err := otp.GenerateOTP(4)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}

	expiredInMins, err := strconv.ParseInt(config.OtpExpiredInMins, 10, 64)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}
	expiredAt := time.Now().UTC().Add(time.Minute * time.Duration(expiredInMins))

	otpData.Code = newOtp
	otpData.ExpiredAt = sql.NullTime{Time: expiredAt, Valid: true}
	otpData.Trial = otpData.Trial + 1
	uc.otpRepo.UpdateByIdentifier(ctx, phoneNumber, otpData)

	uc.otpService.SendOTP(phoneNumber, newOtp)

	return nil
}

// PhoneVerify implements authdomain.Usecase.
func (uc *authUsecase) PhoneVerify(ctx context.Context, phoneNumber string) error {
	errTrace := "authUsecase.PhoneVerify"

	userData, err := uc.userRepo.Find(ctx, phoneNumber)
	if err != nil {
		return errors.WithMessage(err, errTrace)
	}

	if userData.PhoneNumber != "" {
		return errors.New("phone number already used")
	}

	return nil
}
