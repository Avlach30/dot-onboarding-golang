package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

func ValidateRequestJSON[T any]() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		obj := new(T)

		if err := httpContext.ShouldBind(obj); err != nil {

			if validationErrors, isNotValid := err.(validator.ValidationErrors); isNotValid {
				// Convert validation errors to ErrorValidation struct
				errors := make([]pkg.ErrorValidation, 0)

				for _, validationErr := range validationErrors {
					errors = append(errors, pkg.ErrorValidation{
						Key:     utils.StringToSnakeCase(validationErr.Field()), // Fully qualified field name
						Message: fmt.Sprintf("Error %s", validationErr.Tag()),
					})
				}

				// Print errors as JSON
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(http.StatusBadRequest, errors))
			} else {
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(http.StatusBadRequest, nil))
			}

			httpContext.Abort()
			return
		}

		httpContext.Set(constant.RequestBodyJSONKey, obj)
	}
}

func ValidateRequestFormData[T any](obj *T) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		if err := httpContext.ShouldBind(obj); err != nil {
			if validationErrors, isNotValid := err.(validator.ValidationErrors); isNotValid {
				errors := make([]pkg.ErrorValidation, len(validationErrors))

				for i, validationErr := range validationErrors {
					errors[i] = pkg.ErrorValidation{
						Key:     utils.StringToSnakeCase(validationErr.Field()),
						Message: fmt.Sprintf("Error %s", validationErr.Tag()),
					}
				}

				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(http.StatusBadRequest, errors))
			} else {
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(http.StatusBadRequest, nil))
			}

			httpContext.Abort()
			return
		}

		httpContext.Set(constant.RequestBodyJSONKey, obj)
		httpContext.Next()
	}
}
