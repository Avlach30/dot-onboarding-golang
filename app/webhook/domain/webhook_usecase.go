package domain

import (
	"context"
	webhookDto "github.com/codespace-id/codespace-x/app/webhook/dto"
)

type WebhookUsecase interface {
	Disbursement(ctx context.Context, reqDto webhookDto.WebhookDisbursementReq) (err error)
}
