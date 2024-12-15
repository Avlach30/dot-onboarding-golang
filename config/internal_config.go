package config

var (
	// General
	AppMode = Get("APP_MODE", "PROD")
	GinMode = Get("GIN_MODE", "release")

	AppPort = Get("APP_PORT", "8080")

	SentryDSN         = Get("SENTRY_DSN", "")
	SentrySampleTrace = Get("SENTRY_SAMPLE_TRACE", "0.1")

	Secret           = GetRequired("JWT_SECRET")
	JwtExpiredInDays = GetRequired("JWT_EXPIRED_IN_DAYS")
	ServiceAuthToken = GetRequired("SERVICE_AUTH_TOKEN")

	// DATABASE
	DBHost     = GetRequired("DB_HOST")
	DBPort     = GetRequired("DB_PORT")
	DBUsername = GetRequired("DB_USER")
	DBPassword = GetRequired("DB_PASS")
	DBName     = GetRequired("DB_NAME")
	DBTimeZone = GetRequired("DB_TIMEZONE")

	GlobalStateDriver = Get("GLOBAl_STATE_DRIVER", "runtime")

	// REDIS
	RedisHost     = Get("REDIS_HOST", "")
	RedisPassword = Get("REDIS_PASS", "")
	RedisPort     = Get("REDIS_PORT", "")
)
