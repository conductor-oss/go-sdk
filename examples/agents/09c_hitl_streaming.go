package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func checkService(_ context.Context, in serviceIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "status": "unhealthy", "uptime": "0m"}, nil
}

func restartService(_ context.Context, in serviceIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "status": "restarted", "new_uptime": "0m"}, nil
}

func deleteServiceData(_ context.Context, in deleteDataIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "data_type": in.DataType, "status": "deleted"}, nil
}

// HITLStreaming is the Python SDK's examples/agents/09c_hitl_streaming.py:
// several tools, one needing approval, with the run's events streamed.
func HITLStreaming(model string) *ai.Agent {
	return &ai.Agent{
		Name:  "ops_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("check_service", checkService, "Check the health of a service."),
			tool.Func("restart_service", restartService, "Restart a service. Safe operation, no approval needed."),
			tool.Func("delete_service_data", deleteServiceData,
				"Delete service data. Destructive — requires human approval.", tool.RequiresApproval()),
		),
		Instructions: "You are an operations assistant. Work through the request one tool call at a time, in this order:\n1. Check the service with check_service.\n2. If it is unhealthy, restart it with restart_service.\n3. Last, if the user asked you to clear or delete data, call delete_service_data.\nA human approves the deletion, not you — delete_service_data pauses for that approval by itself, so never ask for approval in your own reply.",
	}
}

func runHITLStreaming(ctx context.Context, rt *ai.Runtime, in io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := runWithApproval(ctx, rt, HITLStreaming(Model()),
		"The payments service is down. Check it, restart it, and clear its stale cache data.", in, out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
