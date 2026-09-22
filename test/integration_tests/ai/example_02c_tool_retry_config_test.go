//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type retryQueryIn struct {
	Query string
}

type retrySQLIn struct {
	SQL string
}

type retryDataIn struct {
	Data string
}

// Tool retry configuration — the Python SDK's
// examples/agents/02c_tool_retry_config.py as a test.
//
// The example gives three tools different retry policies and asks the agent
// to look something up. This is that flow, copied, with the print replaced by
// validation: the run completes with the recorded answer, the API tool ran
// and the other two did not, and each tool's retry settings are what the
// server holds in its task definition — which is the example's point, and
// something the model's answer cannot show.
func TestExample02cToolRetryConfig(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "02c_tool_retry_config")

	var apiCalls, dbCalls, processCalls atomic.Int32
	callExternalAPI := func(ctx context.Context, in retryQueryIn) (map[string]any, error) {
		apiCalls.Add(1)
		return map[string]any{"result": "Data for: " + in.Query, "source": "external_api"}, nil
	}
	queryDatabase := func(ctx context.Context, in retrySQLIn) (map[string]any, error) {
		dbCalls.Add(1)
		return map[string]any{"rows": []map[string]any{{"id": 1, "value": in.SQL}}, "count": 1}, nil
	}
	processData := func(ctx context.Context, in retryDataIn) (map[string]any, error) {
		processCalls.Add(1)
		return map[string]any{"processed": in.Data, "status": "ok"}, nil
	}

	agent := &ai.Agent{
		Name:  "retry_config_demo",
		Model: mockModel,
		Tools: ai.Tools(
			tool.Func("call_external_api", "Call an unreliable external API that may need aggressive retries.",
				callExternalAPI, tool.WithRetry(5, 1, ai.RetryExponentialBackoff)),
			tool.Func("query_database", "Run a database query with fixed-interval retries for transient connection issues.",
				queryDatabase, tool.WithRetry(3, 5, ai.RetryFixed)),
			tool.Func("process_data", "Process data locally — light retries with linear backoff.",
				processData, tool.WithRetry(2, 2, ai.RetryLinearBackoff)),
		),
		Instructions: "You help users fetch and process data. Use the appropriate tool for each request.",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, agent, "Look up the latest Python release info from the API.")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if apiCalls.Load() == 0 {
		t.Error("call_external_api never ran")
	}
	if n := dbCalls.Load() + processCalls.Load(); n != 0 {
		t.Errorf("query_database/process_data ran %d times; the request was for the API", n)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}

	// The retry configuration reached the server: one task definition per tool.
	for _, want := range []struct {
		name   string
		count  int
		delay  int
		policy string
	}{
		{"call_external_api", 5, 1, "EXPONENTIAL_BACKOFF"},
		{"query_database", 3, 5, "FIXED"},
		{"process_data", 2, 2, "LINEAR_BACKOFF"},
	} {
		resp, err := http.Get(strings.TrimRight(os.Getenv("CONDUCTOR_SERVER_URL"), "/") + "/metadata/taskdefs/" + want.name)
		if err != nil {
			t.Fatalf("read task definition %s: %v", want.name, err)
		}
		var def struct {
			RetryCount        int    `json:"retryCount"`
			RetryDelaySeconds int    `json:"retryDelaySeconds"`
			RetryLogic        string `json:"retryLogic"`
		}
		derr := json.NewDecoder(resp.Body).Decode(&def)
		resp.Body.Close()
		if derr != nil {
			t.Fatalf("decode task definition %s: %v", want.name, derr)
		}
		if def.RetryCount != want.count || def.RetryDelaySeconds != want.delay || def.RetryLogic != want.policy {
			t.Errorf("%s: server holds retry %d x %ds %s, want %d x %ds %s",
				want.name, def.RetryCount, def.RetryDelaySeconds, def.RetryLogic, want.count, want.delay, want.policy)
		}
	}
}
