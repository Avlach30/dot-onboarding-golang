package middleware

import (
	"context"
	"net/http"
	"strings"

	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/jwt"
	"github.com/julienschmidt/httprouter"
)

type MiddlewareType struct {
	CheckTokenAuth bool
}

// locals
type contextKey string

const (
	PhoneNumber contextKey = "phoneNumber"
	Role        contextKey = "role"
)

func Wrapper(next httprouter.Handle, middlewareType MiddlewareType) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var ctx context.Context

		if middlewareType.CheckTokenAuth {
			userAgent := r.Header.Get("Authorization")
			splitToken := strings.Split(userAgent, " ")
			if len(splitToken) < 2 {
				httperror.SetResponse(w, 401, "invalid token")
				return
			}

			claims, err := jwt.ParseToken(splitToken[1])
			if err != nil {
				httperror.SetResponse(w, 401, "invalid token")
				return
			}

			ctx = context.WithValue(r.Context(), PhoneNumber, claims.PhoneNumber)
			ctx = context.WithValue(ctx, Role, claims.Role)

		}

		next(w, r.WithContext(ctx), ps)
	}
}
