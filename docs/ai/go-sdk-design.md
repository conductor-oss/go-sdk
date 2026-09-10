# Agents in the Go SDK — design as built

## Goals

- Produce the **same `agentConfig` document** the Python serializer produces for the same agent,
  so the server compiles identical workflows regardless of SDK. Enforced by the golden files.
- Reuse the SDK's existing worker framework rather than build a second polling loop.
- Be idiomatic Go where that costs nothing on the wire: typed structs, `context.Context` on every
  callback, reflection-derived schemas, errors instead of exceptions.
- Verify behaviour, not just shape: every worker-backed feature has an e2e test asserting a
  worker-side effect against a live server.

## Non-goals

- Framework adapters (LangChain, LangGraph, Claude Agent SDK). See `framework-support.md`.
- Server-side orchestration logic. The server compiles and runs the loop; the SDK describes the
  agent and serves workers.
- A parallel SSE-first runtime. `Run` polls status; `AgentHandle.Events` streams, but the control
  flow never depends on a particular event arriving.

## Package layout

```
sdk/ai/
  agent.go          Agent, PromptTemplate, ConversationMemory, ReasoningEffort, Validate
  serializer.go     Agent → agentConfig (toConfig)
  strategy.go       Strategy constants
  tooldef.go        ToolDef, ToolType, the sealed Guardrail interface
  tool/             constructors: Func, HTTP, Human, Agent, MCP + options
  handoff.go        OnToolResult, OnTextMention, OnCondition, HandoffFunc, HandoffState
  guardrail.go      RegexGuardrail, LLMGuardrail, CustomGuardrail, NewCustomGuardrail
  termination.go    TextMention/StopMessage/MaxMessage/TokenUsage, And/Or
  worker_ref.go     StopWhenFunc, RouterFunc, derived worker names
  plan.go           Plan, Step, Op, Ref, Generate — caller-supplied plans
  execution.go      CodeExecutionConfig, CLIConfig; code_executor.go, cli_runner.go
  secret.go         Secret, SecretsEnv, ErrCredentialNotFound
  runtime.go        Runtime, Config, NewRuntime, Start, Run, registerWorkers, Shutdown
  handle.go         AgentHandle, Event, EventType
  result.go         AgentResult, Status, TokenUsage
  dispatch.go       reflection-based tool executor; task → context plumbing
sdk/client/
  agent_client.go   REST: compile / start / status / respond / stop
  sse.go            StreamSSE — its own request path, because APIClient buffers bodies
```

## Core types

### `Agent`

A plain struct, built with a literal, validated by `Agent.Validate()` before anything is sent.
Field names follow Python's snake_case → Go's exported CamelCase; the serializer hand-writes the
wire key for each, as Python does. Two things Go cannot do that Python can, and how they are handled:

- **Defaults from constructors.** Go has none, so `MaxTurns == 0` means "default (25)", and
  `StopMessageTermination{}` means `"TERMINATE"`. Substituted at serialization, not construction.
- **Validation at construction.** `Validate` walks the tree instead: sub-agents, router,
  termination, handoffs (target must be a sub-agent; one `OnCondition` per target), guardrail
  enums, `Instructions` xor `InstructionsTemplate`.

### Strategy

`Strategy` is a string type whose constants match Python's enum. It is emitted **only when
the agent has sub-agents** — `hasSubAgents()` — because Python omits it for leaf agents even
though the default is `handoff`. `StrategyRouter` requires `Router` or `RouterFunc`, never both.

### Tools — generics at the boundary, reflection behind it

```go
func Func[In, Out any](name, description string,
	fn func(context.Context, In) (Out, error), opts ...Option) ai.ToolDef
```

The type parameters make the handler's shape a compile-time contract and give the constructor
the `In` and `Out` types to derive `inputSchema` and `outputSchema` from — struct fields and
`json` tags, types and required-ness only, to match what Python and Java emit. `ToolDef` itself is
not generic: the handler is stored as `any` so tools of different types share one slice, and
dispatch binds the task input to `In` and calls the handler by reflection. Deriving `Out` is the
one place Go gives the planner more than Python can, which chained plan steps need.

Non-worker tools (`tool.HTTP`, `tool.Human`, `tool.Agent`, `tool.MCP`) have no handler; the
server dispatches them itself. Their type-specific settings live in `ToolDef.Config`, as on the
wire; `Credentials` is a typed field that lands under `config` when serialized, which is where the
server's compiler looks.

### Callbacks take `context.Context` and a state struct

Every Go function the server can dispatch to has the shape `func(ctx context.Context, in State) (Out, error)`:

| callback | state | returns | Python equivalent |
|---|---|---|---|
| `HandoffFunc` | `HandoffState{Result, Conversation, Context, ActiveAgent, ToolResults}` | `(bool, error)` | `OnCondition(lambda ctx: ...)` |
| `GuardrailFunc` | `GuardrailInput{Content, Iteration, ToolCalls}` | `(GuardrailResult, error)` | `(content: str) -> GuardrailResult` |
| `StopWhenFunc` | `StopWhenState{Result, Messages, Iteration}` | `(bool, error)` | `stop_when(...)` |
| `RouterFunc` | `prompt string` | `(string, error)` | router callable |

