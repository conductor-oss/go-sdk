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
	api       *client.APIClient
	agents    client.AgentClient
	metadata  client.MetadataClient
	workflow  client.WorkflowClient
	scheduler client.SchedulerClient
	runner    *worker.TaskRunner
	config    Config
	mu        sync.Mutex
	// started records which (task name, domain) pairs already have a worker.
	// A stateful run polls its own domain, so one task name can have several.
	started map[workerKey]bool
	// defs are the task definitions for started workers, registered after a
	// run starts; registered records which ones have been sent.
	defs       map[string]model.TaskDef
	registered map[string]bool
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

// Plan compiles the agent into a Conductor workflow definition without
// running it, and returns the server's response: the workflow under
// "workflowDef" and the worker task names it needs under "requiredWorkers".
//
// It is the Python SDK's runtime.plan(): the same agentConfig document goes
// to /agent/compile, and nothing is registered or started. Use it to inspect
// what an agent compiles to, or to compare compiled output across SDKs.
func (r *Runtime) Plan(ctx context.Context, agent *Agent) (map[string]any, error) {
	if err := agent.Validate(); err != nil {
		return nil, err
	}
	// A skill is compiled from its raw document, the same shape a run sends;
	// as agentConfig the server would see an ordinary agent and none of the
	// skill's scripts, resources or sub-agents.
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
// changing the stored agent. It is the counterpart of the Python SDK's
// RunSettings: the set fields are merged into the agentConfig sent to the
// server before the run starts. Apply it with WithRunSettings.
type RunSettings struct {
	Model                string
	Temperature          *float64
	MaxTokens            *int
	ReasoningEffort      ReasoningEffort
	ThinkingBudgetTokens *int
}

// configOverrides is the wire map merged onto the agentConfig, field for field
// with Python's RunSettings.to_config_overrides.
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

// WithMedia attaches media inputs to the run, such as image or document paths
// the server can read, or URLs. It is the counterpart of the Python SDK's
// run(..., media=[...]). The server reads the paths itself, so a local path
// must sit under the server's allowed media directory.
func WithMedia(media ...string) RunOption {
	return func(o *runOptions) { o.media = append(o.media, media...) }
}

// WithRunSettings overrides the agent's model parameters for this run only.
func WithRunSettings(rs RunSettings) RunOption {
	return func(o *runOptions) { o.settings = &rs }
}

// startPayload validates the agent and any options, registers the agent's
// workers, and builds the body of the /agent/start request. Run and Start
// share it so a plan reaches the server the same way from either.
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
	// A stateful agent's workers poll a per-run domain, which only exists
	// once the run has started, so those are registered afterwards by the
	// caller. Everything else starts polling before the run does, so a tool
	// call cannot arrive before something is listening.
	if runID == "" {
		if err := r.registerWorkers(agent, nil); err != nil {
			return nil, err
		}
	}

	// A skill does not travel as agentConfig. The server's SkillNormalizer
	// compiles the raw document, which /agent/start accepts under
	// framework and rawConfig — the same request the Python SDK sends.
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

	// Per-run settings mutate a copy of the agentConfig before it is sent, so
	// they reach the LLM tasks without a new server field, as in Python.
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
		// The server builds the run's task-to-domain map from this, which is
		// how a stateful run's tasks reach this process's own workers.
		payload["runId"] = runID
	}
	if o.plan != nil {
		// The server reads workflow.input.static_plan ahead of the planner's
		// output; the key matches AgentRequest in Java and runtime.run(plan=)
		// in Python.
		payload["static_plan"] = o.plan.toPayload()
	}
	return payload, nil
}

// mediaWire renders the media list for the request, always as a JSON array
// even when empty, which is what the server expects.
func mediaWire(media []string) []any {
	out := make([]any, 0, len(media))
	for _, m := range media {
		out = append(out, m)
	}
	return out
}

// Start begins a run and returns at once.
//
// Workers for the agent's tools are registered before the run starts, so a tool
// call cannot arrive before something is polling for it.
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

// Run starts an agent and blocks until the execution finishes.
//
// It validates the agent, registers a worker for every tool that has a Go
// handler, starts the run, then polls until the server reports a terminal
// state. Tool calls arrive as Conductor tasks while this is waiting.
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
// handler, on this agent and on every agent nested under it. Tools the server
// dispatches itself, and external tools served elsewhere, have no handler and
// are skipped.
//
// Registration is idempotent per task name so repeated runs of the same agent
// do not stack workers.
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

// walkWorkers registers one agent's workers, then those of the agents nested
// under it.
func (r *Runtime) walkWorkers(a *Agent, domains map[string]string) error {
	if err := r.startWorkers(a.workerTools(), domains); err != nil {
		return err
	}
	// Everything nested under the agent is part of the same run, so it reads
	// the same routing.
	for _, sub := range a.nestedAgents() {
		if err := r.walkWorkers(sub, domains); err != nil {
			return err
		}
	}
	return nil
}

