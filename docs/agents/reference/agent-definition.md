# Agent definition reference

Every exported field of `ai.Agent` and `ai.ToolDef`, the two structs that make up the agent document the server compiles into a Conductor workflow.

`Agent` is a configuration document, not a constructor call: set the fields you need and leave the rest at their zero value, which is what "unset" means on the wire.

```go
agent := &ai.Agent{
    Name:         "weather_bot",
    Model:        "openai/gpt-4o",
    Instructions: "Use tools to answer questions.",
    Tools:        ai.Tools(tool.Func("get_weather", getWeather, "Current weather for a city.")),
}
```

## Pointer fields and `ai.Ptr`

Optional numeric and boolean settings are pointers because zero is a meaningful value: `Temperature: ai.Ptr(0.0)` is sent as `0`, while a nil `Temperature` is omitted and the model's own default applies. `func Ptr[T any](v T) *T` exists to set them inline.

Pointer fields: `Synthesize`, `MaxTokens`, `Temperature`, `ThinkingBudgetTokens`, `ContextWindowBudget` on `Agent`; `TimeoutSeconds`, `MaxCalls`, `RetryCount`, `RetryDelaySeconds` on `ToolDef`; `Enabled` on `CodeExecutionConfig` and `CLIConfig`; `Version` on `PromptTemplate`.

## Agent

### Identity and prompt

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Name` | `string` | yes | — | `name` | Becomes the Conductor workflow name. Must match `^[a-zA-Z_][a-zA-Z0-9_-]*$`. |
| `Model` | `string` | no | — | `model` | `provider/model`, e.g. `openai/gpt-4o`. Omitted when empty, so a sub-agent inherits its parent's model at compile time. |
| `BaseURL` | `string` | no | server's configured URL | `baseUrl` | Points this agent's provider at another endpoint, such as a proxy. |
| `Instructions` | `string` | no | — | `instructions` | The system prompt. Mutually exclusive with `InstructionsTemplate`. |
| `InstructionsTemplate` | `*PromptTemplate` | no | — | `instructions` | Uses a server-stored prompt template instead of a literal string. Serialized as `{"type": "prompt_template", "name": ..., "variables"?, "version"?}`. |
| `Introduction` | `string` | no | — | `introduction` | The agent's opening message to a user. |
| `Metadata` | `map[string]any` | no | — | `metadata` | Arbitrary data carried with the definition. |

`PromptTemplate`: `Name string` (required), `Variables map[string]any` (values may be expressions such as `${workflow.input.user_tier}`), `Version *int` (nil is latest). The SDK never creates templates.

### Tools

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Tools` | `[]ToolDef` | no | — | `tools` | Tools the agent may call. Build them with `sdk/ai/tool`; wrap with `ai.Tools(...)`. |
| `RequiredTools` | `[]string` | no | — | `requiredTools` | Tools the run should call before finishing. Kept for parity but unusable on current servers: it compiles to nested `DO_WHILE`s and deadlocks. |
| `PrefillTools` | `[]PrefillToolCall` | no | — | `prefillTools` | Tools run with fixed arguments before the first LLM turn, their results opening the conversation. Build with `ai.Prefill(tool, args)`. A prefill tool need not appear in `Tools`; its worker is registered either way. |
| `Credentials` | `[]string` | no | — | `credentials` | Secret names available to every tool on this agent, as a fallback for tools that cannot name their own. Prefer `tool.WithCredentials`. |
| `CodeExecution` | `*CodeExecutionConfig` | no | — | `codeExecution` | Adds a derived tool `{agent}_execute_code` that runs code the model writes. |
| `CLI` | `*CLIConfig` | no | — | `cliConfig` | Adds a derived tool `{agent}_run_command` that runs shell commands. |

`CodeExecutionConfig`: `Enabled *bool` (nil is true), `AllowedLanguages []string` (empty is `["python"]`), `AllowedCommands []string` (empty is no restriction), `TimeoutSeconds int` (zero is 30), `Executor CodeExecutor` (nil is `LocalExecutor`; the executor is local and never travels to the server).

`CLIConfig`: `Enabled *bool` (nil is true), `AllowedCommands []string`, `TimeoutSeconds int` (zero is 30), `AllowShell bool` (false).

