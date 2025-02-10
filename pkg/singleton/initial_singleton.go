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
	internalCbs       *CircuitBreakerSingleton
	externalCbs       *CircuitBreakerSingleton
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

		internalCbs = newDefaultCircuitBreaker(InternalCircuitBreaker)

		externalCbs = newDefaultCircuitBreaker(ExternalCircuitBreaker)

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

func newDefaultCircuitBreaker(typeCbs string) *CircuitBreakerSingleton {

	// init circuit breaker
	var circuitBreakerOpenDuration float64
	var circuitBreakerHalfOpenDuration float64
	var circuitBreakerRequestFailedTreshold int64
	var circuitBreakerIgnoreFailureEndpoints []string
	var circuitBreakerFailureRatioTreshold float64

	if typeCbs == InternalCircuitBreaker {
		circuitBreakerIgnoreFailureEndpoints = strings.Split(config.CircuitBreakerInternalIgnoreFailureEndpoints, ",")
		circuitBreakerOpenDuration, _ = strconv.ParseFloat(config.CircuitBreakerInternalOpenDuration, 32)
		circuitBreakerHalfOpenDuration, _ = strconv.ParseFloat(config.CircuitBreakerInternalHalfOpenDuration, 32)
		circuitBreakerRequestFailedTreshold, _ = strconv.ParseInt(config.CircuitBreakerInternalRequestFailedTreshold, 10, 64)
		circuitBreakerFailureRatioTreshold, _ = strconv.ParseFloat(config.CircuitBreakerInternalFailureRatioTreshold, 64)
	} else {
		circuitBreakerIgnoreFailureEndpoints = strings.Split(config.CircuitBreakerExternalIgnoreFailureEndpoints, ",")
		circuitBreakerOpenDuration, _ = strconv.ParseFloat(config.CircuitBreakerExternalOpenDuration, 32)
		circuitBreakerHalfOpenDuration, _ = strconv.ParseFloat(config.CircuitBreakerExternalHalfOpenDuration, 32)
		circuitBreakerRequestFailedTreshold, _ = strconv.ParseInt(config.CircuitBreakerExternalRequestFailedTreshold, 10, 64)
		circuitBreakerFailureRatioTreshold, _ = strconv.ParseFloat(config.CircuitBreakerExternalFailureRatioTreshold, 64)
	}

	return &CircuitBreakerSingleton{
		CircuitBreaker: CircuitBreaker{
			IgnoreFailureEndpoints: circuitBreakerIgnoreFailureEndpoints,
			ReadyToTrip: func(currentCondition *CurrentErrorCondition) bool {

				failureRatio := float64(currentCondition.TotalFailures) / float64(currentCondition.Requests)
				isConditionOpenReached := currentCondition.TotalFailures >= circuitBreakerRequestFailedTreshold && failureRatio >= circuitBreakerFailureRatioTreshold

				openDuration := time.Duration(circuitBreakerOpenDuration * float64(time.Minute))
				maxTimeOpenDuration := currentCondition.LastOpen.Add(openDuration)
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
}
