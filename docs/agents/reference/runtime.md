# Runtime reference

`ai.Runtime` is the entry point of `sdk/ai`: it starts agents, hosts the Conductor workers their Go tools need, deploys agents, and manages their schedules.

```go
rt := ai.NewRuntime(ai.Config{})
defer rt.Shutdown()

result, err := rt.Run(context.Background(), agent, "What is the capital of France?")
```

One `Runtime` serves many agents from a single `TaskRunner`; a worker per task name is reused across runs.

## Construction

| Constructor | Description |
|---|---|
| `func NewRuntime(cfg Config) *Runtime` | Builds an `APIClient` from the environment (`client.NewAPIClientFromEnv`). |
| `func NewRuntimeWithClient(apiClient *client.APIClient, cfg Config) *Runtime` | Reuses an existing `*client.APIClient`, sharing its auth, retry and connection settings. The way to take credentials from outside the environment. |

### Environment

Read by `client.NewAPIClientFromEnv`, not by `Runtime` itself.

| Variable | Default | Description |
|---|---|---|
| `CONDUCTOR_SERVER_URL` | `http://localhost:8080/api` | Server API base URL. All an unauthenticated open-source Conductor needs. |
| `CONDUCTOR_AUTH_KEY` | _(unset)_ | Orkes application access key ID. Leave unset for open source, which has no token endpoint. |
| `CONDUCTOR_AUTH_SECRET` | _(unset)_ | Access key secret; exchanged with the key for a token sent on every request. |
| `CONDUCTOR_CLIENT_HTTP_TIMEOUT` | `30` | HTTP timeout, in seconds. |

Proxy and TLS variables are documented in [proxy configuration](../../api_client/proxy_configuration.md) and [TLS configuration](../../api_client/tls_configuration.md).

## Config

`ai.Config` tunes worker polling only; the server URL and auth live on the `APIClient`. The zero value is usable.

| Field | Type | Default | Description |
|---|---|---|---|
| `WorkerPollInterval` | `time.Duration` | `100ms` | How often tool workers poll for tasks. |
| `WorkerBatchSize` | `int` | `1` | How many tasks a worker takes per poll. |
| `StatusPollInterval` | `time.Duration` | `500ms` | How often `Run` and `AgentHandle.Result` check for completion. |

Zero or negative values take the default.

## Methods

### Running

| Signature | Description |
|---|---|
| `func (r *Runtime) Run(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentResult, error)` | Validates the agent, starts its workers, starts the run, and blocks until it is terminal. A `FAILED` run returns both the result and an error. |
| `func (r *Runtime) Start(ctx context.Context, agent *Agent, prompt string, opts ...RunOption) (*AgentHandle, error)` | Same setup, but returns as soon as the server reports an `executionId`, with the agent's workers already polling. |
| `func (r *Runtime) Plan(ctx context.Context, agent *Agent) (map[string]any, error)` | Compiles the agent without registering or running anything. The map holds `workflowDef` and `requiredWorkers`. `POST /agent/compile`. |

### Deployment

| Signature | Description |
|---|---|
| `func (r *Runtime) Deploy(ctx context.Context, agent *Agent) (string, error)` | Compiles and registers the agent without starting a run; returns the server's `agentName`, falling back to `agent.Name`. Starts no workers. `POST /agent/deploy`. |
| `func (r *Runtime) Serve(ctx context.Context, agents ...*Agent) error` | Deploys each agent, registers its task definitions and starts its workers on the domainless queue, then blocks until `ctx` is cancelled, calls `Shutdown` and returns `ctx.Err()`. Errors with no agents. |
| `func (r *Runtime) Shutdown()` | Stops every worker this runtime started. |

`Deploy` then `Serve` is how another process, or a schedule, starts runs by agent name; see [deploy, serve, run](../concepts/deploy-serve-run.md) and [`client.md`](./client.md) for starting a deployed agent by name.

### Steering a run

| Signature | Description |
|---|---|
| `func (r *Runtime) Signal(ctx context.Context, executionID, message string) error` | Sets a persistent signal the agent prepends to its next LLM turn. It persists until overwritten; an empty message clears it. Works on any agent. `POST /agent/{id}/signal`. |
| `func (r *Runtime) SendMessage(ctx context.Context, executionID string, message any) error` | Pushes a message into the execution's workflow message queue, for an agent waiting on `tool.WaitForMessage`. A non-map value is wrapped as `{"message": value}`. |
| `func (r *Runtime) Pause(ctx context.Context, executionID string) error` | Suspends the execution, keeping its state. |
| `func (r *Runtime) Resume(ctx context.Context, executionID string) error` | Continues a paused execution. This is the inverse of `Pause`, not Python's `resume(id, agent)` re-attach. |

