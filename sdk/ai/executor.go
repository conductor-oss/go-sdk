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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// A CodeExecutor runs code the model wrote and reports what happened. The
// executors here are the counterparts of the Python SDK's LocalCodeExecutor,
// DockerCodeExecutor, JupyterCodeExecutor and ServerlessCodeExecutor.
//
// An executor is used in two ways: set on CodeExecutionConfig.Executor, where
// it replaces the local subprocess behind the agent's derived execute_code
// tool, or wrapped as a standalone tool with ExecutorTool. Executors run on
// the worker host; only Docker and a remote service isolate the code from it.
type CodeExecutor interface {
	// Execute runs code and returns the outcome. It reports failures of the
	// code in the result rather than as an error, so the model can read the
	// stderr and try again; the error return is for the executor itself
	// being unusable.
	Execute(ctx context.Context, code string) ExecutionResult
}

// ExecutionResult is what an executor observed: the code's output, its
// stderr, its exit code, and whether it was killed for running too long.
type ExecutionResult struct {
	Output   string
	Error    string
	ExitCode int
	TimedOut bool
}

// Success is true when the code exited 0 without timing out.
func (r ExecutionResult) Success() bool { return r.ExitCode == 0 && !r.TimedOut }

// toolOutput renders a result the way the Python SDK's tool entries do: a
// success carries stdout and stderr as they were; a failure folds the error,
// a timeout note and the exit code into stderr, so the model sees one text.
func (r ExecutionResult) toolOutput(timeoutSeconds int) codeExecOut {
	if r.Success() {
		return codeExecOut{Status: "success", Stdout: r.Output, Stderr: r.Error}
	}
	var parts []string
	if s := strings.TrimRight(r.Error, "\n"); s != "" {
		parts = append(parts, s)
	}
	if r.TimedOut {
		parts = append(parts, fmt.Sprintf("TIMED OUT after %ds", timeoutSeconds))
	}
	parts = append(parts, fmt.Sprintf("Exit code: %d", r.ExitCode))
	return codeExecOut{Status: "error", Stdout: r.Output, Stderr: strings.Join(parts, "\n")}
}

// LocalExecutor runs code in a subprocess on the worker host, with no sandbox.
// It is what CodeExecutionConfig uses when no Executor is set.
type LocalExecutor struct {
	// Language selects the interpreter; empty means python.
	Language string
	// TimeoutSeconds kills the process after this long; zero means 30.
	TimeoutSeconds int
	// WorkingDir is the process's working directory; empty means the
	// directory the code file is written to.
	WorkingDir string
}

func (e LocalExecutor) Execute(ctx context.Context, code string) ExecutionResult {
	if code == "" {
		return ExecutionResult{Output: "No code provided. Nothing to execute."}
	}
	language := e.Language
	if language == "" {
		language = defaultLanguage
	}
	interpreter, ok := interpreters[language]
	if !ok {
		return ExecutionResult{Error: "Unsupported language: " + language, ExitCode: 1}
	}
	return runInterpreter(ctx, interpreter, code, fileExtensions[language],
		timeoutOrDefault(e.TimeoutSeconds), e.WorkingDir)
}

// DockerExecutor runs each snippet in a fresh container, so the code cannot
// reach the host's filesystem or, unless enabled, the network. It needs the
// docker command and a running daemon on the worker host.
type DockerExecutor struct {
	// Image is the container image; empty means python:3.12-slim.
	Image string
	// Language selects the interpreter inside the image: python, bash or
	// node. Empty means python.
	Language string
	// TimeoutSeconds bounds the run; zero means 30. The container gets ten
	// seconds more, for startup.
	TimeoutSeconds int
	// NetworkEnabled gives the container network access. Off by default.
	NetworkEnabled bool
	// MemoryLimit is passed to docker --memory, for example "256m".
	MemoryLimit string
	// Volumes mounts host paths read-only at container paths.
	Volumes map[string]string
}

// dockerInterpreters maps a language to the command inside the image.
var dockerInterpreters = map[string]string{"python": "python3", "bash": "bash", "node": "node"}

// args is the docker command line for one snippet, without the leading
// "docker". Volumes are listed in sorted order so the command is stable.
func (e DockerExecutor) args(code string) []string {
	args := []string{"run", "--rm"}
	if !e.NetworkEnabled {
		args = append(args, "--network=none")
	}
	if e.MemoryLimit != "" {
		args = append(args, "--memory", e.MemoryLimit)
	}
	for _, host := range sortedVolumeKeys(e.Volumes) {
		args = append(args, "-v", host+":"+e.Volumes[host]+":ro")
	}
	image := e.Image
	if image == "" {
		image = "python:3.12-slim"
	}
	interpreter, ok := dockerInterpreters[e.Language]
	if !ok {
		interpreter = "python3"
	}
	return append(args, image, interpreter, "-c", code)
}

func (e DockerExecutor) Execute(ctx context.Context, code string) ExecutionResult {
	timeout := timeoutOrDefault(e.TimeoutSeconds)
	runCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout+10)*time.Second)
	defer cancel()
	// The image and interpreter are operator configuration; the model's code
	// travels as one argument.
	cmd := exec.CommandContext(runCtx, "docker", e.args(code)...) //nolint:gosec // see above
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	res := ExecutionResult{Output: stdout.String(), Error: stderr.String()}
	switch {
	case err == nil:
		return res
	case runCtx.Err() != nil:
		return ExecutionResult{Error: fmt.Sprintf("Docker execution timed out after %ds", timeout),
			ExitCode: -1, TimedOut: true}
	default:
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			res.ExitCode = ee.ExitCode()
			return res
		}
		if errors.Is(err, exec.ErrNotFound) {
			return ExecutionResult{Error: "Docker not found. Is Docker installed and running?", ExitCode: 127}
		}
		return ExecutionResult{Error: err.Error(), ExitCode: 1}
	}
}

