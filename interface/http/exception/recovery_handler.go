package exception

import (
	"io"
	"log"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
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
				c.JSON(500, err)
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
