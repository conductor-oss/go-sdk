package validation

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError represents a single validation error
type ValidationError struct {
	FieldPath string
	Message   string
	Value     interface{}
}

func (e *ValidationError) Error() string {
	if e.FieldPath == "" {
		return e.Message
	}
	return fmt.Sprintf("%s: %s", e.FieldPath, e.Message)
}

// MultiValidationError represents multiple validation errors
type MultiValidationError struct {
	Errors []*ValidationError
}

func (e *MultiValidationError) Error() string {
	if len(e.Errors) == 0 {
		return "validation failed"
	}
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	return fmt.Sprintf("validation failed with %d errors: [%s]",
		len(e.Errors), strings.Join(messages, ", "))
}

// Validator provides fluent validation API
type Validator struct {
	errors []*ValidationError
}

func NewValidator() *Validator {
	return &Validator{
		errors: []*ValidationError{},
	}
}

func (v *Validator) RequiredString(field, value string) *Validator {
	if value == "" {
		v.errors = append(v.errors, &ValidationError{
			FieldPath: field,
			Message:   "cannot be null or empty",
			Value:     value,
		})
	}
	return v
}

func (v *Validator) SliceNotEmpty(field string, slice interface{}) *Validator {
	if slice == nil {
		v.errors = append(v.errors, &ValidationError{
			FieldPath: field,
			Message:   "cannot be null or empty",
			Value:     slice,
		})
		return v
	}

	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		v.errors = append(v.errors, &ValidationError{
			FieldPath: field,
			Message:   "expected slice",
			Value:     slice,
		})
		return v
	}

	if val.Len() == 0 {
		v.errors = append(v.errors, &ValidationError{
			FieldPath: field,
			Message:   "cannot be null or empty",
			Value:     slice,
		})
	}
	return v
}

func (v *Validator) Error() error {
	if len(v.errors) == 0 {
		return nil
	}
	return &MultiValidationError{Errors: v.errors}
}
