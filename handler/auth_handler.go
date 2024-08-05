package handler

import (
	"encoding/json"
	"github.com/codespace-id/codespace-x/dto"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type AuthHandler struct {
}

func NewAuthHandler(router *httprouter.Router) {
	basePath := "/api/v1/auth"
	authHandler := &AuthHandler{}

	router.POST(basePath+"/register", authHandler.Register())
	router.POST(basePath+"/exchange-token", authHandler.ExchangeToken())

}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param body-payload body dto.RegisterRequest true "payload"
// @Success 200 {object} pkg.BaseResponse{data=nil} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var payloadReq dto.RegisterRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payloadReq); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"code":400,"message":"body payload required"}`))
			return
		}
		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			errByte, _ := json.Marshal(pkg.BaseResponse{
				Code:    400,
				Message: "error",
				Data:    errMsgs,
			})

			w.Header().Set("Content-Type", "application/json")
			w.Write(errByte)
			return
		}
		defer r.Body.Close()

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

// @Summary Exchange Token
// @Description Exchange Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body-payload body dto.ExchangeRequest true "payload exchange token"
// @Success 200 {object} pkg.BaseResponse{data=dto.ExchangeResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/auth/exchange-token [post]
func (h *AuthHandler) ExchangeToken() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var payloadReq dto.ExchangeRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payloadReq); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"code":400,"message":"body payload required"}`))
			return
		}
		// validate payload
		errMsgs := pkg.ValidateStruct(payloadReq)
		if len(errMsgs) > 0 {
			errByte, _ := json.Marshal(pkg.BaseResponse{
				Code:    400,
				Message: "error",
				Data:    errMsgs,
			})

			w.Header().Set("Content-Type", "application/json")
			w.Write(errByte)
			return
		}
		defer r.Body.Close()

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    dto.ExchangeResponse{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.08Kl3VCSoYS0T4jDEjxaTjff10yx_YC8END-X0ARU1o"},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
