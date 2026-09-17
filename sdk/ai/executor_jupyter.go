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
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// JupyterExecutor runs code in a Jupyter kernel that stays alive between
// calls, so variables and imports persist from one snippet to the next, as in
// a notebook. It is the counterpart of the Python SDK's JupyterCodeExecutor
// and, like it, drives the kernel through the jupyter_client package: a small
// helper written in Python is started once and holds the kernel; this
// executor sends it code over a pipe and reads the outcome back. The worker
// host therefore needs a python3 with jupyter_client and ipykernel installed.
//
// Close shuts the kernel down. Execute calls are serialized: a kernel runs one
// cell at a time.
type JupyterExecutor struct {
	// KernelName is the Jupyter kernel to start; empty means python3.
	KernelName string
	// TimeoutSeconds bounds each cell; zero means 30.
	TimeoutSeconds int
	// StartupCode runs once when the kernel starts.
	StartupCode string
	// Python is the interpreter that hosts the helper; empty means python3
	// from PATH. Point it at an environment that has jupyter_client.
	Python string

	mu     sync.Mutex
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout *bufio.Reader
	err    error // a failed start, reported on every call
}

// jupyterBridge is the helper. It reads one JSON object per line from stdin,
// {"code": ...}, runs it on the kernel the way JupyterCodeExecutor.execute
// does, and writes one JSON object per line back: output, error, timed_out.
const jupyterBridge = `
import json, sys
try:
    from jupyter_client import KernelManager
except ImportError:
    print(json.dumps({"fatal": "JupyterCodeExecutor requires jupyter_client. Install with: pip install jupyter_client ipykernel"}), flush=True)
    sys.exit(0)
cfg = json.loads(sys.stdin.readline())
km = KernelManager(kernel_name=cfg["kernel"])
km.start_kernel()
kc = km.client()
kc.start_channels()
kc.wait_for_ready(timeout=30)
def run(code, timeout):
    kc.execute(code)
    outputs, errors = [], []
    try:
        while True:
            msg = kc.get_iopub_msg(timeout=timeout)
            t, c = msg.get("msg_type", ""), msg.get("content", {})
            if t == "stream":
                (errors if c.get("name") == "stderr" else outputs).append(c.get("text", ""))
            elif t == "execute_result":
                outputs.append(c.get("data", {}).get("text/plain", ""))
            elif t == "error":
                errors.append("\n".join(str(l) for l in c.get("traceback", [])))
            elif t == "status" and c.get("execution_state") == "idle":
                break
    except Exception:
        if not outputs and not errors:
            return {"output": "", "error": "", "timed_out": True}
    return {"output": "".join(outputs), "error": "".join(errors), "timed_out": False}
if cfg.get("startup"):
    run(cfg["startup"], cfg["timeout"])
print(json.dumps({"ready": True}), flush=True)
for line in sys.stdin:
    req = json.loads(line)
    print(json.dumps(run(req["code"], req["timeout"])), flush=True)
km.shutdown_kernel(now=True)
`

type jupyterReply struct {
	Output   string `json:"output"`
	Error    string `json:"error"`
	TimedOut bool   `json:"timed_out"`
	Ready    bool   `json:"ready"`
	Fatal    string `json:"fatal"`
}

// start launches the helper and waits for the kernel to be ready.
func (e *JupyterExecutor) start() error {
	python := e.Python
	if python == "" {
		python = "python3"
	}
	cmd := exec.Command(python, "-c", jupyterBridge) //nolint:gosec // operator configuration
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = nil
	if startErr := cmd.Start(); startErr != nil {
		return fmt.Errorf("start jupyter helper: %w", startErr)
	}
	e.cmd, e.stdin, e.stdout = cmd, stdin, bufio.NewReader(stdout)
	cfg, err := json.Marshal(map[string]any{
		"kernel":  orDefault(e.KernelName, "python3"),
		"startup": e.StartupCode,
		"timeout": timeoutOrDefault(e.TimeoutSeconds),
	})
	if err != nil {
		return err
	}
	if _, werr := e.stdin.Write(append(cfg, '\n')); werr != nil {
		return werr
	}
	reply, err := e.read(60 * time.Second)
	if err != nil {
		return fmt.Errorf("kernel startup failed: %w", err)
	}
	if reply.Fatal != "" {
		return fmt.Errorf("%s", reply.Fatal)
	}
	return nil
}

// read waits for the helper's next line.
func (e *JupyterExecutor) read(limit time.Duration) (jupyterReply, error) {
	type result struct {
		line string
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		line, err := e.stdout.ReadString('\n')
		ch <- result{line, err}
	}()
	select {
	case r := <-ch:
		if r.err != nil {
			return jupyterReply{}, r.err
		}
		var reply jupyterReply
		if err := json.Unmarshal([]byte(strings.TrimSpace(r.line)), &reply); err != nil {
			return jupyterReply{}, fmt.Errorf("unexpected helper output: %q", r.line)
		}
		return reply, nil
	case <-time.After(limit):
		return jupyterReply{}, fmt.Errorf("no reply from the kernel helper within %s", limit)
	}
}

func (e *JupyterExecutor) Execute(ctx context.Context, code string) ExecutionResult {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.cmd == nil && e.err == nil {
		e.err = e.start()
	}
	if e.err != nil {
		return ExecutionResult{Error: e.err.Error(), ExitCode: 1}
	}
	timeout := timeoutOrDefault(e.TimeoutSeconds)
	req, err := json.Marshal(map[string]any{"code": code, "timeout": timeout})
	if err != nil {
		return ExecutionResult{Error: "kernel helper: " + err.Error(), ExitCode: 1}
	}
	if _, werr := e.stdin.Write(append(req, '\n')); werr != nil {
		return ExecutionResult{Error: "kernel helper: " + werr.Error(), ExitCode: 1}
	}
	reply, err := e.read(time.Duration(timeout+30) * time.Second)
	if err != nil {
		return ExecutionResult{Error: "kernel helper: " + err.Error(), ExitCode: 1}
	}
	if reply.TimedOut {
		return ExecutionResult{Error: fmt.Sprintf("Execution timed out after %ds", timeout), ExitCode: -1, TimedOut: true}
	}
	res := ExecutionResult{Output: reply.Output, Error: reply.Error}
	if reply.Error != "" {
		res.ExitCode = 1
	}
	return res
}

// Close shuts the kernel and its helper down and reports any trouble doing
// so. The executor can be used again afterwards; the next Execute starts a
// fresh kernel.
func (e *JupyterExecutor) Close() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.cmd == nil {
		return nil
	}
	// Closing stdin ends the helper's loop, which shuts the kernel down.
	err := e.stdin.Close()
	done := make(chan error, 1)
	go func() { done <- e.cmd.Wait() }()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		err = errors.Join(err, e.cmd.Process.Kill())
	}
	e.cmd, e.stdin, e.stdout, e.err = nil, nil, nil, nil
	return err
}
