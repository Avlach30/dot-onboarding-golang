package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/middleware"
	"github.com/julienschmidt/httprouter"
)

type NotificationHandler struct {
}

func NewNotificationHandler(router *httprouter.Router) {
	basePath := "/api/v1/notifications"
	notificationHandler := &NotificationHandler{}

	router.POST(basePath, middleware.Wrapper(notificationHandler.ListNotification(), middleware.MiddlewareType{TokenAuth: true, XServiceAuthToken: true}))

}

// @Summary List Notif
// @Description List Notif
// @Tags Notifications
// @Accept json
// @Produce json
// @Param X-Service-Auth-Token header string true "X-Service-Auth-Token"
// @Param authorization header string true "Authorization value"
// @Success 200 {object} pkg.BaseResponse{} "success"
// @Failure default {object} pkg.BaseResponse "error"
// @Router /api/v1/notifications [get]
func (h *NotificationHandler) ListNotification() httprouter.Handle {
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
