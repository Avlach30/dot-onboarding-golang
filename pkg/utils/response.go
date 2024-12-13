package utils

import "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"

func SucessResponse(data interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		Message: "Success",
		Data:    &data,
	})
}

func ErrorResponse(errors []interface{}) *pkg.BaseResponse {
	return &(pkg.BaseResponse{
		Message: "Failed",
		Errors:  &errors,
	})
}

func PaginationBuilder(items []interface{}, meta *pkg.MetaResponse) *pkg.PaginationResponse {
	return &(pkg.PaginationResponse{
		Items: &items,
		Meta:  meta,
	})
}