// ServerlessExecutor sends the code to an HTTP service and reads the outcome
// back, for AWS Lambda, Cloud Functions, or a hosted execution API. The
// request is a JSON object with code, language and timeout; the response is
// read for output or stdout, error or stderr, and exit_code.
type ServerlessExecutor struct {
	// Endpoint is the URL the code is POSTed to. Required.
	Endpoint string
	// APIKey, when set, is sent as a bearer token.
	APIKey string
	// Language is sent with the code; empty means python.
	Language string
	// TimeoutSeconds is sent with the code and bounds the request, which
	// gets five seconds more. Zero means 30.
	TimeoutSeconds int
	// Headers are added to the request.
	Headers map[string]string
	// Client makes the request; nil means http.DefaultClient.
	Client *http.Client
}

func (e ServerlessExecutor) Execute(ctx context.Context, code string) ExecutionResult {
	timeout := timeoutOrDefault(e.TimeoutSeconds)
	language := e.Language
	if language == "" {
		language = defaultLanguage
	}
	body, err := json.Marshal(map[string]any{"code": code, "language": language, "timeout": timeout})
	if err != nil {
		return ExecutionResult{Error: "Request failed: " + err.Error(), ExitCode: 1}
	}
	reqCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout+5)*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(reqCtx, http.MethodPost, e.Endpoint, bytes.NewReader(body))
	if err != nil {
		return ExecutionResult{Error: "Request failed: " + err.Error(), ExitCode: 1}
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}
	if e.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+e.APIKey)
	}
	client := e.Client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return ExecutionResult{Error: "Request failed: " + err.Error(), ExitCode: 1}
	}
	raw, readErr := io.ReadAll(resp.Body)
	if closeErr := resp.Body.Close(); readErr == nil {
		readErr = closeErr
	}
	if readErr != nil {
		return ExecutionResult{Error: "Request failed: " + readErr.Error(), ExitCode: 1}
	}
	if resp.StatusCode >= 400 {
		return ExecutionResult{Error: fmt.Sprintf("Request failed: HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(raw))), ExitCode: 1}
	}
	var data map[string]any
	if err := json.Unmarshal(raw, &data); err != nil {
		return ExecutionResult{Error: "Request failed: " + err.Error(), ExitCode: 1}
	}
	res := ExecutionResult{Output: firstString(data, "output", "stdout"), Error: firstString(data, "error", "stderr")}
	if code, ok := data["exit_code"].(float64); ok {
		res.ExitCode = int(code)
	}
	return res
}

// ExecutorTool wraps an executor as a tool the model can call directly, the
// counterpart of the Python SDK's executor.as_tool(). The tool takes one
// argument, code, and returns status, stdout and stderr. An empty name means
// execute_code; an empty description is derived from the language and
// timeout. Unlike CodeExecutionConfig, there is no language or command
// allow-list here: the executor runs whatever it is given.
func ExecutorTool(exec CodeExecutor, name, description string) ToolDef {
	if name == "" {
		name = "execute_code"
	}
	language, timeout := executorLanguage(exec), executorTimeout(exec)
	if description == "" {
		description = fmt.Sprintf("Execute %s code. Returns stdout, stderr, and exit code. Timeout: %ds.", language, timeout)
	}
	return ToolDef{
		Name:        name,
		Description: description,
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{"code": map[string]any{"type": "string", "description": "The code to execute."}},
			"required":   []string{"code"},
		},
		Handler: func(ctx context.Context, in struct {
			Code string `json:"code"`
		}) (codeExecOut, error) {
			if in.Code == "" {
				return codeExecOut{Status: "success", Stdout: "No code provided. Nothing to execute."}, nil
			}
			return exec.Execute(ctx, in.Code).toolOutput(timeout), nil
		},
	}
}

// executorLanguage and executorTimeout read the settings the concrete
// executors expose, for descriptions and messages.
func executorLanguage(e CodeExecutor) string {
	switch x := e.(type) {
	case LocalExecutor:
		return orDefault(x.Language, defaultLanguage)
	case *LocalExecutor:
		return orDefault(x.Language, defaultLanguage)
	case DockerExecutor:
		return orDefault(x.Language, defaultLanguage)
	case *DockerExecutor:
		return orDefault(x.Language, defaultLanguage)
	case ServerlessExecutor:
		return orDefault(x.Language, defaultLanguage)
	case *ServerlessExecutor:
		return orDefault(x.Language, defaultLanguage)
	}
	return defaultLanguage
}

func executorTimeout(e CodeExecutor) int {
	switch x := e.(type) {
	case LocalExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case *LocalExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case DockerExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case *DockerExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case ServerlessExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case *ServerlessExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	case *JupyterExecutor:
		return timeoutOrDefault(x.TimeoutSeconds)
	}
	return timeoutOrDefault(0)
}

func orDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func sortedVolumeKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
