package config

var (
	// SENTRY
	SentryDSN         = Get("SENTRY_DSN", "")
	SentrySampleTrace = Get("SENTRY_SAMPLE_TRACE", "0.1")
)
