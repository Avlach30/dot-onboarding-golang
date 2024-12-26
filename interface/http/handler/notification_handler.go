package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/notification/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

type NotificationHandler struct {
	notificationUseCase domain.NotificationUseCase
}

func NewNotificationHandler(router *gin.Engine, notificationUseCase domain.NotificationUseCase) {
	notificationHandlerRoute := router.Group("/v1/api/notifications", guard.AuthGuard())

	notificationHandler := &NotificationHandler{
		notificationUseCase: notificationUseCase,
	}

	notificationHandlerRoute.GET("/", notificationHandler.Pagination())
	notificationHandlerRoute.GET("/has-unread", notificationHandler.HasUnread())
	notificationHandlerRoute.PATCH("/mark-as-read/:id", notificationHandler.MarkAsRead())
	notificationHandlerRoute.GET("/:id", notificationHandler.Detail())
}

func (notificationHandler *NotificationHandler) Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimToken := ctx.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		data, total := notificationHandler.notificationUseCase.Pagination(ctx, userId)

		meta := utils.PaginationMetaBuilder(ctx, total)

		ctx.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (notificationHandler *NotificationHandler) HasUnread() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimToken := ctx.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		hasUnread := notificationHandler.notificationUseCase.HasUnread(ctx, userId)

		ctx.JSON(http.StatusOK, utils.SucessResponse(hasUnread))
	}
}

func (notificationHandler *NotificationHandler) MarkAsRead() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimToken := ctx.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		notificationHandler.notificationUseCase.MarkAsRead(ctx, id, userId)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (notificationHandler *NotificationHandler) Detail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claimToken := ctx.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)

		data := notificationHandler.notificationUseCase.Detail(ctx, id, userId)

		ctx.JSON(http.StatusOK, utils.SucessResponse(data))
	}
}
