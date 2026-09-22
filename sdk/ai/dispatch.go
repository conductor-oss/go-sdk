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

	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// toolExecutor adapts a typed tool handler to the worker signature Conductor
// expects. The handler is `any` because ToolDef cannot carry every tool's
// generic parameters, so reflection recovers them here; the shape was already
// validated when tool.Func built the ToolDef, so a mismatch is a programming
// error. inputData is JSON round-tripped into the handler's input type, through
// the same `json` tags the schema was derived from.
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
		// The model sends the wire names the schema advertised; untagged fields
		// need them mapped back to Go names before encoding/json binds them.
		input := schema.RekeyInput(task.InputData, inType)
		if len(input) > 0 {
			raw, err := json.Marshal(input)
			if err != nil {
				return nil, fmt.Errorf("tool %q: encode input: %w", t.Name, err)
			}
			if err := json.Unmarshal(raw, in.Interface()); err != nil {
				return nil, fmt.Errorf("tool %q: bind input: %w", t.Name, err)
			}
		}

		ctx := withTaskContext(context.Background(), task)

		// Guardrails run around the handler, as in the Python worker.
		if blocked, err := checkToolInput(ctx, t, task.InputData); err != nil {
			return nil, err
		} else if blocked != nil {
			return blocked, nil
		}

		out := fn.Call([]reflect.Value{reflect.ValueOf(ctx), in.Elem()})
		if err, ok := out[1].Interface().(error); ok && err != nil {
			return nil, err
		}
		result := out[0].Interface()
		if renamed, ok := snakeCaseOutput(result); ok {
			result = renamed
		}
		value, err := checkToolOutput(ctx, t, result)
		if err != nil {
			return nil, err
		}
		return toolOutput(value), nil
	}, nil
}

// snakeCaseOutput is RekeyInput's counterpart for a result: an untagged output
// field is advertised as snake_case but encoding/json would write its Go name,
// so a result whose type has such fields is round-tripped through JSON and
// renamed. A result with none is left alone, in declaration order. Renaming
// happens before the output guardrails so they judge what the model will see.
func snakeCaseOutput(v any) (any, bool) {
	if v == nil {
		return nil, false
	}
	rt := reflect.TypeOf(v)
	if !schema.NeedsRekey(rt) {
		return nil, false
	}
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, false
	}
	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, false
	}
	return schema.RekeyOutputValue(decoded, rt), true
}

// toolOutput shapes a handler's return value as task output: maps and structs
// travel as is, anything else becomes {"result": value}, as the Python worker
// does with a non-dict return. Without it the worker's JSON round-trip into a
// map drops a scalar silently and the model sees an empty result.
func toolOutput(v any) any {
	if v == nil {
		return map[string]any{"result": nil}
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return map[string]any{"result": nil}
		}
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.Struct:
		return v
	case reflect.Map:
		if rv.Type().Key().Kind() == reflect.String {
			return v
		}
	}
	return map[string]any{"result": v}
}

type taskContextKey struct{}

// withTaskContext attaches the task so task-scoped lookups can reach it.
func withTaskContext(ctx context.Context, task *model.Task) context.Context {
	return context.WithValue(ctx, taskContextKey{}, task)
}

// taskFromContext returns the task a tool is currently serving.
func taskFromContext(ctx context.Context) (*model.Task, bool) {
	t, ok := ctx.Value(taskContextKey{}).(*model.Task)
	return t, ok
}
