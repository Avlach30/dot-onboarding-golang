package pkg

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateStruct(
	data interface{},
) []string {
	errMsgs := []string{}

	// Use reflection to iterate over fields and check tags
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		// Check for validate:"required" tag
		if tag := field.Tag.Get("validate"); tag == "required" {
			if value.IsZero() || (value.Kind() == reflect.String && value.Len() == 0) {
				errMsgs = append(errMsgs, fmt.Sprintf("%s is required", toSnakeCase(field.Name)))
			}
		}
	}

	return errMsgs
}

func toSnakeCase(s string) string {
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
