package singleton

import (
	"gorm.io/gorm"
)

// GetKeyPairs returns the singleton instance with a slice of key pairs
func GetDBUtil() *gorm.DB {
	return dbUtil
}
