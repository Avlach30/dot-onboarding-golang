package httperror

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/pkg"
)

func SetResponse(w http.ResponseWriter, code int, message interface{}) {
	errByte, _ := json.Marshal(pkg.BaseResponse{
		Code:    code,
		Message: "error",
		Data:    message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(errByte)
}
