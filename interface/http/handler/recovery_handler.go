package handler

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

// Recovery ...
func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f, gin.DefaultErrorWriter)
}

// Recovery500 ...
func Recovery500() gin.HandlerFunc {
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

				errorResponse := utils.ErrorResponse(panicException.StatusCode, panicException.ErrorMessage)
				c.JSON(panicException.StatusCode, errorResponse)
				c.Abort()
			}
		}()

		c.Next()
	}
}

// RecoveryWithWriter ...
func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	var logger *log.Logger
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if logger != nil {
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					reset := string([]byte{27, 91, 48, 109})
					logger.Printf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s", httprequest, err, reset)
				}

				f(c, err)
			}
		}()

		c.Next()
	}
}
