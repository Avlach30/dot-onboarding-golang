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
	AppMode          = Get("APP_MODE", "PROD")
	Secret           = GetRequired("JWT_SECRET")
	JwtExpiredInDays = GetRequired("JWT_EXPIRED_IN_DAYS")
	ServiceAuthToken = GetRequired("SERVICE_AUTH_TOKEN")

	// MySQL
	Host     = GetRequired("MYSQL_DB_HOST")
	Username = GetRequired("MYSQL_DB_USER")
	Password = GetRequired("MYSQL_DB_PASS")
	Database = GetRequired("MYSQL_DB_NAME")

	// OTP
	ZenzivaBaseURL   = Get("ZENZIVA_BASE_URL", "")
	ZenzivaUserKey   = Get("ZENZIVA_USER_KEY", "")
	ZenzivaPassKey   = Get("ZENZIVA_PASS_KEY", "")
	OtpExpiredInMins = Get("OTP_EXPIRED_IN_MINS", "5")

	// Discord
	WebhookNewInquiry     = Get("DISCORD_WEBHOOK_NEW_INQUIRY", "")
	WebhookNewOutPayments = Get("DISCORD_WEBHOOK_NEW_OUT_PAYMENTS", "")

	// Xendit
	XenditWebhookToken = Get("XENDIT_WEBHOOK_TOKEN", "")
)
