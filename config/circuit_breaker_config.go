package config

var (
	// OUTER CIRCUIT BREAKER
	IsCircuitBreakerEnabled              = Get("IS_CIRCUIT_BREAKER_ENABLE", "false")
	CircuitBreakerOpenDuration           = Get("CIRCUIT_BREAKER_OPEN_DURATION", "0.1")
	CircuitBreakerHalfOpenDuration       = Get("CIRCUIT_BREAKER_HALF_OPEN_DURATION", "0.1")
	CircuitBreakerIgnoreFailureEndpoints = Get("CIRCUIT_BREAKER_IGNORE_FAILURE_ENDPOINTS", "")
	CircuitBreakerRequestFailedTreshold  = Get("CIRCUIT_BREAKER_REQUEST_FAILED_TRESHOLD", "3")
	CircuitBreakerFailureRatioTreshold   = Get("CIRCUIT_BREAKER_FAILURE_RATIO_TRESHOLD", "0.5")
)
