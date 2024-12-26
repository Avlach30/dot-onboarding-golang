package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
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
	roleHandlerRoute := router.Group("/v1/api/roles", guard.AuthGuard())

	roleHandler := &RoleHandler{
		roleUsecase: roleUsecase,
	}

	roleHandlerRoute.GET("/", roleHandler.Pagination())
	roleHandlerRoute.DELETE("/:id", roleHandler.Delete())
	roleHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.RoleUpdateRequest{}), roleHandler.Update())
	roleHandlerRoute.GET("/:id", roleHandler.Detail())
	roleHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.RoleCreateRequest{}), roleHandler.Create())
}

func (roleHandler *RoleHandler) Pagination() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		data, total := roleHandler.roleUsecase.Pagination(httpContext)

		meta := utils.PaginationMetaBuilder(httpContext, total)

		httpContext.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (roleHandler *RoleHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := singleton.GetHTTPRequest[dto.RoleCreateRequest](httpContext)
		newRole := domain.RoleEntity{
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
		updateRole := domain.RoleEntity{
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
