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
	"os"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The Python SDK's e2e/test_suite7_media_tools.py as Go tests, one per Python
// test and under the same name. Image and audio tools compile with their
// provider model passed through, and the server runs the matching media task.
//
// Live only: a media task hands the model the URL of a freshly generated
// file, which is different on every run, so no recording can match. The
// tests gate on the provider key the Python suite gates on.

// assertToolCompiled is the suite's _assert_tool_compiled.
func assertToolCompiled(t *testing.T, rt *ai.Runtime, agent *ai.Agent, toolType, wantModel, step string) {
	t.Helper()
	ad := agentDef(t, planAgent(t, rt, agent))
	for _, tl := range asList(ad["tools"]) {
		m, ok := tl.(map[string]any)
		if !ok || m["toolType"] != toolType {
			continue
		}
		cfg, _ := m["config"].(map[string]any)
		if cfg["model"] != wantModel {
			t.Errorf("[%s] compiled %s tool model = %v, want %s", step, toolType, cfg["model"], wantModel)
		}
		return
	}
	t.Fatalf("[%s] no %s tool in the plan: %v", step, toolType, ad["tools"])
}

// assertMediaGenerated is the suite's _assert_media_generated: the run
// completed, the media task ran to completion, and it produced output.
func assertMediaGenerated(t *testing.T, res *ai.AgentResult, step, prefix string) {
	t.Helper()
	assertRunCompleted(t, res, step)
	wf := getWorkflow(t, res.ExecutionID)
	task := mediaTask(wf, prefix)
	if task == nil {
		var types []string
		for _, tk := range wf.Tasks {
			types = append(types, tk.TaskType)
		}
		t.Fatalf("[%s] no %s task; task types: %v", step, prefix, types)
	}
	if taskStatus(*task) != "COMPLETED" {
		t.Fatalf("[%s] %s status = %q (reason=%.300q)", step, prefix, taskStatus(*task), task.ReasonForIncompletion)
	}
	if len(task.OutputData) == 0 {
		t.Fatalf("[%s] %s task produced no output", step, prefix)
	}
}

// The Python test is marked xfail: OpenAI removed the dall-e-2 default and
// the model passthrough is a known issue in the runtime. Here a failing run
// skips with that reason, and a completing one passes.
func TestImageOpenai(t *testing.T) {
	skipInPlayback(t, "a media task hands the model a fresh file URL every run")
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set")
	}
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_image_openai", Model: model(t),
		Instructions: "Generate images when asked. Call the gen_image tool.",
		Tools:        ai.Tools(tool.Image("gen_image", "Generate an image from a text prompt.", "openai", "dall-e-3"))}
	assertToolCompiled(t, rt, agent, "generate_image", "dall-e-3", "Image/OpenAI")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Generate an image of a red circle on a white background. Use size "1024x1024".`)
	task := mediaTask(getWorkflow(t, res.ExecutionID), "GENERATE_IMAGE")
	if res.Status != ai.StatusCompleted || task == nil || taskStatus(*task) != "COMPLETED" {
		reason := res.Error
		if task != nil {
			reason = task.ReasonForIncompletion
		}
		t.Skipf("known failure, xfail in the Python suite: OpenAI removed the dall-e-2 default; model passthrough issue in the runtime (status=%s reason=%.300q)", res.Status, reason)
	}
	assertMediaGenerated(t, res, "Image/OpenAI", "GENERATE_IMAGE")
}

func TestImageGemini(t *testing.T) {
	skipInPlayback(t, "a media task hands the model a fresh file URL every run")
	if os.Getenv("GOOGLE_AI_API_KEY") == "" {
		t.Skip("GOOGLE_AI_API_KEY not set")
	}
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_image_gemini", Model: model(t),
		Instructions: "Generate images when asked. Call the gen_image_gemini tool.",
		Tools:        ai.Tools(tool.Image("gen_image_gemini", "Generate an image using Gemini Imagen.", "google_gemini", "imagen-3.0-generate-002"))}
	assertToolCompiled(t, rt, agent, "generate_image", "imagen-3.0-generate-002", "Image/Gemini")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, "Generate an image of a blue square on a white background.")
	assertMediaGenerated(t, res, "Image/Gemini", "GENERATE_IMAGE")
}

func TestAudioOpenai(t *testing.T) {
	skipInPlayback(t, "a media task hands the model a fresh file URL every run")
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set")
	}
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_audio_openai", Model: model(t),
		Instructions: "Convert text to speech when asked. Call the gen_audio tool.",
		Tools:        ai.Tools(tool.Audio("gen_audio", "Convert text to speech audio.", "openai", "tts-1"))}
	assertToolCompiled(t, rt, agent, "generate_audio", "tts-1", "Audio/OpenAI")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent, `Convert this text to speech: "Hello, this is an end to end test."`)
	assertMediaGenerated(t, res, "Audio/OpenAI", "GENERATE_AUDIO")
}
