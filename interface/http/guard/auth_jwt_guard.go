package guard

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
)

// Custom guard to check JWT token in Authorization header
func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(*exception.UnathorizedException("Authorization header is required"))
		}

		// Check if the token starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			panic(*exception.UnathorizedException("Token not valid"))
		}

		// Extract the token from the "Bearer " prefix
		tokenString := authHeader[len("Bearer "):]

		// Here, you would normally verify the token (e.g., using JWT)
		_, err := jwt.ParseToken(tokenString)
		if err != nil {
			panic(*exception.UnathorizedException("Token not valid"))
		}

		// Token is valid, continue processing the request
		c.Next()
	}
}
