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
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type opsServiceIn struct {
	ServiceName string `json:"service_name"`
}

type opsDeleteIn struct {
	ServiceName string `json:"service_name"`
	DataType    string `json:"data_type"`
}

// Human in the loop with streaming — the Python SDK's
// examples/agents/09c_hitl_streaming.py as a test.
//
// An operations agent checks a service, restarts it and deletes its data;
// only the deletion requires approval. This is that flow, copied, with the
// terminal prompt replaced by an automatic approval and the print by
// validation: the run completes with the recorded answer, the three tools
// ran once each in the instructed order, and the deletion ran only after the
// approval was given.
func TestExample09cHITLStreaming(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "09c_hitl_streaming")

	var mu sync.Mutex
	var order []string
	var deleteAt atomic.Int64
	record := func(name string) {
		mu.Lock()
		order = append(order, name)
		mu.Unlock()
	}
	checkService := func(ctx context.Context, in opsServiceIn) (map[string]any, error) {
		record("check_service")
		return map[string]any{"service": in.ServiceName, "status": "unhealthy", "uptime": "0m"}, nil
	}
	restartService := func(ctx context.Context, in opsServiceIn) (map[string]any, error) {
		record("restart_service")
		return map[string]any{"service": in.ServiceName, "status": "restarted", "new_uptime": "0m"}, nil
	}
	deleteServiceData := func(ctx context.Context, in opsDeleteIn) (map[string]any, error) {
		record("delete_service_data")
		deleteAt.Store(time.Now().UnixNano())
		return map[string]any{"service": in.ServiceName, "data_type": in.DataType, "status": "deleted"}, nil
	}

	agent := &ai.Agent{
		Name:  "ops_agent",
		Model: mockModel,
		Tools: []ai.ToolDef{
			tool.Func("check_service", "Check the health of a service.", checkService),
			tool.Func("restart_service", "Restart a service. Safe operation, no approval needed.", restartService),
			tool.Func("delete_service_data", "Delete service data. Destructive — requires human approval.",
				deleteServiceData, tool.RequiresApproval()),
		},
		Instructions: "You are an operations assistant. Work through the request one tool call at a " +
			"time, in this order:\n" +
			"1. Check the service with check_service.\n" +
			"2. If it is unhealthy, restart it with restart_service.\n" +
			"3. Last, if the user asked you to clear or delete data, call " +
			"delete_service_data.\n" +
			"A human approves the deletion, not you — delete_service_data pauses for that " +
			"approval by itself, so never ask for approval in your own reply.",
	}

	result, approvedAt := runWithApproval(t, runtime, agent,
		"The payments service is down. Check it, restart it, and clear its stale cache data.")

	// Validation, in place of the example's event printing.
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	mu.Lock()
	got := strings.Join(order, " > ")
	mu.Unlock()
	if want := "check_service > restart_service > delete_service_data"; got != want {
		t.Errorf("tools ran as %q, want %q", got, want)
	}
	if approvedAt.IsZero() {
		t.Error("the run never paused for approval")
	} else if deleteAt.Load() != 0 && deleteAt.Load() < approvedAt.UnixNano() {
		t.Error("delete_service_data ran before the approval was given")
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}
