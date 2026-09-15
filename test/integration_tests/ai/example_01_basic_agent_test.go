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
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Basic Agent — the Python SDK's examples/agents/01_basic_agent.py as a test.
//
// The example defines an agent, runs it, and prints the result. This is that
// flow, copied, with the print replaced by validation: the run must complete
// and the answer must be the one recorded when the Python example ran. The
// agent, its instructions and the prompt are the Python example's, character
// for character; the recording only matches an identical request, so a run
// that completes with the recorded answer shows the Go SDK sent the same
// request the Python SDK did.
func TestExample01BasicAgent(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "01_basic_agent")

	agent := &ai.Agent{
		Name:         "greeter",
		Model:        mockModel,
		Instructions: "You are a friendly assistant. Keep responses brief.",
	}

	prompt := "Say hello and tell me a fun fact about Python."

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, err := runtime.Run(ctx, agent, prompt)
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if result.ExecutionID == "" {
		t.Error("no execution ID returned")
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}