### LLM parameters

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `MaxTokens` | `*int` | no | model's own | `maxTokens` | Cap on tokens generated per LLM call. |
| `Temperature` | `*float64` | no | model's own | `temperature` | Sampling temperature. |
| `ReasoningEffort` | `ReasoningEffort` | no | — | `reasoningEffort` | `ReasoningEffortMinimal`, `Low`, `Medium`, `High`. OpenAI reasoning models only; others ignore it. Validated against the shared schema's enum. |
| `ThinkingBudgetTokens` | `*int` | no | — | `thinkingConfig` | Enables extended thinking; expands to `{"enabled": true, "budgetTokens": n}`. |
| `ContextWindowBudget` | `*int` | no | server's | `contextWindowBudget` | Token threshold for proactive context condensing. |
| `IncludeContents` | `string` | no | inherit | `includeContents` | `"none"` starts a sub-agent with fresh context; empty inherits. |
| `Memory` | `*ConversationMemory` | no | — | `memory` | Seeds prior conversation. Non-nil means stateful memory, so an empty value still sends `"memory": {}`. |
| `OutputType` | `any` | no | — | `outputType` | Constrains the final answer to a struct's shape: pass a zero value, as in `OutputType: Ticket{}`. The schema comes from its `json` tags. Serialized as `{"schema": ..., "className": ...}`. |

`ConversationMemory`: `Messages []map[string]any` (prior turns, each a map of `role` to a role name and `message` to the text), `MaxMessages int` (zero omits the cap).

### Limits

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `MaxTurns` | `int` | no | `25` | `maxTurns` | Bounds the agent loop. Always sent; zero is serialized as 25. Negative is rejected. |
| `TimeoutSeconds` | `int` | no | `0` | `timeoutSeconds` | Bounds the whole execution. Always sent; zero means no explicit limit. |

### Multi-agent

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Agents` | `[]*Agent` | no | — | `agents` | Sub-agents. A non-empty value makes this a multi-agent system and is what causes `strategy` to be sent. |
| `Strategy` | `Strategy` | no | `StrategyHandoff` | `strategy` | How sub-agents are orchestrated. Emitted only when the agent has sub-agents (`Agents`, `Planner` or `Fallback`). |
| `Router` | `*Agent` | with `StrategyRouter` | — | `router` | Sub-agent-picking agent, run by the server as a nested agent document. |
| `RouterFunc` | `RouterFunc` | with `StrategyRouter` | — | `router` | Picks a sub-agent in Go, run as a worker named `<agent>_router_fn`. Mutually exclusive with `Router`; `StrategyRouter` requires one of the two. |
| `Handoffs` | `[]HandoffCondition` | no | — | `handoffs` | Conditions that move control between sub-agents, usually with `StrategySwarm`. |
| `AllowedTransitions` | `map[string][]string` | no | no restriction | `allowedTransitions` | Sub-agent name to the targets it may hand off to. |
| `Synthesize` | `*bool` | no | `true` | `synthesize` | The final LLM step that combines specialists' results. Only the opt-out, `Ptr(false)`, travels. |

`Strategy` values: `StrategyHandoff`, `StrategySequential`, `StrategyParallel`, `StrategyRouter`, `StrategyRoundRobin`, `StrategyRandom`, `StrategySwarm`, `StrategyManual`, `StrategyPlanExecute`. `HandoffCondition` implementations: `OnCondition` (a Go predicate, run as a worker `<agent>_handoff_<target>`, so one condition per target), `OnTextMention`, `OnToolResult`. See [multi-agent](../concepts/multi-agent.md).

### Plan and execute

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Planner` | `*Agent` | with `StrategyPlanExecute` | — | `planner` | Emits a plan over the parent's `Tools`. A full nested agent document, so it may have its own tools. |
| `Fallback` | `*Agent` | no | — | `fallback` | Runs when the plan cannot be carried out. |
| `FallbackMaxTurns` | `int` | no | server's | `fallbackMaxTurns` | Bounds the fallback agent; zero omits the key. |
| `PlannerContext` | `[]PlanContext` | no | — | `plannerContext` | Extra material handed to the planner. Rejected unless `Strategy` is `StrategyPlanExecute`. |
| `PlanSource` | `map[string]any` | no | — | `planSource` | Supplies the plan from a server-evaluated expression instead of the planner. Sent as written; a non-nil empty map is still sent. |
| `EnablePlanning` | `bool` | no | `false` | `enablePlanning` | The server appends a fixed "plan first, then execute step by step" paragraph to this agent's instructions. Prompt text only; not plan-execute, and unrelated to `Planner`. |

