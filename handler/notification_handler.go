package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
	"github.com/julienschmidt/httprouter"
)

type NotificationHandler struct {
}

func NewNotificationHandler(router *httprouter.Router) {
	basePath := "/api/v1/notifications"
	notificationHandler := &NotificationHandler{}

	router.POST(basePath, notificationHandler.ListNotification())

}

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
