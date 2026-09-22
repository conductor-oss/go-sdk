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
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite4_mcp_tools.py as Go tests. The server
// discovers and calls an agent's MCP tools, first on an open MCP server and
// then on one that wants a bearer token the server resolves from its store.
//
// test_mcp_result_reaches_the_answer is TestMCPToolResultReachesTheAnswer in
// mcp_test.go, ported earlier against the shared testkit on port 3001; it is
// not repeated here. test_mcp_lifecycle is TestMcpLifecycle below, on its
// own testkit on port 3002 so it can restart it in auth mode.

const (
	mcpLifecyclePort = 3002
	mcpLifecycleURL  = "http://localhost:3002/mcp"
	mcpAuthKeyName   = "MCP_AUTH_KEY"
	mcpAuthKeyValue  = "e2e-test-secret-key-12345"

	// Shared by suites 4 and 5.
	toolAgentInstructions = "You have access to MCP tools. Call exactly the tools specified in each prompt.\n" +
		"Report each tool's result verbatim. Do not skip any tool.\n"
	promptUseThreeTools = "Call exactly these three tools with these exact arguments:\n" +
		"1. math_add with a=3 and b=4\n" +
		"2. string_reverse with text=\"hello\"\n" +
		"3. encoding_base64_encode with text=\"test\"\n" +
		"Report each result.\n"
)

// validateToolExecution is the suites' _validate_tool_execution: the run
// completed, and each of the three tools has a COMPLETED task whose output
// holds the expected value.
func validateToolExecution(t *testing.T, res *ai.AgentResult, step string, tasks map[string]taskmodel.Task) {
	t.Helper()
	assertRunCompleted(t, res, step)
	for _, name := range testkitToolNames {
		task, ok := tasks[name]
		if !ok {
			t.Errorf("[%s] no task ran for %s", step, name)
			continue
		}
		if taskStatus(task) != "COMPLETED" {
			t.Errorf("[%s] %s status = %q (reason=%q)", step, name, taskStatus(task), task.ReasonForIncompletion)
		}
		if want := testkitToolExpected[name]; !strings.Contains(outputString(task), want) {
			t.Errorf("[%s] %s output lacks %q: %s", step, name, want, outputString(task))
		}
	}
}

func TestMcpLifecycle(t *testing.T) {
	requireMCPTestkitBinary(t)
	rt := newRuntime(t)
	store := newSecretStore(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Phase 1: an open MCP server.
	stop := startMCPTestkit(t, mcpLifecyclePort, "")
	discovered, err := mcpListTools(mcpLifecycleURL, "")
	if err != nil {
		t.Fatalf("[Phase 1: Discovery] %v", err)
	}
	assertSameToolSet(t, "Phase 1: Discovery", discovered)

	agent := &ai.Agent{Name: "e2e_mcp_unauth", Model: model(t), Instructions: toolAgentInstructions,
		Tools: ai.Tools(tool.MCP("test_mcp", "Deterministic test tools via MCP", mcpLifecycleURL))}
	res := runTolerant(t, rt, ctx, agent, promptUseThreeTools)
	validateToolExecution(t, res, "Phase 1: Unauthenticated execution",
		findMCPToolTasks(getWorkflow(t, res.ExecutionID), testkitToolNames...))

	// Phase 2: the same server wanting the bearer token the store holds.
	authKey, created := ensureCredential(t, store, mcpAuthKeyName, mcpAuthKeyValue)
	if created {
		t.Cleanup(func() { store.delete(mcpAuthKeyName) })
	}
	stop()
	time.Sleep(time.Second)
	startMCPTestkit(t, mcpLifecyclePort, authKey)

	if _, err := mcpListTools(mcpLifecycleURL, ""); err == nil {
		t.Fatal("[Phase 2: Auth check] the server in auth mode answered a request without the token")
	}
	discoveredAuth, err := mcpListTools(mcpLifecycleURL, authKey)
	if err != nil {
		t.Fatalf("[Phase 2: Auth Discovery] %v", err)
	}
	assertSameToolSet(t, "Phase 2: Auth Discovery", discoveredAuth)

	authAgent := &ai.Agent{Name: "e2e_mcp_auth", Model: model(t), Instructions: toolAgentInstructions,
		Tools: ai.Tools(tool.MCP("test_mcp_auth", "Authenticated MCP test tools", mcpLifecycleURL,
			tool.WithHeaders(map[string]string{"Authorization": "Bearer ${" + mcpAuthKeyName + "}"}),
			tool.WithCredentials(mcpAuthKeyName)))}
	resAuth := runTolerant(t, rt, ctx, authAgent, promptUseThreeTools)
	validateToolExecution(t, resAuth, "Phase 2: Authenticated execution",
		findMCPToolTasks(getWorkflow(t, resAuth.ExecutionID), testkitToolNames...))
}
