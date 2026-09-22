//go:build ignore

// Human in the Loop — a tool that pauses the run for approval.
//
// Run with:  go run agents/09_human_in_the_loop.go
//
// transfer_funds is marked as requiring approval, so the runtime pauses the
// run before it executes and the person at the terminal decides. Events are
// streamed while the run progresses.
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

type accountIn struct {
	AccountID string
}

type transferIn struct {
	FromAcct string
	ToAcct   string
	Amount   float64
}

func checkBalance(ctx context.Context, in accountIn) (map[string]any, error) {
	return map[string]any{"account_id": in.AccountID, "balance": 15000.00}, nil
}

func transferFunds(ctx context.Context, in transferIn) (map[string]any, error) {
	return map[string]any{"status": "completed", "from": in.FromAcct, "to": in.ToAcct, "amount": in.Amount}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	agent := &ai.Agent{
		Name:  "banker",
		Model: model,
		Tools: ai.Tools(
			tool.Func("check_balance", "Check the balance of an account.", checkBalance),
			tool.Func("transfer_funds", "Request a funds transfer; runtime pauses for human approval before execution.",
				transferFunds, tool.RequiresApproval()),
		),
		Instructions: "You are a banking assistant. Use check_balance for balance inquiries. " +
			"When asked to transfer money, first check the balance, then call " +
			"transfer_funds to request the transfer. The runtime will pause for " +
			"human approval before the transfer executes.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()
	ctx := context.Background()

	handle, err := runtime.Start(ctx, agent, "Transfer $500 from ACC-789 to ACC-456. Check the balance first.")
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
