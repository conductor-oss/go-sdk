//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"encoding/json"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/log"
)

// wrapGeneratedError wraps generated errors to maintain backward compatibility
func wrapGeneratedError(err error, resp *http.Response) error {
	if err == nil {
		return nil
	}

	// Check if it's already a GenericOpenAPIError
	if genErr, ok := err.(*orkes.GenericOpenAPIError); ok {
		// Wrap it in our GenericSwaggerError with status code
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		return NewGenericSwaggerError(genErr.Body(), genErr.Error(), genErr.Model(), statusCode)
	}

	// For other errors, just return as is
	return err
}

// GetPointerValue safely gets value from pointer
func GetPointerValue[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

// ToPointer returns pointer to value
func ToPointer[T any](value T) *T {
	return &value
}

// ToPointerIfNotEmpty returns pointer to value only if it's not empty (zero value)
// For primitive types like string, int, etc.
func ToPointerIfNotEmpty[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

// ToPointerIfNotEmptyString returns pointer to string only if it's not empty
func ToPointerIfNotEmptyString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

// ToPointerIfNotZero returns pointer to numeric value only if it's not zero
func ToPointerIfNotZero[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

func GetInputAsMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	casted, ok := input.(map[string]interface{})
	if ok {
		return casted
	}
	data, err := json.Marshal(input)
	if err != nil {
		log.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil
	}
	var parsedInput map[string]interface{}
	if err := json.Unmarshal(data, &parsedInput); err != nil {
		return nil
	}
	return parsedInput
}