`PlanContext`: exactly one of `Text string` or `URL string`, plus `Headers map[string]string` (sent when fetching), `Optional bool` (planning proceeds when the fetch fails), `MaxBytes int` (zero is the server default).

A ready-made `ai.Plan` may be passed per run with `ai.WithPlan`, skipping the planner LLM; see [`runtime.md`](./runtime.md).

### Control and safety

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Guardrails` | `[]Guardrail` | no | — | `guardrails` | Checks on this agent's input or output. Implementations: `RegexGuardrail`, `LLMGuardrail` (both server-side), `CustomGuardrail` (a Go function run as a worker named after the guardrail). |
| `Termination` | `TerminationCondition` | no | — | `termination` | Stops the loop on a server-evaluated condition tree: `MaxMessageTermination`, `TextMentionTermination`, `StopMessageTermination`, `TokenUsageTermination`, combined with `AndTermination`/`OrTermination`. |
| `StopWhen` | `StopWhenFunc` | no | — | `stopWhen` | Stops the loop on Go logic, run as a worker named `<agent>_stop_when`. Independent of `Termination`. |
| `Gate` | `GateCondition` | no | — | `gate` | Decides, after this agent finishes inside a sequential pipeline, whether the pipeline continues. `TextGate` is evaluated server-side; `GateFunc` runs as a worker named `<agent>_gate`. |
| `Callbacks` | `*Callbacks` | no | — | `callbacks` | Lifecycle hooks. Each set hook runs as a worker named `<agent>_<position>`. |
| `MaskedFields` | `[]string` | no | — | `maskedFields` | Input and output fields the server redacts from execution history and the UI. |

`Callbacks` fields and their wire positions: `OnAgentStart` → `before_agent`, `OnAgentEnd` → `after_agent`, `OnModelStart` → `before_model`, `OnModelEnd` → `after_model`, `OnToolStart` → `before_tool`, `OnToolEnd` → `after_tool`. A hook returns a non-empty map to override what happens next, or nil to continue; an error counts as nil, so a broken hook never blocks the run.

### Execution and hosting

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Stateful` | `bool` | no | `false` | _(none)_ | Routes every tool on this agent to a per-execution worker domain, so one run's calls reach the same process. It has no wire key of its own: it stamps `stateful` onto each tool. |
| `External` | `bool` | no | `false` | `external` | The agent is served elsewhere; no workers are started locally. Always sent. |

## Validation

`func (a *Agent) Validate() error` reports the first configuration error in the agent tree, and every `Runtime` method calls it before serializing. It checks:

- `Name` non-empty and matching `^[a-zA-Z_][a-zA-Z0-9_-]*$`; `MaxTurns >= 0`.
- `Strategy` and `ReasoningEffort` against the shared schema's enums.
- `Instructions` xor `InstructionsTemplate`, and `InstructionsTemplate.Name` non-empty.
- `Router` xor `RouterFunc`, and `StrategyRouter` requiring one of them.
- Every tool and prefill tool (see `ToolDef.Validate`), every guardrail, every handoff, the gate, and the termination tree.
- `PlannerContext` only under `StrategyPlanExecute`, and each entry's `Text` xor `URL`.
- `Planner`, `Fallback`, `Router` and every sub-agent, recursively.

Open ranges — timeouts, temperature bounds, model names — are left to the server, which can change them without an SDK release.

## ToolDef

`ToolDef` is a tool as the server sees it. The constructors in [`sdk/ai/tool`](../concepts/tools.md) are the supported way to build one: they derive `InputSchema` and `OutputSchema` by reflection over the handler's argument and return types, keeping the wire format identical to the other SDKs. Writing a `ToolDef` literal is for cases the constructors do not cover.

```go
tools := ai.Tools(
    tool.Func("get_weather", getWeather, "Current weather for a city."),
    tool.HTTP("lookup_order", "https://api.example.com/orders/{id}", "Look up an order."),
)
```

