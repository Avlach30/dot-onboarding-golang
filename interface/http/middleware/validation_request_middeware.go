package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
)

func ValidateRequestJSON[T any](obj *T) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		if err := httpContext.ShouldBindJSON(obj); err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			httpContext.Abort()
			return
		}

		httpContext.Set(constant.RequestBodyJSONKey, obj) // Pass the parsed struct to the handler
	}
}
