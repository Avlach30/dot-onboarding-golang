package singleton

import (
	"fmt"
	"os"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/storage"
)

// Singleton struct to hold a slice of RSA key pairs
type StorageSingleton struct {
	StorageManager *storage.StorageManager
}

// GetKeyPairs returns the singleton instance with a slice of key pairs
func GetGlobalStorage() *StorageSingleton {
	return storageSingleton
}

func MoveFile(sourcePath string, targetPath string) error {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	storageType := config.Storage
	switch storageType {

	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		err = storageManager.StoreData(targetPath, data)
		if err != nil {
			return err
		}
		return nil
	default:
		return StoreFileLocal(data, targetPath)
	}
}

func GetPresignURLUpload(targetPath string) (string, error) {
	storageType := config.Storage
	uploadDuration := (time.Duration(1) * time.Minute)
	switch storageType {
	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		presignedURL, err := storageManager.GeneratePresignURL("put", targetPath, uploadDuration)
		return presignedURL, err
	default:
		return "", fmt.Errorf("not support for presignURL")
	}
}

func ReadFile(filePath string) ([]byte, error) {
	storageType := config.Storage
	switch storageType {
	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		return storageManager.GetData(filePath)
	default:
		return os.ReadFile(filePath)
	}
}

func DeleteFile(filePath string) error {
	storageType := config.Storage
	switch storageType {
	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		return storageManager.DeleteData(filePath)
	default:
		return os.Remove(filePath)
	}
}

func GetPresignURLDownload(targetPath string) (string, error) {
	storageType := config.Storage
	uploadDuration := (time.Duration(1) * time.Minute)
	switch storageType {
	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		presignedURL, err := storageManager.GeneratePresignURL("get", targetPath, uploadDuration)
		return presignedURL, err
	default:
		return "", fmt.Errorf("not support for presignURL")
	}
}

func StoreFileLocal(fileSource []byte, targetPath string) error {
	err := os.WriteFile(targetPath, fileSource, 0o666)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}

func StoreFileBuff(fileSource []byte, targetPath string) error {
	storageType := config.Storage
	switch storageType {
	case "gcs", "s3", "minio":
		storageManager := *storageSingleton.StorageManager
		return storageManager.StoreData(targetPath, fileSource)
	default:
		return StoreFileLocal(fileSource, targetPath)
	}
}
