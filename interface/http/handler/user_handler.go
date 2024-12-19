package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase domain.UserUsecase) {
	userHandlerRoute := router.Group("/v1/api/users", guard.AuthGuard())

	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	userHandlerRoute.GET("/", userHandler.Pagination())
	userHandlerRoute.DELETE("/:id", userHandler.Delete())
	userHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.UserUpdateRequest{}), userHandler.Update())
	userHandlerRoute.GET("/:id", userHandler.FindById())
	userHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.UserCreateRequest{}), userHandler.Create())
}

func (userHandler *UserHandler) Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, total := userHandler.userUsecase.Pagination(ctx)

		meta := utils.PaginationMetaBuilder(ctx, total)

		ctx.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (userHandler *UserHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.UserCreateRequest)
		userHandler.userUsecase.Create(ctx, request)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (userHandler *UserHandler) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		userData, _ := userHandler.userUsecase.FindById(ctx, id, false)

		ctx.JSON(http.StatusOK, utils.SucessResponse(userData))
	}
}

func (userHandler *UserHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRequest := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.UserUpdateRequest)
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		updateUser := domain.UserEntity{
			Name:  userRequest.Name,
			Email: userRequest.Email,
		}
		userHandler.userUsecase.Update(ctx, id, &updateUser)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (userHandler *UserHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		userHandler.userUsecase.Delete(ctx, id)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
