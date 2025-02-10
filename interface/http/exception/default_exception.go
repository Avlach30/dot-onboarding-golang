package exception

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

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

func ServerErrorException(err error) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = err.Error()
	panic.StatusCode = http.StatusInternalServerError

	// Capture error to sentry
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "error",
		Message:  err.Error(),
		Level:    sentry.LevelError,
		Data: map[string]interface{}{
			"timestamp": time.Now().Unix(),
		},
	})
	sentry.CaptureException(err)

	return panic
}

func ServiceUnavailable() *Exception {
	panic := &Exception{}
	panic.ErrorMessage = "Service Unavailable"
	panic.StatusCode = http.StatusServiceUnavailable

	return panic
}

func UnauthorizedException(errorMessage string) *Exception {
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

func BadRequestException(errorMessage string) *Exception {
	panic := &Exception{}
	panic.ErrorMessage = errorMessage
	panic.StatusCode = http.StatusBadRequest

	return panic
}