### Control plane

| Signature | Description |
|---|---|
| `func (r *Runtime) AgentClient() client.AgentClient` | The lower-level `/agent` control plane, for operations `Runtime` does not wrap. See [`client.md`](./client.md). |

### Schedules

Every schedule method takes the agent's name; the short schedule name is stored server-side as `<agentName>-<name>`.

| Signature | Description |
|---|---|
| `func (r *Runtime) SaveSchedule(ctx context.Context, agentName string, s Schedule) error` | Creates or updates one schedule. Validates `s` first. |
| `func (r *Runtime) GetSchedule(ctx context.Context, agentName, name string) (*Schedule, error)` | Reads one schedule by short name. |
| `func (r *Runtime) ListSchedules(ctx context.Context, agentName string) ([]Schedule, error)` | Every schedule whose stored name carries this agent's prefix. |
| `func (r *Runtime) DeleteSchedule(ctx context.Context, agentName, name string) error` | Deletes one schedule. |
| `func (r *Runtime) PauseSchedule(ctx context.Context, agentName, name string) error` | Pauses one schedule. |
| `func (r *Runtime) ResumeSchedule(ctx context.Context, agentName, name string) error` | Resumes a paused schedule. |
| `func (r *Runtime) ReconcileSchedules(ctx context.Context, agentName string, desired []Schedule) error` | Makes the agent's schedules match `desired` exactly, deleting any other schedule of this agent. `nil` is a no-op; an empty slice deletes all. Rejects duplicate names. |

Schedules run a deployed agent by name, so `Deploy` the agent first. See [scheduling](../concepts/scheduling.md).

## RunOption

Applies to both `Run` and `Start`.

| Option | Description |
|---|---|
| `func WithPlan(plan *Plan) RunOption` | Supplies the plan a `StrategyPlanExecute` agent carries out instead of one from its `Planner`. Errors unless `agent.Strategy == StrategyPlanExecute`; the plan is validated, then sent as `static_plan`. The `Planner` slot is still required, but the planner never runs. |
| `func WithMedia(media ...string) RunOption` | Attaches media inputs as paths or URLs. The server reads paths itself, so a local path must sit under its allowed media directory. Repeatable. |
| `func WithRunSettings(rs RunSettings) RunOption` | Overrides the agent's model parameters for this run only, by merging onto a copy of the `agentConfig`. The stored agent is unchanged. |

An agent loaded with `ai.LoadSkill` travels as `framework` plus `rawConfig`, so `WithPlan` and `WithRunSettings` do not reach the server on that path; `WithMedia` does.

### RunSettings

| Field | Type | Wire key | Description |
|---|---|---|---|
| `Model` | `string` | `model` | Model for this run, `provider/model`. Omitted when empty. |
| `Temperature` | `*float64` | `temperature` | Sampling temperature. Pointer so `Ptr(0)` is sent. |
| `MaxTokens` | `*int` | `maxTokens` | Cap on generated tokens per LLM call. |
| `ReasoningEffort` | `ReasoningEffort` | `reasoningEffort` | `minimal`, `low`, `medium`, `high`. Omitted when empty. |
| `ThinkingBudgetTokens` | `*int` | `thinkingConfig` | Expands to `{"enabled": true, "budgetTokens": n}`. |

## AgentResult

Returned by `Run`, `AgentHandle.Result` and `AgentHandle.Status`.

| Field | Type | Description |
|---|---|---|
| `ExecutionID` | `string` | Conductor workflow ID. |
| `Status` | `Status` | Execution state; see below. |
| `Output` | `string` | The agent's final answer. |
| `FinishReason` | `string` | Why the loop stopped, e.g. `STOP`, `MAX_TURNS`. |
| `Error` | `string` | Failure reason when `Status` is `StatusFailed`. |
| `TokenUsage` | `TokenUsage` | `PromptTokens`, `CompletionTokens`, `TotalTokens`; zero when the server does not report them. |
| `Raw` | `map[string]any` | The untouched status document. |

`func (r *AgentResult) PrintResult()` writes the result to stdout in a boxed layout, the counterpart of Python's `print_result`.

