package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	localConfig "gitlab.dot.co.id/playground/boilerplates/golang-service/config"
)

// LocalManager handles local file system storage.
type LocalManager struct {
	BasePath string
}

// NewLocalManager initializes a new LocalManager with the given base directory.
func NewLocalManager() (StorageManager, error) {
	basePath := localConfig.LocalStoragePath 

	if basePath == "" {
		return nil, fmt.Errorf("base path for local storage not configured")
	}

	// Ensure the base directory exists.
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %v", err)
	}

	return &LocalManager{BasePath: basePath}, nil
}

// StoreData saves the data to a local file.
func (lm *LocalManager) StoreData(key string, data []byte) error {
	fullPath := filepath.Join(lm.BasePath, key)

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return os.WriteFile(fullPath, data, 0644)
}

// GetData retrieves the data from a local file.
func (lm *LocalManager) GetData(key string) ([]byte, error) {
	fullPath := filepath.Join(lm.BasePath, key)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return data, nil
}

// DeleteData deletes the file from local storage.
func (lm *LocalManager) DeleteData(key string) error {
	fullPath := filepath.Join(lm.BasePath, key)
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	return nil
}

// GeneratePresignURL simulates a presigned URL by returning the absolute path.
func (lm *LocalManager) GeneratePresignURL(method string, key string, expiration time.Duration) (string, error) {
	fullPath := filepath.Join(lm.BasePath, key)
	return fullPath, nil // Simulasi lokal path
}
