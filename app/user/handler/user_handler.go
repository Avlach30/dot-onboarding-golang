package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	userdomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/userdomain"

	httperror "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/error"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/middleware"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
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

	router.GET(basePath+"/profile", middleware.Wrapper(userHandler.Profile(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.POST(basePath+"/register", middleware.Wrapper(userHandler.Register(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))
	router.POST(basePath+"/delete", middleware.Wrapper(userHandler.Delete(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))

}

// @Summary Get Profile
// @Description Get Profile
// @Tags User
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Success 200 {object} pkg.BaseResponse{data=dto.GetProfileResponse} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/users/profile [get]
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
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
			Data: dto.GetProfileResponse{
				Fullname:    res.Fullname,
				ImageURL:    res.ImageURL,
				PhoneNumber: res.PhoneNumber,
				Email:       res.Email,
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
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Param body-payload body dto.RegisterRequest true "dto.RegisterRequest"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/users/register [post]
func (h *UserHandler) Register() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var err error
		var payloadReq dto.RegisterRequest

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

// @Summary Delete
// @Description Delete
// @Tags User
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/users/delete [post]
func (h *UserHandler) Delete() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
		var err error

		phoneNumber, ok := r.Context().Value(middleware.PhoneNumber).(string)
		if !ok {
			httperror.SetResponse(w, 500, "invalid token")
			return
		}

		err = h.userUsecase.Delete(r.Context(), phoneNumber)
		if err != nil {
			httperror.SetResponse(w, 500, "internal server error")
			return
		}

		dataByte, _ := json.Marshal(pkg.BaseResponse{
			Code:    200,
			Message: "success",
		})

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(dataByte)
		if err != nil {
			return
		}
	}
}
