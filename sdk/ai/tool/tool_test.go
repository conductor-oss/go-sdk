package tool_test

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type WeatherIn struct {
	City string
	Days int `json:"days,omitempty"`
}

func getWeather(ctx context.Context, in WeatherIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "tempF": 72}, nil
}

// The example from the plan must actually build, and its schema must match what
// the Python SDK emits for the equivalent @tool function.
func TestFuncMatchesPythonSchema(t *testing.T) {
	td := tool.Func("get_weather", "Get the current weather for a city.", getWeather)

	if td.Name != "get_weather" || td.ToolType != ai.ToolTypeWorker {
		t.Fatalf("unexpected tool: %+v", td)
	}
	if td.Handler == nil {
		t.Error("Handler must be retained so the runtime can register a worker")
	}

	want := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"city": map[string]any{"type": "string"},
			"days": map[string]any{"type": "integer"},
		},
		"required": []any{"city"}, // days has omitempty, mirroring Python's default
	}
	if got := roundTrip(t, td.InputSchema); !reflect.DeepEqual(got, want) {
		t.Errorf("inputSchema = %v\nwant %v", got, want)
	}

	wantOut := map[string]any{"type": "object", "additionalProperties": map[string]any{}}
	if got := roundTrip(t, td.OutputSchema); !reflect.DeepEqual(got, wantOut) {
		t.Errorf("outputSchema = %v\nwant %v", got, wantOut)
	}
}

// An external tool carries the same schema as a local one and no handler, so
// the runtime starts no worker for it and leaves the calls to whoever polls.
func TestExternalHasSchemaButNoHandler(t *testing.T) {
	td := tool.External[WeatherIn, map[string]any]("get_weather", "Get the current weather for a city.",
		tool.RequiresApproval())
	local := tool.Func("get_weather", "Get the current weather for a city.", getWeather)

	if td.Handler != nil {
		t.Error("External must not carry a handler; the worker runs elsewhere")
	}
	if td.ToolType != ai.ToolTypeWorker || !td.ApprovalRequired {
		t.Errorf("unexpected tool: %+v", td)
	}
	if got, want := roundTrip(t, td.InputSchema), roundTrip(t, local.InputSchema); !reflect.DeepEqual(got, want) {
		t.Errorf("inputSchema = %v\nwant the same as Func: %v", got, want)
	}
	if got, want := roundTrip(t, td.OutputSchema), roundTrip(t, local.OutputSchema); !reflect.DeepEqual(got, want) {
		t.Errorf("outputSchema = %v\nwant the same as Func: %v", got, want)
	}
}

func TestOptions(t *testing.T) {
	td := tool.Func("open_pr", "Open a pull request", getWeather,
		tool.WithCredentials("GH_TOKEN"),
		tool.RequiresApproval(),
		tool.WithTimeout(45),
		tool.WithMaxCalls(2),
	)
	if !reflect.DeepEqual(td.Credentials, []string{"GH_TOKEN"}) {
		t.Errorf("credentials = %v", td.Credentials)
	}
	if !td.ApprovalRequired || td.TimeoutSeconds == nil || *td.TimeoutSeconds != 45 ||
		td.MaxCalls == nil || *td.MaxCalls != 2 {
		t.Errorf("options not applied: %+v", td)
	}
}

func roundTrip(t *testing.T, v any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}
