//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Worker represents a configurable worker definition that can be registered with TaskRunner.
// This entity enables an options pattern for configuration.
type Worker struct {
	taskName string
	handler  model.ExecuteTaskFunction

	batchSize    int
	pollInterval time.Duration
	pollTimeout  time.Duration
	domain       string

	binder InputBinder
}

// Provider exposes a Worker instance for registration with a TaskRunner.
// It is implemented by types that can provide a base Worker (e.g., Worker and TypedWorker).
type Provider interface{ Worker() *Worker }

// NewWorker constructs a Worker.
func NewWorker(taskName string, f func(t *model.Task) (interface{}, error), options ...Option) *Worker {
	w := newBaseWorker(taskName, options...)
	w.handler = f
	return w
}

// newBaseWorker initializes a Worker with default configuration.
func newBaseWorker(taskName string, options ...Option) *Worker {
	w := &Worker{
		taskName:     taskName,
		batchSize:    1,
		pollInterval: 500 * time.Millisecond,
		pollTimeout:  -1 * time.Millisecond,
		domain:       "",
		binder:       JSONBinder{},
	}
	for _, opt := range options {
		if opt != nil {
			opt(w)
		}
	}
	return w
}

// Worker allows base Worker to satisfy the Provider interface used by registration.
func (w *Worker) Worker() *Worker { return w }
