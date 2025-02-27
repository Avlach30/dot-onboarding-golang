package guard

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
)

// Custom guard to check JWT token in Authorization header
func AuthGuard() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		authHeader := httpContext.GetHeader("Authorization")
		if authHeader == "" {
			panic(*exception.UnauthorizedException("Authorization header is required"))
		}

		// Check if the token starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			panic(*exception.UnauthorizedException("Token not valid"))
		}

		// Extract the token from the "Bearer " prefix
		tokenString := authHeader[len("Bearer "):]

		// Here, you would normally verify the token (e.g., using JWT)
		claimToken, err := jwt.ParseToken(tokenString)
		if err != nil {
			panic(*exception.UnauthorizedException("Token not valid (Claim Token)"))
		}

		// Set a boolean flag to indicate that the user is authorized and user info
		httpContext.Set(constant.IsAuthorizedHeaderKey, true)

		// Attach the user info to the context for further use
		httpContext.Set(constant.AuthUserInfoKey, claimToken)

		// Token is valid, continue processing the request
		httpContext.Next()
	}
}
