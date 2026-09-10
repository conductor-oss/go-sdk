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
	"slices"
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

// RunOption adjusts how Run and Start begin a run.
type RunOption func(*runOptions)

type runOptions struct {
	plan *Plan
}

// WithPlan supplies the plan a StrategyPlanExecute agent carries out, in place
// of one written by its Planner.
//
// The server still requires the Planner slot to be set — the strategy is
// compiled around it — but with a plan supplied the planner never runs, so
// its instructions can say as much. The plan is validated before the run
// starts; see Plan.Validate for what is checked.
func WithPlan(plan *Plan) RunOption {
	return func(o *runOptions) { o.plan = plan }
}

// startPayload validates the agent and any options, registers the agent's
// workers, and builds the body of the /agent/start request. Run and Start
// share it so a plan reaches the server the same way from either.
func (r *Runtime) startPayload(agent *Agent, prompt string, opts []RunOption) (map[string]any, error) {
	if err := agent.Validate(); err != nil {
		return nil, err
	}
	var o runOptions
	for _, opt := range opts {
		opt(&o)
	}
	if o.plan != nil {
		if agent.Strategy != StrategyPlanExecute {
			return nil, fmt.Errorf(
				"agent %q: WithPlan requires StrategyPlanExecute, got %q", agent.Name, agent.Strategy)
		}
		if err := o.plan.Validate(); err != nil {
			return nil, fmt.Errorf("agent %q: %w", agent.Name, err)
		}
	}
	if err := r.registerWorkers(agent); err != nil {
		return nil, err
	}

	// A skill does not travel as agentConfig. The server's SkillNormalizer
	// compiles the raw document, which /agent/start accepts under
	// framework and rawConfig — the same request the Python SDK sends.
	if agent.skill != nil {
		return map[string]any{
			"framework": skillFramework,
			"rawConfig": agent.skill.rawConfig(),
			"prompt":    prompt,
			"sessionId": "",
			"media":     []any{},
			"context":   map[string]any{},
		}, nil
	}

	payload := map[string]any{
		"agentConfig": agent.toConfig(),
		"prompt":      prompt,
		"sessionId":   "",
		"media":       []any{},
	}
	if o.plan != nil {
		// The server reads workflow.input.static_plan ahead of the planner's
		// output; the key matches AgentRequest in Java and runtime.run(plan=)
		// in Python.
		payload["static_plan"] = o.plan.toPayload()
	}
	return payload, nil
}

// Start begins a run and returns at once.
//
// Workers for the agent's tools are registered before the run starts, so a tool
// call cannot arrive before something is polling for it.
func (r *Runtime) Start(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentHandle, error) {
	payload, err := r.startPayload(agent, prompt, opts)
	if err != nil {
		return nil, err
	}
	started, err := r.agents.Start(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("start agent %q: %w", agent.Name, err)
	}
	executionID, ok := started["executionId"].(string)
	if !ok || executionID == "" {
		return nil, fmt.Errorf("start agent %q: server returned no executionId", agent.Name)
	}
	return &AgentHandle{ExecutionID: executionID, rt: r}, nil
}

// Run starts an agent and blocks until the execution finishes.
//
// It validates the agent, registers a worker for every tool that has a Go
// handler, starts the run, then polls until the server reports a terminal
// state. Tool calls arrive as Conductor tasks while this is waiting.
func (r *Runtime) Run(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentResult, error) {
	payload, err := r.startPayload(agent, prompt, opts)
	if err != nil {
		return nil, err
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
		// The derived execution tools are serialized as worker tools, so they
		// need workers too. Without this the server queues a task nothing
		// polls for and the run stalls.
		tools := slices.Clone(a.Tools)
		if a.CodeExecution != nil && enabledOrDefault(a.CodeExecution.Enabled) {
			t := a.CodeExecution.codeTool(a.Name)
			t.Handler = a.CodeExecution.executeCode
			tools = append(tools, t)
		}
		if a.CLI != nil && enabledOrDefault(a.CLI.Enabled) {
			t := a.CLI.cliTool(a.Name)
			t.Handler = a.CLI.runCommand
			tools = append(tools, t)
		}
		// on_condition handoffs are worker-backed too. The server compiles each
		// into a SIMPLE task named "<agent>_handoff_<target>", so without a worker
		// under exactly that name the run stalls the way code execution used to.
		if conds := a.onConditions(); len(conds) > 0 {
			// The server reports the active agent as an index into
			// [parent, sub-agents...] — MultiAgentCompiler builds allSwarmAgents
			// with the parent first — so index 0 is this agent, not its first
			// child. The names go in that order so the handler resolves the
			// index the same way.
			names := make([]string, 0, 1+len(a.Agents))
			names = append(names, a.Name)
			for _, sub := range a.Agents {
				if sub != nil {
					names = append(names, sub.Name)
				}
			}
			for _, c := range conds {
				tools = append(tools, ToolDef{
					Name:    handoffTaskName(a.Name, c.Target),
					Handler: c.handoffHandler(names),
				})
			}
		}

		// Custom guardrails are worker-backed as well: the server compiles each
		// into a SIMPLE task named after the guardrail, agent-level and
		// tool-level alike.
		for _, g := range a.customGuardrails() {
			tools = append(tools, ToolDef{Name: g.Name, Handler: g.guardrailHandler()})
		}

		// A skill's scripts and its read_skill_file tool run here too; the
		// server emits worker tools under these names when it normalizes the
		// skill document.
		if a.skill != nil {
			tools = append(tools, a.skill.workers(a.Name)...)
		}

		for _, t := range tools {
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
		// An agent exposed as a tool runs as its own workflow, and its worker
		// tools are dispatched the same way as the parent's, so it needs its
		// workers registered too. This is what lets a skill serve as a tool.
		for _, t := range a.Tools {
			if sub, ok := t.Config["agent"].(*Agent); ok {
				if err := walk(sub); err != nil {
					return err
				}
			}
		}
		if a.Router != nil {
			if err := walk(a.Router); err != nil {
				return err
			}
		}
		if a.Planner != nil {
			if err := walk(a.Planner); err != nil {
				return err
			}
		}
		if a.Fallback != nil {
			return walk(a.Fallback)
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
