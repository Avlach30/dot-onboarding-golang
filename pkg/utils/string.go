package utils

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/exp/rand"
)

func StringToSnakeCase(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 5) // Approximate additional space for underscores

	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result = make([]byte, length)
	rand.Seed(uint64(time.Now().UnixNano())) // Seed the random number generator

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // Pick random character from charset
	}

	return string(result)
}

func UUIDChecker(uuidString string) uuid.UUID {
	id, err := uuid.Parse(uuidString)
	if err != nil {
		panic(*exception.BussinessException("Invalid UUID format"))
	}

	return id
}

// convertPayload attempts to convert the payload string to an appropriate type
func StringToInterface(payload string) interface{} {
	var result interface{}

	// Try to unmarshal as JSON
	if err := json.Unmarshal([]byte(payload), &result); err == nil {
		return result
	}

	// Try converting to int, float, or bool
	if intVal, err := strconv.Atoi(payload); err == nil {
		return intVal
	}
	if floatVal, err := strconv.ParseFloat(payload, 64); err == nil {
		return floatVal
	}
	if boolVal, err := strconv.ParseBool(payload); err == nil {
		return boolVal
	}

	// If all else fails, return the original string
	return payload
}
