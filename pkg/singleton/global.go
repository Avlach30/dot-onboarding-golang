package singleton

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/storage"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
	"gorm.io/gorm"
)

var (
	once              sync.Once
	utilsSingleton    *UtilsSingleton
	storageSingleton  *StorageSingleton
	dbUtil            *gorm.DB
	singletonInstance *KeyPairSingleton
	redisClient       *redis.Client
)

// GetKeyPairs returns the singleton instance with a slice of key pairs
func InitGlobal(workers *task.Workers, db *gorm.DB, storageManager *storage.StorageManager) {
	once.Do(func() {

		// global worker
		utilsSingleton = &UtilsSingleton{
			Workers:           workers,
			ListRegisteredJob: &task.ListRegisteredJob{},
		}

		// global storage
		storageSingleton = &StorageSingleton{
			StorageManager: storageManager,
		}

		// ini db for global util
		dbUtil = db

		// global presistence state
		stateDriver := config.GlobalStateDriver
		switch stateDriver {
		case "redis":
			redisClient = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
				Password: config.RedisPassword,
				DB:       0,
				PoolSize: 100,
			})
		default:
			singletonInstance = &KeyPairSingleton{}
		}
	})
}