| Field | Type | Required | Default | Wire key | Description |
|---|---|---|---|---|---|
| `Name` | `string` | yes | — | `name` | The name the model calls and the task name the worker registers under. Must match `^[a-zA-Z_][a-zA-Z0-9_-]*$`. |
| `Description` | `string` | no | — | `description` | What the model reads to decide whether to call the tool. Always sent. |
| `InputSchema` | `map[string]any` | no | — | `inputSchema` | JSON Schema for the handler's argument type. Always sent; set by the constructors. |
| `OutputSchema` | `map[string]any` | no | — | `outputSchema` | JSON Schema for the return type. Omitted when empty. |
| `ToolType` | `ToolType` | no | `ToolTypeWorker` | `toolType` | How the server dispatches the call. |
| `Config` | `map[string]any` | depends on type | — | `config` | Type-specific settings: an `url` for HTTP/API, a `server_url` for MCP, an `agent` for an agent-as-tool (serialized as a nested `agentConfig`). |
| `Credentials` | `[]string` | no | — | `config.credentials` | Secret names this tool may read. They ride inside `config`, where the server's compiler looks for them, and are repeated as the task definition's runtime metadata. Set with `tool.WithCredentials`. |
| `ApprovalRequired` | `bool` | no | `false` | `approvalRequired` | Pauses the run for a human before the call is dispatched. Emitted only when true. |
| `Stateful` | `bool` | no | `false` | `stateful` | Routes this tool to a per-execution worker domain. Emitted when either the tool or its agent is stateful. |
| `TimeoutSeconds` | `*int` | no | server's | `timeoutSeconds` | Bounds one call. |
| `MaxCalls` | `*int` | no | unlimited | `maxCalls` | Caps how many times the tool may be called in a run. |
| `Guardrails` | `[]Guardrail` | no | — | `guardrails` | Checks on this tool's input or output. |
| `RetryCount` | `*int` | no | `2` | _(task definition)_ | Retries on failure. |
| `RetryDelaySeconds` | `*int` | no | `2` | _(task definition)_ | Delay between retries. |
| `RetryPolicy` | `RetryPolicy` | no | `RetryLinearBackoff` | _(task definition)_ | `RetryFixed`, `RetryLinearBackoff`, `RetryExponentialBackoff`. |
| `Handler` | `any` | for worker tools | — | _(never sent)_ | The Go worker function, set by the constructors. Nil for server-dispatched tool types and for workers served from another process. |

The three retry fields configure the Conductor task definition the runtime registers after a run starts, not `agentConfig`. Set them with `tool.WithRetry(count, delaySeconds, policy)`. The registered definition also carries `timeoutSeconds: 0` (the agent bounds the run), `responseTimeoutSeconds: 10` and `timeoutPolicy: RETRY`.

### ToolType

| Constant | Wire value | Dispatched to |
|---|---|---|
| `ToolTypeWorker` | `worker` | A Go function in this SDK. The only type that needs a `Handler`. |
| `ToolTypeHTTP` | `http` | Server-side HTTP call. |
| `ToolTypeAPI` | `api` | Server-side API call; `Config["url"]` required. |
| `ToolTypeMCP` | `mcp` | An MCP server; `Config["server_url"]` required. |
| `ToolTypeHuman` | `human` | A human, through the respond endpoint. |
| `ToolTypeAgent` | `agent_tool` | Another agent, run as its own workflow. |
| `ToolTypeGenerateImage` | `generate_image` | Server-side media generation. |
| `ToolTypeGenerateAudio` | `generate_audio` | Server-side media generation. |
| `ToolTypeGenerateVideo` | `generate_video` | Server-side media generation. |
| `ToolTypeGeneratePDF` | `generate_pdf` | Server-side document generation. |
| `ToolTypeRAGIndex` | `rag_index` | Vector-store indexing. |
| `ToolTypeRAGSearch` | `rag_search` | Vector-store search. |
| `ToolTypePullWorkflowMessages` | `pull_workflow_messages` | The workflow message queue. |

### ToolDef validation

`func (t ToolDef) Validate() error` checks the name pattern, that `ToolType` and `RetryPolicy` are known values, that an MCP tool has `Config["server_url"]` and an API tool has `Config["url"]`, and that every `${NAME}` placeholder in `Config["headers"]` is declared in `Credentials` — otherwise the placeholder would reach the remote server as literal text.
