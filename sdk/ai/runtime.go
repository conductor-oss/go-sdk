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
	"fmt"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
)

// Config holds runtime settings. The zero value is usable; ConfigFromEnv reads
// the CONDUCTOR_AGENT_* variables the other SDKs use.
type Config struct {
	// WorkerPollInterval is how often tool workers poll. Zero means 100ms.
	WorkerPollInterval time.Duration
	// WorkerBatchSize is how many tasks a worker takes per poll. Zero means 1.
	WorkerBatchSize int
	// StatusPollInterval is how often Run checks whether a run has finished.
	// Zero means 500ms.
	StatusPollInterval time.Duration
}

func (c Config) workerPoll() time.Duration {
	if c.WorkerPollInterval <= 0 {
		return 100 * time.Millisecond
	}
	return c.WorkerPollInterval
}

func (c Config) batchSize() int {
	if c.WorkerBatchSize <= 0 {
		return 1
	}
	return c.WorkerBatchSize
}

func (c Config) statusPoll() time.Duration {
	if c.StatusPollInterval <= 0 {
		return 500 * time.Millisecond
	}
	return c.StatusPollInterval
}

// Runtime starts agents and hosts the workers their tools need.
//
// It owns a TaskRunner, so one Runtime can serve many agents; workers are
// registered once per task name and reused across runs.
type Runtime struct {
	api     *client.APIClient
	agents  client.AgentClient
	runner  *worker.TaskRunner
	config  Config
	mu      sync.Mutex
	started map[string]bool // task names already registered
}

// NewRuntime builds a Runtime from CONDUCTOR_SERVER_URL and, on a secured
// server, CONDUCTOR_AUTH_KEY and CONDUCTOR_AUTH_SECRET.
func NewRuntime(cfg Config) *Runtime {
	apiClient := client.NewAPIClientFromEnv()
	return NewRuntimeWithClient(apiClient, cfg)
}

// NewRuntimeWithClient builds a Runtime over an existing APIClient, so agents
// share auth and connection settings with the rest of an application.
func NewRuntimeWithClient(apiClient *client.APIClient, cfg Config) *Runtime {
	return &Runtime{
		api:     apiClient,
		agents:  client.NewAgentClient(apiClient),
		runner:  worker.NewTaskRunnerWithApiClient(apiClient),
		config:  cfg,
		started: map[string]bool{},
	}
}

// AgentClient exposes the control plane for operations Runtime does not wrap.
func (r *Runtime) AgentClient() client.AgentClient { return r.agents }

// Run starts an agent and blocks until the execution finishes.
//
// It validates the agent, registers a worker for every tool that has a Go
// handler, starts the run, then polls until the server reports a terminal
// state. Tool calls arrive as Conductor tasks while this is waiting.
func (r *Runtime) Run(ctx context.Context, agent *Agent, prompt string) (*AgentResult, error) {
	if err := agent.Validate(); err != nil {
		return nil, err
	}
	if err := r.registerWorkers(agent); err != nil {
		return nil, err
	}

	payload := map[string]any{
		"agentConfig": agent.toConfig(),
		"prompt":      prompt,
		"sessionId":   "",
		"media":       []any{},
	}
	started, err := r.agents.Start(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("start agent %q: %w", agent.Name, err)
	}
	executionID, ok := started["executionId"].(string)
	if !ok || executionID == "" {
		return nil, fmt.Errorf("start agent %q: server returned no executionId", agent.Name)
	}

	return r.awaitResult(ctx, executionID)
}

// awaitResult polls until the run reaches a terminal state.
//
// Polling rather than SSE: the existing APIClient reads a whole response body
// before returning, so a streaming read needs its own request path. That is
// worth building for Stream; Run only needs the terminal state.
func (r *Runtime) awaitResult(ctx context.Context, executionID string) (*AgentResult, error) {
	ticker := time.NewTicker(r.config.statusPoll())
	defer ticker.Stop()

	for {
		status, err := r.agents.Status(ctx, executionID)
		if err != nil {
			return nil, fmt.Errorf("poll %s: %w", executionID, err)
		}
		res := resultFrom(executionID, status)
		if res.Status.Terminal() {
			if res.Status == StatusFailed {
				return res, fmt.Errorf("execution %s failed: %s", executionID, res.Error)
			}
			return res, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
		}
	}
}

// registerWorkers starts a Conductor worker for every tool carrying a Go
// handler. Tools the server dispatches itself, and external tools served
// elsewhere, have no handler and are skipped.
//
// Registration is idempotent per task name so repeated runs of the same agent
// do not stack workers.
func (r *Runtime) registerWorkers(agent *Agent) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var walk func(a *Agent) error
	walk = func(a *Agent) error {
		for _, t := range a.Tools {
			if t.Handler == nil || r.started[t.Name] {
				continue
			}
			fn, err := toolExecutor(t)
			if err != nil {
				return err
			}
			if err := r.runner.StartWorker(
				t.Name, fn, r.config.batchSize(), r.config.workerPoll()); err != nil {
				return fmt.Errorf("start worker %q: %w", t.Name, err)
			}
			r.started[t.Name] = true
		}
		for _, sub := range a.Agents {
			if err := walk(sub); err != nil {
				return err
			}
		}
		if a.Router != nil {
			return walk(a.Router)
		}
		return nil
	}
	return walk(agent)
}

// Shutdown stops every worker this runtime started.
func (r *Runtime) Shutdown() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for name := range r.started {
		r.runner.Shutdown(name)
		delete(r.started, name)
	}
}
