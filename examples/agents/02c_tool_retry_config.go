//go:build ignore

// Tool retry configuration — customizing retry behavior per tool.
//
// Run with:  go run agents/02c_tool_retry_config.go
//
// Demonstrates:
//   - retry policies: fixed, linear backoff, or exponential backoff
//   - retry count: number of retry attempts
//   - retry delay: base delay between retries
//   - mixing different retry strategies across tools
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

type queryIn struct {
	Query string `json:"query"`
}

type sqlIn struct {
	SQL string `json:"sql"`
}

type dataIn struct {
	Data string `json:"data"`
}

func callExternalAPI(ctx context.Context, in queryIn) (map[string]any, error) {
	return map[string]any{"result": "Data for: " + in.Query, "source": "external_api"}, nil
}

func queryDatabase(ctx context.Context, in sqlIn) (map[string]any, error) {
	return map[string]any{"rows": []map[string]any{{"id": 1, "value": in.SQL}}, "count": 1}, nil
}

func processData(ctx context.Context, in dataIn) (map[string]any, error) {
	return map[string]any{"processed": in.Data, "status": "ok"}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	agent := &ai.Agent{
		Name:  "retry_config_demo",
		Model: model,
		Tools: []ai.ToolDef{
			tool.Func("call_external_api", "Call an unreliable external API that may need aggressive retries.",
				callExternalAPI, tool.WithRetry(5, 1, ai.RetryExponentialBackoff)),
			tool.Func("query_database", "Run a database query with fixed-interval retries for transient connection issues.",
				queryDatabase, tool.WithRetry(3, 5, ai.RetryFixed)),
			tool.Func("process_data", "Process data locally — light retries with linear backoff.",
				processData, tool.WithRetry(2, 2, ai.RetryLinearBackoff)),
		},
		Instructions: "You help users fetch and process data. Use the appropriate tool for each request.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "Look up the latest Python release info from the API.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
