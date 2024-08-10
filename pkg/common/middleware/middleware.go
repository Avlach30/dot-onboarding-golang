package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/codespace-id/codespace-x/config"
	httperror "github.com/codespace-id/codespace-x/pkg/common/error"
	"github.com/codespace-id/codespace-x/pkg/jwt"
	"github.com/julienschmidt/httprouter"
)

type MiddlewareType struct {
	TokenAuth         bool
	XServiceAuthToken bool
}

// locals
type contextKey string

const (
	App         contextKey = "app"
	PhoneNumber contextKey = "phoneNumber"
	Role        contextKey = "role"
)

func Wrapper(next httprouter.Handle, middlewareType MiddlewareType) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), App, "codespace-x")

		if middlewareType.TokenAuth {
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

			ctx = context.WithValue(ctx, PhoneNumber, claims.PhoneNumber)
			ctx = context.WithValue(ctx, Role, claims.Role)

		}

		if middlewareType.XServiceAuthToken {
			serviceTokenReq := r.Header.Get("X-Service-Auth-Token")
			if serviceTokenReq == "" {
				httperror.SetResponse(w, 401, "unauthorized")
				return
			}

			serviceToken := config.ServiceAuthToken
			if serviceTokenReq != serviceToken {
				httperror.SetResponse(w, 401, "unauthorized")
				return
			}

		}

		next(w, r.WithContext(ctx), ps)
	}
}
