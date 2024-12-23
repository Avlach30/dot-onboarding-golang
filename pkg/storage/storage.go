package storage

import "time"

type StorageManager interface {
	StoreData(key string, data []byte) error
	GeneratePresignURL(method string, key string, expiration time.Duration) (string, error)
	GetData(key string) ([]byte, error)
	DeleteData(key string) error
}
