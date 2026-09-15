//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Sequential pipeline — the Python SDK's examples/agents/06_sequential_pipeline.py
// as a test.
//
// A researcher, a writer and an editor run in order, each seeing the previous
// output. This is that flow, copied, with the print replaced by validation: the
// run completes and the pipeline's answer is the editor's recorded answer, the
// last of the three model calls. Each specialist's prompt embeds the previous
// specialist's text, so replaying all three shows the outputs were threaded
// through in the same order as in Python.
func TestExample06SequentialPipeline(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "06_sequential_pipeline")

	researcher := &ai.Agent{
		Name:  "researcher",
		Model: mockModel,
		Instructions: "You are a researcher. Given a topic, provide key facts and data points. " +
			"Be thorough but concise. Output raw research findings.",
	}
	writer := &ai.Agent{
		Name:  "writer",
		Model: mockModel,
		Instructions: "You are a writer. Take research findings and write a clear, engaging " +
			"article. Use headers and bullet points where appropriate.",
	}
	editor := &ai.Agent{
		Name:  "editor",
		Model: mockModel,
		Instructions: "You are an editor. Review the article for clarity, grammar, and tone. " +
			"Make improvements and output the final polished version.",
	}
	// researcher >> writer >> editor
	pipeline := &ai.Agent{
		Name:     "researcher_writer_editor",
		Model:    mockModel,
		Agents:   []*ai.Agent{researcher, writer, editor},
		Strategy: ai.StrategySequential,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, pipeline, "The impact of AI agents on software development in 2025")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the editor's recorded answer\n--- got ---\n%.300s\n--- recorded ---\n%.300s", got, want)
	}
}
