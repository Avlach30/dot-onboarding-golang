package config

var (
	// General
	AppMode = Get("APP_MODE", "PROD")
	GinMode = Get("GIN_MODE", "release")

	AppPort = Get("APP_PORT", "8080")
	Storage = Get("STORAGE", "local")

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

	// DATABASE CONNECTION POOL
	DBMaxIdleConn     = Get("DB_MAX_IDLE_CONNS", "10")
	DBMaxOpenConn     = Get("DB_MAX_OPEN_CONNS", "20")
	DBConnMaxIdleTime = Get("DB_CONN_MAX_IDLETIME_IN_MINUTES", "60")
	DBConnMaxLifetime = Get("DB_CONN_MAX_LIFETIME_IN_MINUTES", "10")

	GlobalStateDriver = Get("GLOBAl_STATE_DRIVER", "runtime")

	// REDIS
	RedisHost     = Get("REDIS_HOST", "")
	RedisPassword = Get("REDIS_PASS", "")
	RedisPort     = Get("REDIS_PORT", "")

	MaxWorkerQueue        = Get("MAX_WORKER_QUEUE", "3")
	MaxParalelWorkerQueue = Get("MAX_PARALEL_WORKER_QUEUE", "3")
	MaxTriesQueue         = Get("MAX_TRIES", "3")

	MaxSizeUploadFile = Get("MAX_SIZE_UPLOAD_FILE", "10240000") // 10mb

	S3BucketName      = Get("S3_BUCKET_NAME", "")
	S3AccessKeyID     = Get("S3_ACCESS_KEY", "")
	S3SecretAccessKey = Get("S3_SECRET_ACCESS_KEY", "")
	S3Region          = Get("S3_REGION", "")
	S3Endpoint        = Get("S3_ENDPOINT", "")

	MinIOBucketName      = Get("MINIO_BUCKET_NAME", "")
	MinIOAccessKeyID     = Get("MINIO_ACCESS_KEY", "")
	MinIOSecretAccessKey = Get("MINIO_SECRET_ACCESS_KEY", "")
	MinIORegion          = Get("MINIO_REGION", "")
	MinIOEndpoint        = Get("MINIO_ENDPOINT", "")
	MinIOPort            = Get("MINIO_PORT", "")

	GCSCredentialsFilePath = Get("GCS_CREDENTIAL_FILE_PATH", "")
	GCSBucketName          = Get("GCS_BUCKET_NAME", "")

	LogDriver     = Get("LOG_DRIVER", "file")
	LogFolderPath = Get("LOG_FOLDER_PATH", "log")
)
