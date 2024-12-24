package handler

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
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
	authHandlerRoute.POST("/ldap", middleware.ValidateRequestJSON(&dto.AuthSignLDAPRequest{}), authHandler.SignLDAP())
	authHandlerRoute.POST("/oidc", middleware.ValidateRequestJSON(&dto.AuthSignOIDCRequest{}), authHandler.SignOIDC())
}

func (authHandler *AuthHandler) SignIn() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		authRequest := singleton.GetHTTPRequest[dto.AuthSignInRequest](httpContext)
		token, expirationTime := authHandler.authUsecase.SignInBasic(httpContext, authRequest.Email, authRequest.Password)

		data := &dto.AuthSignInResponse{
			Token:     token,
			ExpiredAt: expirationTime,
			Type:      "Bearer",
		}

		response := utils.SucessResponse(data)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (authHandler *AuthHandler) SignLDAP() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		authRequest := singleton.GetHTTPRequest[dto.AuthSignLDAPRequest](httpContext)
		token, expirationTime := authHandler.authUsecase.SignInLDAP(httpContext, authRequest.Username, authRequest.Password)

		data := &dto.AuthSignInResponse{
			Token:     token,
			ExpiredAt: expirationTime,
			Type:      "Bearer",
		}

		response := utils.SucessResponse(data)

		httpContext.JSON(http.StatusOK, response)
	}
}

func (authHandler *AuthHandler) SignOIDC() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		authRequest := singleton.GetHTTPRequest[dto.AuthSignOIDCRequest](httpContext)
		token, expirationTime := authHandler.authUsecase.SignInByOIDCCode(httpContext, authRequest.Code)

		data := &dto.AuthSignInResponse{
			Token:     token,
			ExpiredAt: expirationTime,
			Type:      "Bearer",
		}

		response := utils.SucessResponse(data)

		httpContext.JSON(http.StatusOK, response)
	}
}
