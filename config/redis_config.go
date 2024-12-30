package config

var (
	// REDIS
	RedisHost     = Get("REDIS_HOST", "")
	RedisPassword = Get("REDIS_PASS", "")
	RedisPort     = Get("REDIS_PORT", "")
)
