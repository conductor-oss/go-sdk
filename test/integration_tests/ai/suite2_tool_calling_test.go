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
	"os"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite2_tool_calling.py as a Go test, under the
// same name. A worker tool's credentials come only from the server's store,
// delivered with the task at poll time: a missing credential fails the task
// terminally, the process environment is never read, and a stored value and
// its later update reach the tool on the next run with a fresh worker.

type s2XIn struct {
	X string
}

// skipInPlayback skips a test that can only run against a live model.
func skipInPlayback(t *testing.T, reason string) {
	t.Helper()
	if model(t) == mockModel {
		t.Skip("runs live only: " + reason)
	}
}

// requireRuntimeMetadata is the conftest fixture of the same name: skip unless
// the server persists a task definition's runtimeMetadata, which credential
// delivery depends on.
func requireRuntimeMetadata(t *testing.T) {
	t.Helper()
	base := strings.TrimRight(os.Getenv("CONDUCTOR_SERVER_URL"), "/")
	if base == "" {
		t.Skip("CONDUCTOR_SERVER_URL is not set")
	}
	name := fmt.Sprintf("_rtmd_cap_probe_%d", time.Now().UnixNano()%100000000)
	body := fmt.Sprintf(`[{"name":%q,"runtimeMetadata":["PROBE"],"retryCount":0,"timeoutSeconds":0}]`, name)
	hc := &http.Client{Timeout: 8 * time.Second}
	defer func() {
		req, _ := http.NewRequest(http.MethodDelete, base+"/metadata/taskdefs/"+name, nil)
		if resp, err := hc.Do(req); err == nil {
			resp.Body.Close()
		}
	}()
	resp, err := hc.Post(base+"/metadata/taskdefs", "application/json", strings.NewReader(body))
	if err != nil || (resp.StatusCode != 200 && resp.StatusCode != 204) {
		t.Skip("server does not persist TaskDef.runtimeMetadata (conductor-oss PR #1255) — worker credential injection requires it")
	}
	resp.Body.Close()
	got, err := hc.Get(base + "/metadata/taskdefs/" + name)
	if err != nil || got.StatusCode != 200 {
		t.Skip("server does not persist TaskDef.runtimeMetadata (conductor-oss PR #1255) — worker credential injection requires it")
	}
	defer got.Body.Close()
	var def struct {
		RuntimeMetadata []string `json:"runtimeMetadata"`
	}
	if decodeJSON(got, &def) != nil || len(def.RuntimeMetadata) != 1 || def.RuntimeMetadata[0] != "PROBE" {
		t.Skip("server does not deliver TaskDef.runtimeMetadata (conductor-oss PR #1255) — worker credential injection requires it")
	}
}

// decodeJSON reads a JSON response body into v.
func decodeJSON(resp *http.Response, v any) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

// paidTool returns the first three characters of the named credential, the
// way paid_tool_a and paid_tool_b do. The Python dispatcher fails the task
// terminally before the function runs when a declared credential was not
// delivered; here the handler asks for it and returns a non-retryable error,
// which ends the task the same way.
func paidTool(name, prefix, credential string) ai.ToolDef {
	return tool.Func(name, func(ctx context.Context, in s2XIn) (string, error) {
		v, err := ai.Secret(ctx, credential)
		if err != nil {
			return "", taskmodel.NewNonRetryableError(err)
		}
		if len(v) > 3 {
			v = v[:3]
		}
		return prefix + ":" + v, nil
	},
		fmt.Sprintf("A tool that needs %s. Returns first 3 chars of credential.", credential), tool.WithCredentials(credential))
}

