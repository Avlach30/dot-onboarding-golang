package handler

import (
	"encoding/json"
	"github.com/codespace-id/codespace-x/app/webhook/domain"
	webhookDto "github.com/codespace-id/codespace-x/app/webhook/dto"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type webhookHandler struct {
	webhookUC domain.WebhookUsecase
}

func NewWebhookHandler(router *httprouter.Router, webhookUC domain.WebhookUsecase) {
	basePath := "/webhook"
	webhookHandler := &webhookHandler{
		webhookUC,
	}

	router.POST(basePath+"/disbursement", webhookHandler.XenditDisbursementCallback())
}

func (h *webhookHandler) XenditDisbursementCallback() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// validate, the request is originally from xendit
		callbackToken := r.Header.Get("x-callback-token")
		if callbackToken != config.XenditWebhookToken {
			httperror.SetResponse(w, 401, "unauthorized")
			return
		}

		var err error
		var payloadReq webhookDto.WebhookDisbursementReq

		decoder := json.NewDecoder(r.Body)
		if err = decoder.Decode(&payloadReq); err != nil {
			httperror.SetResponse(w, 400, "body payload required")
			return
		}
		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			httperror.SetResponse(w, 400, errMsgs)
			return
		}
		defer r.Body.Close()

		err = h.webhookUC.Disbursement(r.Context(), payloadReq)
		if err != nil {
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    nil,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
