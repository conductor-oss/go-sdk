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
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
)

// The Python SDK's e2e/test_suite10_code_execution.py as Go tests, one per
// Python test and under the same name: the derived execute_code tool in the
// plan, and code running through the local, Docker and Jupyter executors.

func requireDocker(t *testing.T) {
	t.Helper()
	if err := exec.Command("docker", "info").Run(); err != nil {
		t.Skip("Docker not available")
	}
}

func requireJupyter(t *testing.T) {
	t.Helper()
	if err := exec.Command("python3", "-c", "import jupyter_client").Run(); err != nil {
		t.Skip("jupyter_client not installed")
	}
}

// executeCodeTasks is the suite's _find_execute_code_tasks.
func executeCodeTasks(wf taskmodel.Workflow) []taskmodel.Task {
	var out []taskmodel.Task
	for _, task := range wf.Tasks {
		if strings.Contains(task.ReferenceTaskName, "execute_code") || strings.Contains(task.TaskDefName, "execute_code") {
			out = append(out, task)
		}
	}
	return out
}

func anyTaskOutputContains(tasks []taskmodel.Task, want string) bool {
	for _, task := range tasks {
		if strings.Contains(outputString(task), want) {
			return true
		}
	}
	return false
}

// The suite's agents. Three share the name e2e_ce_local on purpose, as the
// Python suite does.

func s10AgentLocalCode(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_local", Model: m,
		Instructions: "You can execute code. When asked to compute something, write code in the specified language " +
			"that prints the result and execute it using the execute_code tool.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python", "bash"}, TimeoutSeconds: 30}}
}

func s10AgentPythonOnly(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_local", Model: m,
		Instructions:  "You can execute code. When asked to run code, execute it using your execute_code tool. You MUST use the tool.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}, TimeoutSeconds: 30}}
}

func s10AgentShortTimeout(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_local", Model: m, MaxTurns: 2,
		Instructions: "You can execute Python code. When asked to run code, execute it using your execute_code tool " +
			"exactly as provided. Do not modify the code.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}, TimeoutSeconds: 3,
			Executor: ai.LocalExecutor{Language: "python", TimeoutSeconds: 3}}}
}

func s10AgentDockerPython(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_docker_py", Model: m,
		Instructions: "You can execute Python code in a Docker container. When asked to compute something, " +
			"write Python code that prints the result and execute it.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}, TimeoutSeconds: 30,
			Executor: ai.DockerExecutor{Image: "python:3.12-slim", TimeoutSeconds: 30}}}
}

func s10AgentDockerNoNetwork(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_docker_nonet", Model: m,
		Instructions: "You can execute Python code in a Docker container with no network. When asked to run code, " +
			"execute it using your execute_code tool.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}, TimeoutSeconds: 30,
			Executor: ai.DockerExecutor{Image: "python:3.12-slim", TimeoutSeconds: 30, NetworkEnabled: false}}}
}

func s10AgentJupyter(m string) *ai.Agent {
	return &ai.Agent{Name: "e2e_ce_jupyter", Model: m,
		Instructions: "You can execute Python code in a Jupyter kernel. State persists across calls. " +
			"When asked to run code, execute it using your execute_code tool exactly as provided.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}, TimeoutSeconds: 30,
			Executor: &ai.JupyterExecutor{TimeoutSeconds: 30}}}
}

