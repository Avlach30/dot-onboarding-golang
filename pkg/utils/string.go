package utils

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/exp/rand"
)

func ToSnakeCase(s string) string {
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
		panic(exception.BussinessException("Invalid UUID format"))
	}

	return id
}
