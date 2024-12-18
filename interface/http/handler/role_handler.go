package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
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
	roleHandlerRoute.GET("/:id", roleHandler.FindById())
	roleHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.RoleCreateRequest{}), roleHandler.Create())
}

func (roleHandler *RoleHandler) Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, total := roleHandler.roleUsecase.Pagination(ctx)

		meta := utils.PaginationMetaBuilder(ctx, total)

		ctx.JSON(http.StatusOK, utils.PaginationBuilder(data, *meta))
	}
}

func (roleHandler *RoleHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleRequest := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.RoleCreateRequest)
		newRole := domain.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}

		roleHandler.roleUsecase.Create(ctx, &newRole)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (roleHandler *RoleHandler) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		roleData, _ := roleHandler.roleUsecase.FindById(ctx, id)

		ctx.JSON(http.StatusOK, utils.SucessResponse(roleData))
	}
}

func (roleHandler *RoleHandler) FindByKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramKey := ctx.Param("key")
		roleData, _ := roleHandler.roleUsecase.FindByKey(ctx, paramKey)

		ctx.JSON(http.StatusOK, utils.SucessResponse(roleData))
	}
}

func (roleHandler *RoleHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleRequest := ctx.MustGet(constant.RequestBodyJSONKey).(*dto.RoleUpdateRequest)
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		updateRole := domain.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}
		roleHandler.roleUsecase.Update(ctx, id, &updateRole)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (roleHandler *RoleHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramId := ctx.Param("id")
		id := utils.UUIDChecker(paramId)
		roleHandler.roleUsecase.Delete(ctx, id)

		ctx.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}
