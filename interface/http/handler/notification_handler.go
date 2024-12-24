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
}

func (notificationHandler *NotificationHandler) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		claimToken := httpContext.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		data, total := notificationHandler.notificationUseCase.Pagination(httpContext, userId)

		meta := utils.PaginationMetaBuilder(httpContext, total)

		httpContext.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (notificationHandler *NotificationHandler) HasUnread() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		claimToken := httpContext.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		hasUnread := notificationHandler.notificationUseCase.HasUnread(httpContext, userId)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(hasUnread))
	}
}

func (notificationHandler *NotificationHandler) MarkAsRead() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		claimToken := httpContext.MustGet(constant.AuthUserInfoKey)
		userId := claimToken.(*jwt.CustomClaims).ID
		id := httpContext.Param("id")
		notificationHandler.notificationUseCase.MarkAsRead(httpContext, id, userId)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
