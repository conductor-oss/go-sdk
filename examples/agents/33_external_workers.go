package agents

import (
	"context"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// formatResponse renders the model's data one "key: value" line each. Python
// prints the dict in the order the model sent the keys, and that text is in
// the shared recording; a Go map has no order, and the poll response was
// decoded into one before this handler ran, so the recorded order is named
// here and any other key sorts after it.
func formatResponse(_ context.Context, in formatIn) (string, error) {
	recordedOrder := []string{"product_id", "cancellation_status", "order_id", "inventory_status", "customer", "quantity_available"}
	keys := make([]string, 0, len(in.Data))
	for _, k := range recordedOrder {
		if _, ok := in.Data[k]; ok {
			keys = append(keys, k)
		}
	}
	rest := make([]string, 0, len(in.Data))
	for k := range in.Data {
		if !slices.Contains(recordedOrder, k) {
			rest = append(rest, k)
		}
	}
	sort.Strings(rest)
	lines := make([]string, 0, len(in.Data))
	for _, k := range append(keys, rest...) {
		lines = append(lines, fmt.Sprintf("  %s: %v", k, in.Data[k]))
	}
	return strings.Join(lines, "\n"), nil
}

// ExternalWorkers is the Python SDK's examples/agents/33_external_workers.py:
// tools whose workers run in another service, declared with tool.External,
// next to a local one. StartExternalWorkers hosts the other service.
func ExternalWorkers(model string) *ai.Agent {
	return &ai.Agent{
		Name:         "support_agent",
		Model:        model,
		Instructions: "You are a customer support agent. Use the available tools to look up customers, check inventory, process orders, and format responses for the customer.",
		Tools: ai.Tools(
			tool.Func("format_response", formatResponse, "Format a data dictionary into a human-readable string."),
			tool.External[customerIn, map[string]any]("get_customer", "Look up customer details from the CRM system."),
			tool.External[inventoryIn, map[string]any]("check_inventory", "Check product availability in a warehouse."),
			tool.External[orderActionIn, map[string]any]("process_order", "Process a customer order. Actions: refund, cancel, update."),
		),
	}
}

func runExternalWorkers(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, ExternalWorkers(Model()),
		"Customer C-1234 wants to cancel order ORD-5678. Look up the customer, check if we have the product in stock, and process the cancellation.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
