package validation

import (
	"errors"
	"reflect"
	"testing"
)

type TestStruct struct {
	RequiredField   string `validate:"required"`
	OptionalField   string
	MinValueField   int `validate:"min=10"`
	MaxValueField   int `validate:"max=100"`
	MultiValidation int `validate:"required,min=5,max=50"`
}

func TestValidateObjectsByTags(t *testing.T) {
	tests := []struct {
		name          string
		obj           TestStruct
		wantErr       bool
		expectedError error
	}{
		{
			name: "valid struct",
			obj: TestStruct{
				RequiredField:   "not empty",
				MinValueField:   15,
				MaxValueField:   50,
				MultiValidation: 25,
			},
			wantErr: false,
		},
		{
			name: "missing required field",
			obj: TestStruct{
				MinValueField: 15,
				MaxValueField: 50,
			},
			wantErr:       true,
			expectedError: ErrRequiredField,
		},
		{
			name: "value below minimum",
			obj: TestStruct{
				RequiredField: "not empty",
				MinValueField: 5,
				MaxValueField: 50,
			},
			wantErr:       true,
			expectedError: ErrMinRangeValidation,
		},
		{
			name: "value above maximum",
			obj: TestStruct{
				RequiredField: "not empty",
				MinValueField: 15,
				MaxValueField: 150,
			},
			wantErr:       true,
			expectedError: ErrMaxRangeValidation,
		},
		{
			name: "multiple validation failure",
			obj: TestStruct{
				RequiredField:   "not empty",
				MultiValidation: 2,
			},
			wantErr:       true,
			expectedError: ErrMinRangeValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateObjectsByTags(tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateObjectsByTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && !errors.Is(err, tt.expectedError) {
				t.Errorf("ValidateObjectsByTags() expected error = %v, got = %v", tt.expectedError, err)
			}
		})
	}
}

func TestValidateFieldByTag(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() (reflect.Value, string)
		wantErr error
	}{
		{
			name: "required field - valid",
			setup: func() (reflect.Value, string) {
				str := "not empty"
				return reflect.ValueOf(str), "required"
			},
			wantErr: nil,
		},
		{
			name: "required field - invalid",
			setup: func() (reflect.Value, string) {
				var str string
				return reflect.ValueOf(str), "required"
			},
			wantErr: ErrRequiredField,
		},
		{
			name: "min value - valid",
			setup: func() (reflect.Value, string) {
				num := 15
				return reflect.ValueOf(num), "min=10"
			},
			wantErr: nil,
		},
		{
			name: "min value - invalid",
			setup: func() (reflect.Value, string) {
				num := 5
				return reflect.ValueOf(num), "min=10"
			},
			wantErr: ErrMinRangeValidation,
		},
		{
			name: "max value - valid",
			setup: func() (reflect.Value, string) {
				num := 50
				return reflect.ValueOf(num), "max=100"
			},
			wantErr: nil,
		},
		{
			name: "max value - invalid",
			setup: func() (reflect.Value, string) {
				num := 150
				return reflect.ValueOf(num), "max=100"
			},
			wantErr: ErrMaxRangeValidation,
		},
		{
			name: "min value - wrong type",
			setup: func() (reflect.Value, string) {
				str := "not a number"
				return reflect.ValueOf(str), "min=10"
			},
			wantErr: ErrMaxValidationForIntOnly,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, tag := tt.setup()
			err := ValidateFieldByTag(field, tag)
			if !errors.Is(err, tt.wantErr) && (err != nil) != (tt.wantErr != nil) {
				t.Errorf("ValidateFieldByTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
