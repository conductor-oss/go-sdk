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
	"strings"
	"testing"
)

func routerFn() RouterFunc {
	return func(context.Context, string) (string, error) { return "billing", nil }
}

// A router takes one of two forms on the same wire key: a nested agent the
// server runs, or a reference to a worker.
func TestRouterFormsShareOneKey(t *testing.T) {
	asAgent := (&Agent{
		Name:     "r",
		Model:    testModel,
		Strategy: StrategyRouter,
		Agents:   []*Agent{billing()},
		Router:   &Agent{Name: "classifier", Model: testModel, Instructions: "Classify."},
	}).toConfig()

	nested, ok := asAgent["router"].(map[string]any)
	if !ok {
		t.Fatalf("router as agent must serialize to a nested config, got %#v", asAgent["router"])
	}
	if nested["name"] != "classifier" {
		t.Errorf("nested router name = %v, want classifier", nested["name"])
	}

	asFunc := (&Agent{
		Name:       "r",
		Model:      testModel,
		Strategy:   StrategyRouter,
		Agents:     []*Agent{billing()},
		RouterFunc: routerFn(),
	}).toConfig()

	want := map[string]any{"taskName": "r_router_fn"}
	if !reflect.DeepEqual(asFunc["router"], want) {
		t.Errorf("router as func = %v, want %v", asFunc["router"], want)
	}
}

// Python emits the router whatever the strategy is, and even with no
// sub-agents; only the strategy key itself depends on sub-agents. Verified
// against the Python serializer.
func TestRouterEmittedRegardlessOfStrategy(t *testing.T) {
	parallel := (&Agent{
		Name:       "a",
		Model:      testModel,
		Strategy:   StrategyParallel,
		Agents:     []*Agent{billing(), tech()},
		RouterFunc: routerFn(),
	}).toConfig()
	if _, ok := parallel["router"]; !ok {
		t.Error("router must be sent even under a non-router strategy")
	}
	if parallel["strategy"] != string(StrategyParallel) {
		t.Errorf("strategy = %v, want parallel", parallel["strategy"])
	}

	// No sub-agents: the router is still sent, but strategy is not, because
	// strategy rides on the presence of sub-agents.
	lone := (&Agent{Name: "a", Model: testModel, RouterFunc: routerFn()}).toConfig()
	if _, ok := lone["router"]; !ok {
		t.Error("router must be sent even with no sub-agents")
	}
	if _, ok := lone["strategy"]; ok {
		t.Errorf("strategy must be omitted with no sub-agents, got %v", lone["strategy"])
	}
}

func TestRouterValidation(t *testing.T) {
	cases := []struct {
		name    string
		agent   *Agent
		wantErr string
	}{
		{
			name: "router strategy with no router",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyRouter,
				Agents: []*Agent{billing()}},
			wantErr: "StrategyRouter requires Router or RouterFunc",
		},
		{
			// Both forms target the same wire key, so only one can win. Python
			// cannot reach this state: router is a single field there.
			name: "both router forms set",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyRouter,
				Agents:     []*Agent{billing()},
				Router:     &Agent{Name: "classifier", Model: testModel},
				RouterFunc: routerFn()},
			wantErr: "set either Router or RouterFunc, not both",
		},
		{
			name: "invalid nested router agent",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyRouter,
				Agents: []*Agent{billing()},
				Router: &Agent{Name: "2bad", Model: testModel}},
			wantErr: `agent "a" router: invalid agent name "2bad"`,
		},
		{
			name: "router func satisfies the strategy",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyRouter,
				Agents: []*Agent{billing()}, RouterFunc: routerFn()},
		},
		{
			name: "router agent satisfies the strategy",
			agent: &Agent{Name: "a", Model: testModel, Strategy: StrategyRouter,
				Agents: []*Agent{billing()},
				Router: &Agent{Name: "classifier", Model: testModel}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.agent.Validate()
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("want valid, got %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("want error containing %q, got %v", tc.wantErr, err)
			}
		})
	}
}
