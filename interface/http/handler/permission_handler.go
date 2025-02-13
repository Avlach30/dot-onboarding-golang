package handler

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/constant"
	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type PermissionHandler struct {
	permissionUsecase domain.PermissionUsecase
}

func NewPermissionHandler(router *gin.Engine, permissionUsecase domain.PermissionUsecase) {
	permissionHandlerRoute := router.Group("/api/v1/permissions", guard.AuthGuard())

	permissionHandler := &PermissionHandler{
		permissionUsecase: permissionUsecase,
	}

	permissionHandlerRoute.GET("/", guard.PermissionGuard(constant.ReadPermission), permissionHandler.Pagination())
	permissionHandlerRoute.DELETE("/:id", guard.PermissionGuard(constant.DeletePermission), permissionHandler.Delete())
	permissionHandlerRoute.PATCH("/:id", guard.PermissionGuard(constant.UpdatePermission), middleware.ValidateRequestJSON[dto.PermissionUpdateRequest](), permissionHandler.Update())
	permissionHandlerRoute.GET("/:id", guard.PermissionGuard(constant.UpdatePermission), permissionHandler.Detail())
	permissionHandlerRoute.POST("/", guard.PermissionGuard(constant.ReadPermission), middleware.ValidateRequestJSON[dto.PermissionCreateRequest](), permissionHandler.Create())
}

func (permissionHandler *PermissionHandler) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		data, total := permissionHandler.permissionUsecase.Pagination(httpContext)

		meta := utils.PaginationMetaBuilder(httpContext, total)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(data, *meta)))
	}
}

func (permissionHandler *PermissionHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		permissionRequest := singleton.GetHTTPRequest[dto.PermissionCreateRequest](httpContext)
		newPermission := entities.PermissionEntity{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Create(httpContext, &newPermission)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (permissionHandler *PermissionHandler) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		permissionData := permissionHandler.permissionUsecase.FindOneById(httpContext, id)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(permissionData))
	}
}

func (permissionHandler *PermissionHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		permissionRequest := singleton.GetHTTPRequest[dto.PermissionUpdateRequest](httpContext)
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		updatePermission := entities.PermissionEntity{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Update(httpContext, id, &updatePermission)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (permissionHandler *PermissionHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		permissionHandler.permissionUsecase.Delete(httpContext, id)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
