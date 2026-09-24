package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// SequentialPipeline is the Python SDK's examples/agents/06_sequential_pipeline.py:
// researcher, writer and editor run in order, each seeing the previous output.
// Python writes it as researcher >> writer >> editor, which names the pipeline
// after its members.
func SequentialPipeline(model string) *ai.Agent {
	researcher := &ai.Agent{
		Name:         "researcher",
		Model:        model,
		Instructions: "You are a researcher. Given a topic, provide key facts and data points. Be thorough but concise. Output raw research findings.",
	}
	writer := &ai.Agent{
		Name:         "writer",
		Model:        model,
		Instructions: "You are a writer. Take research findings and write a clear, engaging article. Use headers and bullet points where appropriate.",
	}
	editor := &ai.Agent{
		Name:         "editor",
		Model:        model,
		Instructions: "You are an editor. Review the article for clarity, grammar, and tone. Make improvements and output the final polished version.",
	}
	return &ai.Agent{
		Name:     "researcher_writer_editor",
		Model:    model,
		Agents:   []*ai.Agent{researcher, writer, editor},
		Strategy: ai.StrategySequential,
	}
}

func runSequentialPipeline(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, SequentialPipeline(Model()), "The impact of AI agents on software development in 2025", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