The struct carries only fields the server actually sends. It lets a check use data Python's
scalar signatures discard — a guardrail can see which retry it is on, or which tools the model
called — and lets fields be added without breaking callers. `context.Context` is the one thing Go
convention is strict about for anything that may block; the earlier `GuardrailContext` struct that
lacked it was replaced for that reason.

`HandoffState.ActiveAgent` is a **name**. The server sends an index into
`[parent, sub-agents...]` as a string; the adapter resolves it and falls back to the raw value if
it cannot, so a mismatch is visible rather than silently wrong.

## Wire compatibility — the bar

The invariant is Python's: **absent optionals and empty collections are omitted.** The serializer
builds a `map[string]any` and only sets keys that have a value, which is easier to keep honest than
struct tags with `omitempty` (a zero `int` and an unset one are indistinguishable there).

`TestGoldenAgentConfig` is the contract test. Each fixture maps to a Go agent
constructor; the test serializes it and compares to the Python document parsed into
`map[string]any`, so key order is free but every key and value must match. The fixtures cover the
parts that are easy to get subtly wrong: omitted-vs-zero, strategy only with sub-agents, derived
task names, nested termination trees, credentials under a tool's `config`.

Regenerate fixtures only when the wire format itself changes, from a python-sdk checkout, with
`sdk/ai/testdata/agent_config/generate_fixtures.py`. A fixture change means every SDK must follow.

## `Runtime` — composition over the existing worker framework

```go
rt := ai.NewRuntime(ai.Config{ /* WorkerPollInterval 100ms, WorkerBatchSize 1, StatusPollInterval 500ms */ })
defer rt.Shutdown()

res, err := rt.Run(ctx, agent, prompt)             // start, then poll status to a terminal state
h, err   := rt.Start(ctx, agent, prompt)           // start, return at once
res, err  = rt.Run(ctx, agent, prompt, ai.WithPlan(p))   // caller-supplied plan
```

`Runtime` owns one `worker.TaskRunner` and registers every worker-backed thing onto it, so agent
workers get the same polling, batching, and shutdown as every other worker in this SDK. **Workers
run as goroutines in the calling process** — unlike Python's forked processes — so a test can
assert on a counter the tool closure incremented.

`Start` does three things in order: `Validate`; `registerWorkers`; `POST /agent/start`. Workers
are registered **before** the run starts so a task cannot be enqueued with nothing polling for it.
`Run` is `Start` followed by `awaitResult`.

### Derived worker names — the contract with the server

The server compiles each callable into a `SIMPLE` task and dispatches by name. The SDK must be
polling under exactly that name. `worker_ref.go` owns the suffixes; they match Python's serializer.

| what | task name | registered by |
|---|---|---|
| a `tool.Func` | the tool's name | `registerWorkers` |
| `StopWhen` | `{agent}_stop_when` | " |
| `RouterFunc` | `{agent}_router_fn` | " |
| `OnCondition` handoff | `{agent}_handoff_{target}` — **one worker per condition** | " |
| `CustomGuardrail` | its `Name` | " |
| code execution / CLI | derived execution tools | " |

The handoff row is where Go differs from python-sdk and java-sdk today. The server changed this
contract on 2026-07-20: `{agent}_handoff_check` became an INLINE resolver and each
`on_condition` gets its own SIMPLE task named from `taskName`. Both other SDKs still register
`_handoff_check`, so their `OnCondition` handoffs sit `SCHEDULED` with `pollCount: 0` forever.
Go's `TestOnConditionHandoff` fails within two minutes if this contract drifts again.

## `AgentHandle` — a run in flight

`Events(ctx)` streams SSE; `Status`, `Waiting`, `Result` poll; `Respond`, `Approve`, `Reject`,
`Stop` act. `Respond` confirms the server reports the run as waiting before posting, because the
`waiting` event arrives slightly before the server will accept a response — posting on the event
alone races it. `EventType` has the same values python-sdk and java-sdk declare; `Event.Name` keeps
the raw wire name so an event the server adds is not lost.

## Error handling

Errors are returned, never panicked. `Validate` errors name the agent and the field. A tool
handler's error becomes a failed task with the message; a guardrail's error is a **failure, not a
pass**; a handoff predicate's error means **no handoff**, matching Python's `should_handoff`.
`ErrCredentialNotFound` is a sentinel so a tool can `errors.Is` it and report to the model rather
than fail the task.

## Base-SDK model changes

`model.Task` gained `RuntimeMetadata map[string]string` — the wire-only field the server
delivers credentials on. `sdk/client/sse.go` adds a streaming request path because `APIClient`
reads whole bodies before returning.

## Testing

- Unit: `go test ./sdk/ai/` — validation, serialization, adapters, golden conformance.
- E2E: `go test -tags e2e ./test/ai_e2e/` with `CONDUCTOR_SERVER_URL` (and
  `CONDUCTOR_AUTH_KEY`/`SECRET` if secured; `CONDUCTOR_AGENT_LLM_MODEL` to override
  `openai/gpt-4o-mini`). Each scenario asserts a worker-side effect. The streaming test
  retries up to three times because the model sometimes stalls or runs away — measured, and
  documented in the test.
