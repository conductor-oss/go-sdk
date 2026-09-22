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
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite6_pdf_tools.py as a Go test, under the same
// name. A generate_pdf tool compiles to a GENERATE_PDF task the server runs
// on the model's markdown, and the file it produces can be downloaded and
// read back.
//
// Live only: the tool hands the model the URL of a freshly generated file,
// which is different on every run, so no recording can match.

const sampleMarkdown = "# Conductor Agents E2E Test Report\n\n" +
	"## Overview\n\nThis document validates the PDF generation pipeline.\n\n" +
	"## Key Metrics\n\n" +
	"| Metric       | Value |\n|-------------|-------|\n| Tests Run   | 12    |\n| Passed      | 11    |\n| Skipped     | 1     |\n\n" +
	"## Features Tested\n\n- MCP tool discovery and execution\n- HTTP tool with OpenAPI spec\n- Credential lifecycle management\n- CLI command whitelisting\n\n" +
	"## Code Example\n\n```python\nfrom conductor.ai.agents import Agent, pdf_tool\n\nagent = Agent(\n    name=\"pdf_generator\",\n    tools=[pdf_tool()],\n)\n```\n\n" +
	"## Conclusion\n\nAll critical paths validated successfully.\n"

var expectedPDFPhrases = []string{
	"Conductor Agents E2E Test Report", "Overview", "Key Metrics", "Tests Run", "12",
	"Features Tested", "MCP tool discovery", "Credential lifecycle", "Code Example", "Conclusion",
}

// serverBaseURL is CONDUCTOR_SERVER_URL without its /api suffix.
func serverBaseURL() string {
	return strings.TrimSuffix(strings.TrimRight(os.Getenv("CONDUCTOR_SERVER_URL"), "/"), "/api")
}

// mediaTask is the suites' _find_media_task: the first task whose type or
// definition carries prefix.
func mediaTask(wf taskmodel.Workflow, prefix string) *taskmodel.Task {
	for i := range wf.Tasks {
		if strings.Contains(wf.Tasks[i].TaskType, prefix) || strings.Contains(wf.Tasks[i].TaskDefName, prefix) {
			return &wf.Tasks[i]
		}
	}
	return nil
}

// extractPDFURL is the suite's _extract_pdf_url over the task output, then
// its regex fallback over the answer.
func extractPDFURL(task *taskmodel.Task, answer string) string {
	for _, key := range []string{"url", "pdfUrl", "pdf_url", "fileUrl", "file_url", "result"} {
		if s, ok := task.OutputData[key].(string); ok && (strings.HasPrefix(s, "http") || strings.HasPrefix(s, "/")) {
			return s
		}
	}
	// The OSS server stores the file locally and reports it as a file:// URL
	// under result.location and media[].location.
	if r, ok := task.OutputData["result"].(map[string]any); ok {
		if s, ok := r["location"].(string); ok && s != "" {
			return s
		}
	}
	if resp, ok := task.OutputData["response"].(map[string]any); ok {
		if body, ok := resp["body"].(map[string]any); ok {
			for _, key := range []string{"url", "pdfUrl", "fileUrl", "result"} {
				if s, ok := body[key].(string); ok {
					return s
				}
			}
		}
	}
	return regexp.MustCompile(`(https?://[^\s\)"]+\.pdf[^\s\)"]*)`).FindString(answer)
}

func TestPdfGenerationAndRoundtrip(t *testing.T) {
	skipInPlayback(t, "the generated file's URL is new on every run")
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_pdf_gen", Model: model(t),
		Instructions: "You generate PDF documents from markdown. When asked, call the generate_pdf tool with the exact markdown provided. Do not modify the markdown content.",
		Tools:        ai.Tools(tool.PDF("generate_pdf", "Generate a PDF document from markdown text."))}

	// Step 0: the plan carries the tool.
	ad := agentDef(t, planAgent(t, rt, agent))
	var pdfTools int
	for _, tl := range asList(ad["tools"]) {
		if m, ok := tl.(map[string]any); ok && m["toolType"] == "generate_pdf" {
			pdfTools++
		}
	}
	if pdfTools != 1 {
		t.Fatalf("[PDF Plan] %d generate_pdf tools in the plan, want 1", pdfTools)
	}

	// Step 1: the run.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent,
		"Convert the following markdown to a PDF document. Pass it exactly as-is to the generate_pdf tool:\n\n"+sampleMarkdown)
	assertRunCompleted(t, res, "PDF Gen")

	// Step 2: the GENERATE_PDF task.
	wf := getWorkflow(t, res.ExecutionID)
	task := mediaTask(wf, "GENERATE_PDF")
	if task == nil {
		var types []string
		for _, tk := range wf.Tasks {
			types = append(types, tk.TaskType)
		}
		t.Fatalf("[PDF Task] no GENERATE_PDF task; task types: %v", types)
	}
	if taskStatus(*task) != "COMPLETED" {
		t.Fatalf("[PDF Task] status = %q (reason=%q)", taskStatus(*task), task.ReasonForIncompletion)
	}

	// Step 3: the file.
	url := extractPDFURL(task, res.Output)
	if url == "" {
		t.Skipf("Could not extract PDF URL from task output or agent response. Task: %v. Agent output: %.300s", task.OutputData, res.Output)
	}
	var content []byte
	if strings.HasPrefix(url, "file://") {
		// A local server wrote the file to its own disk; read it from there.
		path := strings.TrimPrefix(url, "file://")
		b, err := os.ReadFile(path)
		if err != nil {
			t.Skipf("the server reports the PDF at %s, which this test cannot read: %v", url, err)
		}
		content = b
	} else {
		if strings.HasPrefix(url, "/") {
			url = serverBaseURL() + url
		}
		resp, err := (&http.Client{Timeout: 30 * time.Second}).Get(url)
		if err != nil {
			t.Fatalf("[PDF Download] %v", err)
		}
		defer resp.Body.Close()
		content, _ = io.ReadAll(resp.Body)
		if resp.StatusCode != 200 {
			t.Fatalf("[PDF Download] GET %s = %d", url, resp.StatusCode)
		}
	}
	if len(content) <= 100 {
		t.Fatalf("[PDF Download] %s is only %d bytes", url, len(content))
	}

	// Round trip, with pdftotext standing in for Python's markitdown.
	if _, err := exec.LookPath("pdftotext"); err != nil {
		t.Skip("pdftotext not installed — skipping PDF round-trip validation. Install poppler.")
	}
	pdfPath := filepath.Join(t.TempDir(), "report.pdf")
	if err := os.WriteFile(pdfPath, content, 0o644); err != nil {
		t.Fatal(err)
	}
	text, err := exec.Command("pdftotext", pdfPath, "-").Output()
	if err != nil {
		t.Fatalf("pdftotext: %v", err)
	}
	extracted := strings.ToLower(string(text))
	if len(extracted) <= 50 {
		t.Fatalf("[PDF Round-trip] only %d characters of text extracted", len(extracted))
	}
	var missing []string
	for _, p := range expectedPDFPhrases {
		if !strings.Contains(extracted, strings.ToLower(p)) {
			missing = append(missing, p)
		}
	}
	if len(missing) > 2 {
		t.Errorf("[PDF Round-trip] %d phrases missing from the PDF text: %v", len(missing), missing)
	}
}
