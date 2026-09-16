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
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/antihax/optional"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// Helpers shared by the ports of the Python SDK's e2e suites 2–10. They are
// the module-level helpers each Python suite repeats: fetch the workflow,
// find a tool's task, read and write the secret store, run mcp-testkit, and
// the checks every suite makes on a result.

// getWorkflow fetches an execution with its tasks, as the suites' _get_workflow does.
func getWorkflow(t *testing.T, executionID string) taskmodel.Workflow {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	svc := &client.WorkflowResourceApiService{APIClient: newAPIClient(t)}
	wf, _, err := svc.GetExecutionStatus(ctx, executionID,
		&client.WorkflowResourceApiGetExecutionStatusOpts{IncludeTasks: optional.NewBool(true)})
	if err != nil {
		t.Fatalf("get workflow %s: %v", executionID, err)
	}
	return wf
}

// systemTaskTypes are the task types that are never a tool's own task.
var systemTaskTypes = map[string]bool{
	"LLM_CHAT_COMPLETE": true, "SWITCH": true, "DO_WHILE": true, "INLINE": true,
	"SET_VARIABLE": true, "FORK": true, "FORK_JOIN_DYNAMIC": true, "JOIN": true,
	"SUB_WORKFLOW": true, "TERMINATE": true, "WAIT": true, "EVENT": true, "DECISION": true,
}

// findToolTasks returns the first task for each name whose reference carries
// the name or whose definition or type is the name (suite 2's _find_tool_tasks_for).
func findToolTasks(wf taskmodel.Workflow, names ...string) map[string]taskmodel.Task {
	found := map[string]taskmodel.Task{}
	for _, name := range names {
		for _, task := range wf.Tasks {
			if strings.Contains(task.ReferenceTaskName, name) || task.TaskDefName == name || task.TaskType == name {
				found[name] = task
				break
			}
		}
	}
	return found
}

// findMCPToolTasks is suite 4's _find_mcp_tool_tasks: a CALL_MCP_TOOL task is
// matched by the tool name in its input, any other task by definition, type
// or reference. An MCP task's reference is the model's call id, not the tool.
func findMCPToolTasks(wf taskmodel.Workflow, names ...string) map[string]taskmodel.Task {
	found := map[string]taskmodel.Task{}
	for _, name := range names {
		for _, task := range wf.Tasks {
			var hit bool
			if task.TaskType == "CALL_MCP_TOOL" {
				toolName := fmt.Sprint(task.InputData["toolName"])
				if v, ok := task.InputData["tool_name"]; ok && toolName == "<nil>" {
					toolName = fmt.Sprint(v)
				}
				hit = toolName == name || strings.Contains(fmt.Sprint(task.InputData), name)
			} else {
				hit = task.TaskDefName == name || task.TaskType == name || strings.Contains(task.ReferenceTaskName, name)
			}
			if hit {
				found[name] = task
				break
			}
		}
	}
	return found
}

// findHTTPToolTasks is suite 5's _find_http_tool_tasks: by definition or
// type, then by reference, then by the name appearing in a non-system task's input.
func findHTTPToolTasks(wf taskmodel.Workflow, names ...string) map[string]taskmodel.Task {
	found := map[string]taskmodel.Task{}
	for _, name := range names {
		for _, task := range wf.Tasks {
			hit := task.TaskDefName == name || task.TaskType == name ||
				strings.Contains(task.ReferenceTaskName, name) ||
				(!systemTaskTypes[task.TaskType] && strings.Contains(fmt.Sprint(task.InputData), name))
			if hit {
				found[name] = task
				break
			}
		}
	}
	return found
}

// tasksOfType returns the tasks whose type is one of types, in workflow order.
func tasksOfType(wf taskmodel.Workflow, types ...string) []taskmodel.Task {
	var out []taskmodel.Task
	for _, task := range wf.Tasks {
		for _, typ := range types {
			if task.TaskType == typ {
				out = append(out, task)
				break
			}
		}
	}
	return out
}

// taskStatus is a task's status as the Python suites compare it, a string.
func taskStatus(task taskmodel.Task) string { return string(task.Status) }

// outputString is str(task["outputData"]) for substring checks.
func outputString(task taskmodel.Task) string { return fmt.Sprint(task.OutputData) }

