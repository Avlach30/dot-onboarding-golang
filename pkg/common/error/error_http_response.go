package httperror

import (
	"encoding/json"
	"github.com/codespace-id/codespace-x/pkg"
	"net/http"
)

func BadRequest(w http.ResponseWriter, code int, message interface{}) {
	errByte, _ := json.Marshal(pkg.BaseResponse{
		Code:    code,
		Message: "error",
		Data:    message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(errByte)
}
