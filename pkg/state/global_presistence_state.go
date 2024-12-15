package state

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
)

// KeyPair struct to hold a single RSA key pair
type KeyPair struct {
	Key   string
	Value any
}

// Singleton struct to hold a slice of RSA key pairs
type KeyPairSingleton struct {
	KeyPairs []KeyPair
}

var (
	singletonInstance *KeyPairSingleton
	once              sync.Once
	redisGlobalState  *redis.Client
)

// GetKeyPairs returns the singleton instance with a slice of key pairs
func GetGlobalState() *KeyPairSingleton {
	once.Do(func() {
		stateDriver := config.GlobalStateDriver
		switch stateDriver {
		case "redis":
			redisGlobalState = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
				Password: config.RedisPassword,
				DB:       0,
				PoolSize: 100,
			})
		default:
			singletonInstance = &KeyPairSingleton{}
		}
	})

	return singletonInstance
}

func Get[T any](key string, keyPairSingleton *KeyPairSingleton) any {
	switch config.GlobalStateDriver {
	case "redis":
		value, err := redisGlobalState.Get(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
		if err != nil {
			return nil
		}

		var result T
		json.Unmarshal([]byte(value), &result)

		return result
	default:
		for _, person := range keyPairSingleton.KeyPairs {
			if person.Key == constant.RuntimeGlobalStateKey+key {
				return person.Value // Return the pointer to the found person
			}
		}
	}

	return nil
}

func (keyPairSingleton *KeyPairSingleton) Delete(key string) {
	switch config.GlobalStateDriver {
	case "redis":
		redisGlobalState.Del(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
	default:
		for i, person := range keyPairSingleton.KeyPairs {
			if person.Key == constant.RuntimeGlobalStateKey+key {
				keyPairSingleton.KeyPairs = append(keyPairSingleton.KeyPairs[:i], keyPairSingleton.KeyPairs[i+1:]...)
				break
			}
		}
	}
}

func (keyPairSingleton *KeyPairSingleton) Set(key string, value any) {
	switch config.GlobalStateDriver {
	case "redis":
		jsonData, _ := json.Marshal(value)
		redisGlobalState.Set(context.Background(), constant.RedisGlobalStatePrefixKey+key, jsonData, 0)
	default:
		keyPairSingleton.Delete(constant.RuntimeGlobalStateKey + key)
		keyPairSingleton.KeyPairs = append(keyPairSingleton.KeyPairs, KeyPair{
			Key:   key,
			Value: value,
		})
	}
}

func (keyPairSingleton *KeyPairSingleton) IsExist(key string) bool {
	switch config.GlobalStateDriver {
	case "redis":
		result, err := redisGlobalState.Exists(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
		if err == nil {
			return false
		}
		return result > 0
	default:
		result := Get[any](constant.RuntimeGlobalStateKey+key, keyPairSingleton)
		return result != nil
	}
}
