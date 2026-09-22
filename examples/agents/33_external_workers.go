//go:build ignore

// External Worker Tools — reference workers running in other services.
//
// Run with:  go run agents/33_external_workers.go
//
// Demonstrates tool.External for referencing Conductor workers that exist in
// another repository, service, or language. The type parameters provide the
// schema and the arguments the description, but no local worker is started:
// Conductor dispatches the task to whatever worker is polling for that task
// definition name.
//
// This is useful when:
//   - Workers are written in Java, Python, or another language
//   - Workers run in a separate microservice
//   - You want to reuse existing Conductor task definitions without duplicating code
//
// Requirements:
//   - Conductor server with LLM support
//   - The referenced workers must be running somewhere
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type orderActionIn struct {
	OrderID string
	Action  string
}

type deleteAccountIn struct {
	UserID string
	Reason string
}

type formatIn struct {
	Data map[string]any
}

type customerIn struct {
	CustomerID string
}

type inventoryIn struct {
	ProductID string
	Warehouse string `json:"warehouse,omitempty"`
}

// Example 1: basic external worker reference. The type parameters define the
// schema; no implementation is needed. Conductor dispatches "process_order"
// tasks to whatever worker is polling.
var processOrder = tool.External[orderActionIn, map[string]any]("process_order",
	"Process a customer order. Actions: refund, cancel, update.")

// Example 2: external worker with an approval gate. Dangerous operations can
// require human approval before execution. Declared here to show the option;
// the agent below does not use it.
var deleteAccount = tool.External[deleteAccountIn, map[string]any]("delete_account",
	"Permanently delete a user account. Requires manager approval.",
	tool.RequiresApproval())

// Example 3: mix local and external tools. Local tool.Func tools and external
// references work side by side.
func formatResponse(ctx context.Context, in formatIn) (string, error) {
	keys := make([]string, 0, len(in.Data))
	for k := range in.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	lines := make([]string, 0, len(keys))
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("  %s: %v", k, in.Data[k]))
	}
	return strings.Join(lines, "\n"), nil
}

var getCustomer = tool.External[customerIn, map[string]any]("get_customer",
	"Look up customer details from the CRM system.")

var checkInventory = tool.External[inventoryIn, map[string]any]("check_inventory",
	"Check product availability in a warehouse.")

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	// Agent: combines local and external tools.
	supportAgent := &ai.Agent{
		Name:  "support_agent",
		Model: model,
		Instructions: "You are a customer support agent. Use the available tools to " +
			"look up customers, check inventory, process orders, and format " +
			"responses for the customer.",
		Tools: ai.Tools(
			tool.Func("format_response", formatResponse, "Format a data dictionary into a human-readable string."), // local — runs in this process
			getCustomer,    // external — runs in the CRM service
			checkInventory, // external — runs in the inventory service
			processOrder,   // external — runs in the order service
		),
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	fmt.Println("=== External Worker Tools ===")
	fmt.Println("Agent has 1 local tool + 3 external worker references.")
	fmt.Println()

	result, err := runtime.Run(context.Background(), supportAgent,
		"Customer C-1234 wants to cancel order ORD-5678. "+
			"Look up the customer, check if we have the product in stock, "+
			"and process the cancellation.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()

	// Production pattern:
	// 1. Deploy once during CI/CD: runtime.Deploy(ctx, supportAgent)
	// 2. In a separate long-lived worker process: runtime.Serve(ctx, supportAgent)
}
