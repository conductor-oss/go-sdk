//go:build ignore

// Guardrails — a custom output check that makes the model revise its answer.
//
// Run with:  go run agents/10_guardrails.go
//
// A support agent has two tools, one of which returns a card number. A custom
// guardrail on the output rejects any response containing a card number or
// SSN and, with OnFailRetry, sends the model back to revise; the final answer
// must be redacted.
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
	"regexp"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type orderIn struct {
	OrderID string
}

type customerIn struct {
	CustomerID string
}

func getOrderStatus(ctx context.Context, in orderIn) (map[string]any, error) {
	return map[string]any{
		"order_id":           in.OrderID,
		"status":             "shipped",
		"tracking":           "1Z999AA10123456784",
		"estimated_delivery": "2026-02-22",
	}, nil
}

func getCustomerInfo(ctx context.Context, in customerIn) (map[string]any, error) {
	return map[string]any{
		"customer_id":  in.CustomerID,
		"name":         "Alice Johnson",
		"email":        "alice@example.com",
		"card_on_file": "4532-0150-1234-5678", // PII!
		"membership":   "gold",
	}, nil
}

var (
	ccPattern  = regexp.MustCompile(`\b\d{4}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}\b`)
	ssnPattern = regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`)
)

// noPII rejects responses that contain credit card numbers or SSNs.
func noPII(ctx context.Context, in ai.GuardrailInput) (ai.GuardrailResult, error) {
	if ccPattern.MatchString(in.Content) || ssnPattern.MatchString(in.Content) {
		return ai.GuardrailResult{
			Passed: false,
			Message: "Your response contains PII (credit card or SSN). " +
				"Redact all card numbers and SSNs before responding.",
		}, nil
	}
	return ai.GuardrailResult{Passed: true}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	guard := ai.NewCustomGuardrail("no_pii", noPII)
	guard.Position = ai.PositionOutput
	guard.OnFail = ai.OnFailRetry

	agent := &ai.Agent{
		Name:  "support_agent",
		Model: model,
		Tools: ai.Tools(
			tool.Func("get_order_status", "Look up the current status of an order.", getOrderStatus),
			tool.Func("get_customer_info", "Retrieve customer details including payment info on file.", getCustomerInfo),
		),
		Instructions: "You are a customer support assistant. Use the available tools to " +
			"answer questions about orders and customers. Always include all " +
			"details from the tool results in your response.",
		Guardrails: []ai.Guardrail{guard},
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent,
		"I need a full summary: What's the status of order ORD-42, "+
			"and what's the profile for customer CUST-7?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
	if strings.Contains(result.Output, "4532-0150-1234-5678") {
		fmt.Println("[WARN] PII leaked through the guardrail!")
	} else {
		fmt.Println("[OK] PII was redacted from the final output.")
	}
}
