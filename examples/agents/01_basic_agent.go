package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// BasicAgent is the Python SDK's examples/agents/01_basic_agent.py: define an
// agent, run it, print the answer.
func BasicAgent(model string) *ai.Agent {
	return &ai.Agent{
		Name:         "greeter",
		Model:        model,
		Instructions: "You are a friendly assistant. Keep responses brief.",
	}
}

func runBasicAgent(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, BasicAgent(Model()), "Say hello and tell me a fun fact about Python.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
