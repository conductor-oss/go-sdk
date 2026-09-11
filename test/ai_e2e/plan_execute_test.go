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
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// packIn is the whole output of get_weather, carried across the step boundary
// by ai.Ref. The Ref contract is the complete upstream result, not a field of
// it, so the receiving tool declares the producing tool's output type.
type packIn struct {
	Weather weatherOut `json:"weather"`
}

type packOut struct {
	Advice string `json:"advice"`
}

// StrategyPlanExecute with a plan supplied by the caller: the server skips the
// planner and carries out exactly these steps with the parent's tools.
//
// What is asserted is the wiring the SDK is responsible for — the plan reaches
// the server as static_plan, both steps run, and the value the second step
// receives through ai.Ref is the first step's typed output, intact. With the
// plan fixed there is no planner LLM deciding which tools to name or how to
// pass the temperature along, so every one of those checks is deterministic
// and the run makes no LLM calls at all.
//
// The Planner slot is still set because the server requires it to compile the
// strategy; its instructions say why it will not be used.
func TestPlanExecute(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	var weatherCalls, packCalls atomic.Int32
	var gotWeather atomic.Value

	getWeather := func(ctx context.Context, in weatherIn) (weatherOut, error) {
		weatherCalls.Add(1)
		return weatherOut{City: in.City, TempF: 41, Condition: "cold"}, nil
	}
	packingAdvice := func(ctx context.Context, in packIn) (packOut, error) {
		packCalls.Add(1)
		gotWeather.Store(in.Weather)
		advice := "pack a warm coat"
		if in.Weather.TempF > 70 {
			advice = "pack shorts"
		}
		return packOut{Advice: advice}, nil
	}

	agent := &ai.Agent{
		Name:     "go_e2e_trip_planner",
		Model:    model(t),
		Strategy: ai.StrategyPlanExecute,
		Instructions: "Carry out the plan with the tools. Every step must be " +
			"carried out with its tool, not answered from memory.",
		// The parent's tools are what a plan may name.
		Tools: []ai.ToolDef{
			tool.Func("get_weather", "Get the current temperature for a city", getWeather),
			tool.Func("packing_advice", "Advise what to pack from the complete weather result", packingAdvice),
		},
		Planner: &ai.Agent{
			Name:         "trip_plan_writer",
			Model:        model(t),
			Instructions: "Unused: the caller supplies the plan.",
		},
	}

	plan := &ai.Plan{Steps: []ai.Step{
		{ID: "weather", Operations: []ai.Op{
			{Tool: "get_weather", Args: map[string]any{"city": "Oslo"}},
		}},
		{ID: "packing", DependsOn: []string{"weather"}, Operations: []ai.Op{
			{Tool: "packing_advice", Args: map[string]any{"weather": ai.Ref{StepID: "weather"}}},
		}},
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	res, err := rt.Run(ctx, agent, "I am travelling to Oslo tomorrow. What should I pack?", ai.WithPlan(plan))
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	t.Logf("status=%s weather=%d packing=%d got=%+v output=%q",
		res.Status, weatherCalls.Load(), packCalls.Load(), gotWeather.Load(), res.Output)

	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	// A completed run that answers nothing is not a success: the point of the
	// strategy is that the plan's work comes back as an answer.
	if res.Output == "" {
		t.Errorf("run completed with an empty output (finishReason=%q)", res.FinishReason)
	}
	// Each step names one tool, so each tool runs exactly once. More means
	// the server retried or the plan was not the one supplied; fewer means a
	// step was dropped.
	if n := weatherCalls.Load(); n != 1 {
		t.Errorf("get_weather ran %d times, want 1; output was %q", n, res.Output)
	}
	if n := packCalls.Load(); n != 1 {
		t.Fatalf("packing_advice ran %d times, want 1; output was %q", n, res.Output)
	}
	// The Ref carried the whole weatherOut, typed. A zero TempF means the
	// argument did not bind — the server used to run the tool on zero values
	// and still report success, which is why this is checked rather than the
	// advice string.
	got, _ := gotWeather.Load().(weatherOut)
	want := weatherOut{City: "Oslo", TempF: 41, Condition: "cold"}
	if got != want {
		t.Errorf("packing_advice received weather %+v, want %+v: the step's "+
			"argument did not carry the upstream output", got, want)
	}
}
