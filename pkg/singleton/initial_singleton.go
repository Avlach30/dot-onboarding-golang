package singleton

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

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
	cbs               *CircuitBreakerSingleton
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

		// init circuit breaker
		circuitBreakerOpenDuration, _ := strconv.ParseFloat(config.CircuitBreakerOpenDuration, 10)
		circuitBreakerHalfOpenDuration, _ := strconv.ParseFloat(config.CircuitBreakerHalfOpenDuration, 10)
		circuitBreakerRequestFailedTreshold, _ := strconv.ParseInt(config.CircuitBreakerRequestFailedTreshold, 10, 64)
		circuitBreakerFailureRatioTreshold, _ := strconv.ParseFloat(config.CircuitBreakerFailureRatioTreshold, 64)

		cbs = &CircuitBreakerSingleton{
			CircuitBreaker: CircuitBreaker{
				IgnoreFailureEndpoints: strings.Split(config.CircuitBreakerIgnoreFailureEndpoints, ","),
				ReadyToTrip: func(currentCondition *CurrentErrorCondition) bool {

					failureRatio := float64(currentCondition.TotalFailures) / float64(currentCondition.Requests)
					isConditionOpenReached := currentCondition.TotalFailures >= circuitBreakerRequestFailedTreshold && failureRatio >= circuitBreakerFailureRatioTreshold

					openDuration := time.Duration(circuitBreakerOpenDuration * float64(time.Minute))
					maxTimeOpenDuration := cbs.CurrentCondition.LastOpen.Add(openDuration)
					isOpenDurationReached := time.Now().After(maxTimeOpenDuration)

					halfOpenDuration := time.Duration(circuitBreakerHalfOpenDuration * float64(time.Minute))
					maxTimeHalfOpenDuration := maxTimeOpenDuration.Add(halfOpenDuration)
					isHalfOpenDurationReached := time.Now().After(maxTimeHalfOpenDuration)

					if !isOpenDurationReached && currentCondition.State == StateOpen {
						return false
					}

					if isOpenDurationReached && !isHalfOpenDurationReached && currentCondition.State == StateOpen && currentCondition.State != StateHalfOpen {
						currentCondition.State = StateHalfOpen
					} else if isConditionOpenReached {
						currentCondition.LastOpen = time.Now()
						currentCondition.State = StateOpen
					} else if !isOpenDurationReached {
						currentCondition.State = StateOpen
					} else if isHalfOpenDurationReached && (currentCondition.State == StateOpen || currentCondition.State == StateHalfOpen) {
						currentCondition.State = StateClose
					}

					return currentCondition.State == StateClose || currentCondition.State == StateHalfOpen
				},
			},
			CurrentCondition: CurrentErrorCondition{
				State:         "Close",
				Requests:      0,
				TotalFailures: 0,
			},
		}

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
