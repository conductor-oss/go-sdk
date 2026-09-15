//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func TestLocalExecutor(t *testing.T) {
	if _, err := exec.LookPath("python3"); err != nil {
		t.Skip("python3 not on PATH")
	}
	ctx := context.Background()
	res := LocalExecutor{}.Execute(ctx, "print('hi')")
	if !res.Success() || strings.TrimSpace(res.Output) != "hi" {
		t.Errorf("python run = %+v", res)
	}
	res = LocalExecutor{}.Execute(ctx, "import sys; sys.exit(3)")
	if res.Success() || res.ExitCode != 3 {
		t.Errorf("exit 3 not reported: %+v", res)
	}
	res = LocalExecutor{TimeoutSeconds: 1}.Execute(ctx, "import time; time.sleep(5)")
	if !res.TimedOut || res.ExitCode != -1 {
		t.Errorf("timeout not reported: %+v", res)
	}
	res = LocalExecutor{Language: "cobol"}.Execute(ctx, "x")
	if res.Success() || !strings.Contains(res.Error, "Unsupported language") {
		t.Errorf("unsupported language = %+v", res)
	}
	if res := (LocalExecutor{}).Execute(ctx, ""); !res.Success() || !strings.Contains(res.Output, "No code provided") {
		t.Errorf("empty code = %+v", res)
	}
}

// The docker command line matches the Python executor's, volume order made stable.
func TestDockerExecutorArgs(t *testing.T) {
	e := DockerExecutor{Image: "python:3.12-alpine", Language: "bash", MemoryLimit: "256m",
		Volumes: map[string]string{"/b": "/mnt/b", "/a": "/mnt/a"}}
	want := []string{"run", "--rm", "--network=none", "--memory", "256m",
		"-v", "/a:/mnt/a:ro", "-v", "/b:/mnt/b:ro", "python:3.12-alpine", "bash", "-c", "echo hi"}
	if got := e.args("echo hi"); !reflect.DeepEqual(got, want) {
		t.Errorf("args = %q\nwant %q", got, want)
	}
	if got := (DockerExecutor{NetworkEnabled: true}).args("x"); got[2] != "python:3.12-slim" || got[3] != "python3" {
		t.Errorf("defaults = %q", got)
	}
}

// Live: needs a docker daemon; skipped otherwise.
func TestDockerExecutorLive(t *testing.T) {
	if exec.Command("docker", "info").Run() != nil {
		t.Skip("docker not available")
	}
	res := DockerExecutor{TimeoutSeconds: 60}.Execute(context.Background(), "import sys; print(sys.version_info[0])")
	if !res.Success() || strings.TrimSpace(res.Output) != "3" {
		t.Fatalf("docker run = %+v", res)
	}
	res = DockerExecutor{TimeoutSeconds: 60}.Execute(context.Background(), "raise SystemExit(4)")
	if res.Success() || res.ExitCode != 4 {
		t.Errorf("exit code not reported through docker: %+v", res)
	}
}

func TestServerlessExecutor(t *testing.T) {
	var seen map[string]any
	var auth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		json.NewDecoder(r.Body).Decode(&seen)
		if strings.Contains(seen["code"].(string), "boom") {
			json.NewEncoder(w).Encode(map[string]any{"stdout": "", "stderr": "kaboom", "exit_code": 2})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"output": "cloud says hi"})
	}))
	defer srv.Close()
	e := ServerlessExecutor{Endpoint: srv.URL, APIKey: "sk-test", Language: "node", TimeoutSeconds: 7}
	res := e.Execute(context.Background(), "console.log(1)")
	if !res.Success() || res.Output != "cloud says hi" {
		t.Errorf("result = %+v", res)
	}
	if seen["language"] != "node" || seen["timeout"] != float64(7) || auth != "Bearer sk-test" {
		t.Errorf("request = %v auth=%q", seen, auth)
	}
	res = e.Execute(context.Background(), "boom")
	if res.Success() || res.ExitCode != 2 || res.Error != "kaboom" {
		t.Errorf("error result = %+v", res)
	}
	res = ServerlessExecutor{Endpoint: "http://127.0.0.1:1/nowhere", TimeoutSeconds: 1}.Execute(context.Background(), "x")
	if res.Success() || !strings.HasPrefix(res.Error, "Request failed") {
		t.Errorf("unreachable endpoint = %+v", res)
	}
}

type recordingExecutor struct {
	code string
	res  ExecutionResult
}

func (r *recordingExecutor) Execute(_ context.Context, code string) ExecutionResult {
	r.code = code
	return r.res
}

