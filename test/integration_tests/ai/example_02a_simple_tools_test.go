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

// weatherIn is shared with agent_test.go.

type stockIn struct {
	Symbol string
}

// Simple Tool Calling — the Python SDK's examples/agents/02a_simple_tools.py
// as a test.
//
// The example gives an agent two tools, weather and stock price, asks about
// the weather, and prints the result; the LLM is expected to pick
// get_weather. This is that flow, copied, with the print replaced by
// validation: the run completes, get_weather ran exactly once for San
// Francisco, get_stock_price never ran, and the final message is the one
// recorded when the Python example ran. The tools' names, descriptions and
// parameters are the Python example's, character for character, because the
// recorded requests carry the tool schemas and only match an identical set.
func TestExample02aSimpleTools(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "02a_simple_tools")

	var weatherCalls, stockCalls atomic.Int32
	var weatherCity atomic.Value

	getWeather := func(ctx context.Context, in weatherIn) (map[string]any, error) {
		weatherCalls.Add(1)
		weatherCity.Store(in.City)
		return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
	}
	getStockPrice := func(ctx context.Context, in stockIn) (map[string]any, error) {
		stockCalls.Add(1)
		return map[string]any{"symbol": in.Symbol, "price": 182.50, "change": "+1.2%"}, nil
	}

	agent := &ai.Agent{
		Name:  "weather_stock_agent",
		Model: mockModel,
		Tools: ai.Tools(
			tool.Func("get_weather", "Get the current weather for a city.", getWeather),
			tool.Func("get_stock_price", "Get the current stock price for a ticker symbol.", getStockPrice),
		),
		Instructions: "You are a helpful assistant. Use tools to answer questions.",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	// The LLM will call get_weather (not get_stock_price).
	result, err := runtime.Run(ctx, agent, "What's the weather like in San Francisco?")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if n := weatherCalls.Load(); n != 1 {
		t.Errorf("get_weather ran %d times, want exactly 1", n)
	}
	if city, _ := weatherCity.Load().(string); city != "San Francisco" {
		t.Errorf("get_weather was asked about %q, want San Francisco", city)
	}
	if n := stockCalls.Load(); n != 0 {
		t.Errorf("get_stock_price ran %d times; the LLM should have picked get_weather only", n)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%s\n--- recorded ---\n%s", got, want)
	}
}
