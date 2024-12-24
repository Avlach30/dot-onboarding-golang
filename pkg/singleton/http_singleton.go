package singleton

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
)

func GetContextFromGinContext(httpContext *gin.Context) *context.Context {
	context := httpContext.Request.Context()
	return &context
}

func GetAuthUserID(httpContext *gin.Context) uuid.UUID {
	authEntityInfo, isExists := httpContext.Get(constant.AuthUserInfoKey)
	if !isExists {
		return uuid.Nil
	}

	return authEntityInfo.(*jwt.CustomClaims).ID
}

func GetHTTPRequest[T any](httpContext *gin.Context) *T {
	return httpContext.MustGet(constant.RequestBodyJSONKey).(*T)
}
