package handler

import (
	"log"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

func isCheckCircuitBreakerOpen(httpContext *gin.Context) bool {
	isCircuitBreakerEnable, _ := strconv.ParseBool(config.IsCircuitBreakerIneternalEnabled)
	if isCircuitBreakerEnable {
		singleton.CountRequestCircuitBreaker(singleton.InternalCircuitBreaker)
		cbs := singleton.GetCircuitBreaker(singleton.InternalCircuitBreaker)
		if !cbs.IsReadyToTrip() {
			httpContext.JSON(http.StatusServiceUnavailable, utils.ErrorResponse(http.StatusServiceUnavailable, "Service Unavailable", ""))
			httpContext.Abort()
			return true
		}
	}

	return false
}

func RecoverPanic() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		if isCheckCircuitBreakerOpen(httpContext) {
			return
		}

		defer handlePanic(httpContext)
		httpContext.Next()
	}
}

func handlePanic(httoContext *gin.Context) {
	if err := recover(); err != nil {
		panicException := createPanicException(err)
		stackTrace := getStackTrace()

		cbs := singleton.GetCircuitBreaker(singleton.InternalCircuitBreaker)
		if cbs != nil && panicException.StatusCode == http.StatusInternalServerError {
			cbs.FailureHappend(httoContext.Request.URL.Path)
		}

		log.Println(stackTrace)

		errorResponse := utils.ErrorResponse(panicException.StatusCode, panicException.ErrorMessage, stackTrace)
		httoContext.JSON(panicException.StatusCode, errorResponse)
		httoContext.Abort()
	}
}

func createPanicException(err interface{}) exception.Exception {
	if ex, ok := err.(exception.Exception); ok {
		return ex
	}

	return exception.Exception{
		ErrorMessage: "Internal Server Error",
		StatusCode:   http.StatusInternalServerError,
	}
}

func getStackTrace() string {
	if config.AppMode != "PROD" {
		return string(debug.Stack())
	}

	return ""
}