### Status

`StatusRunning`, `StatusCompleted`, `StatusFailed`, `StatusTerminated`, `StatusTimedOut`, `StatusPaused`, `StatusWaiting` (blocked on a human).

`func (s Status) Terminal() bool` is true for `COMPLETED`, `FAILED`, `TERMINATED` and `TIMED_OUT`.

## AgentHandle

Returned by `Start`. Its only exported field is `ExecutionID string`. For the narrative — streaming loops, approvals on nested executions — see [streaming and human input](../concepts/streaming-hitl.md).

| Signature | Description |
|---|---|
| `func (h *AgentHandle) Events(ctx context.Context) (<-chan Event, error)` | Streams updates over SSE (`GET /agent/stream/{id}`) until the run ends or `ctx` is cancelled, closing the channel with the stream. An already finished run yields no events. |
| `func (h *AgentHandle) Status(ctx context.Context) (*AgentResult, error)` | Current state without waiting. |
| `func (h *AgentHandle) Result(ctx context.Context) (*AgentResult, error)` | Blocks until the run is terminal, polling every `Config.StatusPollInterval`. |
| `func (h *AgentHandle) Waiting(ctx context.Context) (bool, error)` | Whether the run is blocked on a human. |
| `func (h *AgentHandle) For(ev Event) (*AgentHandle, error)` | A handle on the execution an event came from, so a nested agent's request for input is answered on its own execution. Errors when the event names no execution. |
| `func (h *AgentHandle) Respond(ctx context.Context, output map[string]any) error` | Answers a run waiting on a human. Waits until the server itself reports the run waiting, because the `waiting` event arrives slightly before the server accepts a response. Errors on a run that has already finished. |
| `func (h *AgentHandle) Approve(ctx context.Context) error` | `Respond` with `{"approved": true}`. |
| `func (h *AgentHandle) Reject(ctx context.Context, reason string) error` | `Respond` with `{"approved": false, "reason": reason}`. |
| `func (h *AgentHandle) Stop(ctx context.Context) error` | Terminates the run. |
| `func (h *AgentHandle) Signal(ctx context.Context, message string) error` | `Runtime.Signal` on this execution. |
| `func (h *AgentHandle) SendMessage(ctx context.Context, message any) error` | `Runtime.SendMessage` on this execution. |
| `func (h *AgentHandle) Pause(ctx context.Context) error` | `Runtime.Pause` on this execution. |
| `func (h *AgentHandle) Resume(ctx context.Context) error` | `Runtime.Resume` on this execution. |

### Event

| Field | Type | Description |
|---|---|---|
| `Type` | `EventType` | Normalized event kind, a direct conversion of the wire name. |
| `Name` | `string` | The server's raw event name, kept so an unmapped name is not lost. |
| `Text` | `string` | Streamed output or message text, taken from `text`, `content`, `message`, `result` or `delta`. |
| `Data` | `map[string]any` | The decoded payload. |
| `ExecutionID` | `string` | The event's own execution; a nested agent's is a sub-execution. Pass the event to `For`. |

`EventType` values: `EventThinking`, `EventToolCall`, `EventToolResult`, `EventHandoff`, `EventWaiting`, `EventMessage`, `EventError`, `EventDone`, `EventGuardrailPass`, `EventGuardrailFail`.

## Schedule

| Field | Type | Required | Default | Description |
|---|---|---|---|---|
| `Name` | `string` | yes | — | Short name, unique within the agent. Stored as `<agentName>-<Name>`. |
| `Cron` | `string` | yes | — | Cron expression, e.g. `0 0 * * *`. |
| `Timezone` | `string` | no | `UTC` | Zone the cron runs in. |
| `Input` | `map[string]any` | no | `{}` | Input the scheduled run starts the agent with. |
| `Catchup` | `bool` | no | `false` | Run occurrences missed while the scheduler was down. |
| `Paused` | `bool` | no | `false` | Create the schedule paused. |
| `StartAt` | `int64` | no | `0` | Epoch millis the schedule becomes active; zero is unbounded. |
| `EndAt` | `int64` | no | `0` | Epoch millis it stops; zero is unbounded. Must be after `StartAt` when both are set. |
| `Description` | `string` | no | — | Free text. |

`func (s Schedule) Validate() error` reports the first problem: missing name, missing cron, or `StartAt >= EndAt`.
