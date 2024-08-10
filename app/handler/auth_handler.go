package handler

import (
	"encoding/json"
	"net/http"

	userdomain "github.com/codespace-id/codespace-x/app/domain/user"
	"github.com/codespace-id/codespace-x/app/dto"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"
	"github.com/codespace-id/codespace-x/pkg/jwt"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type AuthHandler struct {
	userUsecase userdomain.Usecase
}

func NewAuthHandler(router *httprouter.Router, userUsecase userdomain.Usecase) {
	basePath := "/api/v1/auth"
	authHandler := &AuthHandler{
		userUsecase: userUsecase,
	}

	router.POST(basePath+"/exchange-token", middleware.Wrapper(authHandler.ExchangeToken(), middleware.MiddlewareType{XServiceAuthToken: true}))

}

// @Summary Exchange Token
// @Description Exchange Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param body-payload body dto.ExchangeRequest true "payload exchange token"
// @Success 200 {object} pkg.BaseResponse{data=dto.ExchangeResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/exchange-token [post]
func (h *AuthHandler) ExchangeToken() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		payloadReq := dto.ExchangeRequest{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payloadReq); err != nil {
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

		token, err := jwt.CreateToken(payloadReq.PhoneNumber, "CLIENT", payloadReq.FirebaseIdToken)
		if err != nil {
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    dto.ExchangeResponse{Token: token},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