// assertRunCompleted mirrors the suites' _assert_run_completed: an execution
// id, not stalled at the tool-calling stage, and COMPLETED.
func assertRunCompleted(t *testing.T, res *ai.AgentResult, step string) {
	t.Helper()
	if res == nil || res.ExecutionID == "" {
		t.Fatalf("[%s] no execution id", step)
	}
	if res.FinishReason == "TOOL_CALLS" {
		t.Fatalf("[%s] run stalled at the tool-calling stage: status=%s error=%q", step, res.Status, res.Error)
	}
	if res.Status != ai.StatusCompleted {
		t.Fatalf("[%s] status = %q, want %q (error=%q)", step, res.Status, ai.StatusCompleted, res.Error)
	}
}

// assertTerminal mirrors the suites' "status in (COMPLETED, FAILED, TERMINATED)".
func assertTerminal(t *testing.T, res *ai.AgentResult, step string) {
	t.Helper()
	if res == nil || res.ExecutionID == "" {
		t.Fatalf("[%s] no execution id", step)
	}
	switch res.Status {
	case ai.StatusCompleted, ai.StatusFailed, ai.StatusTerminated:
	default:
		t.Fatalf("[%s] status = %q, want a terminal status (error=%q)", step, res.Status, res.Error)
	}
}

// runTolerant runs the agent and returns whatever result there is. A FAILED
// run makes Run return an error alongside the result; the Python suites look
// only at the result, so callers here do too.
func runTolerant(t *testing.T, rt *ai.Runtime, ctx context.Context, agent *ai.Agent, prompt string) *ai.AgentResult {
	t.Helper()
	res, err := rt.Run(ctx, agent, prompt)
	if res == nil {
		t.Fatalf("run of %s returned no result: %v", agent.Name, err)
	}
	return res
}

// ── secret store ────────────────────────────────────────────────────

// secretStore is the server's /secrets API as the suites use it: PUT a
// text/plain value, GET it back, DELETE it. The OSS server's store is
// read-only: it answers 501 to writes and serves the values it was started
// with as CONDUCTOR_SECRET_<NAME>.
type secretStore struct{ base string }

func newSecretStore(t *testing.T) secretStore {
	t.Helper()
	url := os.Getenv("CONDUCTOR_SERVER_URL")
	if url == "" {
		t.Skip("CONDUCTOR_SERVER_URL is not set")
	}
	return secretStore{base: strings.TrimRight(url, "/") + "/secrets/"}
}

func (s secretStore) do(method, name, body string) (int, string) {
	var rdr io.Reader
	if body != "" {
		rdr = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, s.base+name, rdr)
	if err != nil {
		return 0, err.Error()
	}
	if body != "" {
		req.Header.Set("Content-Type", "text/plain")
	}
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return 0, err.Error()
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, strings.TrimSpace(string(b))
}

func (s secretStore) get(name string) (string, int) {
	code, body := s.do(http.MethodGet, name, "")
	return body, code
}
func (s secretStore) put(name, value string) int {
	code, _ := s.do(http.MethodPut, name, value)
	return code
}
func (s secretStore) delete(name string) int {
	code, _ := s.do(http.MethodDelete, name, "")
	return code
}

// putSecretOrSkip is suites 2 and 3's _put_secret: a store that rejects the
// write ends the test as a skip, since the suite needs a writable store.
func putSecretOrSkip(t *testing.T, store secretStore, name, value string) {
	t.Helper()
	if code := store.put(name, value); code < 200 || code >= 300 {
		t.Skipf("server credential store rejected a write (HTTP %d) — this suite needs a writable store to set and update credentials", code)
	}
}

// ensureCredential is suites 4 and 5's _ensure_credential: store the
// preferred value if the store is writable; otherwise use the value the
// server already holds; otherwise skip. created says whether the test wrote it.
func ensureCredential(t *testing.T, store secretStore, name, preferred string) (string, bool) {
	t.Helper()
	if code := store.put(name, preferred); code >= 200 && code < 300 {
		return preferred, true
	}
	if value, code := store.get(name); code >= 200 && code < 300 && value != "" {
		return value, false
	}
	t.Skipf("no credential available for %s: the store rejected the write and the name is not provisioned. "+
		"Set CONDUCTOR_SECRET_%s in the server environment to run this phase.", name, name)
	return "", false
}

