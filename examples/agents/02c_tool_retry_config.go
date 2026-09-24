package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func callExternalAPI(_ context.Context, in queryIn) (map[string]any, error) {
	return map[string]any{"result": "Data for: " + in.Query, "source": "external_api"}, nil
}

func queryDatabase(_ context.Context, in sqlIn) (map[string]any, error) {
	return map[string]any{"rows": []map[string]any{{"id": 1, "value": in.SQL}}, "count": 1}, nil
}

func processData(_ context.Context, in dataIn) (map[string]any, error) {
	return map[string]any{"processed": in.Data, "status": "ok"}, nil
}

// ToolRetryConfig is the Python SDK's examples/agents/02c_tool_retry_config.py:
// a retry policy, count and delay per tool.
func ToolRetryConfig(model string) *ai.Agent {
	return &ai.Agent{
		Name:  "retry_config_demo",
		Model: model,
		Tools: ai.Tools(
			tool.Func("call_external_api", callExternalAPI,
				"Call an unreliable external API that may need aggressive retries.",
				tool.WithRetry(5, 1, ai.RetryExponentialBackoff)),
			tool.Func("query_database", queryDatabase,
				"Run a database query with fixed-interval retries for transient connection issues.",
				tool.WithRetry(3, 5, ai.RetryFixed)),
			tool.Func("process_data", processData,
				"Process data locally — light retries with linear backoff.",
				tool.WithRetry(2, 2, ai.RetryLinearBackoff)),
		),
		Instructions: "You help users fetch and process data. Use the appropriate tool for each request.",
	}
}

func runToolRetryConfig(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, ToolRetryConfig(Model()), "Look up the latest Python release info from the API.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
