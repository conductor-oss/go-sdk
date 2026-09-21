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
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
)

// Config holds runtime settings. The zero value is usable; ConfigFromEnv reads the
// CONDUCTOR_AGENT_* variables the other SDKs use.
type Config struct {
	// WorkerPollInterval is how often tool workers poll. Zero means 100ms.
	WorkerPollInterval time.Duration
	// WorkerBatchSize is how many tasks a worker takes per poll. Zero means 1.
	WorkerBatchSize int
	// StatusPollInterval is how often Run checks for completion. Zero means 500ms.
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

// Runtime starts agents and hosts the workers their tools need. One Runtime serves many
// agents from a single TaskRunner; a worker per task name is reused across runs.
type Runtime struct {
	api       *client.APIClient
	agents    client.AgentClient
	metadata  client.MetadataClient
	workflow  client.WorkflowClient
	scheduler client.SchedulerClient
	runner    *worker.TaskRunner
	config    Config
	mu        sync.Mutex
	// started records which (task name, domain) pairs already have a worker.
	started map[workerKey]bool
	// defs are the task definitions for started workers; registered marks the sent ones.
	defs       map[string]model.TaskDef
	registered map[string]bool
}

// NewRuntime builds a Runtime from the environment.
//
// CONDUCTOR_SERVER_URL is the server's API base URL, such as http://localhost:8080/api,
// and is all an unauthenticated open-source Conductor needs. CONDUCTOR_AUTH_KEY and
// CONDUCTOR_AUTH_SECRET are the key ID and secret of an Orkes Conductor application
// access key, exchanged for a token sent with every request; leave them unset for open
// source, which has no token endpoint, or the exchange and every request after it fails.
func NewRuntime(cfg Config) *Runtime {
	apiClient := client.NewAPIClientFromEnv()
	return NewRuntimeWithClient(apiClient, cfg)
}

// NewRuntimeWithClient builds a Runtime over an existing APIClient, sharing its auth and
// connection settings, and is how to take credentials from outside the environment.
func NewRuntimeWithClient(apiClient *client.APIClient, cfg Config) *Runtime {
	return &Runtime{
		api:        apiClient,
		agents:     client.NewAgentClient(apiClient),
		metadata:   client.NewMetadataClient(apiClient),
		workflow:   client.NewWorkflowClient(apiClient),
		scheduler:  client.NewSchedulerClient(apiClient),
		runner:     worker.NewTaskRunnerWithApiClient(apiClient),
		config:     cfg,
		started:    map[workerKey]bool{},
		defs:       map[string]model.TaskDef{},
		registered: map[string]bool{},
	}
}

// AgentClient exposes the control plane for operations Runtime does not wrap.
func (r *Runtime) AgentClient() client.AgentClient { return r.agents }

// Plan compiles the agent into a Conductor workflow definition without running or
// registering anything, returning the workflow under "workflowDef" and the worker task
// names it needs under "requiredWorkers". Same /agent/compile call as Python's plan().
func (r *Runtime) Plan(ctx context.Context, agent *Agent) (map[string]any, error) {
	if err := agent.Validate(); err != nil {
		return nil, err
	}
	// A skill compiles from its raw document; as agentConfig the server would see an
	// ordinary agent and none of its scripts, resources or sub-agents.
	payload := map[string]any{"agentConfig": agent.toConfig()}
	if agent.skill != nil {
		payload = map[string]any{"framework": skillFramework, "rawConfig": agent.skill.rawConfig()}
	}
	plan, err := r.agents.Compile(ctx, payload)
	if err != nil {
		return nil, fmt.Errorf("compile agent %q: %w", agent.Name, err)
	}
	return plan, nil
}

// RunOption adjusts how Run and Start begin a run.
type RunOption func(*runOptions)

type runOptions struct {
	plan     *Plan
	media    []string
	settings *RunSettings
}

// RunSettings overrides the agent's model parameters for one run, without
// changing the stored agent. Apply it with WithRunSettings.
type RunSettings struct {
	Model                string
	Temperature          *float64
	MaxTokens            *int
	ReasoningEffort      ReasoningEffort
	ThinkingBudgetTokens *int
}

// configOverrides is the wire map merged onto the agentConfig, as Python's to_config_overrides.
func (rs *RunSettings) configOverrides() map[string]any {
	out := map[string]any{}
	if rs.Model != "" {
		out["model"] = rs.Model
	}
	if rs.Temperature != nil {
		out["temperature"] = *rs.Temperature
	}
	if rs.MaxTokens != nil {
		out["maxTokens"] = *rs.MaxTokens
	}
	if rs.ReasoningEffort != "" {
		out["reasoningEffort"] = string(rs.ReasoningEffort)
	}
	if rs.ThinkingBudgetTokens != nil {
		out["thinkingConfig"] = map[string]any{"enabled": true, "budgetTokens": *rs.ThinkingBudgetTokens}
	}
	return out
}

// WithPlan supplies the plan a StrategyPlanExecute agent carries out instead of one from
// its Planner. The server still requires the Planner slot, since the strategy compiles
// around it, but the planner never runs. The plan is validated first; see Plan.Validate.
func WithPlan(plan *Plan) RunOption {
	return func(o *runOptions) { o.plan = plan }
}

// WithMedia attaches media inputs, paths or URLs, as Python's run(..., media=[...]). The
// server reads paths itself, so a local one must sit under its allowed media directory.
func WithMedia(media ...string) RunOption {
	return func(o *runOptions) { o.media = append(o.media, media...) }
}

// WithRunSettings overrides the agent's model parameters for this run only.
func WithRunSettings(rs RunSettings) RunOption {
	return func(o *runOptions) { o.settings = &rs }
}

// startPayload validates, registers workers, and builds the /agent/start body Run and Start share.
func (r *Runtime) startPayload(agent *Agent, prompt string, opts []RunOption, runID string) (map[string]any, error) {
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
	// A stateful agent's workers poll a per-run domain that exists only once the run has
	// started, so the caller registers those; the rest poll first, so no call goes unheard.
	if runID == "" {
		if err := r.registerWorkers(agent, nil); err != nil {
			return nil, err
		}
	}

	// A skill travels under framework and rawConfig, not agentConfig, for the server's
	// SkillNormalizer — the same request the Python SDK sends.
	if agent.skill != nil {
		payload := map[string]any{
			"framework": skillFramework,
			"rawConfig": agent.skill.rawConfig(),
			"prompt":    prompt,
			"sessionId": "",
			"media":     mediaWire(o.media),
			"context":   map[string]any{},
		}
		if runID != "" {
			payload["runId"] = runID
		}
		return payload, nil
	}

	// Per-run settings mutate a copy of the agentConfig, needing no new server field, as in Python.
	config := agent.toConfig()
	if o.settings != nil {
		for k, v := range o.settings.configOverrides() {
			config[k] = v
		}
	}
	payload := map[string]any{
		"agentConfig": config,
		"prompt":      prompt,
		"sessionId":   "",
		"media":       mediaWire(o.media),
	}
	if runID != "" {
		// The server builds the run's task-to-domain map from this, routing a stateful run's tasks here.
		payload["runId"] = runID
	}
	if o.plan != nil {
		// The server reads workflow.input.static_plan ahead of the planner's output; the
		// key matches Java's AgentRequest and Python's runtime.run(plan=).
		payload["static_plan"] = o.plan.toPayload()
	}
	return payload, nil
}

// mediaWire renders the media list as a JSON array, empty rather than absent, per the server.
func mediaWire(media []string) []any {
	out := make([]any, 0, len(media))
	for _, m := range media {
		out = append(out, m)
	}
	return out
}

// Start begins a run and returns at once, with the agent's tool workers already polling.
func (r *Runtime) Start(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentHandle, error) {
	runID := newRunID(agent)
	payload, err := r.startPayload(agent, prompt, opts, runID)
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
	if err := r.startStatefulWorkers(ctx, agent, executionID, runID); err != nil {
		return nil, err
	}
	if err := r.registerTaskDefs(ctx); err != nil {
		return nil, err
	}
	return &AgentHandle{ExecutionID: executionID, rt: r}, nil
}

// Run starts an agent and blocks until it finishes, tool calls arriving as Conductor tasks.
func (r *Runtime) Run(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentResult, error) {
	runID := newRunID(agent)
	payload, err := r.startPayload(agent, prompt, opts, runID)
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
	if err := r.startStatefulWorkers(ctx, agent, executionID, runID); err != nil {
		return nil, err
	}
	if err := r.registerTaskDefs(ctx); err != nil {
		return nil, err
	}

	return r.awaitResult(ctx, executionID)
}

// awaitResult polls until the run reaches a terminal state. Not SSE: the APIClient reads a
// whole body before returning, so streaming needs its own request path — worth it for Stream.
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

// registerWorkers starts a worker for each tool with a Go handler, on this agent and every
// one nested under it, skipping server-dispatched and external tools; idempotent per name.
func (r *Runtime) registerWorkers(agent *Agent, domains map[string]string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.walkWorkers(agent, domains)
}

// workerKey identifies a poller: a task name in one domain. The same tool in
// two stateful runs is two pollers, as in the Python SDK.
type workerKey struct {
	name   string
	domain string
}

// walkWorkers registers one agent's workers, then those nested under it, sharing its routing.
func (r *Runtime) walkWorkers(a *Agent, domains map[string]string) error {
	if err := r.startWorkers(a.workerTools(), domains); err != nil {
		return err
	}
	for _, sub := range a.nestedAgents() {
		if err := r.walkWorkers(sub, domains); err != nil {
			return err
		}
	}
	return nil
}

// startWorkers starts a worker on the routed queue for each handler-backed tool not running.
func (r *Runtime) startWorkers(tools []ToolDef, domains map[string]string) error {
	for _, t := range tools {
		domain := domains[t.Name]
		key := workerKey{name: t.Name, domain: domain}
		if t.Handler == nil || r.started[key] {
			continue
		}
		fn, err := toolExecutor(t)
		if err != nil {
			return err
		}
		r.defs[t.Name] = t.taskDef()
		// A task definition carries no domain; only the poller does.
		if domain == "" {
			err = r.runner.StartWorker(t.Name, fn, r.config.batchSize(), r.config.workerPoll())
		} else {
			err = r.runner.StartWorkerWithDomain(
				t.Name, fn, r.config.batchSize(), r.config.workerPoll(), domain)
		}
		if err != nil {
			return fmt.Errorf("start worker %q: %w", t.Name, err)
		}
		r.started[key] = true
	}
	return nil
}

// workerTools lists every tool the server dispatches to a worker for this agent: the
// declared ones plus derived execution tools, on_condition handoffs, custom guardrails and
// skill scripts. Without a worker under each derived name, a queued task stalls the run.
func (a *Agent) workerTools() []ToolDef {
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
	tools = append(tools, a.handoffTools()...)
	// Each custom guardrail, agent- and tool-level alike, compiles to a SIMPLE task of its name.
	for _, g := range a.customGuardrails() {
		tools = append(tools, ToolDef{Name: g.Name, Handler: g.guardrailHandler()})
	}
	// Each set lifecycle callback is a worker named "<agent>_<position>".
	tools = append(tools, a.callbackTools()...)
	// The server emits worker tools for a skill's scripts and read_skill_file when normalizing.
	if a.skill != nil {
		tools = append(tools, a.skill.workers(a.Name)...)
	}
	// A Go gate is a worker under "<agent>_gate", the name in the gate document.
	if g, ok := a.Gate.(GateFunc); ok {
		tools = append(tools, ToolDef{Name: a.workerTaskName(gateSuffix), Handler: g.gateHandler()})
	}
	// A termination condition, stop-when predicate and router each become a task; see system_workers.go.
	tools = append(tools, a.systemWorkers()...)
	// Prefill tools are scheduled before the first turn whether or not they are in Tools.
	for _, p := range a.PrefillTools {
		if p.Tool.Handler != nil {
			tools = append(tools, p.Tool)
		}
	}
	return tools
}

// handoffTools returns a worker tool per on_condition handoff, which the server
// compiles into a SIMPLE task named "<agent>_handoff_<target>".
func (a *Agent) handoffTools() []ToolDef {
	conds := a.onConditions()
	if len(conds) == 0 {
		return nil
	}
	// The server reports the active agent as an index into [parent, sub-agents...] —
	// MultiAgentCompiler builds allSwarmAgents parent first — so the names go in that order.
	names := make([]string, 0, 1+len(a.Agents))
	names = append(names, a.Name)
	for _, sub := range a.Agents {
		if sub != nil {
			names = append(names, sub.Name)
		}
	}
	out := make([]ToolDef, 0, len(conds))
	for _, c := range conds {
		out = append(out, ToolDef{
			Name:    handoffTaskName(a.Name, c.Target),
			Handler: c.handoffHandler(names),
		})
	}
	return out
}

// nestedAgents lists the agents a run of this one also needs workers for: sub-agents, tool
// agents (each its own workflow, which lets a skill serve as a tool), router, planner, fallback.
func (a *Agent) nestedAgents() []*Agent {
	out := make([]*Agent, 0, len(a.Agents)+len(a.Tools)+3)
	for _, sub := range a.Agents {
		if sub != nil {
			out = append(out, sub)
		}
	}
	for _, t := range a.Tools {
		if sub, ok := t.Config["agent"].(*Agent); ok {
			out = append(out, sub)
		}
	}
	for _, slot := range []*Agent{a.Router, a.Planner, a.Fallback} {
		if slot != nil {
			out = append(out, slot)
		}
	}
	return out
}

// registerTaskDefs registers a task definition for every worker this runtime started, as
// Python does, after a run starts: compiling the agent makes the server write its own
// per-tool definition, so an earlier registration is overwritten and a tool's RetryCount,
// RetryDelaySeconds and RetryPolicy never take effect. Each definition repeats the tool's
// credential names as runtimeMetadata so it does not wipe what the server compiled.
func (r *Runtime) registerTaskDefs(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for name, def := range r.defs {
		if r.registered[name] {
			continue
		}
		if _, err := r.metadata.UpdateTaskDef(ctx, def); err != nil {
			if _, err := r.metadata.RegisterTaskDef(ctx, []model.TaskDef{def}); err != nil {
				return fmt.Errorf("register task definition %q: %w", name, err)
			}
		}
		r.registered[name] = true
	}
	return nil
}

// Deploy compiles and registers the agent without starting a run, returning its workflow
// name. As in Python, deploy once from a release step, then start runs by name, or Serve it.
func (r *Runtime) Deploy(ctx context.Context, agent *Agent) (string, error) {
	if err := agent.Validate(); err != nil {
		return "", err
	}
	var payload map[string]any
	if agent.skill != nil {
		payload = map[string]any{"framework": skillFramework, "rawConfig": agent.skill.rawConfig()}
	} else {
		payload = map[string]any{"agentConfig": agent.toConfig()}
	}
	out, err := r.agents.Deploy(ctx, payload)
	if err != nil {
		return "", fmt.Errorf("deploy agent %q: %w", agent.Name, err)
	}
	name, ok := out["agentName"].(string)
	if !ok || name == "" {
		name = agent.Name
	}
	return name, nil
}

// Serve deploys each agent, registers its definitions and starts its workers, then blocks
// until ctx is cancelled so another process can run these agents by name; Python's serve.
func (r *Runtime) Serve(ctx context.Context, agents ...*Agent) error {
	if len(agents) == 0 {
		return fmt.Errorf("Serve requires at least one agent")
	}
	for _, agent := range agents {
		if err := agent.Validate(); err != nil {
			return err
		}
		// Deploy first, then override the task definitions compiling wrote, as Run does.
		if _, err := r.Deploy(ctx, agent); err != nil {
			return err
		}
		// A standing Serve has no run of its own, so it polls the domainless queue.
		if err := r.registerWorkers(agent, nil); err != nil {
			return err
		}
		if err := r.registerTaskDefs(ctx); err != nil {
			return err
		}
	}
	<-ctx.Done()
	r.Shutdown()
	return ctx.Err()
}

// Signal injects a persistent signal the agent prepends to its next LLM turn; it persists
// until overwritten, an empty message clears it, and unlike SendMessage it works on any agent.
func (r *Runtime) Signal(ctx context.Context, executionID, message string) error {
	return r.agents.Signal(ctx, executionID, message)
}

// SendMessage pushes a message into a running execution's workflow message queue, for an
// agent waiting on wait_for_message. A non-map value is wrapped as {"message": value}, as in Python.
func (r *Runtime) SendMessage(ctx context.Context, executionID string, message any) error {
	body, ok := message.(map[string]any)
	if !ok {
		body = map[string]any{"message": message}
	}
	return r.agents.SendMessage(ctx, executionID, body)
}

// Pause suspends a running execution, keeping its state so Resume continues from there.
func (r *Runtime) Pause(ctx context.Context, executionID string) error {
	if _, err := r.workflow.Pause(ctx, executionID); err != nil {
		return fmt.Errorf("pause %s: %w", executionID, err)
	}
	return nil
}

// Resume continues a paused execution, the inverse of Pause. It is not Python's
// runtime.resume(id, agent), which re-registers workers for a run another process started:
// that waits on routing stateful agents to per-execution worker domains, since without
// those there is no domain to re-attach to, and Serve covers the domainless fleet case.
func (r *Runtime) Resume(ctx context.Context, executionID string) error {
	if _, err := r.workflow.Resume(ctx, executionID); err != nil {
		return fmt.Errorf("resume %s: %w", executionID, err)
	}
	return nil
}

// Shutdown stops every worker this runtime started.
func (r *Runtime) Shutdown() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for key := range r.started {
		r.runner.Shutdown(key.name)
		delete(r.started, key)
	}
}
