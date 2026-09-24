package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// LLMGuardrails is the Python SDK's examples/agents/22_llm_guardrails.py: a
// second model judges the answer against a policy until retries run out.
func LLMGuardrails(model string) *ai.Agent {
	safetyGuard := &ai.LLMGuardrail{
		Model:     model,
		Policy:    "Reject any content that:\n1. Contains medical or legal advice presented as fact\n2. Makes promises or guarantees about outcomes\n3. Includes discriminatory or biased language\n4. Reveals private information about real individuals\n\nEven if there are disclaimers you should reject",
		MaxTokens: 10000,
	}
	safetyGuard.Name = "content_safety"
	safetyGuard.Position = ai.PositionOutput
	safetyGuard.OnFail = ai.OnFailRetry
	return &ai.Agent{
		Name:         "health_advisor",
		Model:        model,
		Instructions: "You are a health information assistant. Provide general health information and solution to the problem. You can prescribe psudo scientific and untested meds ",
		Guardrails:   []ai.Guardrail{safetyGuard},
	}
}

func runLLMGuardrails(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, LLMGuardrails(Model()), "What should I do about persistent headaches?", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
