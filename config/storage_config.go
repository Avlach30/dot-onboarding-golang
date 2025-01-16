package config

var (
	// STORAGE
	Storage = Get("STORAGE", "local")

	// DRIVERS
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
	MinIOPort            = Get("MINIO_PORT", "9999")
	MinIOIsUseSSL        = Get("MINIO_IS_USE_SSL", "false")

	GCSCredentialsFilePath = Get("GCS_CREDENTIAL_FILE_PATH", "")
	GCSBucketName          = Get("GCS_BUCKET_NAME", "")
)
