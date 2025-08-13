//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import "time"

// Option defines a functional option for configuring a Worker.
type Option func(*Worker)

// WithBatchSize sets the number of tasks to fetch per poll for the worker.
func WithBatchSize(size int) Option {
	return func(w *Worker) {
		if size > 0 {
			w.batchSize = size
		}
	}
}

// WithPollInterval sets the polling interval for the worker.
func WithPollInterval(interval time.Duration) Option {
	return func(w *Worker) {
		if interval > 0 {
			w.pollInterval = interval
		}
	}
}

// WithPollTimeout sets the polling timeout for the worker. Negative values mean server default.
func WithPollTimeout(timeout time.Duration) Option {
	return func(w *Worker) {
		w.pollTimeout = timeout
	}
}

// WithDomain sets the task domain for the worker.
func WithDomain(domain string) Option {
	return func(w *Worker) {
		w.domain = domain
	}
}
