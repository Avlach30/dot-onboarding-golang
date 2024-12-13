package utils

import (
	"strings"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
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

func UUIDChecker(uuidString string) uuid.UUID {
	id, err := uuid.Parse(uuidString)
	if err != nil {
		panic(exception.BussinessException("Invalid UUID format"))
	}

	return id
}
