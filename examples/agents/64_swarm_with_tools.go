//go:build ignore

// Swarm with Tools — sub-agents have their own domain tools.
//
// Run with:  go run agents/64_swarm_with_tools.go
//
// Extends the basic swarm pattern (example 17) by giving each specialist its
// own tools. The swarm transfer mechanism works alongside the tools: the model
// can call domain tools and transfer tools in the same turn.
//
// Flow:
//  1. Front-line support triages the request
//  2. Calls transfer_to_billing_specialist or transfer_to_order_specialist
//  3. Specialist uses its domain tool (check_balance / lookup_order)
//  4. Specialist responds with the result
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type accountIn struct {
	AccountID string `json:"account_id"`
}

type orderIn struct {
	OrderID string `json:"order_id"`
}

// Domain tools.

func checkBalance(ctx context.Context, in accountIn) (map[string]any, error) {
	return map[string]any{"account_id": in.AccountID, "balance": 5432.10, "currency": "USD"}, nil
}

func lookupOrder(ctx context.Context, in orderIn) (map[string]any, error) {
	return map[string]any{"order_id": in.OrderID, "status": "shipped", "eta": "2 days"}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	// Specialist agents with tools.
	billingSpecialist := &ai.Agent{
		Name:  "billing_specialist",
		Model: model,
		Instructions: "You are a billing specialist. Use the check_balance tool to look up " +
			"account balances. Include the balance amount in your response.",
		Tools: []ai.ToolDef{tool.Func("check_balance", "Check the balance of a bank account.", checkBalance)},
	}
	orderSpecialist := &ai.Agent{
		Name:  "order_specialist",
		Model: model,
		Instructions: "You are an order specialist. Use the lookup_order tool to check " +
			"order status. Include the shipping status and ETA in your response.",
		Tools: []ai.ToolDef{tool.Func("lookup_order", "Look up the status of an order.", lookupOrder)},
	}

	// Front-line support with swarm handoffs.
	support := &ai.Agent{
		Name:  "support",
		Model: model,
		Instructions: "You are front-line customer support. Triage customer requests. " +
			"Transfer to billing_specialist for account/payment questions, " +
			"order_specialist for shipping/order questions.",
		Agents:   []*ai.Agent{billingSpecialist, orderSpecialist},
		Strategy: ai.StrategySwarm,
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "billing", Target: "billing_specialist"},
			&ai.OnTextMention{Text: "order", Target: "order_specialist"},
		},
		MaxTurns: 3,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()
	ctx := context.Background()

	// Scenario 1: billing question → billing specialist uses check_balance.
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  Scenario 1: Billing question (swarm → billing + tool)")
	fmt.Println(strings.Repeat("=", 60))
	result, err := runtime.Run(ctx, support, "What's the balance on account ACC-456?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
	if strings.Contains(result.Output, "5432") {
		fmt.Println("[OK] Billing specialist used check_balance tool")
	} else {
		fmt.Println("[WARN] Expected balance amount in output")
	}

	// Scenario 2: order question → order specialist uses lookup_order.
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  Scenario 2: Order question (swarm → order + tool)")
	fmt.Println(strings.Repeat("=", 60))
	result2, err := runtime.Run(ctx, support, "Where is my order ORD-789?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result2.PrintResult()
	if strings.Contains(strings.ToLower(result2.Output), "shipped") {
		fmt.Println("[OK] Order specialist used lookup_order tool")
	} else {
		fmt.Println("[WARN] Expected shipping status in output")
	}

	// Production pattern:
	// 1. Deploy once during CI/CD: runtime.Deploy(ctx, support)
	// 2. In a separate long-lived worker process: runtime.Serve(ctx, support)
}
