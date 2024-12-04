package httperror

import (
	"encoding/json"
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
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
