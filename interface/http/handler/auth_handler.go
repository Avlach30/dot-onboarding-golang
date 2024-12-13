package handler

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(router *gin.Engine, authUsecase domain.AuthUsecase) {
	authHandlerRoute := router.Group("/api/v1/auth")

	authHandler := &AuthHandler{
		authUsecase: authUsecase,
	}

	authHandlerRoute.POST("/sign-in", middleware.ValidateRequestJSON(&dto.AuthSignInRequest{}), authHandler.SignIn())
}

func (authHandler *AuthHandler) SignIn() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		authRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.AuthSignInRequest)
		token := authHandler.authUsecase.SignInJWT(authRequest.Email, authRequest.Password)

		data := &dto.AuthSignInResponse{
			Token: token,
			Type:  "Bearer",
		}

		response := utils.SucessResponse(data)

		httpContext.JSON(http.StatusOK, response)
	}
}
