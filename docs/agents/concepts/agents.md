# Agents

`ai.Agent` is the single orchestration primitive: an LLM, its instructions, its tools, and
optionally other agents. This page covers the struct's fields, how they are validated, and how a
defined agent is run.

## An agent is a document, not a constructor

`ai.Agent` is a declarative description sent to the Conductor server, which compiles it into a
workflow. There is no builder and no `NewAgent`: set the fields you need and leave the rest at
their zero value, which is what "unset" means on the wire. Optional numeric settings are pointers
because zero is a meaningful value for them; set those with `ai.Ptr`.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type weatherIn struct {
	City string
}

func getWeather(ctx context.Context, in weatherIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
}

func main() {
	agent := &ai.Agent{
		Name:         "weather_bot",
		Model:        "openai/gpt-4o",
		Instructions: "Use tools to answer questions. Keep answers short.",
		Tools: ai.Tools(
			tool.Func("get_weather", getWeather, "Get the current weather for a city."),
		),
		Temperature: ai.Ptr(0.2),
		MaxTurns:    10,
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "What's the weather in San Francisco?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
```

`ai.NewRuntime` reads `CONDUCTOR_SERVER_URL` (plus `CONDUCTOR_AUTH_KEY` and
`CONDUCTOR_AUTH_SECRET` on Orkes Conductor). `Run` compiles the agent, starts a worker for every
worker tool, and blocks until the run finishes.

## Identity

| Field | Type | Notes |
|---|---|---|
| `Name` | `string` | Required. Becomes the Conductor workflow name; must match `^[a-zA-Z_][a-zA-Z0-9_-]*$`. |
| `Model` | `string` | `"provider/model"`, e.g. `"openai/gpt-4o"`. Empty on a sub-agent inherits the parent's model at compile time. |
| `BaseURL` | `string` | Points this agent's provider at another endpoint, such as a proxy. |
| `Introduction` | `string` | The agent's opening message to the user. |
| `Metadata` | `map[string]any` | Arbitrary data carried with the definition. |

## Instructions

`Instructions` is the system prompt. `InstructionsTemplate` names a prompt template already stored
on the server; the SDK never creates one. They are mutually exclusive, and `Validate` rejects an
agent that sets both.

```go
agent := &ai.Agent{
	Name:  "support",
	Model: "openai/gpt-4o",
	InstructionsTemplate: &ai.PromptTemplate{
		Name:      "support_prompt",
		Variables: map[string]any{"tier": "${workflow.input.user_tier}"},
		Version:   ai.Ptr(3),
	},
}
```

`Version` nil selects the latest version. Variables may hold server-evaluated expressions.

## Tools

`Tools` is a `[]ai.ToolDef`. `ai.Tools` is a variadic helper that saves the slice literal:

```go
Tools: ai.Tools(
	tool.Func("lookup_order", lookupOrder, "Look up an order by id."),
	tool.HTTP("search", "https://api.example.com/search", "Search the product catalog."),
),
```

Only worker tools — those built with `tool.Func` — execute in your process. See
[Tools](./tools.md).

## Sub-agents

A non-empty `Agents` makes this a multi-agent system and causes `Strategy` to be sent; `Strategy`
defaults to `ai.StrategyHandoff`.

```go
team := &ai.Agent{
	Name:     "research_team",
	Model:    "openai/gpt-4o",
	Agents:   []*ai.Agent{researcher, writer},
	Strategy: ai.StrategySequential,
}
```

The strategies are `StrategyHandoff`, `StrategySequential`, `StrategyParallel`, `StrategyRouter`,
`StrategyRoundRobin`, `StrategyRandom`, `StrategySwarm`, `StrategyManual` and
`StrategyPlanExecute`. `Router`, `RouterFunc`, `Handoffs`, `AllowedTransitions`, `Planner`,
`Fallback`, `FallbackMaxTurns`, `PlannerContext` and `Synthesize` fill in the strategy-specific
slots. See [Multi-agent](./multi-agent.md).

An agent can also be exposed as a plain tool with `tool.Agent`, which the parent's LLM calls
inline instead of handing control over.

## Loop and model settings

| Field | Type | Unset behaviour |
|---|---|---|
| `MaxTurns` | `int` | `0` sends the default of 25. Negative is rejected. |
| `TimeoutSeconds` | `int` | `0` means no explicit limit. |
| `MaxTokens` | `*int` | nil leaves it to the model. |
| `Temperature` | `*float64` | nil is unset; `ai.Ptr(0.0)` is sent as `0`. |
| `ReasoningEffort` | `ai.ReasoningEffort` | Empty is unset. One of `ReasoningEffortMinimal`, `ReasoningEffortLow`, `ReasoningEffortMedium`, `ReasoningEffortHigh`; OpenAI reasoning models only. |
| `ThinkingBudgetTokens` | `*int` | nil disables extended thinking. |
| `ContextWindowBudget` | `*int` | nil disables proactive context condensing. |

`ai.Ptr` is the helper for all of these: `MaxTokens: ai.Ptr(4096)`.

To override the model parameters for a single run without changing the definition, pass
`ai.WithRunSettings(ai.RunSettings{...})` to `Run` or `Start`.

## Conversation memory

`Memory` seeds the agent with prior turns and bounds retained history. A non-nil `Memory` means
the agent is stateful even when both fields are empty.

```go
Memory: &ai.ConversationMemory{
	Messages: []map[string]any{
		{"role": "user", "content": "My order is ACC-1."},
		{"role": "assistant", "content": "Noted."},
	},
	MaxMessages: 40,
},
```

`IncludeContents: "none"` starts a sub-agent with a fresh context instead of inheriting the
parent's.

## Code execution and CLI

`CodeExecution` gives the agent a derived tool named `{agent}_execute_code`; `CLI` gives it
`{agent}_run_command`. Both are enforced by the server.

```go
agent := &ai.Agent{
	Name:  "analyst",
	Model: "openai/gpt-4o",
	CodeExecution: &ai.CodeExecutionConfig{
		AllowedLanguages: []string{"python"},
		TimeoutSeconds:   30,
	},
	CLI: &ai.CLIConfig{
		AllowedCommands: []string{"git status", "git log"},
		TimeoutSeconds:  30,
	},
}
```

`CodeExecutionConfig.Enabled` and `CLIConfig.Enabled` default to true when nil. Empty
`AllowedLanguages` defaults to `["python"]`; an empty `AllowedCommands` means no restriction, so
set one. `CodeExecutionConfig.Executor` selects how the code runs locally — nil means
`ai.LocalExecutor`, a subprocess on the worker host; `ai.DockerExecutor`, `ai.JupyterExecutor` and
`ai.ServerlessExecutor` are the alternatives. The executor never travels to the server.

`ai.ExecutorTool` wraps an executor as an ordinary tool instead, with no language or command
allow-list.

## Credentials

`Credentials` names secrets available to every tool on the agent, as a fallback for tools that
cannot declare their own. Prefer `tool.WithCredentials` on the tool itself.

## Remaining fields

| Field | Purpose |
|---|---|
| `OutputType` | Constrains the final answer to a struct's shape; pass a zero value, as in `OutputType: Ticket{}`. |
| `Guardrails` | Checks on this agent's input or output. Tools carry their own. See [Guardrails](./guardrails.md). |
| `Callbacks` | Hooks around the agent, each LLM call and each tool call; each set hook runs as a worker. See [Callbacks](./callbacks.md). |
| `Termination`, `StopWhen` | Stop the loop on a server-evaluated condition tree, or on Go logic run as a worker. See [Termination](./termination.md). |
| `Gate` | Decides whether a sequential pipeline continues past this agent. |
| `PrefillTools` | Tools run with fixed arguments before the first LLM turn; build each with `ai.Prefill`. |
| `MaskedFields` | Input and output fields the server redacts from execution history and the UI. |
| `Stateful` | Routes every tool on this agent to a per-execution worker domain. See [Stateful agents](./stateful.md). |
| `External` | Marks the agent as served elsewhere; no workers start locally. |

`RequiredTools` and `PlanSource` exist for parity with the Python SDK; `RequiredTools` is not
usable on current servers.

## Validation

`Validate() error` reports the first configuration error in the agent tree and is called by the
runtime before anything is serialized, so a bad definition fails locally rather than mid-run. It
checks the name pattern, the `Strategy` and `ReasoningEffort` enums, `MaxTurns >= 0`,
`Instructions` xor `InstructionsTemplate`, `Router` xor `RouterFunc` (and that
`StrategyRouter` has one), every tool, prefill, guardrail, handoff, gate and termination
condition, and recurses into `Agents`, `Router`, `Planner` and `Fallback`.

Open-ended values — model names, timeouts, temperature bounds — are left to the server, which can
change them without an SDK release.

```go
if err := agent.Validate(); err != nil {
	log.Fatal(err)
}
```

## Running

`ai.NewRuntime(ai.Config{})` returns a `*ai.Runtime`; one runtime serves many agents and should be
shut down with `defer runtime.Shutdown()`.

| Call | Use |
|---|---|
| `runtime.Run(ctx, agent, prompt, opts...)` | Blocks, returns `*ai.AgentResult`. |
| `runtime.Start(ctx, agent, prompt, opts...)` | Returns an `*ai.AgentHandle` for streaming, approval and control; see [Streaming and human input](./streaming-hitl.md). |
| `runtime.Deploy(ctx, agent)` | Compiles and registers without running; returns the workflow name. |
| `runtime.Serve(ctx, agents...)` | Deploys, starts workers, and blocks so another process can start runs by name. |

`AgentResult` carries `ExecutionID`, `Status`, `Output`, `FinishReason`, `Error`, `TokenUsage` and
the raw status document; `PrintResult` writes it to stdout. See
[Deploy · Serve · Run](./deploy-serve-run.md).

## Next steps

Continue with [tools](./tools.md), [multi-agent](./multi-agent.md), or the [agent definition reference](../reference/agent-definition.md).
