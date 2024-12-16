package singleton

import (
	"context"
	"encoding/json"

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

// GetKeyPairs returns the singleton instance with a slice of key pairs
func GetGlobalState() *KeyPairSingleton {
	return singletonInstance
}

func Get[T any](key string, keyPairSingleton *KeyPairSingleton) any {
	switch config.GlobalStateDriver {
	case "redis":
		value, err := redisClient.Get(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
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
		redisClient.Del(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
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
		redisClient.Set(context.Background(), constant.RedisGlobalStatePrefixKey+key, jsonData, 0)
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
		result, err := redisClient.Exists(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
		if err == nil {
			return false
		}
		return result > 0
	default:
		result := Get[any](constant.RuntimeGlobalStateKey+key, keyPairSingleton)
		return result != nil
	}
}
