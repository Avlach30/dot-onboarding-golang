package utils

import (
	"reflect"
	"strconv"
)

func FindAnyIntersect(slice1, slice2 []string) []string {
	// Create a map to store the values from slice1
	valueMap := make(map[string]bool)
	common := []string{}

	// Add all elements of slice1 to the map
	for _, value := range slice1 {
		valueMap[value] = true
	}

	// Check for common elements in slice2
	for _, value := range slice2 {
		if valueMap[value] {
			common = append(common, value)
			// Remove to avoid duplicate results if needed
			delete(valueMap, value)
		}
	}

	return common
}

func IsAnyIntersect(slice1, slice2 []string) bool {
	return len(FindAnyIntersect(slice1, slice2)) > 0
}

// Function to map struct to map[string]string
func StructToMap(input interface{}, setKeyStringToSnakeCase bool) map[string]string {
	result := make(map[string]string)

	// Use reflection to iterate over struct fields
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil // input must be a struct
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Convert field value to string
		var value string
		switch field.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(field.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = strconv.FormatUint(field.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		case reflect.Bool:
			value = strconv.FormatBool(field.Bool())
		default:
			continue // Skip unsupported types
		}

		key := fieldType.Name
		if setKeyStringToSnakeCase {
			key = StringToSnakeCase(fieldType.Name)
		}

		// Add the field value to the map using the field name as the key
		result[key] = value
	}

	return result
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {

			return true

		}
	}

	return false
}
