//go:build ignore

// Regex Guardrails — server-side pattern checks on the output.
//
// Run with:  go run agents/21_regex_guardrails.go
//
// Two regex guardrails block email addresses (retrying the model) and Social
// Security numbers (failing the run). They run on the server; no worker is
// involved. Scenario 1 asks for a profile that contains both; scenario 2
// asks a clean question the guardrails pass.
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

type userIn struct {
	UserID string
}

func getUserProfile(ctx context.Context, in userIn) (map[string]any, error) {
	return map[string]any{
		"name":       "Alice Johnson",
		"email":      "alice.johnson@example.com", // PII - should be blocked
		"ssn":        "123-45-6789",               // PII - should be blocked
		"department": "Engineering",
		"role":       "Senior Developer",
	}, nil
}

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	noEmails := &ai.RegexGuardrail{
		Patterns: []string{`[\w.+-]+@[\w-]+\.[\w.-]+`},
		Mode:     "block",
		Message:  "Response must not contain email addresses. Redact them.",
	}
	noEmails.Name = "no_email_addresses"
	noEmails.Position = ai.PositionOutput
	noEmails.OnFail = ai.OnFailRetry

	noSSN := &ai.RegexGuardrail{
		Patterns: []string{`\b\d{3}-\d{2}-\d{4}\b`},
		Mode:     "block",
		Message:  "Response must not contain Social Security Numbers.",
	}
	noSSN.Name = "no_ssn"
	noSSN.Position = ai.PositionOutput
	noSSN.OnFail = ai.OnFailRaise

	agent := &ai.Agent{
		Name:  "hr_assistant",
		Model: model,
		Tools: ai.Tools(tool.Func("get_user_profile", getUserProfile, "Retrieve a user's profile from the database.")),
		Instructions: "You are an HR assistant. When asked about employees, look up their " +
			"profile and share ALL the details you find.",
		Guardrails: []ai.Guardrail{noEmails, noSSN},
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()
	ctx := context.Background()

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  Scenario 1: Request PII — guardrails trigger")
	fmt.Println(strings.Repeat("=", 60))
	result, err := runtime.Run(ctx, agent, "Tell me everything about user U-001.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
	if strings.Contains(result.Output, "alice.johnson@example.com") {
		fmt.Println("[FAIL] Email leaked!")
	} else {
		fmt.Println("[OK] Email was blocked by RegexGuardrail")
	}
	if strings.Contains(result.Output, "123-45-6789") {
		fmt.Println("[FAIL] SSN leaked!")
	} else {
		fmt.Println("[OK] SSN was blocked by RegexGuardrail")
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("  Scenario 2: Non-PII question — guardrails pass")
	fmt.Println(strings.Repeat("=", 60))
	cleanAgent := &ai.Agent{
		Name:         "dept_assistant",
		Model:        model,
		Instructions: "You are an HR assistant. Answer questions about departments.",
		Guardrails:   []ai.Guardrail{noEmails, noSSN},
	}
	result2, err := runtime.Run(ctx, cleanAgent, "What departments exist at the company?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result2.PrintResult()
	if result2.Status == ai.StatusCompleted {
		fmt.Println("[OK] Clean response passed guardrails successfully")
	} else {
		fmt.Printf("[WARN] Unexpected status: %s\n", result2.Status)
	}
}
