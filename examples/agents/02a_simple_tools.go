//go:build ignore

// Simple Tool Calling — two tools, the LLM picks the right one.
//
// Run with:  go run agents/02a_simple_tools.go
//
// The agent has two tools: one for weather, one for stock prices. Based on
// the user's question, the LLM decides which tool to call. In the Conductor
// UI you'll see each tool call as a separate task with its inputs and
// outputs clearly visible.
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional; defaults to
//     openai/gpt-4o, as the Python example's settings do)
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type weatherIn struct {
	City string
}

type stockIn struct {
	Symbol string
}

func getWeather(ctx context.Context, in weatherIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
}

func getStockPrice(ctx context.Context, in stockIn) (map[string]any, error) {
	return map[string]any{"symbol": in.Symbol, "price": 182.50, "change": "+1.2%"}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	agent := &ai.Agent{
		Name:  "weather_stock_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("get_weather", "Get the current weather for a city.", getWeather),
			tool.Func("get_stock_price", "Get the current stock price for a ticker symbol.", getStockPrice),
		),
		Instructions: "You are a helpful assistant. Use tools to answer questions.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	// The LLM will call get_weather (not get_stock_price).
	result, err := runtime.Run(context.Background(), agent, "What's the weather like in San Francisco?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()

	// Production pattern:
	// 1. Deploy once during CI/CD, from a release script:
	//    runtime.Deploy(ctx, agent)
	//
	// 2. In a separate long-lived worker process:
	//    runtime.Serve(ctx, agent)
}
