package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	userdomain "github.com/codespace-id/codespace-x/app/domain/user"
	userdto "github.com/codespace-id/codespace-x/app/dto/user"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/common/middleware"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	userUsecase userdomain.Usecase
}

func NewUserHandler(router *httprouter.Router, userUsecase userdomain.Usecase) {
	basePath := "/api/v1/users"
	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	router.GET(basePath+"/profile", middleware.Wrapper(userHandler.Profile(), middleware.MiddlewareType{CheckTokenAuth: true}))
	router.POST(basePath+"/register", middleware.Wrapper(userHandler.Register(), middleware.MiddlewareType{CheckTokenAuth: true}))

}

func (h *UserHandler) Profile() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var err error

		// Retrieve values from context (locals)
		phoneNumber, ok := r.Context().Value(middleware.PhoneNumber).(string)
		if !ok {
			httperror.SetResponse(w, 400, "invalid token, please relogin")
			return
		}

		res, err := h.userUsecase.Profile(r.Context(), phoneNumber)
		if err != nil {
			httperror.SetResponse(w, 400, "body payload required")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: userdto.GetProfileResponse{
				Fullname: res.Fullname,
				Role:     res.Role,
				ImageURL: res.ImageURL,
			},
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}

// @Summary Register
// @Description Register
// @Tags User
// @Accept json
// @Produce json
// @Param body-payload body dto.RegisterRequest true "payload"
// @Success 200 {object} pkg.BaseResponse{data=nil} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/users/register [post]
func (h *UserHandler) Register() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var err error
		var payloadReq userdto.RegisterRequest

		// Retrieve values from context (locals)
		phoneNumber, _ := r.Context().Value(middleware.PhoneNumber).(string)

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

		if payloadReq.PhoneNumber != phoneNumber {
			httperror.SetResponse(w, 400, "invalid token")
			return
		}

		if err = h.userUsecase.Create(r.Context(), payloadReq); err != nil {

			if strings.Contains(err.Error(), "DuplicatePhone") {
				httperror.SetResponse(w, 400, "already register, please login")
				return
			}

			httperror.SetResponse(w, 500, err.Error())
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
