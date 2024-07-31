package handler

import (
	"encoding/json"
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

func (h *AuthHandler) Register() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *AuthHandler) ExchangeToken() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		// payload
		type payload struct {
			FirebaseIdToken string `json:"firebase_id_token" validate:"required"`
		}
		var payloadReq payload
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

		data := map[string]interface{}{}
		data["token"] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.08Kl3VCSoYS0T4jDEjxaTjff10yx_YC8END-X0ARU1o"

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data:    data,
		})

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
