//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// toolExecutor adapts a typed tool handler to the worker signature Conductor
// expects.
//
// The handler is stored as `any` because ToolDef cannot carry the generic
// parameters of every tool in one slice. Reflection recovers them here: the
// shape was already validated when tool.Func built the ToolDef, so a mismatch
// is a programming error rather than a runtime condition.
//
// The task's inputData is JSON round-tripped into the handler's input type,
// which keeps binding consistent with how the schema was derived — both go
// through `json` tags.
func toolExecutor(t ToolDef) (model.ExecuteTaskFunction, error) {
	fn := reflect.ValueOf(t.Handler)
	ft := fn.Type()
	if ft.Kind() != reflect.Func || ft.NumIn() != 2 || ft.NumOut() != 2 {
		return nil, fmt.Errorf("tool %q: handler must be func(context.Context, In) (Out, error)", t.Name)
	}
	if !ft.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		return nil, fmt.Errorf("tool %q: first argument must be context.Context", t.Name)
	}
	if !ft.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return nil, fmt.Errorf("tool %q: second return value must be error", t.Name)
	}
	inType := ft.In(1)

	return func(task *model.Task) (any, error) {
		in := reflect.New(inType)
		if len(task.InputData) > 0 {
			raw, err := json.Marshal(task.InputData)
			if err != nil {
				return nil, fmt.Errorf("tool %q: encode input: %w", t.Name, err)
			}
			if err := json.Unmarshal(raw, in.Interface()); err != nil {
				return nil, fmt.Errorf("tool %q: bind input: %w", t.Name, err)
			}
		}

		// The task context carries anything the server delivered with the task,
		// so a handler can read it without a second call.
		ctx := withTaskContext(context.Background(), task)

		out := fn.Call([]reflect.Value{reflect.ValueOf(ctx), in.Elem()})
		if err, ok := out[1].Interface().(error); ok && err != nil {
			return nil, err
		}
		return out[0].Interface(), nil
	}, nil
}

type taskContextKey struct{}

// withTaskContext attaches the task to the context so credential lookups and
// anything else task-scoped can reach it.
func withTaskContext(ctx context.Context, task *model.Task) context.Context {
	return context.WithValue(ctx, taskContextKey{}, task)
}

// taskFromContext returns the task a tool is currently serving.
func taskFromContext(ctx context.Context) (*model.Task, bool) {
	t, ok := ctx.Value(taskContextKey{}).(*model.Task)
	return t, ok
}
