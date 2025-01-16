package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	localConfig "gitlab.dot.co.id/playground/boilerplates/golang-service/config"
)

// MinIOManager manages MinIO S3 operations.
type MinIOManager struct {
	Client *minio.Client
	Bucket string
}

// NewMinIOManager initializes a new MinIOManager.
func NewMinIOManager() (StorageManager, error) {

	bucket := localConfig.MinIOBucketName
	s3AccessKeyID := localConfig.MinIOAccessKeyID
	s3SecretAccessKey := localConfig.MinIOSecretAccessKey
	minioEndpoint := localConfig.MinIOEndpoint
	minioPort := localConfig.MinIOPort
	minioIsUseSSL := localConfig.MinIOIsUseSSL

	// Create a custom MinIO configuration.
	minioClient, err := minio.New(minioEndpoint+":"+minioPort, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKeyID, s3SecretAccessKey, ""),
		Secure: minioIsUseSSL == "true",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load MinIO configuration: %v", err)
	}

	return &MinIOManager{
		Client: minioClient,
		Bucket: bucket,
	}, nil
}

// StoreData uploads an object to the MinIO S3 bucket.
func (minioManager *MinIOManager) StoreData(key string, data []byte) error {
	objectOptions := minio.PutObjectOptions{}
	info, err := minioManager.Client.PutObject(context.TODO(), minioManager.Bucket, key, bytes.NewReader(data), int64(len(data)), objectOptions)
	if err != nil {
		return err
	}

	log.Println("info.Expiration : ", info.Expiration)
	log.Println("info.Location : ", info.Location)
	log.Println("info.VersionID : ", info.VersionID)
	log.Println("info.ExpirationRuleID : ", info.ExpirationRuleID)
	log.Println("info.Key : ", info.Key)

	return nil
}

// GeneratePresignURL generates a presigned URL for an object.
func (minioManager *MinIOManager) GeneratePresignURL(method string, key string, expiration time.Duration) (string, error) {
	var req *url.URL
	var err error

	switch method {
	case "get":
		req, err = minioManager.Client.PresignedGetObject(context.TODO(), minioManager.Bucket, key, expiration, nil)
	case "put":
		log.Println("put")
		req, err = minioManager.Client.PresignedPutObject(context.TODO(), minioManager.Bucket, key, expiration)
	default:
		err = fmt.Errorf("unsupported method: %s", method)
	}

	return req.String(), err
}

// GetData retrieves an object from the MinIO S3 bucket.
func (minioManager *MinIOManager) GetData(key string) ([]byte, error) {
	output, err := minioManager.Client.GetObject(context.TODO(), minioManager.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %v", err)
	}
	defer output.Close()

	data, err := io.ReadAll(output)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %v", err)
	}
	return data, nil
}

// DeleteData deletes an object from the MinIO S3 bucket.
func (minioManager *MinIOManager) DeleteData(key string) error {
	err := minioManager.Client.RemoveObject(context.TODO(), minioManager.Bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete data: %v", err)
	}
	return nil
}
