package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return nil
}

func Get(key, defaultValue string) string {
	LoadConfig()

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetRequired(key string) string {
	LoadConfig()

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

var (
	// General
	AppMode = Get("APP_MODE", "PROD")

	SentryDSN         = Get("SENTRY_DSN", "PROD")
	SentrySampleTrace = Get("SENTRY_SAMPLE_TRACE", "1.0")

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

	// REDIS
	RedisHost     = Get("REDIS_HOST", "")
	RedisPassword = Get("REDIS_Password", "")
	RedisPort     = Get("REDIS_PORT", "")

	// OTP
	ZenzivaBaseURL   = Get("ZENZIVA_BASE_URL", "")
	ZenzivaUserKey   = Get("ZENZIVA_USER_KEY", "")
	ZenzivaPassKey   = Get("ZENZIVA_PASS_KEY", "")
	OtpExpiredInMins = Get("OTP_EXPIRED_IN_MINS", "5")

	// Xendit
	XenditWebhookToken = Get("XENDIT_WEBHOOK_TOKEN", "")

	// Tester Account
	TestAccountPhoneNumber = Get("TEST_ACCOUNT_PHONE_NUMBER", "")
	TestOtpNumber          = Get("TEST_ACCOUNT_OTP", "")
)
