package handler

import (
	"encoding/json"
	authdomain "github.com/codespace-id/codespace-x/app/auth/domain"
	"github.com/codespace-id/codespace-x/app/auth/dto"
	userdomain "github.com/codespace-id/codespace-x/app/user/userdomain"
	"net/http"

	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"
	"github.com/codespace-id/codespace-x/pkg/jwt"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type AuthHandler struct {
	userUsecase userdomain.Usecase
	authUsecase authdomain.Usecase
}

func NewAuthHandler(router *httprouter.Router, userUsecase userdomain.Usecase, authUsecase authdomain.Usecase) {
	basePath := "/api/v1/auth"
	authHandler := &AuthHandler{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}

	router.POST(basePath+"/exchange-token", middleware.Wrapper(authHandler.ExchangeToken(), middleware.MiddlewareType{XServiceAuthToken: true}))
	router.POST(basePath+"/otp/request", middleware.Wrapper(authHandler.OtpRequest(), middleware.MiddlewareType{XServiceAuthToken: true}))
	router.POST(basePath+"/otp/validate", middleware.Wrapper(authHandler.OtpValidate(), middleware.MiddlewareType{XServiceAuthToken: true}))
	router.POST(basePath+"/otp/resend", middleware.Wrapper(authHandler.OtpResend(), middleware.MiddlewareType{XServiceAuthToken: true}))
	router.POST(basePath+"/phone/verify", middleware.Wrapper(authHandler.PhoneVerify(), middleware.MiddlewareType{XServiceAuthToken: true}))
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

// @Summary OTP Request
// @Description OTP Request
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param body-payload body dto.OtpRequest true "payload otp request"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/otp/request [post]
func (h *AuthHandler) OtpRequest() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		payloadReq := dto.OtpRequest{}
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

		if err := h.authUsecase.OtpRequest(r.Context(), payloadReq.PhoneNumber); err != nil {
			httperror.SetResponse(w, 400, err.Error())
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    nil,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary OTP Validate
// @Description OTP Validate
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param body-payload body dto.OtpValidateRequest true "payload otp request"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/otp/validate [post]
func (h *AuthHandler) OtpValidate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		payloadReq := dto.OtpValidateRequest{}
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

		if err := h.authUsecase.OtpValidate(r.Context(), dto.OtpValidateRequest{
			PhoneNumber: payloadReq.PhoneNumber,
			Otp:         payloadReq.Otp,
		}); err != nil {
			httperror.SetResponse(w, 400, err.Error())
			return
		}

		token, err := jwt.CreateToken(payloadReq.PhoneNumber, "CLIENT", "0")
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

// @Summary OTP Resend
// @Description OTP Resend
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param body-payload body dto.OtpRequest true "payload otp request resend"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/otp/resend [post]
func (h *AuthHandler) OtpResend() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		payloadReq := dto.OtpRequest{}
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

		if err := h.authUsecase.OtpResend(r.Context(), payloadReq.PhoneNumber); err != nil {
			httperror.SetResponse(w, 400, err.Error())
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    nil,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary Phone Verification
// @Description Used when register before hit register button
// @Tags Auth
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param body-payload body dto.OtpRequest true "payload otp request resend"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/phone/verify [post]
func (h *AuthHandler) PhoneVerify() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		payloadReq := dto.OtpRequest{}
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

		if err := h.authUsecase.PhoneVerify(r.Context(), payloadReq.PhoneNumber); err != nil {
			httperror.SetResponse(w, 400, err.Error())
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    "phone number available",
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
