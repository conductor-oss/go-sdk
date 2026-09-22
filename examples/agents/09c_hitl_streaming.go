//go:build ignore

// Human in the Loop with streaming — several tools, one needing approval.
//
// Run with:  go run agents/09c_hitl_streaming.go
//
// An operations agent checks a service, restarts it, and clears its data.
// Only the deletion is marked as requiring approval; the run streams its
// events and pauses at that point for the person at the terminal.
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type serviceIn struct {
	ServiceName string
}

type deleteIn struct {
	ServiceName string
	DataType    string
}

func checkService(ctx context.Context, in serviceIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "status": "unhealthy", "uptime": "0m"}, nil
}

func restartService(ctx context.Context, in serviceIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "status": "restarted", "new_uptime": "0m"}, nil
}

func deleteServiceData(ctx context.Context, in deleteIn) (map[string]any, error) {
	return map[string]any{"service": in.ServiceName, "data_type": in.DataType, "status": "deleted"}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	agent := &ai.Agent{
		Name:  "ops_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("check_service", "Check the health of a service.", checkService),
			tool.Func("restart_service", "Restart a service. Safe operation, no approval needed.", restartService),
			tool.Func("delete_service_data", "Delete service data. Destructive — requires human approval.",
				deleteServiceData, tool.RequiresApproval()),
		),
		Instructions: "You are an operations assistant. Work through the request one tool call at a " +
			"time, in this order:\n" +
			"1. Check the service with check_service.\n" +
			"2. If it is unhealthy, restart it with restart_service.\n" +
			"3. Last, if the user asked you to clear or delete data, call " +
			"delete_service_data.\n" +
			"A human approves the deletion, not you — delete_service_data pauses for that " +
			"approval by itself, so never ask for approval in your own reply.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()
	ctx := context.Background()

	handle, err := runtime.Start(ctx, agent, "The payments service is down. Check it, restart it, and clear its stale cache data.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "start failed:", err)
		os.Exit(1)
	}
	fmt.Printf("Started: %s\n\n", handle.ExecutionID)

	// Follow the run as it streams. When it pauses for a human, ask on the
	// terminal and answer through the handle; the run continues from there.
	events, err := handle.Events(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "stream failed:", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	for event := range events {
		switch event.Type {
		case ai.EventThinking:
			fmt.Printf("  [thinking] %s\n", event.Text)
		case ai.EventToolCall:
			fmt.Printf("  [tool_call] %v\n", event.Data)
		case ai.EventToolResult:
			fmt.Printf("  [tool_result] %.100v\n", event.Data)
		case ai.EventWaiting:
			fmt.Println("\n--- Human input required ---")
			fmt.Print("  Approve? (y/n): ")
			answer, _ := reader.ReadString('\n')
			if a := strings.ToLower(strings.TrimSpace(answer)); a == "y" || a == "yes" {
				err = handle.Approve(ctx)
			} else {
				err = handle.Reject(ctx, "declined by operator")
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, "respond failed:", err)
				os.Exit(1)
			}
			fmt.Println()
		case ai.EventDone:
			fmt.Printf("\nDone: %s\n", event.Text)
		}
	}
}
