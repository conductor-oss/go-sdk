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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The Python SDK's e2e/test_suite5_http_tools.py as Go tests, one per Python
// test and under the same name. Server-side HTTP tools against the testkit's
// REST API, open and then behind a bearer token from the store; and an API
// tool built from a public OpenAPI document.

const (
	httpLifecyclePort = 3003
	httpLifecycleBase = "http://localhost:3003"
	httpLifecycleSpec = httpLifecycleBase + "/api-docs"
	httpAuthKeyName   = "HTTP_AUTH_KEY"
	httpAuthKeyValue  = "e2e-http-test-secret-key-67890"
	orkesSpecURL      = "https://developer.orkescloud.com/api-docs"

	httpAgentInstructions = "You have access to HTTP API tools. Call exactly the tools specified in each prompt.\n" +
		"Report each tool's result verbatim. Do not skip any tool.\n"
)

// s5HTTPTools is the suite's _make_http_tools: three of the testkit's
// endpoints as HTTP tools, with the given headers and credentials.
func s5HTTPTools(base string, headers map[string]string, credentials ...string) []ai.ToolDef {
	opts := func(method string, schema map[string]any) []tool.Option {
		o := []tool.Option{tool.WithMethod(method), tool.WithInputSchema(schema)}
		if headers != nil {
			o = append(o, tool.WithHeaders(headers))
		}
		if len(credentials) > 0 {
			o = append(o, tool.WithCredentials(credentials...))
		}
		return o
	}
	str := func(desc string) map[string]any { return map[string]any{"type": "string", "description": desc} }
	num := func(desc string) map[string]any { return map[string]any{"type": "number", "description": desc} }
	return ai.Tools(
		tool.HTTP("math_add", "Add two numbers (a + b)", base+"/api/math/add", opts("GET", map[string]any{
			"type": "object", "properties": map[string]any{"a": num("First number"), "b": num("Second number")},
			"required": []string{"a", "b"}})...),
		tool.HTTP("string_reverse", "Reverse a string", base+"/api/string/reverse", opts("POST", map[string]any{
			"type": "object", "properties": map[string]any{"text": str("Text to reverse")}, "required": []string{"text"}})...),
		tool.HTTP("encoding_base64_encode", "Base64-encode a string", base+"/api/encoding/base64-encode", opts("POST", map[string]any{
			"type": "object", "properties": map[string]any{"text": str("Text to encode")}, "required": []string{"text"}})...),
	)
}

func TestHttpLifecycle(t *testing.T) {
	requireMCPTestkitBinary(t)
	rt := newRuntime(t)
	store := newSecretStore(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Phase 1: the open REST API.
	stop := startMCPTestkit(t, httpLifecyclePort, "")
	discovered, err := openAPIOperationIDs(httpLifecycleSpec, "")
	if err != nil {
		t.Fatalf("[Phase 1: Discovery] %v", err)
	}
	assertSameToolSet(t, "Phase 1: Discovery", discovered)

	agent := &ai.Agent{Name: "e2e_http_unauth", Model: model(t), Instructions: httpAgentInstructions,
		Tools: s5HTTPTools(httpLifecycleBase, nil)}
	res := runTolerant(t, rt, ctx, agent, promptUseThreeTools)
	validateToolExecution(t, res, "Phase 1: Unauthenticated execution",
		findHTTPToolTasks(getWorkflow(t, res.ExecutionID), testkitToolNames...))

	// Phase 2: the same API behind the bearer token the store holds.
	authKey, created := ensureCredential(t, store, httpAuthKeyName, httpAuthKeyValue)
	if created {
		t.Cleanup(func() { store.delete(httpAuthKeyName) })
	}
	stop()
	time.Sleep(time.Second)
	startMCPTestkit(t, httpLifecyclePort, authKey)

	if code := httpStatus(httpLifecycleSpec); code != 401 && code != 403 {
		t.Fatalf("[Phase 2: Auth check] GET %s without the token = %d, want 401 or 403", httpLifecycleSpec, code)
	}
	discoveredAuth, err := openAPIOperationIDs(httpLifecycleSpec, authKey)
	if err != nil {
		t.Fatalf("[Phase 2: Auth Discovery] %v", err)
	}
	assertSameToolSet(t, "Phase 2: Auth Discovery", discoveredAuth)

	authAgent := &ai.Agent{Name: "e2e_http_auth", Model: model(t), Instructions: httpAgentInstructions,
		Tools: s5HTTPTools(httpLifecycleBase, map[string]string{"Authorization": "Bearer ${" + httpAuthKeyName + "}"}, httpAuthKeyName)}
	resAuth := runTolerant(t, rt, ctx, authAgent, promptUseThreeTools)
	validateToolExecution(t, resAuth, "Phase 2: Authenticated execution",
		findHTTPToolTasks(getWorkflow(t, resAuth.ExecutionID), testkitToolNames...))
}

// An API tool built from Orkes' public OpenAPI document compiles with its
// config intact and exposes the whitelisted startWorkflow operation. The run
// is allowed to fail, since the agent has no Orkes credentials; when it
// completes, the answer names the operation. In playback the whole spec is
// part of the request, so a change to the public document shows up as a
// failed run, which this test tolerates the way the Python one does.
func TestExternalOpenapiSpec(t *testing.T) {
	rt := newRuntime(t)
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Get(orkesSpecURL)
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
		t.Skipf("Orkes API spec not reachable at %s: %v", orkesSpecURL, err)
	}
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string `json:"operationId"`
		} `json:"paths"`
	}
	err = json.NewDecoder(resp.Body).Decode(&spec)
	resp.Body.Close()
	if err != nil {
		t.Skipf("Orkes API spec not parseable: %v", err)
	}
	var found, total int
	for path, ops := range spec.Paths {
		for _, op := range ops {
			total++
			if op.OperationID == "startWorkflow" {
				found++
				if !strings.Contains(path, "/workflow") {
					t.Errorf("startWorkflow is at %s, expected a /workflow path", path)
				}
			}
		}
	}
	if found == 0 {
		t.Fatalf("no startWorkflow operation among %d operations in the spec", total)
	}

	agent := &ai.Agent{Name: "e2e_orkes_api", Model: model(t),
		Instructions: "You have access to the Orkes Conductor API tools. Answer questions about available API operations.",
		Tools:        ai.Tools(tool.API("orkes_api", "Orkes Conductor API", orkesSpecURL, tool.WithToolNames("startWorkflow")))}
	ad := agentDef(t, planAgent(t, rt, agent))
	var apiTools []map[string]any
	for _, tl := range asList(ad["tools"]) {
		if m, ok := tl.(map[string]any); ok && m["toolType"] == "api" {
			apiTools = append(apiTools, m)
		}
	}
	if len(apiTools) == 0 {
		t.Fatalf("no api tool in the plan: %v", ad["tools"])
	}
	if !strings.Contains(fmt.Sprint(apiTools[0]["config"]), "orkescloud") {
		t.Errorf("api tool config lacks the spec URL: %v", apiTools[0]["config"])
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, agent,
		"What is the API endpoint to start a new workflow? Give me the HTTP method, path, and operationId.")
	assertTerminal(t, res, "External OpenAPI spec")
	if res.Status == ai.StatusCompleted && !strings.Contains(res.Output, "startWorkflow") {
		t.Errorf("completed answer does not name startWorkflow: %.400s", res.Output)
	}
}
