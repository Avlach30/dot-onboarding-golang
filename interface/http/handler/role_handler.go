package handler

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/constant"
	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase domain.RoleUsecase
}

func NewRoleHandler(router *gin.Engine, roleUsecase domain.RoleUsecase) {
	roleHandlerRoute := router.Group("/api/v1/roles", guard.AuthGuard())

	roleHandler := &RoleHandler{
		roleUsecase: roleUsecase,
	}

	roleHandlerRoute.GET("", guard.PermissionGuard(constant.ReadRole), roleHandler.Pagination())
	roleHandlerRoute.DELETE(":id", guard.PermissionGuard(constant.DeleteRole), roleHandler.Delete())
	roleHandlerRoute.PATCH(":id", guard.PermissionGuard(constant.UpdateRole), middleware.ValidateRequestJSON[dto.RoleUpdateRequest](), roleHandler.Update())
	roleHandlerRoute.GET(":id", guard.PermissionGuard(constant.ReadRole), roleHandler.Detail())
	roleHandlerRoute.POST("", guard.PermissionGuard(constant.CreateRole), middleware.ValidateRequestJSON[dto.RoleCreateRequest](), roleHandler.Create())
}

func (roleHandler *RoleHandler) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		data, total := roleHandler.roleUsecase.Pagination(httpContext)

		meta := utils.PaginationMetaBuilder(httpContext, total)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(data, *meta)))
	}
}

func (roleHandler *RoleHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := singleton.GetHTTPRequest[dto.RoleCreateRequest](httpContext)
		newRole := entities.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}

		roleHandler.roleUsecase.Create(httpContext, &newRole)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (roleHandler *RoleHandler) Detail() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleData := roleHandler.roleUsecase.FindOneById(httpContext, id)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(roleData))
	}
}

func (roleHandler *RoleHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := singleton.GetHTTPRequest[dto.RoleUpdateRequest](httpContext)
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		updateRole := entities.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}
		roleHandler.roleUsecase.Update(httpContext, id, &updateRole)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (roleHandler *RoleHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleHandler.roleUsecase.Delete(httpContext, id)

		httpContext.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
