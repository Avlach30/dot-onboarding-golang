package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	localConfig "gitlab.dot.co.id/playground/boilerplates/golang-service/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Manager manages S3 operations.
type S3Manager struct {
	Client *s3.Client
	Bucket string
}

// NewS3Manager initializes a new S3Manager.
func NewS3Manager() (StorageManager, error) {

	bucket := localConfig.S3BucketName
	s3AccessKeyID := localConfig.S3AccessKeyID
	s3SecretAccessKey := localConfig.S3SecretAccessKey
	s3Region := localConfig.S3Region

	// Create a custom AWS configuration.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s3AccessKeyID),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s3SecretAccessKey,
			s3Region,
			"",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %v", err)
	}
	client := s3.NewFromConfig(cfg)

	return &S3Manager{
		Client: client,
		Bucket: bucket,
	}, nil
}

// StoreData uploads an object to the S3 bucket.
func (s3Manager *S3Manager) StoreData(key string, data []byte) error {
	_, err := s3Manager.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Manager.Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return err
	}

	return nil
}

// GeneratePresignURL generates a presigned URL for an object.
func (s3Manager *S3Manager) GeneratePresignURL(method string, key string, expiration time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s3Manager.Client)

	var req *v4.PresignedHTTPRequest
	var err error
	switch method {
	case "put":
		req, err = presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(s3Manager.Bucket),
			Key:    aws.String(key),
		}, s3.WithPresignExpires(expiration))
	case "get":
		req, err = presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(s3Manager.Bucket),
			Key:    aws.String(key),
		}, s3.WithPresignExpires(expiration))
	default:
		return "", fmt.Errorf("unknown method: %s", method)
	}

	return req.URL, err
}

// GetData retrieves an object from the S3 bucket.
func (s3Manager *S3Manager) GetData(key string) ([]byte, error) {
	output, err := s3Manager.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s3Manager.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %v", err)
	}
	defer output.Body.Close()

	data, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %v", err)
	}
	return data, nil
}

// DeleteData deletes an object from the S3 bucket.
func (s3Manager *S3Manager) DeleteData(key string) error {
	_, err := s3Manager.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s3Manager.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete data: %v", err)
	}
	return nil
}
