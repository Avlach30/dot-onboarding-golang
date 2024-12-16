package handler

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/job/start_job"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
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

		singleton.Delegate(start_job.TaskName, map[string]any{
			"email": "test@test.com",
			"name":  "test123",
			"jobs":  []string{"job1", "job2", "job3"},
			"origanization": map[string]any{
				"a": map[string]any{
					"b": "c",
				},
				"d": "e",
			},
		})

		authRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.AuthSignInRequest)
		token, expirationTime := authHandler.authUsecase.SignInBasic(authRequest.Email, authRequest.Password)

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
		authRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.AuthSignLDAPRequest)
		token, expirationTime := authHandler.authUsecase.SignInLDAP(authRequest.Username, authRequest.Password)

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
		authRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.AuthSignOIDCRequest)
		token, expirationTime := authHandler.authUsecase.SignInByOIDCCode(authRequest.Code)

		data := &dto.AuthSignInResponse{
			Token:     token,
			ExpiredAt: expirationTime,
			Type:      "Bearer",
		}

		response := utils.SucessResponse(data)

		httpContext.JSON(http.StatusOK, response)
	}
}
