package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	permissionUsecase domain.PermissionUsecase
}

func NewPermissionHandler(router *gin.Engine, permissionUsecase domain.PermissionUsecase) {
	permissionHandlerRoute := router.Group("/v1/api/permissions", guard.AuthGuard())

	permissionHandler := &PermissionHandler{
		permissionUsecase: permissionUsecase,
	}

	permissionHandlerRoute.GET("/", permissionHandler.Pagination())
	permissionHandlerRoute.DELETE("/:id", permissionHandler.Delete())
	permissionHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.PermissionUpdateRequest{}), permissionHandler.Update())
	permissionHandlerRoute.GET("/:id", permissionHandler.FindById())
	permissionHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.PermissionCreateRequest{}), permissionHandler.Create())
}

func (permissionHandler *PermissionHandler) Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, total := permissionHandler.permissionUsecase.Pagination(ctx)

		meta := utils.PaginationMetaBuilder(ctx, total)

		ctx.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (permissionHandler *PermissionHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permissionRequest := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.PermissionCreateRequest)
		newPermission := domain.Permission{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Create(ctx, &newPermission)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (permissionHandler *PermissionHandler) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		permissionData, _ := permissionHandler.permissionUsecase.FindById(ctx, id)

		ctx.JSON(http.StatusOK, utils.SucessResponse(permissionData))
	}
}

func (permissionHandler *PermissionHandler) FindByKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramKey := ctx.Param("key")
		permissionData, _ := permissionHandler.permissionUsecase.FindByKey(ctx, paramKey)

		ctx.JSON(http.StatusOK, utils.SucessResponse(permissionData))
	}
}

func (permissionHandler *PermissionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		permissionRequest := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.PermissionUpdateRequest)
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		updatePermission := domain.Permission{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Update(ctx, id, &updatePermission)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (permissionHandler *PermissionHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		permissionHandler.permissionUsecase.Delete(ctx, id)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
