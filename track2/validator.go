package track2

import (
	"fmt"
	"reflect"
	"regexp"
)

const (
	tagMinLength   = "minLength"
	tagMaxLength   = "maxLength"
	tagExactLength = "exactLength"
)

var (
	regexpNumeric = regexp.MustCompile("^[0-9]+$")

	validators = map[string]func(string, int) bool{
		tagMinLength:   func(input string, x int) bool { return len(input) >= x },
		tagMaxLength:   func(input string, x int) bool { return len(input) <= x },
		tagExactLength: func(input string, x int) bool { return len(input) == x },
	}
)

func validateNumericFields(obj any) error {
	objValue := reflect.ValueOf(obj)
	for objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return fmt.Errorf("validator: expected a struct, got %s", objValue.Kind())
	}

	objType := objValue.Type()
	for i := range objType.NumField() {
		field := objType.Field(i)

		if field.Type.Kind() != reflect.String {
			continue // Only validate string fields
		}

		value := objValue.Field(i).String()

		if !regexpNumeric.MatchString(value) {
			return fmt.Errorf("validator: field value %s must be numeric, got %s", field.Name, value)
		}

		for tag, fn := range validators {
			tagValue := field.Tag.Get(tag)
			if tagValue == "" {
				continue // Skip if no tag is set
			}

			var length int
			_, err := fmt.Sscanf(tagValue, "%d", &length)
			if err != nil {
				return fmt.Errorf("validator: invalid tag %s for field %s - %w", tag, field.Name, err)
			}

			if !fn(value, length) {
				return fmt.Errorf("validator: %s must satisfy %s condition with value %d", field.Name, tag, length)
			}
		}
	}

	return nil
}
