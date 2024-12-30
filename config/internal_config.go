package config

var (
	// GENERAL
	AppMode = Get("APP_MODE", "PROD")
	GinMode = Get("GIN_MODE", "release")
	AppPort = Get("APP_PORT", "8080")

	// DATABASE
	DBHost     = GetRequired("DB_HOST")
	DBPort     = GetRequired("DB_PORT")
	DBUsername = GetRequired("DB_USER")
	DBPassword = GetRequired("DB_PASS")
	DBName     = GetRequired("DB_NAME")
	DBTimeZone = GetRequired("DB_TIMEZONE")

	// DATABASE CONNECTION POOL
	DBMaxIdleConn     = Get("DB_MAX_IDLE_CONNS", "10")
	DBMaxOpenConn     = Get("DB_MAX_OPEN_CONNS", "20")
	DBConnMaxIdleTime = Get("DB_CONN_MAX_IDLETIME_IN_MINUTES", "60")
	DBConnMaxLifetime = Get("DB_CONN_MAX_LIFETIME_IN_MINUTES", "10")

	GlobalStateDriver = Get("GLOBAl_STATE_DRIVER", "runtime")

	MaxSizeUploadFile = Get("MAX_SIZE_UPLOAD_FILE", "10240000") // 10mb

	LogDriver     = Get("LOG_DRIVER", "file")
	LogFolderPath = Get("LOG_FOLDER_PATH", "log")
)
