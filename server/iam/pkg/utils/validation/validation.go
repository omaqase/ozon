package validation

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrRequiredField           = errors.New("validation error: field is required")
	ErrMaxValidationForIntOnly = errors.New("validation error: maximum validation allowed for int field")
	ErrInvalidMinLength        = errors.New("validation error: invalid minimum length")
	ErrInvalidMaxLength        = errors.New("validation error: invalid maximum length")
	ErrMinRangeValidation      = errors.New("validation error: value is below the minimum allowed range")
	ErrMaxRangeValidation      = errors.New("validation error: value exceeds the maximum allowed range")
)

func ValidateObjectsByTags(object interface{}) error {
	value := reflect.ValueOf(object)
	valueType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := valueType.Field(i)

		validationTag := fieldType.Tag.Get("validate")

		if validationTag != "" {
			validationTags := strings.Split(validationTag, ",")

			for _, tag := range validationTags {
				if err := ValidateFieldByTag(field, tag); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func ValidateFieldByTag(field reflect.Value, tag string) error {
	switch {
	case tag == "required":
		if field.IsZero() {
			return ErrRequiredField
		}

	case strings.HasPrefix(tag, "min="):
		if field.Kind() != reflect.Int {
			return ErrMaxValidationForIntOnly
		}

		minimumStr := strings.TrimPrefix(tag, "min=")
		minimumInt, err := strconv.Atoi(minimumStr)
		if err != nil {
			return ErrInvalidMinLength
		}

		fieldValue := field.Int()
		if int(fieldValue) < minimumInt {
			return ErrMinRangeValidation
		}

	case strings.HasPrefix(tag, "max="):
		if field.Kind() != reflect.Int {
			return ErrMaxValidationForIntOnly
		}

		maximumStr := strings.TrimPrefix(tag, "max=")
		maximumInt, err := strconv.Atoi(maximumStr)
		if err != nil {
			return ErrInvalidMaxLength
		}

		fieldValue := field.Int()
		if int(fieldValue) > maximumInt {
			return ErrMaxRangeValidation
		}
	}

	return nil
}
