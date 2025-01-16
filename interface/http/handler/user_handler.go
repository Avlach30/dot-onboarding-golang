package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase domain.UserUsecase) {
	userHandlerRoute := router.Group("/api/v1/users", guard.AuthGuard())

	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	userHandlerRoute.GET("/", userHandler.Pagination())
	userHandlerRoute.DELETE("/:id", userHandler.Delete())
	userHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.UserUpdateRequest{}), userHandler.Update())
	userHandlerRoute.GET("/:id", userHandler.Detail())
	userHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.UserCreateRequest{}), userHandler.Create())
}

func (userHandler *UserHandler) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		data, total := userHandler.userUsecase.Pagination(httpContext)

		meta := utils.PaginationMetaBuilder(httpContext, total)

		httpContext.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (userHandler *UserHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		request := singleton.GetHTTPRequest[dto.UserCreateRequest](httpContext)
		newUser := domain.UserEntity{
			Name:  request.Name,
			Email: request.Email,
		}
		userHandler.userUsecase.Create(httpContext, &newUser, request.RoleIds)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (userHandler *UserHandler) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		userData := userHandler.userUsecase.FindOneById(httpContext, id, false)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(userData))
	}
}

func (userHandler *UserHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		request := singleton.GetHTTPRequest[dto.UserUpdateRequest](httpContext)
		newUser := domain.UserEntity{
			Name:  request.Name,
			Email: request.Email,
		}
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)

		userHandler.userUsecase.Update(httpContext, id, newUser, request.RoleIds)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (userHandler *UserHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		userHandler.userUsecase.Delete(httpContext, id)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
