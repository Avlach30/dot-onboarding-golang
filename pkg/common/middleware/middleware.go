package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/codespace-id/codespace-x/pkg"
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
	Roles       contextKey = "roles"
)

func Wrapper(next httprouter.Handle, middlewareType MiddlewareType) httprouter.Handle {

	next = emptyArrayInterceptor(next)

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), App, "codespace-x")

		if middlewareType.TokenAuth {
			var shouldReturn bool
			ctx, shouldReturn = authMiddleware(ctx, r, w)
			if shouldReturn {
				return
			}
		}

		if middlewareType.XServiceAuthToken {
			shouldReturn := serviceTokenMiddleware(r, w)
			if shouldReturn {
				return
			}

		}

		next(w, r.WithContext(ctx), ps)
	}
}

func serviceTokenMiddleware(r *http.Request, w http.ResponseWriter) bool {
	serviceTokenReq := r.Header.Get("X-Service-Auth-Token")
	if serviceTokenReq == "" {
		httperror.SetResponse(w, 401, "unauthorized")
		return true
	}

	serviceToken := config.ServiceAuthToken
	if serviceTokenReq != serviceToken {
		httperror.SetResponse(w, 401, "unauthorized")
		return true
	}

	return false
}

func authMiddleware(ctx context.Context, r *http.Request, w http.ResponseWriter) (context.Context, bool) {
	userAgent := r.Header.Get("Authorization")
	if userAgent == "" {
		// threat as guest
		return ctx, false
	}

	splitToken := strings.Split(userAgent, " ")
	if len(splitToken) < 2 {
		httperror.SetResponse(w, 401, "invalid token")
		return nil, true
	}

	claims, err := jwt.ParseToken(splitToken[1])
	if err != nil {
		httperror.SetResponse(w, 401, "invalid token")
		return nil, true
	}

	ctx = context.WithValue(ctx, PhoneNumber, claims.PhoneNumber)
	ctx = context.WithValue(ctx, Roles, claims.Roles)
	return ctx, false
}

func emptyArrayInterceptor(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Create a response writer that captures the response
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}

		// Call the next handler, and call again after finished
		next(rw, r, ps)

		var baseResponse pkg.BaseResponse
		json.Unmarshal(rw.body.Bytes(), &baseResponse)

		if baseResponse.Meta != nil && baseResponse.Data == nil {
			baseResponse.Data = []string{}
			dataByte, _ := json.Marshal(baseResponse)
			w.Write(dataByte)
		} else {
			dataByte, _ := json.Marshal(baseResponse)
			w.Write(dataByte)
		}
	}
}

// Custom response writer to capture the response body
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Reset()
	return rw.body.Write(b)
}
