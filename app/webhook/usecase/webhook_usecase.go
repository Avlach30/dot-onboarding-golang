package usecase

import (
	"context"
	"github.com/codespace-id/codespace-x/app/webhook/domain"
	webhookDto "github.com/codespace-id/codespace-x/app/webhook/dto"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/Integrations/notifications"
	"github.com/codespace-id/codespace-x/pkg/common/formatter"
	"golang.org/x/text/language"
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

	webhookTitle := "**Xendit Disbursement SUCCESS** 💸"
	if reqDto.Status != "COMPLETED" {
		webhookTitle = "**Xendit Disbursement FAILED** ⛔"
	}

	amountAsString := formatter.CurrencyDotSeparator(language.Indonesian, "IDR", reqDto.Amount)

	uc.discordNotif.Send(config.WebhookNewOutPayments, webhookTitle+"\n\n Transfer ke: "+reqDto.AccountHolderName+" \n Bank: "+reqDto.BankCode+" \n Amount: "+amountAsString+" \n Description: "+reqDto.DisbursementDescription+"")

	return nil
}

func (uc *webhookUsecase) BatchDisbursement(ctx context.Context, reqDto webhookDto.WebhookBatchDisbursementReq) (err error) {

	webhookTitle := "**Xendit BATCH Disbursement SUCCESS** 💰"
	if reqDto.Status != "COMPLETED" {
		webhookTitle = "**Xendit BATCH Disbursement FAILED** ⛔"
	}

	totalRequest := strconv.Itoa(reqDto.TotalRequestBatch)
	totalRequestAmount := formatter.CurrencyDotSeparator(language.Indonesian, "IDR", reqDto.TotalRequestBatchAmount)
	totalExecuted := strconv.Itoa(reqDto.TotalExecutedBatch)
	totalExecutedAmount := formatter.CurrencyDotSeparator(language.Indonesian, "IDR", reqDto.TotalExecutedBatchAmount)
	totalError := strconv.Itoa(reqDto.TotalErrorBatch)
	totalErrorAmount := formatter.CurrencyDotSeparator(language.Indonesian, "IDR", reqDto.TotalErrorBatchAmount)

	uc.discordNotif.Send(config.WebhookNewOutPayments, webhookTitle+"\n\nTotal Request: "+totalRequest+"\nTotal Request Amount: "+totalRequestAmount+"\nTotal Executed: "+totalExecuted+"\nTotal Executed Amount: "+totalExecutedAmount+"\nTotal Error : "+totalError+"\nTotal Error Amount: "+totalErrorAmount+"\nReference: "+reqDto.Reference+"")

	return nil
}
