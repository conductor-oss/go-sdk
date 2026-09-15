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
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// A handler may return a struct, a map, or a scalar. The first two are task
// output as they are; a scalar is wrapped as {"result": value}, the shape the
// Python worker gives a non-dict return, so the model sees the value instead
// of an empty object.
func TestToolExecutorShapesOutput(t *testing.T) {
	type out struct {
		City string `json:"city"`
	}
	run := func(handler any) map[string]any {
		fn, err := toolExecutor(ToolDef{Name: "x", Handler: handler})
		if err != nil {
			t.Fatal(err)
		}
		v, err := fn(&model.Task{InputData: map[string]any{"city": "Paris"}})
		if err != nil {
			t.Fatal(err)
		}
		res, err := model.GetTaskResultFromTaskExecutionOutput(&model.Task{}, v)
		if err != nil {
			t.Fatal(err)
		}
		return res.OutputData
	}

	got := run(func(ctx context.Context, in echoIn) (string, error) { return "pong:" + in.City, nil })
	if !reflect.DeepEqual(got, map[string]any{"result": "pong:Paris"}) {
		t.Errorf("string output = %v", got)
	}
	got = run(func(ctx context.Context, in echoIn) (int, error) { return 42, nil })
	if !reflect.DeepEqual(got, map[string]any{"result": float64(42)}) {
		t.Errorf("int output = %v", got)
	}
	got = run(func(ctx context.Context, in echoIn) ([]string, error) { return []string{"a"}, nil })
	if !reflect.DeepEqual(got, map[string]any{"result": []any{"a"}}) {
		t.Errorf("slice output = %v", got)
	}
	got = run(func(ctx context.Context, in echoIn) (out, error) { return out{City: in.City}, nil })
	if !reflect.DeepEqual(got, map[string]any{"city": "Paris"}) {
		t.Errorf("struct output = %v", got)
	}
	got = run(func(ctx context.Context, in echoIn) (*out, error) { return &out{City: in.City}, nil })
	if !reflect.DeepEqual(got, map[string]any{"city": "Paris"}) {
		t.Errorf("pointer-to-struct output = %v", got)
	}
	got = run(func(ctx context.Context, in echoIn) (map[string]any, error) { return map[string]any{"k": 1}, nil })
	if !reflect.DeepEqual(got, map[string]any{"k": float64(1)}) {
		t.Errorf("map output = %v", got)
	}
}
