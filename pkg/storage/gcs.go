package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	localConfig "gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"google.golang.org/api/option"
)

type GCSManager struct {
	Client *storage.Client
	Bucket string
}

func NewGCSManager() (StorageManager, error) {
	backgroundContext := context.Background()

	bucket := localConfig.GCSBucketName
	credentialsFile := localConfig.GCSCredentialsFilePath

	creds := option.WithCredentialsFile(credentialsFile)
	client, err := storage.NewClient(backgroundContext, creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}
	return &GCSManager{
		Client: client,
		Bucket: bucket,
	}, nil
}

func (gcsManager *GCSManager) StoreData(key string, data []byte) error {
	backgroundContext := context.Background()
	bucket := gcsManager.Client.Bucket(gcsManager.Bucket)
	obj := bucket.Object(key)

	writer := obj.NewWriter(backgroundContext)
	_, err := writer.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data: %v", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	return nil
}

func (gcsManager *GCSManager) GeneratePresignURL(method string, key string, expiration time.Duration) (string, error) {
	opts := &storage.SignedURLOptions{
		Method:  method,
		Expires: time.Now().Add(expiration),
	}

	url, err := storage.SignedURL(gcsManager.Bucket, key, opts)
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %v", err)
	}

	return url, nil
}

func (gcsManager *GCSManager) GetData(key string) ([]byte, error) {
	backgroundContext := context.Background()
	bucket := gcsManager.Client.Bucket(gcsManager.Bucket)
	obj := bucket.Object(key)

	reader, err := obj.NewReader(backgroundContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %v", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %v", err)
	}

	return data, nil
}

func (gcsManager *GCSManager) DeleteData(key string) error {
	backgroundContext := context.Background()
	bucket := gcsManager.Client.Bucket(gcsManager.Bucket)
	obj := bucket.Object(key)

	if err := obj.Delete(backgroundContext); err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}
