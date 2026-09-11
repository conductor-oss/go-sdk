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
	"fmt"
)

// GateCondition decides, after an agent inside a sequential pipeline
// finishes, whether the pipeline goes on to the next agent or stops there.
// It is one of TextGate, evaluated by the server, or GateFunc, run as a
// worker. Set it on the agent that produces the output to inspect.
type GateCondition interface {
	gateConfig(agentName string) map[string]any
	validateGate() error
}

// TextGate stops the pipeline when the agent's output contains Text.
// Matching is case-sensitive unless IgnoreCase is set, as in Python's
// TextGate(case_sensitive=True). Compiled entirely server-side.
type TextGate struct {
	Text       string
	IgnoreCase bool
}

func (g TextGate) gateConfig(string) map[string]any {
	return map[string]any{
		"type":          "text_contains",
		"text":          g.Text,
		"caseSensitive": !g.IgnoreCase,
	}
}

func (g TextGate) validateGate() error {
	if g.Text == "" {
		return fmt.Errorf("TextGate.Text is required")
	}
	return nil
}

// GateState is what a GateFunc sees: the agent's output.
type GateState struct {
	Result string
}

// GateFunc decides in Go whether the pipeline continues. Return true to
// continue and false to stop. The runtime registers it as a worker under
// "<agent>_gate"; an error means continue, the safe default the Python
// worker takes as well.
type GateFunc func(ctx context.Context, state GateState) (bool, error)

func (f GateFunc) gateConfig(agentName string) map[string]any {
	return workerRef(agentName + "_" + gateSuffix)
}

func (f GateFunc) validateGate() error {
	if f == nil {
		return fmt.Errorf("GateFunc is nil")
	}
	return nil
}

// The worker contract, from the Python GateEntry: the server sends the
// agent's result and expects {"decision": "continue" | "stop"}.
type gateIn struct {
	Result string `json:"result"`
}

type gateOut struct {
	Decision string `json:"decision"`
}

func (f GateFunc) gateHandler() func(context.Context, gateIn) (gateOut, error) {
	return func(ctx context.Context, in gateIn) (gateOut, error) {
		cont, err := f(ctx, GateState{Result: in.Result})
		if err != nil || cont {
			return gateOut{Decision: "continue"}, nil
		}
		return gateOut{Decision: "stop"}, nil
	}
}
