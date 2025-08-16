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
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Worker is a generic interface that represents a worker.
type Worker interface {
	TaskName() string
	Options() Options
	Handler() model.ExecuteTaskFunction
	With(...Option) Worker
}

// BaseWorker represents a configurable worker implementation.
type BaseWorker struct {
	taskName string
	handler  model.ExecuteTaskFunction

	options Options

	binder InputBinder
}

// NewWorker constructs a Worker.
func NewWorker(taskName string, f func(t *model.Task) (interface{}, error), options ...Option) *BaseWorker {
	opts := applyOptions(defaultOptions(), options...)
	return &BaseWorker{
		taskName: taskName,
		options:  opts,
		binder:   JSONBinder{},
		handler:  f,
	}
}

// TaskName returns the name of the task.
func (w *BaseWorker) TaskName() string { return w.taskName }

// Options returns the options of the worker.
func (w *BaseWorker) Options() Options { return w.options }

// Handler returns the handler of the worker.
func (w *BaseWorker) Handler() model.ExecuteTaskFunction { return w.handler }

// With returns a new worker with the given options.
func (w *BaseWorker) With(options ...Option) Worker {
	return w.withOptions(applyOptions(w.options, options...))
}

func (w *BaseWorker) withOptions(o Options) Worker {
	cp := *w
	cp.options = o
	return &cp
}