// startWorkers starts a worker for each handler-backed tool not already
// running.
func (r *Runtime) startWorkers(tools []ToolDef, domains map[string]string) error {
	for _, t := range tools {
		// Each worker polls the queue the server routes its own task to.
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

// workerTools lists every tool of this agent the server dispatches to a
// worker: the declared tools plus the derived ones. The derived execution
// tools, on_condition handoffs, custom guardrails, and skill scripts are all
// serialized as worker tools, so without a worker under each derived name the
// server queues a task nothing polls for and the run stalls.
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
	// The server compiles each custom guardrail into a SIMPLE task named after
	// the guardrail, agent-level and tool-level alike.
	for _, g := range a.customGuardrails() {
		tools = append(tools, ToolDef{Name: g.Name, Handler: g.guardrailHandler()})
	}
	// Each set lifecycle callback is a worker named "<agent>_<position>".
	tools = append(tools, a.callbackTools()...)
	// A skill's scripts and its read_skill_file tool run here too; the server
	// emits worker tools under these names when it normalizes the skill
	// document.
	if a.skill != nil {
		tools = append(tools, a.skill.workers(a.Name)...)
	}
	// A Go gate is a worker under "<agent>_gate", the name the serializer put
	// in the gate document.
	if g, ok := a.Gate.(GateFunc); ok {
		tools = append(tools, ToolDef{Name: a.workerTaskName(gateSuffix), Handler: g.gateHandler()})
	}
	// A termination condition, a stop-when predicate and a router function
	// each become a task of their own; see system_workers.go.
	tools = append(tools, a.systemWorkers()...)
	// Prefill tools are scheduled by the server before the first turn whether
	// or not they are also in Tools, so their workers start too.
	for _, p := range a.PrefillTools {
		if p.Tool.Handler != nil {
			tools = append(tools, p.Tool)
		}
	}
	return tools
}

// handoffTools returns a worker tool per on_condition handoff. The server
// compiles each into a SIMPLE task named "<agent>_handoff_<target>".
func (a *Agent) handoffTools() []ToolDef {
	conds := a.onConditions()
	if len(conds) == 0 {
		return nil
	}
	// The server reports the active agent as an index into
	// [parent, sub-agents...] — MultiAgentCompiler builds allSwarmAgents with
	// the parent first — so index 0 is this agent, not its first child. The
	// names go in that order so the handler resolves the index the same way.
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

// nestedAgents lists the agents whose workers a run of this agent also needs:
// sub-agents, agents exposed as tools (each runs as its own workflow, which
// is what lets a skill serve as a tool), and the router, planner, and
// fallback slots.
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

// registerTaskDefs registers a task definition for every worker this runtime
// has started, as the Python SDK does for every worker it hosts. It runs after
// a run starts, because compiling the agent makes the server write its own
// definition for each tool, with its own retry and timeout settings; a
// registration before that would be overwritten, and a tool's RetryCount,
// RetryDelaySeconds and RetryPolicy would never take effect. Each definition
// carries the tool's credential names as runtimeMetadata, so it does not wipe
// what the server compiled there. An existing definition is updated in place;
// a missing one is created. Each name is registered once per runtime.
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

// Deploy compiles and registers the agent on the server without starting a
// run, and returns the workflow name it was registered under. It is the
// counterpart of the Python SDK's runtime.deploy for a single agent: deploy
// once from a release step, then start runs against the stored agent by name,
// or bring up workers for it with Serve.
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

// Serve hosts the workers the given agents' tools need and blocks until ctx is
// cancelled. It deploys each agent, registers its task definitions and starts
// its workers, so a separate process can start runs against these agents by
// name while this one answers their tool calls. It is the counterpart of the
// Python SDK's runtime.serve; a caller wanting it non-blocking runs it in a
// goroutine and cancels ctx to stop.
func (r *Runtime) Serve(ctx context.Context, agents ...*Agent) error {
	if len(agents) == 0 {
		return fmt.Errorf("Serve requires at least one agent")
	}
	for _, agent := range agents {
		if err := agent.Validate(); err != nil {
			return err
		}
		// Deploy first: compiling the agent makes the server write its own task
		// definitions, so register ours afterwards to override them, the same
		// order Run uses around a start.
		if _, err := r.Deploy(ctx, agent); err != nil {
			return err
		}
		// A standing Serve has no run of its own, so it polls the
		// domainless queue; a stateful run's own workers are started by
		// whoever starts that run.
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

// Signal injects a persistent signal into a running execution's context. The
// agent prepends it to the next LLM turn; it persists until overwritten, and
// an empty message clears it. This works on any agent, unlike SendMessage.
func (r *Runtime) Signal(ctx context.Context, executionID, message string) error {
	return r.agents.Signal(ctx, executionID, message)
}

// SendMessage pushes a message into a running execution's workflow message
// queue, for an agent waiting on a wait_for_message tool. A non-map value is
// wrapped as {"message": value}, matching the Python SDK, so the agent
// receives it under the message key.
func (r *Runtime) SendMessage(ctx context.Context, executionID string, message any) error {
	body, ok := message.(map[string]any)
	if !ok {
		body = map[string]any{"message": message}
	}
	return r.agents.SendMessage(ctx, executionID, body)
}

// Pause suspends a running execution. It stops advancing but keeps its state,
// so Resume continues it from where it paused.
func (r *Runtime) Pause(ctx context.Context, executionID string) error {
	if _, err := r.workflow.Pause(ctx, executionID); err != nil {
		return fmt.Errorf("pause %s: %w", executionID, err)
	}
	return nil
}

// Resume continues a paused execution, the inverse of Pause.
//
// Note this is not "resume from an instance" — re-registering workers for a
// run started by another process, which the Python SDK spells
// runtime.resume(id, agent). That case is deferred until the Go runtime
// routes stateful agents to per-execution worker domains, since without that
// its main use, re-attaching to a specific execution's domain, cannot be
// exercised; a standing Serve already covers the domainless fleet case.
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
