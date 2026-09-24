package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

func getUserProfile(_ context.Context, _ userIn) (map[string]any, error) {
	return map[string]any{"name": "Alice Johnson", "email": "alice.johnson@example.com", "ssn": "123-45-6789",
		"department": "Engineering", "role": "Senior Developer"}, nil
}

// RegexGuardrails is the Python SDK's examples/agents/21_regex_guardrails.py:
// server-side regex guardrails block emails and SSNs. It returns the HR agent
// whose answer trips them and the department agent whose answer does not.
func RegexGuardrails(model string) []*ai.Agent {
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
		Name:         "hr_assistant",
		Model:        model,
		Tools:        ai.Tools(tool.Func("get_user_profile", getUserProfile, "Retrieve a user's profile from the database.")),
		Instructions: "You are an HR assistant. When asked about employees, look up their profile and share ALL the details you find.",
		Guardrails:   []ai.Guardrail{noEmails, noSSN},
	}
	cleanAgent := &ai.Agent{
		Name:         "dept_assistant",
		Model:        model,
		Instructions: "You are an HR assistant. Answer questions about departments.",
		Guardrails:   []ai.Guardrail{noEmails, noSSN},
	}
	return []*ai.Agent{agent, cleanAgent}
}

func runRegexGuardrails(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	built := RegexGuardrails(Model())
	first, err := run(ctx, rt, built[0], "Tell me everything about user U-001.", out)
	if err != nil {
		return nil, err
	}
	second, err := run(ctx, rt, built[1], "What departments exist at the company?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{first, second}, nil
}
