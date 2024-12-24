package singleton

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GetContextFromGinContext(ctx *gin.Context) *context.Context {
	context := ctx.Request.Context()
	return &context
}