// The tool renders results the way the Python tool entries do.
func TestExecutorTool(t *testing.T) {
	rec := &recordingExecutor{res: ExecutionResult{Output: "42\n"}}
	td := ExecutorTool(rec, "", "")
	if td.Name != "execute_code" || !strings.Contains(td.Description, "python code") || !strings.Contains(td.Description, "Timeout: 30s") {
		t.Errorf("tool = %q / %q", td.Name, td.Description)
	}
	fn, err := toolExecutor(td)
	if err != nil {
		t.Fatal(err)
	}
	run := func(code string) codeExecOut {
		v, err := fn(&model.Task{InputData: map[string]any{"code": code}})
		if err != nil {
			t.Fatal(err)
		}
		return v.(codeExecOut)
	}
	if out := run("print(42)"); out.Status != "success" || out.Stdout != "42\n" || rec.code != "print(42)" {
		t.Errorf("success = %+v (code seen %q)", out, rec.code)
	}
	rec.res = ExecutionResult{Error: "Traceback...\nNameError\n", ExitCode: 1, TimedOut: true}
	if out := run("x"); out.Status != "error" || out.Stderr != "Traceback...\nNameError\nTIMED OUT after 30s\nExit code: 1" {
		t.Errorf("failure stderr = %q", out.Stderr)
	}
	if out := run(""); out.Status != "success" || !strings.Contains(out.Stdout, "No code provided") {
		t.Errorf("empty code = %+v", out)
	}
	named := ExecutorTool(DockerExecutor{Language: "bash", TimeoutSeconds: 5}, "run_shell", "Runs shell.")
	if named.Name != "run_shell" || named.Description != "Runs shell." {
		t.Errorf("named tool = %q / %q", named.Name, named.Description)
	}
}

// CodeExecutionConfig hands the code to its Executor, keeping the language and
// command checks in front of it; a non-local executor ignores the requested
// language, as in Python.
func TestCodeExecutionConfigUsesExecutor(t *testing.T) {
	rec := &recordingExecutor{res: ExecutionResult{Output: "ok"}}
	c := &CodeExecutionConfig{Executor: rec, AllowedLanguages: []string{"python", "bash"}}
	out, err := c.executeCode(context.Background(), codeExecIn{Code: "echo hi", Language: "bash"})
	if err != nil || out.Status != "success" || rec.code != "echo hi" {
		t.Errorf("out=%+v err=%v seen=%q", out, err, rec.code)
	}
	rec.code = ""
	out, _ = c.executeCode(context.Background(), codeExecIn{Code: "x", Language: "ruby"})
	if out.Status != "error" || !strings.Contains(out.Stderr, "not allowed") || rec.code != "" {
		t.Errorf("disallowed language reached the executor: out=%+v seen=%q", out, rec.code)
	}
	// Nil executor: local, per requested language.
	if _, err := exec.LookPath("bash"); err == nil {
		local := &CodeExecutionConfig{AllowedLanguages: []string{"bash"}}
		out, _ = local.executeCode(context.Background(), codeExecIn{Code: "echo from-bash", Language: "bash"})
		if out.Status != "success" || strings.TrimSpace(out.Stdout) != "from-bash" {
			t.Errorf("local bash = %+v", out)
		}
	}
}

// Live: needs a python with jupyter_client and ipykernel. Set
// CONDUCTOR_TEST_JUPYTER_PYTHON to such an interpreter; skipped otherwise.
func TestJupyterExecutorKeepsState(t *testing.T) {
	python := os.Getenv("CONDUCTOR_TEST_JUPYTER_PYTHON")
	if python == "" {
		python = "python3"
	}
	if exec.Command(python, "-c", "import jupyter_client, ipykernel").Run() != nil {
		t.Skipf("%s has no jupyter_client/ipykernel", python)
	}
	e := &JupyterExecutor{Python: python, TimeoutSeconds: 30, StartupCode: "base = 40"}
	defer e.Close()
	ctx := context.Background()
	if res := e.Execute(ctx, "x = base + 2"); !res.Success() {
		t.Fatalf("assignment = %+v", res)
	}
	if res := e.Execute(ctx, "print(x)"); !res.Success() || strings.TrimSpace(res.Output) != "42" {
		t.Fatalf("state did not persist across cells: %+v", res)
	}
	if res := e.Execute(ctx, "undefined_name"); res.Success() || !strings.Contains(res.Error, "NameError") {
		t.Errorf("kernel error not reported: %+v", res)
	}
}
