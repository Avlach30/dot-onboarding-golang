package singleton

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CircuitBreakerSingleton struct {
	CircuitBreaker   CircuitBreaker
	CurrentCondition CurrentErrorCondition
}

type CircuitBreaker struct {
	IgnoreFailureEndpoints []string
	ReadyToTrip            func(currentCondition *CurrentErrorCondition) bool
}

const (
	StateOpen     = "Open"
	StateHalfOpen = "Half Open"
	StateClose    = "Close"
)

type CurrentErrorCondition struct {
	State         string
	LastOpen      time.Time
	TotalFailures int64
	Requests      int64
}

func GetCircuitBreaker() *CircuitBreakerSingleton {
	return cbs
}

func CountRequestCircuitBreaker() {
	cbs.CountRequest()
}

func (cbs *CircuitBreakerSingleton) FailureHappend(httpContext *gin.Context) {
	if cbs.IsEndpointIgnored(httpContext.Request.URL.Path) {
		return
	}

	if cbs.CurrentCondition.LastOpen.IsZero() {
		cbs.CurrentCondition.LastOpen = time.Now()
	}

	cbs.CurrentCondition.TotalFailures++
}

func (cbs *CircuitBreakerSingleton) IsEndpointIgnored(endpoint string) bool {
	for _, ignoredEndpoint := range cbs.CircuitBreaker.IgnoreFailureEndpoints {
		if ignoredEndpoint == endpoint {
			return true
		}
	}

	return false
}

func (cbs *CircuitBreakerSingleton) CountRequest() {
	cbs.CurrentCondition.Requests++
}

func (cbs *CircuitBreakerSingleton) IsReadyToTrip() bool {
	// beforeCondition := cbs.CurrentCondition.State
	isReadyToTrip := cbs.CircuitBreaker.ReadyToTrip(&cbs.CurrentCondition)

	if cbs.CurrentCondition.State == StateHalfOpen {
		return cbs.CurrentCondition.Requests%2 == 0
	} else if cbs.CurrentCondition.State == StateClose {
		cbs.ResetCondition()
	}

	return isReadyToTrip
}

func (cbs *CircuitBreakerSingleton) ResetCondition() {
	cbs.CurrentCondition.State = "Close"
	cbs.CurrentCondition.Requests = 0
	cbs.CurrentCondition.TotalFailures = 0
	cbs.CurrentCondition.LastOpen = time.Time{}

}

func (cbs *CircuitBreakerSingleton) IsRequestAllowedHalfOpen() bool {
	return cbs.CurrentCondition.Requests%2 == 0
}
