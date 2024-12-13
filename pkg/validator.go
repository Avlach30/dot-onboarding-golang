package pkg

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func ValidateStruct(
	data interface{},
) []string {
	errMsgs := []string{}

	// Use reflection to iterate over fields and check tags
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	if fields.Kind() != reflect.Struct {
		errMsgs = append(errMsgs, "must struct type")
		return errMsgs
	}

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fieldValue := values.Field(i)
		tag := field.Tag.Get("binding")

		if tag == "" {
			continue
		}

		// Split the tag values if there are multiple
		tags := strings.Split(tag, ",")
		for _, t := range tags {
			switch t {
			case "required":
				if fieldValue.String() == "" {
					errMsgs = append(errMsgs, fmt.Sprintf("%s is required", toSnakeCase(field.Name)))
				}
			case "e164":
				if fieldValue.String() != "" {
					pattern := `^\+[1-9]?[0-9]{7,14}$`
					re, _ := regexp.Compile(pattern)

					match := re.MatchString(fieldValue.String())
					if !match {
						errMsgs = append(errMsgs, fmt.Sprintf("%s is must use format e164", toSnakeCase(field.Name)))
					}
				}
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
