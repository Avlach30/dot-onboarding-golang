package usecase

import (
	"context"
	"github.com/codespace-id/codespace-x/app/webhook/domain"
	webhookDto "github.com/codespace-id/codespace-x/app/webhook/dto"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/Integrations/notifications"
	"strconv"
)

type webhookUsecase struct {
	discordNotif notifications.NotificationProxy
}

func NewWebhookUsecase(
	discordNotif notifications.NotificationProxy,
) domain.WebhookUsecase {
	return &webhookUsecase{
		discordNotif: discordNotif,
	}
}

func (uc *webhookUsecase) Disbursement(ctx context.Context, reqDto webhookDto.WebhookDisbursementReq) (err error) {

	webhookTitle := "Xendit Disbursement SUCCESS 💸"
	if reqDto.Status != "COMPLETED" {
		webhookTitle = "Xendit Disbursement FAILED ⛔"
	}

	amountAsString := strconv.FormatFloat(reqDto.Amount, 'f', -1, 64)

	uc.discordNotif.Send(config.WebhookNewOutPayments, webhookTitle+"\n\n Transfer ke: "+reqDto.AccountHolderName+" \n Bank: "+reqDto.BankCode+" \n Amount: Rp. "+amountAsString+" \n Description: "+reqDto.DisbursementDescription+"")

	return nil
}