// ── mcp-testkit ─────────────────────────────────────────────────────

// requireMCPTestkitBinary skips unless mcp-testkit is on PATH.
func requireMCPTestkitBinary(t *testing.T) string {
	t.Helper()
	path, err := exec.LookPath("mcp-testkit")
	if err != nil {
		t.Skip("mcp-testkit not installed — required for this test (pip install mcp-testkit)")
	}
	return path
}

// startMCPTestkit starts `mcp-testkit --transport http --port <port>`, with
// `--auth <key>` when authKey is set, waits until it answers, and returns a
// stop function. The process is also stopped at cleanup.
func startMCPTestkit(t *testing.T, port int, authKey string) func() {
	t.Helper()
	args := []string{"--transport", "http", "--port", strconv.Itoa(port)}
	if authKey != "" {
		args = append(args, "--auth", authKey)
	}
	cmd := exec.Command(requireMCPTestkitBinary(t), args...)
	var output bytes.Buffer
	cmd.Stdout, cmd.Stderr = &output, &output
	if err := cmd.Start(); err != nil {
		t.Fatalf("start mcp-testkit: %v", err)
	}
	var once sync.Once
	stop := func() {
		once.Do(func() {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		})
	}
	t.Cleanup(stop)

	// Ready when any HTTP response comes back; a 4xx from auth mode counts.
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		if cmd.ProcessState != nil {
			t.Fatalf("mcp-testkit exited early: %s", output.String())
		}
		resp, err := (&http.Client{Timeout: 2 * time.Second}).Post(
			fmt.Sprintf("http://localhost:%d/mcp", port), "application/json", strings.NewReader("{}"))
		if err == nil {
			resp.Body.Close()
			return stop
		}
		time.Sleep(500 * time.Millisecond)
	}
	stop()
	t.Fatalf("mcp-testkit on port %d did not come up in 15s: %s", port, output.String())
	return stop
}

// expectedTestkitTools are mcp-testkit's 65 tools, the set the Python suites
// derive from the testkit's source at import time.
var expectedTestkitTools = []string{
	"collection_chunk", "collection_filter_gt", "collection_flatten", "collection_group_by",
	"collection_merge", "collection_sort", "collection_unique", "collection_zip",
	"conversion_bytes_to_human", "conversion_celsius_to_fahrenheit", "conversion_decimal_to_binary",
	"conversion_fahrenheit_to_celsius", "conversion_hex_to_rgb", "conversion_km_to_miles",
	"conversion_miles_to_km", "conversion_rgb_to_hex",
	"datetime_add_days", "datetime_day_of_week", "datetime_days_in_month", "datetime_diff",
	"datetime_format", "datetime_is_leap_year", "datetime_parse", "datetime_week_number",
	"echo", "echo_empty", "echo_error", "echo_large", "echo_multiple", "echo_nested", "echo_schema", "echo_types",
	"encoding_base64_decode", "encoding_base64_encode", "encoding_hex_decode", "encoding_hex_encode",
	"encoding_md5", "encoding_sha256", "encoding_url_decode", "encoding_url_encode",
	"get_weather",
	"math_add", "math_divide", "math_factorial", "math_fibonacci", "math_modulo", "math_multiply",
	"math_power", "math_subtract",
	"string_char_count", "string_join", "string_length", "string_lowercase", "string_replace",
	"string_reverse", "string_split", "string_uppercase",
	"validation_is_email", "validation_is_ipv4", "validation_is_ipv6", "validation_is_json",
	"validation_is_palindrome", "validation_is_url", "validation_is_uuid", "validation_matches_regex",
}

// The three deterministic testkit tools the lifecycle tests call, and the
// value each must show in its task output.
var (
	testkitToolNames    = []string{"math_add", "string_reverse", "encoding_base64_encode"}
	testkitToolExpected = map[string]string{"math_add": "7", "string_reverse": "olleh", "encoding_base64_encode": "dGVzdA=="}
)

