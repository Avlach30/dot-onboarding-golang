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

func Get[T any](key string, keyPairSingleton *KeyPairSingleton) (any, error) {
	switch config.GlobalStateDriver {
	case "redis":
		value, err := redisClient.Get(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
		if err != nil {
			return nil, err
		}

		var result T
		json.Unmarshal([]byte(value), &result)

		return result, nil
	default:
		for _, keyPair := range keyPairSingleton.KeyPairs {
			if keyPair.Key == constant.RuntimeGlobalStateKey+key {
				return keyPair.Value, nil // Return the pointer to the found keyPair
			}
		}
	}

	return nil, nil
}

func (keyPairSingleton *KeyPairSingleton) Delete(key string) error {
	switch config.GlobalStateDriver {
	case "redis":
		_, err := redisClient.Del(context.Background(), constant.RedisGlobalStatePrefixKey+key).Result()
		if err == nil {
			return err
		}

		return nil
	default:
		for i, keyPair := range keyPairSingleton.KeyPairs {
			if keyPair.Key == constant.RuntimeGlobalStateKey+key {
				keyPairSingleton.KeyPairs = append(keyPairSingleton.KeyPairs[:i], keyPairSingleton.KeyPairs[i+1:]...)
				break
			}
		}

		return nil
	}
}

func (keyPairSingleton *KeyPairSingleton) Set(key string, value any) error {
	switch config.GlobalStateDriver {
	case "redis":
		jsonData, _ := json.Marshal(value)
		_, err := redisClient.Set(context.Background(), constant.RedisGlobalStatePrefixKey+key, jsonData, 0).Result()
		if err == nil {
			return err
		}

		return nil
	default:
		keyPairSingleton.Delete(constant.RuntimeGlobalStateKey + key)
		keyPairSingleton.KeyPairs = append(keyPairSingleton.KeyPairs, KeyPair{
			Key:   key,
			Value: value,
		})

		return nil
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
		result, err := Get[any](constant.RuntimeGlobalStateKey+key, keyPairSingleton)
		if err != nil {
			return false
		}

		return result != nil
	}
}
