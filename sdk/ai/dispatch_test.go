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
	if !reflect.DeepEqual(got, map[string]any{"result": json.Number("42")}) {
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
	if !reflect.DeepEqual(got, map[string]any{"k": json.Number("1")}) {
		t.Errorf("map output = %v", got)
	}
}

// An untagged input struct binds the snake_case names the schema advertised.
// The schema and the decoder have to agree here: if they drift, encoding/json
// reports no error and every call silently receives zero values.
func TestToolExecutorBindsUntaggedFieldsBySnakeCase(t *testing.T) {
	type untagged struct {
		AccountID string
		MaxItems  int
	}
	var got untagged
	fn, err := toolExecutor(ToolDef{Name: "x", Handler: func(ctx context.Context, in untagged) (string, error) {
		got = in
		return "ok", nil
	}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fn(&model.Task{InputData: map[string]any{"account_id": "ACC-1", "max_items": 3}}); err != nil {
		t.Fatal(err)
	}
	if got.AccountID != "ACC-1" || got.MaxItems != 3 {
		t.Errorf("bound = %+v, want AccountID=ACC-1 MaxItems=3", got)
	}
}

// An untagged output struct reaches the model under the snake_case names its
// schema advertised, not the Go field names encoding/json would write.
func TestToolExecutorEmitsUntaggedOutputInSnakeCase(t *testing.T) {
	type untaggedOut struct {
		TempF     int
		Condition string
	}
	fn, err := toolExecutor(ToolDef{Name: "x", Handler: func(ctx context.Context, in echoIn) (untaggedOut, error) {
		return untaggedOut{TempF: 72, Condition: "Sunny"}, nil
	}})
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
	want := map[string]any{"temp_f": json.Number("72"), "condition": "Sunny"}
	if !reflect.DeepEqual(res.OutputData, want) {
		t.Errorf("output = %v, want %v", res.OutputData, want)
	}
}
