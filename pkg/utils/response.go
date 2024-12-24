package utils

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
)

func SucessResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		StatusCode: http.StatusOK,
		Data:       &data,
		Version:    "1.0.0",
	})
}

func ErrorResponse(statusCode int, errorMessage string, stackTrace string) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		StatusCode:   statusCode,
		StackTrace:   stackTrace,
		ErrorMessage: errorMessage,
		Version:      "1.0.0",
	})
}

func ErrorValidationResponse(statusCode int, errors []pkg.ErrorValidation) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		StatusCode:   statusCode,
		ErrorMessage: "Validation Error",
		Errors:       &errors,
		Version:      "1.0.0",
	})
}

type PaginationResponse[T any] struct {
	Items *[]T              `json:"items"`
	Meta  *pkg.MetaResponse `json:"meta"`
}

func PaginationBuilder[T any](items []T, meta pkg.MetaResponse) *PaginationResponse[T] {
	return &(PaginationResponse[T]{
		Items: &items,
		Meta:  &meta,
	})
}

func PaginationMetaBuilder(httpContext *gin.Context, total int) *pkg.MetaResponse {
	// Get query params
	page, _ := strconv.Atoi(httpContext.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(httpContext.DefaultQuery("per_page", "10"))

	// Calculate total page
	totalPage := int(math.Ceil(float64(total) / float64(perPage)))

	return &(pkg.MetaResponse{
		Page:      page,
		PerPage:   perPage,
		Total:     total,
		TotalPage: totalPage,
	})
}
