package agents

import (
	"context"
	"io"
	"regexp"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func getOrderStatus(_ context.Context, in orderIn) (map[string]any, error) {
	return map[string]any{"order_id": in.OrderID, "status": "shipped", "tracking": "1Z999AA10123456784", "estimated_delivery": "2026-02-22"}, nil
}

func getCustomerInfo(_ context.Context, in customerIn) (map[string]any, error) {
	return map[string]any{"customer_id": in.CustomerID, "name": "Alice Johnson", "email": "alice@example.com",
		"card_on_file": "4532-0150-1234-5678", "membership": "gold"}, nil
}

var piiPattern = regexp.MustCompile(`\b\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}\b|\b\d{3}-\d{2}-\d{4}\b`)

func noPII(_ context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
	if piiPattern.MatchString(in.Content) {
		return ai.GuardrailResult{Passed: false, Message: "Your response contains PII (credit card or SSN). Redact all card numbers and SSNs before responding."}, nil
	}
	return ai.GuardrailResult{Passed: true}, nil
}

// Guardrails is the Python SDK's examples/agents/10_guardrails.py: a custom
// output guardrail makes the model redact PII and revise.
func Guardrails(model string) *ai.Agent {
	guard := ai.NewCustomGuardrail("no_pii", noPII)
	guard.Position = ai.PositionOutput
	guard.OnFail = ai.OnFailRetry
	return &ai.Agent{
		Name:  "support_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("get_order_status", getOrderStatus, "Look up the current status of an order."),
			tool.Func("get_customer_info", getCustomerInfo, "Retrieve customer details including payment info on file."),
		),
		Instructions: "You are a customer support assistant. Use the available tools to answer questions about orders and customers. Always include all details from the tool results in your response.",
		Guardrails:   []ai.Guardrail{guard},
	}
}

func runGuardrails(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, Guardrails(Model()),
		"I need a full summary: What's the status of order ORD-42, and what's the profile for customer CUST-7?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
