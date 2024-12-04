package domain

import (
	"context"
	authdto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/dto"
)

type Usecase interface {
	OtpRequest(ctx context.Context, phoneNumber string) error
	OtpValidate(ctx context.Context, payload authdto.OtpValidateRequest) error
	OtpResend(ctx context.Context, phoneNumber string) error
	PhoneVerify(ctx context.Context, phoneNumber string) error
}
