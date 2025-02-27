package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SkipOptionsRequest wraps a middleware and skips it for OPTIONS requests
func SkipOptionsRequest(middleware gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		middleware(c)
	}
}
