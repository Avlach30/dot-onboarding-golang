package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
}

func NewUserHandler(router *httprouter.Router) {
	basePath := "/api/v1/users"
	userHandler := &UserHandler{}

	router.GET(basePath+"/profile", userHandler.Profile())

}

func (h *UserHandler) Profile() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		data := map[string]interface{}{}
		data["fullname"] = "Ubaidillah Hakim Fadly Bro"
		data["email"] = "ubaidillahhf@gmail.com"
		data["user_type"] = "Personal"

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
