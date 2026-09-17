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
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type bankAccountIn struct {
	AccountID string `json:"account_id"`
}

type transferIn struct {
	FromAcct string  `json:"from_acct"`
	ToAcct   string  `json:"to_acct"`
	Amount   float64 `json:"amount"`
}

// Human in the loop — the Python SDK's examples/agents/09_human_in_the_loop.py
// as a test.
//
// A banking agent checks a balance and then calls a transfer tool that is
// marked as requiring approval, so the run pauses until a person answers.
// This is that flow, copied, with the terminal prompt replaced by an
// automatic approval and the print by validation: the run completes with the
// recorded answer, the balance was checked first, and the transfer ran once,
// with the right arguments, and only after the approval was given.
func TestExample09HumanInTheLoop(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "09_human_in_the_loop")
	// check_balance returns a balance of 15000.00. Python serializes that
	// float as 15000.0; Go's encoding/json writes the same float64 as 15000.
	// The recorder compares tool results as JSON nodes, so those two equal
	// numbers do not match and the second model call finds no recording.
	// Everything else in this run replays (the first call matches). Until
	// the recorder compares numbers by value, this test cannot pass in
	// playback; see the README, "What has to match".
	t.Skip("recording holds 15000.0 where Go sends 15000; the recorder does not yet compare numbers by value")

	var balanceCalls, transferCalls atomic.Int32
	var transferAt atomic.Int64
	var transfer atomic.Value
	checkBalance := func(ctx context.Context, in bankAccountIn) (map[string]any, error) {
		balanceCalls.Add(1)
		return map[string]any{"account_id": in.AccountID, "balance": 15000.00}, nil
	}
	transferFunds := func(ctx context.Context, in transferIn) (map[string]any, error) {
		transferCalls.Add(1)
		transferAt.Store(time.Now().UnixNano())
		transfer.Store(in)
		return map[string]any{"status": "completed", "from": in.FromAcct, "to": in.ToAcct, "amount": in.Amount}, nil
	}

	agent := &ai.Agent{
		Name:  "banker",
		Model: mockModel,
		Tools: []ai.ToolDef{
			tool.Func("check_balance", "Check the balance of an account.", checkBalance),
			tool.Func("transfer_funds", "Request a funds transfer; runtime pauses for human approval before execution.",
				transferFunds, tool.RequiresApproval()),
		},
		Instructions: "You are a banking assistant. Use check_balance for balance inquiries. " +
			"When asked to transfer money, first check the balance, then call " +
			"transfer_funds to request the transfer. The runtime will pause for " +
			"human approval before the transfer executes.",
	}

	result, approvedAt := runWithApproval(t, runtime, agent, "Transfer $500 from ACC-789 to ACC-456. Check the balance first.")

	// Validation, in place of the example's event printing.
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if n := balanceCalls.Load(); n != 1 {
		t.Errorf("check_balance ran %d times, want 1", n)
	}
	if n := transferCalls.Load(); n != 1 {
		t.Fatalf("transfer_funds ran %d times, want 1", n)
	}
	if in, _ := transfer.Load().(transferIn); in.FromAcct != "ACC-789" || in.ToAcct != "ACC-456" || in.Amount != 500 {
		t.Errorf("transfer_funds arguments = %+v", in)
	}
	if approvedAt.IsZero() {
		t.Error("the run never paused for approval")
	} else if transferAt.Load() < approvedAt.UnixNano() {
		t.Error("transfer_funds ran before the approval was given")
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}

// runWithApproval starts the agent, streams its events, approves the first
// time the run pauses for a human, and returns the result and the moment of
// approval (zero when the run never paused). Control flow uses status polling
// rather than the event stream, as the streaming feature test does, so a
// build that does not emit a particular event cannot hang the test.
//
// The approval answers the form exactly as the person recording the Python
// example did: approved, with "y" typed into the reason field. The server
// folds that reason into the conversation as "Human reviewer feedback:
// Reason: y.", so the next model request only matches the recording if the
// same answer is given here.
func runWithApproval(t *testing.T, runtime *ai.Runtime, agent *ai.Agent, prompt string) (*ai.AgentResult, time.Time) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	handle, err := runtime.Start(ctx, agent, prompt)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	streamCtx, stopStream := context.WithCancel(ctx)
	defer stopStream()
	events, err := handle.Events(streamCtx)
	if err != nil {
		t.Fatalf("Events: %v", err)
	}
	var seen atomic.Int32
	go func() {
		for range events {
			seen.Add(1)
		}
	}()

	var approvedAt time.Time
	for i := 0; i < 150; i++ {
		if approvedAt.IsZero() {
			if waiting, _ := handle.Waiting(ctx); waiting {
				approvedAt = time.Now()
				if err := handle.Respond(ctx, map[string]any{"approved": true, "reason": "y"}); err != nil {
					t.Fatalf("Respond: %v", err)
				}
			}
		}
		if st, err := handle.Status(ctx); err == nil && st.Status.Terminal() {
			break
		}
		time.Sleep(time.Second)
	}
	result, err := handle.Result(ctx)
	if err != nil {
		t.Fatalf("Result: %v", err)
	}
	t.Logf("%d events streamed; approved at %v", seen.Load(), approvedAt.Format(time.StampMilli))
	return result, approvedAt
}