func TestCodeExecutionCompiles(t *testing.T) {
	rt := newRuntime(t)
	agent := &ai.Agent{Name: "e2e_ce_compile", Model: model(t), Instructions: "You can execute Python and Bash code.",
		CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python", "bash"}, TimeoutSeconds: 30}}
	ad := agentDef(t, planAgent(t, rt, agent))
	ce, _ := ad["codeExecution"].(map[string]any)
	if ce == nil {
		t.Fatalf("plan has no codeExecution: %v", keys(ad))
	}
	if ce["enabled"] != true || asInt(ce["timeout"]) != 30 {
		t.Errorf("codeExecution = %v", ce)
	}
	langs := asList(ce["allowedLanguages"])
	if !containsAny(langs, "python") || !containsAny(langs, "bash") {
		t.Errorf("allowedLanguages = %v", langs)
	}
	var execTools []map[string]any
	for _, tl := range asList(ad["tools"]) {
		if m, ok := tl.(map[string]any); ok && strings.Contains(fmt.Sprint(m["name"]), "execute_code") {
			execTools = append(execTools, m)
		}
	}
	if len(execTools) == 0 {
		t.Fatalf("no execute_code tool in the plan")
	}
	if execTools[0]["name"] != "e2e_ce_compile_execute_code" || execTools[0]["toolType"] != "worker" {
		t.Errorf("execute_code tool = %v", execTools[0])
	}
}

func TestToolNamingMultiAgent(t *testing.T) {
	rt := newRuntime(t)
	m := model(t)
	toolNames := func(name, instructions string) []string {
		ad := agentDef(t, planAgent(t, rt, &ai.Agent{Name: name, Model: m, Instructions: instructions,
			CodeExecution: &ai.CodeExecutionConfig{AllowedLanguages: []string{"python"}}}))
		var names []string
		for _, tl := range asList(ad["tools"]) {
			if mm, ok := tl.(map[string]any); ok {
				names = append(names, fmt.Sprint(mm["name"]))
			}
		}
		return names
	}
	a, b := toolNames("agent_a", "Agent A."), toolNames("agent_b", "Agent B.")
	if !contains(a, "agent_a_execute_code") || contains(a, "agent_b_execute_code") {
		t.Errorf("agent_a tools = %v", a)
	}
	if !contains(b, "agent_b_execute_code") || contains(b, "agent_a_execute_code") {
		t.Errorf("agent_b tools = %v", b)
	}
}

func TestLocalPythonExecution(t *testing.T) {
	rt := newRuntime(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, s10AgentLocalCode(model(t)), "Run this exact Python code using execute_code: print(42 * 73)")
	assertRunCompleted(t, res, "Local Python execution")
	tasks := executeCodeTasks(getWorkflow(t, res.ExecutionID))
	if len(tasks) < 1 {
		t.Fatal("no execute_code task ran")
	}
	if !anyTaskOutputContains(tasks, "3066") {
		t.Errorf("no execute_code task output holds 3066: %v", tasks[0].OutputData)
	}
}

func TestLocalBashExecution(t *testing.T) {
	rt := newRuntime(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, s10AgentLocalCode(model(t)), "Run a bash script that prints the result of: echo $((17 + 29))")
	assertRunCompleted(t, res, "Local Bash execution")
	tasks := executeCodeTasks(getWorkflow(t, res.ExecutionID))
	if len(tasks) < 1 {
		t.Fatal("no execute_code task ran")
	}
	if !anyTaskOutputContains(tasks, "46") {
		t.Errorf("no execute_code task output holds 46: %v", tasks[0].OutputData)
	}
}

func TestLanguageRestriction(t *testing.T) {
	rt := newRuntime(t)
	ad := agentDef(t, planAgent(t, rt, s10AgentPythonOnly(model(t))))
	ce, _ := ad["codeExecution"].(map[string]any)
	langs := asList(ce["allowedLanguages"])
	if !containsAny(langs, "python") || containsAny(langs, "bash") {
		t.Errorf("allowedLanguages = %v, want python only", langs)
	}
}

// A 3-second executor timeout kills a 30-second sleep: the sleeping task
// never completes successfully. The Python suite tolerates a run still
// RUNNING when its client timeout elapses, so this one starts the run and
// gives it a minute rather than failing on the deadline.
func TestLocalTimeout(t *testing.T) {
	rt := newRuntime(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	prompt := "Run this exact Python code using execute_code, preserving the line breaks exactly:\n" +
		"```python\nimport time\ntime.sleep(30)\nprint(\"done\")\n```"
	h, err := rt.Start(ctx, s10AgentShortTimeout(model(t)), prompt)
	if err != nil {
		t.Fatalf("Start: %v", err)
	}
	deadline := time.Now().Add(60 * time.Second)
	var res *ai.AgentResult
	for {
		res, err = h.Status(ctx)
		if err != nil {
			t.Fatalf("Status: %v", err)
		}
		if res.Status.Terminal() || time.Now().After(deadline) {
			break
		}
		time.Sleep(time.Second)
	}
	switch res.Status {
	case ai.StatusCompleted, ai.StatusFailed, ai.StatusTerminated, ai.StatusRunning:
	default:
		t.Fatalf("status = %q", res.Status)
	}

	var sleepTasks []taskmodel.Task
	for _, task := range executeCodeTasks(getWorkflow(t, h.ExecutionID)) {
		code := fmt.Sprint(task.InputData["code"])
		if code == "<nil>" {
			code = fmt.Sprint(task.InputData["source"])
		}
		if strings.Contains(strings.ToLower(code), "sleep") {
			sleepTasks = append(sleepTasks, task)
		}
	}
	if len(sleepTasks) == 0 {
		t.Fatal("no execute_code task ran the sleeping code")
	}
	for _, task := range sleepTasks {
		data := task.OutputData
		if r, ok := data["result"].(map[string]any); ok {
			data = r
		}
		stdout, status := fmt.Sprint(data["stdout"]), fmt.Sprint(data["status"])
		if strings.Contains(stdout, "done") {
			t.Errorf("the sleeping code printed done; the timeout did not kill it: %v", data)
		}
		if status == "success" {
			t.Errorf("the sleeping code reported success; the timeout did not kill it: %v", data)
		}
	}
}

func TestDockerPythonExecution(t *testing.T) {
	requireDocker(t)
	rt := newRuntime(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, s10AgentDockerPython(model(t)), "Run this exact Python code using execute_code: print(42 * 73)")
	assertRunCompleted(t, res, "Docker Python execution")
	tasks := executeCodeTasks(getWorkflow(t, res.ExecutionID))
	if len(tasks) < 1 {
		t.Fatal("no execute_code task ran")
	}
	if !anyTaskOutputContains(tasks, "3066") {
		t.Errorf("no execute_code task output holds 3066: %v", tasks[0].OutputData)
	}
}

func TestDockerNetworkDisabled(t *testing.T) {
	requireDocker(t)
	rt := newRuntime(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	res := runTolerant(t, rt, ctx, s10AgentDockerNoNetwork(model(t)),
		"Run this exact Python code using execute_code: import urllib.request; r = urllib.request.urlopen('http://example.com'); print(r.read()[:100])")
	assertTerminal(t, res, "Docker network disabled")
	tasks := executeCodeTasks(getWorkflow(t, res.ExecutionID))
	if len(tasks) < 1 {
		t.Fatal("no execute_code task ran")
	}
	var sawError bool
	for _, task := range tasks {
		out := strings.ToLower(outputString(task))
		for _, kw := range []string{"network", "connection", "urlopen", "unreachable", "refused", "errno", "error", "failed", "resolve", "gaierror"} {
			if strings.Contains(out, kw) {
				sawError = true
			}
		}
	}
	if !sawError {
		t.Errorf("no execute_code task reported a network error: %v", tasks[0].OutputData)
	}
}

// The Jupyter kernel lives in the worker for as long as the executor does,
// so a value set in one run is there in the next.
//
// Live, the model often defines its own x rather than using the kernel's, so
// this passes only some of the time; pytest's automatic reruns are what let
// the Python suite live with that. The recording is from a run where the
// model did follow the instruction, which is what makes the kernel's memory
// the only thing this test turns on in playback.
func TestJupyterStateful(t *testing.T) {
	requireJupyter(t)
	rt := newRuntime(t)
	agent := s10AgentJupyter(model(t))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	first := runTolerant(t, rt, ctx, agent, "Run this exact Python code using execute_code: x = 42")
	assertTerminal(t, first, "Jupyter run 1")

	second := runTolerant(t, rt, ctx, agent, "Run this exact Python code using execute_code: print(x * 73)")
	assertRunCompleted(t, second, "Jupyter run 2")
	tasks := executeCodeTasks(getWorkflow(t, second.ExecutionID))
	if len(tasks) < 1 {
		t.Fatal("no execute_code task ran in the second run")
	}
	if !anyTaskOutputContains(tasks, "3066") {
		for _, task := range tasks {
			t.Logf("run 2 executed %q -> %v", fmt.Sprint(task.InputData["code"]), task.OutputData)
		}
		for _, task := range executeCodeTasks(getWorkflow(t, first.ExecutionID)) {
			t.Logf("run 1 executed %q -> %v", fmt.Sprint(task.InputData["code"]), task.OutputData)
		}
		t.Errorf("the kernel forgot x between runs; outputs: %v", tasks[0].OutputData)
	}
}