// assertSameToolSet fails unless discovered is exactly the testkit's tool set.
func assertSameToolSet(t *testing.T, step string, discovered []string) {
	t.Helper()
	if len(discovered) != len(expectedTestkitTools) {
		t.Fatalf("[%s] discovered %d tools, want %d", step, len(discovered), len(expectedTestkitTools))
	}
	want := map[string]bool{}
	for _, n := range expectedTestkitTools {
		want[n] = true
	}
	var missing, extra []string
	got := map[string]bool{}
	for _, n := range discovered {
		got[n] = true
		if !want[n] {
			extra = append(extra, n)
		}
	}
	for _, n := range expectedTestkitTools {
		if !got[n] {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 || len(extra) > 0 {
		t.Fatalf("[%s] tool set differs. Missing: %v Extra: %v", step, missing, extra)
	}
}

// ── MCP and OpenAPI discovery ───────────────────────────────────────

// mcpListTools lists a Streamable-HTTP MCP server's tools the way the Python
// suite does with the mcp client: initialize, then tools/list. A server that
// refuses the request, such as one in auth mode without the key, is an error.
func mcpListTools(url, authKey string) ([]string, error) {
	var session string
	call := func(id int, method string, params any) (map[string]any, error) {
		body, _ := json.Marshal(map[string]any{"jsonrpc": "2.0", "id": id, "method": method, "params": params})
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json, text/event-stream")
		if authKey != "" {
			req.Header.Set("Authorization", "Bearer "+authKey)
		}
		if session != "" {
			req.Header.Set("Mcp-Session-Id", session)
		}
		resp, err := (&http.Client{Timeout: 15 * time.Second}).Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("%s: HTTP %d", method, resp.StatusCode)
		}
		if s := resp.Header.Get("Mcp-Session-Id"); s != "" {
			session = s
		}
		raw, _ := io.ReadAll(resp.Body)
		var msg map[string]any
		if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/event-stream") {
			scanner := bufio.NewScanner(bytes.NewReader(raw))
			scanner.Buffer(make([]byte, 1<<20), 1<<24)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "data: ") {
					var m map[string]any
					if json.Unmarshal([]byte(line[6:]), &m) == nil && (m["result"] != nil || m["error"] != nil) {
						msg = m
					}
				}
			}
		} else if err := json.Unmarshal(raw, &msg); err != nil {
			return nil, fmt.Errorf("%s: %v", method, err)
		}
		if msg == nil {
			return nil, fmt.Errorf("%s: no response message", method)
		}
		if e, ok := msg["error"]; ok && e != nil {
			return nil, fmt.Errorf("%s: %v", method, e)
		}
		result, _ := msg["result"].(map[string]any)
		return result, nil
	}
	if _, err := call(1, "initialize", map[string]any{
		"protocolVersion": "2025-03-26", "capabilities": map[string]any{},
		"clientInfo": map[string]any{"name": "go-sdk-integration", "version": "0"},
	}); err != nil {
		return nil, err
	}
	result, err := call(2, "tools/list", map[string]any{})
	if err != nil {
		return nil, err
	}
	var names []string
	if tools, ok := result["tools"].([]any); ok {
		for _, tl := range tools {
			if m, ok := tl.(map[string]any); ok {
				names = append(names, fmt.Sprint(m["name"]))
			}
		}
	}
	sort.Strings(names)
	return names, nil
}

// openAPIOperationIDs fetches an OpenAPI document and returns every
// operationId in it, sorted, as suite 5's _discover_tools_via_openapi does.
func openAPIOperationIDs(specURL, authKey string) ([]string, error) {
	req, _ := http.NewRequest(http.MethodGet, specURL, nil)
	if authKey != "" {
		req.Header.Set("Authorization", "Bearer "+authKey)
	}
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GET %s: HTTP %d", specURL, resp.StatusCode)
	}
	var spec struct {
		Paths map[string]map[string]struct {
			OperationID string `json:"operationId"`
		} `json:"paths"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&spec); err != nil {
		return nil, err
	}
	var ids []string
	for _, ops := range spec.Paths {
		for _, op := range ops {
			if op.OperationID != "" {
				ids = append(ids, op.OperationID)
			}
		}
	}
	sort.Strings(ids)
	return ids, nil
}

// httpStatus returns the status code of a GET, or 0 when the request failed.
func httpStatus(url string) int {
	resp, err := (&http.Client{Timeout: 5 * time.Second}).Get(url)
	if err != nil {
		return 0
	}
	resp.Body.Close()
	return resp.StatusCode
}
