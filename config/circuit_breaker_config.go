package config

var (
	// OUTER CIRCUIT BREAKER
	IsCircuitBreakerExternalEnabled              = Get("IS_CIRCUIT_BREAKER_EXTERNAL_ENABLE", "false")
	CircuitBreakerExternalOpenDuration           = Get("CIRCUIT_BREAKER_EXTERNAL_OPEN_DURATION", "0.1")
	CircuitBreakerExternalHalfOpenDuration       = Get("CIRCUIT_BREAKER_EXTERNAL_HALF_OPEN_DURATION", "0.1")
	CircuitBreakerExternalIgnoreFailureEndpoints = Get("CIRCUIT_BREAKER_EXTERNAL_IGNORE_FAILURE_ENDPOINTS", "")
	CircuitBreakerExternalRequestFailedTreshold  = Get("CIRCUIT_BREAKER_EXTERNAL_REQUEST_FAILED_TRESHOLD", "3")
	CircuitBreakerExternalFailureRatioTreshold   = Get("CIRCUIT_BREAKER_EXTERNAL_FAILURE_RATIO_TRESHOLD", "0.5")

	IsCircuitBreakerIneternalEnabled             = Get("IS_CIRCUIT_BREAKER_INTERNAL_ENABLE", "false")
	CircuitBreakerInternalOpenDuration           = Get("CIRCUIT_BREAKER_INTERNAL_OPEN_DURATION", "0.1")
	CircuitBreakerInternalHalfOpenDuration       = Get("CIRCUIT_BREAKER_INTERNAL_HALF_OPEN_DURATION", "0.1")
	CircuitBreakerInternalIgnoreFailureEndpoints = Get("CIRCUIT_BREAKER_INTERNAL_IGNORE_FAILURE_ENDPOINTS", "")
	CircuitBreakerInternalRequestFailedTreshold  = Get("CIRCUIT_BREAKER_INTERNAL_REQUEST_FAILED_TRESHOLD", "3")
	CircuitBreakerInternalFailureRatioTreshold   = Get("CIRCUIT_BREAKER_INTERNAL_FAILURE_RATIO_TRESHOLD", "0.5")
)
