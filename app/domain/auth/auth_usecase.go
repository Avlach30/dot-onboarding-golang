package authdomain

import (
	"context"

	authdto "github.com/codespace-id/codespace-x/app/dto/auth"
)

type Usecase interface {
	OtpRequest(ctx context.Context, phoneNumber string) error
	OtpValidate(ctx context.Context, payload authdto.OtpValidateRequest) error
	OtpResend(ctx context.Context, phoneNumber string) error
	PhoneVerify(ctx context.Context, phoneNumber string) error
}
