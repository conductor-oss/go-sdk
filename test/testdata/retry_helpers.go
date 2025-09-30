//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// RetryTimeout executes a function with retries in case of an error
// maxRetries - maximum number of retry attempts
// initialBackoff - initial delay between attempts
// operation - function to execute
func RetryTimeout(maxRetries int, initialBackoff time.Duration, operation func() error) error {
	var err error
	backoff := initialBackoff

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err = operation()
		if err == nil {
			return nil
		}

		// Check if we need to retry the attempt
		if attempt == maxRetries || !isRetryableError(err) {
			return err
		}

		time.Sleep(backoff)

		// Increase delay for next attempt (simple exponential delay)
		backoff = backoff * 2
	}

	return err
}

// RetryCondition executes a function with retries until the result meets a condition
// maxRetries - maximum number of retry attempts
// initialBackoff - initial delay between attempts
// operation - function to execute
// condition - function that checks if the result meets the condition, returns true if satisfied
func RetryCondition(maxRetries int, initialBackoff time.Duration,
	operation func() error,
	condition func() bool) error {

	var err error
	backoff := initialBackoff

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err = operation()

		// If there was an error, check if we can retry
		if err != nil {
			if attempt == maxRetries || !isRetryableError(err) {
				return err
			}
			log.Printf("Attempt %d failed with error: %v. Retrying in %v...\n", attempt+1, err, backoff)
		} else {
			// If condition is met, return the result
			if condition() {
				return nil
			}

			// If condition is not met but we still have attempts, continue
			if attempt < maxRetries {
				log.Printf("Attempt %d: condition not satisfied. Retrying in %v...\n", attempt+1, backoff)
			} else {
				return fmt.Errorf("condition not satisfied after %d attempts", maxRetries+1)
			}
		}

		time.Sleep(backoff)
		// Increase delay for next attempt
		backoff = backoff * 2
	}

	return err
}

func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded") ||
		strings.Contains(errStr, "Client.Timeout exceeded") ||
		strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "connection closed") ||
		strings.Contains(errStr, "EOF")
}
