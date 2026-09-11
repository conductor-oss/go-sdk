//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai_e2e

import (
	"context"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// An on_condition handoff is a Go predicate the server calls as a worker
// after each swarm turn, under the derived name "<parent>_handoff_<target>".
// The fixture proves it serializes; this asks whether the server dispatches
// to it and hands it the state the contract promises.
//
// A swarm gives every agent transfer tools, and left to itself the model
// bounces control around with tool calls and never produces text — so the
// predicate's Result stays empty. Triage is therefore told to answer in text
// and never transfer, so there is a response for the predicate to inspect.
//
// What is asserted is the wiring the SDK owns: the predicate was dispatched,
// ActiveAgent resolved to a real name (the server sends an index), and the run
// reached a terminal state. Whether the predicate saw the marker and billing
// then ran depends on the model, so those are logged rather than asserted.
func TestOnConditionHandoff(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	const marker = "ROUTE-TO-BILLING"
	var predicateCalls atomic.Int32
	var sawMarker atomic.Bool
	var lastActive atomic.Value
	lastActive.Store("")
	known := map[string]bool{"go_e2e_swarm": true, "triage": true, "billing": true}
	var unknownActive atomic.Value
	unknownActive.Store("")

	swarm := &ai.Agent{
		Name:     "go_e2e_swarm",
		Model:    model(t),
		Strategy: ai.StrategySwarm,
		// Bound the bouncing: enough turns for triage to answer and a handoff
		// to land, not enough to burn a minute if the model transfers anyway.
		MaxTurns: 6,
		Agents: []*ai.Agent{
			{
				Name:  "triage",
				Model: model(t),
				Instructions: "You triage requests. Do NOT call any transfer tool. Reply in " +
					"one plain sentence and end it with the exact phrase " + marker + ".",
			},
			{
				Name:         "billing",
				Model:        model(t),
				Instructions: "You handle billing. Reply in one short sentence. Do not transfer.",
			},
		},
		Handoffs: []ai.HandoffCondition{
			&ai.OnCondition{
				Target: "billing",
				Condition: func(_ context.Context, s ai.HandoffState) (bool, error) {
					predicateCalls.Add(1)
					lastActive.Store(s.ActiveAgent)
					if !known[s.ActiveAgent] {
						unknownActive.Store(s.ActiveAgent)
					}
					hit := strings.Contains(s.Result, marker)
					if hit {
						sawMarker.Store(true)
					}
					return hit, nil
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, swarm, "I was charged twice for my order last month.")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	active, _ := lastActive.Load().(string)
	unknown, _ := unknownActive.Load().(string)
	t.Logf("status=%s predicate calls=%d sawMarker=%v lastActive=%q output=%.100q",
		res.Status, predicateCalls.Load(), sawMarker.Load(), active, res.Output)

	if !res.Status.Terminal() {
		t.Fatalf("status = %q, want a terminal state", res.Status)
	}
	if predicateCalls.Load() == 0 {
		t.Fatal("the on_condition worker was never dispatched: the server did not route " +
			"to <parent>_handoff_<target>")
	}
	// The server sends active_agent as an index into [parent, sub-agents...];
	// the SDK must resolve it to one of those names on every call.
	if unknown != "" {
		t.Errorf("ActiveAgent resolved to %q, which is not the parent or a sub-agent: "+
			"index resolution is wrong", unknown)
	}
}
