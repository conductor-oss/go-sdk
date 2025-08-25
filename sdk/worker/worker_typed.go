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
	"context"
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// TypedWorker is a compositional typed worker that embeds base configuration via Worker
// and provides a type-safe function with a WorkflowContext.
type TypedWorker[TIn, TOut any] struct {
	taskName string
	handler  func(TaskContext, TIn) (TOut, error)

	options Options

	binder InputBinder
}

// NewSimpleTypedWorker creates a typed worker entity with a simple function.
func NewSimpleTypedWorker[TIn, TOut any](
	taskName string,
	f func(context.Context, TIn) (TOut, error),
	options ...Option,
) *TypedWorker[TIn, TOut] {
	opts := applyOptions(defaultOptions(), options...)
	adapted := func(ctx TaskContext, in TIn) (TOut, error) {
		return f(context.Context(ctx), in)
	}
	return &TypedWorker[TIn, TOut]{
		taskName: taskName,
		handler:  adapted,
		options:  opts,
		binder:   JSONBinder{},
	}
}

// NewTypedWorker creates a typed worker entity with a TaskContext in the function.
func NewTypedWorker[TIn, TOut any](
	taskName string,
	f func(TaskContext, TIn) (TOut, error),
	options ...Option,
) *TypedWorker[TIn, TOut] {
	opts := applyOptions(defaultOptions(), options...)
	return &TypedWorker[TIn, TOut]{
		taskName: taskName,
		handler:  f,
		options:  opts,
		binder:   JSONBinder{},
	}
}

// adapter returns a legacy ExecuteTaskFunction that invokes the typed handler.
func (tw *TypedWorker[TIn, TOut]) adapter() model.ExecuteTaskFunction {
	return func(t *model.Task) (interface{}, error) {
		// Bind input
		var in TIn
		if err := tw.binder.Bind(&in, t.InputData); err != nil {
			return nil, fmt.Errorf("input binding error for task %s: %w", t.TaskDefName, err)
		}

		// Create a new context with cancellation  for proper lifecycle management
		parentCtx := tw.options.BaseContext
		if parentCtx == nil {
			parentCtx = context.Background()
		}

		ctx, cancel := context.WithCancel(parentCtx)
		defer cancel()

		// Execute typed handler
		return tw.handler(getWorkflowContext(ctx, t), in)
	}
}

// TaskName returns the name of the task.
func (tw *TypedWorker[TIn, TOut]) TaskName() string { return tw.taskName }

// Options returns the options of the worker.
func (tw *TypedWorker[TIn, TOut]) Options() Options { return tw.options }

// Handler returns the handler of the worker.
func (tw *TypedWorker[TIn, TOut]) Handler() model.ExecuteTaskFunction { return tw.adapter() }

// With returns a new worker with the given options.
func (tw *TypedWorker[TIn, TOut]) With(options ...Option) Worker {
	return tw.withOptions(applyOptions(tw.options, options...))
}

func (tw *TypedWorker[TIn, TOut]) withOptions(o Options) Worker {
	cp := *tw
	cp.options = o
	return &cp
}
