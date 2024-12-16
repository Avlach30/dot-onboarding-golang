package handler

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

// RecoverPanic ...
func RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				panicException := exception.Exception{}

				switch v := err.(type) {
				case exception.Exception:
					panicException = v
				default:
					panicException.ErrorMessage = "Internal Server Error"
					panicException.StatusCode = http.StatusInternalServerError
				}

				isDebugMode := config.AppMode != "PROD"
				stackTrace := ""
				if isDebugMode {
					stackTrace = string(debug.Stack())
				}

				log.Println(err)
				log.Println(stackTrace)

				errorResponse := utils.ErrorResponse(panicException.StatusCode, panicException.ErrorMessage, stackTrace)
				c.JSON(panicException.StatusCode, errorResponse)
				c.Abort()
			}
		}()

		c.Next()
	}
}
