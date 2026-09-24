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
	"regexp"
	"testing"
)

// A run id is minted only for an agent tree that asks for per-execution
// worker domains, and it is what the start request carries.

func TestRunIDOnlyForStatefulAgents(t *testing.T) {
	plain := &Agent{Name: "plain", Tools: []ToolDef{{Name: "t"}}}
	if id := newRunID(plain); id != "" {
		t.Errorf("a plain agent got run id %q; it must poll the shared queue", id)
	}

	statefulAgent := &Agent{Name: "a", Stateful: true}
	statefulTool := &Agent{Name: "b", Tools: []ToolDef{{Name: "t"}, {Name: "s", Stateful: true}}}
	nested := &Agent{Name: "parent", Agents: []*Agent{{Name: "child", Stateful: true}}}
	for _, agent := range []*Agent{statefulAgent, statefulTool, nested} {
		id := newRunID(agent)
		if !regexp.MustCompile(`^[0-9a-f]{32}$`).MatchString(id) {
			t.Errorf("agent %q got run id %q, want 32 hex characters as Python's uuid4().hex", agent.Name, id)
		}
	}
	// Each run gets its own, so two runs never share a domain.
	if newRunID(statefulAgent) == newRunID(statefulAgent) {
		t.Error("two runs of one agent got the same run id")
	}
}

func TestStartPayloadCarriesTheRunID(t *testing.T) {
	rt := NewRuntimeWithClient(nil, Config{})
	agent := &Agent{Name: "acme", Model: testModel, Instructions: "Go."}

	payload, err := rt.startPayload(agent, "hi", nil, "")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := payload["runId"]; ok {
		t.Errorf("a domainless run must not send a run id: %v", payload["runId"])
	}

	payload, err = rt.startPayload(agent, "hi", nil, "abc123")
	if err != nil {
		t.Fatal(err)
	}
	if payload["runId"] != "abc123" {
		t.Errorf("runId = %v, want the minted id", payload["runId"])
	}
}

// Workers for a stateful run poll that run's domain, and the same tool in a
// second run is a second poller rather than a no-op.
func TestWorkersAreStartedPerDomain(t *testing.T) {
	rt := NewRuntimeWithClient(nil, Config{})
	agent := &Agent{Name: "acme", Model: testModel, Instructions: "Go.", Stateful: true,
		Tools: []ToolDef{{Name: "ping", Handler: func(ctx context.Context, in struct{}) (string, error) { return "", nil }}}}

	for _, domain := range []string{"run_one", "run_two"} {
		if err := rt.registerWorkers(agent, map[string]string{"ping": domain}); err != nil {
			t.Fatal(err)
		}
	}
	for _, domain := range []string{"run_one", "run_two"} {
		if !rt.started[workerKey{name: "ping", domain: domain}] {
			t.Errorf("no worker polling %q for ping; started = %v", domain, rt.started)
		}
	}
	if rt.started[workerKey{name: "ping"}] {
		t.Error("a stateful run also started an undomained poller")
	}

	// A task the server did not route stays on the shared queue, which is how
	// a nested skill's own tasks are scheduled.
	if err := rt.registerWorkers(agent, map[string]string{"other": "run_three"}); err != nil {
		t.Fatal(err)
	}
	if !rt.started[workerKey{name: "ping"}] {
		t.Error("an unrouted task's worker must poll the shared queue")
	}
}
