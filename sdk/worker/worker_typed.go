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
	base    *Worker
	handler func(TaskContext, TIn) (TOut, error)
	binder  InputBinder
}

// NewSimpleTypedWorker creates a typed worker entity with a simple function.
func NewSimpleTypedWorker[TIn, TOut any](
	taskName string,
	f func(context.Context, TIn) (TOut, error),
	options ...Option,
) *TypedWorker[TIn, TOut] {
	base := newBaseWorker(taskName, options...)
	adapted := func(ctx TaskContext, in TIn) (TOut, error) {
		return f(ctx, in)
	}
	return &TypedWorker[TIn, TOut]{
		base:    base,
		handler: adapted,
		binder:  JSONBinder{},
	}
}

// NewTypedWorker creates a typed worker entity with a TaskContext in the function.
func NewTypedWorker[TIn, TOut any](
	taskName string,
	f func(TaskContext, TIn) (TOut, error),
	options ...Option,
) *TypedWorker[TIn, TOut] {
	base := newBaseWorker(taskName, options...)
	return &TypedWorker[TIn, TOut]{
		base:    base,
		handler: f,
		binder:  JSONBinder{},
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

		// Execute typed handler
		return tw.handler(getWorkflowContext(context.Background(), t), in)
	}
}

// Worker converts the typed worker into a base Worker.
func (tw *TypedWorker[TIn, TOut]) Worker() *Worker {
	w := *tw.base
	w.handler = tw.adapter()
	return &w
}
