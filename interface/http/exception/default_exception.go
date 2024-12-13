package exception

import "net/http"

type Exception struct {
	StatusCode   int
	ErrorMessage string
}

func BussinessException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusUnprocessableEntity

	return panic
}

func ServerErrorException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusInternalServerError

	return panic
}

func UnathorizedException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusUnauthorized

	return panic
}

func ForbiddenException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusForbidden

	return panic
}

func NotFoundException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusNotFound

	return panic
}
