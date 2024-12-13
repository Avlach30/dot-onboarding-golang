package utils

import (
	"net/http"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
)

func SucessResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		StatusCode: http.StatusOK,
		Data:       &data,
		Version:    "1.0.0",
	})
}

func ErrorResponse(statusCode int, errorMessage string) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		StatusCode:   statusCode,
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

func PaginationBuilder(items []interface{}, meta pkg.MetaResponse) *pkg.PaginationResponse {
	return &(pkg.PaginationResponse{
		Items: &items,
		Meta:  &meta,
	})
}