func TestCredentialLifecycle(t *testing.T) {
	requireRuntimeMetadata(t)
	rt := newRuntime(t)
	store := newSecretStore(t)
	const credA, credB = "E2E_CRED_A", "E2E_CRED_B"
	const prompt = "Call all three tools."

	freeTool := tool.Func("free_tool", func(context.Context, s2XIn) (string, error) { return "free:ok", nil },
		"A tool that needs no credentials. Always succeeds.")
	agent := &ai.Agent{Name: "e2e_cred_lifecycle", Model: model(t), MaxTurns: 3,
		Instructions: "You have three tools: free_tool, paid_tool_a, and paid_tool_b.\n" +
			"You MUST call all three tools exactly once each, with the argument \"test\".\n" +
			"After calling all three, report each tool's output verbatim in this format:\n" +
			"  free_tool: <output>\n  paid_tool_a: <output>\n  paid_tool_b: <output>\n" +
			"Do not skip any tool. Do not add commentary.\n",
		Tools: ai.Tools(freeTool, paidTool("paid_tool_a", "paid_a", credA), paidTool("paid_tool_b", "paid_b", credB))}
	toolNames := []string{"free_tool", "paid_tool_a", "paid_tool_b"}

	t.Cleanup(func() {
		store.delete(credA)
		store.delete(credB)
		os.Unsetenv(credA)
		os.Unsetenv(credB)
	})
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	// Step 1: clean slate.
	store.delete(credA)
	store.delete(credB)

	// Step 2: no credentials. The paid tools must fail terminally, not retry.
	res := runTolerant(t, rt, ctx, agent, prompt)
	assertTerminal(t, res, "Step 2: No credentials")
	for name, task := range findToolTasks(getWorkflow(t, res.ExecutionID), toolNames...) {
		if name == "free_tool" {
			continue
		}
		if s := taskStatus(task); s != "FAILED_WITH_TERMINAL_ERROR" && s != "COMPLETED_WITH_ERRORS" {
			t.Errorf("[Step 2] %s status = %q, want a terminal failure (reason=%q)", name, s, task.ReasonForIncompletion)
		}
	}

	// Step 3: values in the process environment must not be read.
	os.Setenv(credA, "from-env-aaa")
	os.Setenv(credB, "from-env-bbb")
	resEnv := runTolerant(t, rt, ctx, agent, prompt)
	os.Unsetenv(credA)
	os.Unsetenv(credB)
	if strings.Contains(resEnv.Output, "from-env") {
		t.Fatalf("[Step 3] a tool read the credential from the environment: %.300s", resEnv.Output)
	}

	// Step 4: store the credentials and run with a fresh worker.
	rt = restartRuntime(t, rt)
	putSecretOrSkip(t, store, credA, "secret-aaa-value")
	putSecretOrSkip(t, store, credB, "secret-bbb-value")
	resCreds := runTolerant(t, rt, ctx, agent, prompt)
	assertRunCompleted(t, resCreds, "Step 4: With credentials")
	tasks := findToolTasks(getWorkflow(t, resCreds.ExecutionID), toolNames...)
	if task, ok := tasks["free_tool"]; !ok || taskStatus(task) != "COMPLETED" {
		t.Errorf("[Step 4] free_tool = %v", tasks["free_tool"])
	}
	for _, name := range []string{"paid_tool_a", "paid_tool_b"} {
		task, ok := tasks[name]
		if !ok || taskStatus(task) != "COMPLETED" || !strings.Contains(outputString(task), "sec") {
			t.Errorf("[Step 4] %s should have completed with the stored value: %v", name, task)
		}
	}
	if out := resCreds.Output; !strings.Contains(strings.ToLower(out), "free") || !strings.Contains(out, "sec") {
		t.Errorf("[Step 4] answer should report free_tool and the stored value: %.300s", out)
	}

	// Step 5: update the credentials; the next fresh worker sees the new values.
	rt = restartRuntime(t, rt)
	putSecretOrSkip(t, store, credA, "newval-xxx-updated")
	putSecretOrSkip(t, store, credB, "newval-yyy-updated")
	resUpdated := runTolerant(t, rt, ctx, agent, prompt)
	assertRunCompleted(t, resUpdated, "Step 5: Updated credentials")
	task, ok := findToolTasks(getWorkflow(t, resUpdated.ExecutionID), "paid_tool_a")["paid_tool_a"]
	if !ok || taskStatus(task) != "COMPLETED" || !strings.Contains(outputString(task), "new") {
		t.Errorf("[Step 5] paid_tool_a should have completed with the updated value: %v", task)
	}
	if !strings.Contains(resUpdated.Output, "new") {
		t.Errorf("[Step 5] answer should report the updated value: %.300s", resUpdated.Output)
	}
}

// restartRuntime is the suite's restart_runtime: stop the pollers, wait for
// them to drain, and start a fresh runtime that registers its workers anew.
func restartRuntime(t *testing.T, rt *ai.Runtime) *ai.Runtime {
	t.Helper()
	rt.Shutdown()
	time.Sleep(2 * time.Second)
	return newRuntime(t)
}
