# Callbacks

Callbacks are lifecycle hooks the server calls at fixed points in an agent run. This page covers
the six hooks, what they receive, and what returning a value does.

```go
agent := &ai.Agent{
    Name:  "observed_agent",
    Model: "openai/gpt-4o",
    Tools: ai.Tools(tool.Func("echo", echo, "Echo the input text back.")),
    Callbacks: &ai.Callbacks{
        OnToolStart: func(ctx context.Context, in ai.CallbackInput) (map[string]any, error) {
            log.Printf("tool call: %v", in.ToolCalls)
            return nil, nil
        },
    },
}
```

## Hooks

`Agent.Callbacks` is a `*ai.Callbacks`. Every field is a `CallbackFunc` and every one is optional.

| Field | Fires | Worker task |
|---|---|---|
| `OnAgentStart` | Before the agent starts | `<agent>_before_agent` |
| `OnAgentEnd` | After the agent finishes | `<agent>_after_agent` |
| `OnModelStart` | Before each LLM call | `<agent>_before_model` |
| `OnModelEnd` | After each LLM call | `<agent>_after_model` |
| `OnToolStart` | Before each tool call | `<agent>_before_tool` |
| `OnToolEnd` | After each tool call | `<agent>_after_tool` |

Each hook that is set runs as a Conductor worker in your process, registered under the task name
above — the agent's `Name` plus its wire position. Only the hooks you set are serialized and only
those get a worker.

## CallbackFunc

```go
type CallbackFunc func(ctx context.Context, in ai.CallbackInput) (map[string]any, error)
```

`CallbackInput`:

| Field | Type | Meaning |
|---|---|---|
| `Position` | `string` | The wire position, e.g. `"before_tool"`. |
| `AgentName` | `string` | The agent the hook fired for. |
| `LLMResult` | `any` | The model result as decoded JSON; its shape varies by turn. |
| `ToolCalls` | `any` | The turn's tool calls as decoded JSON. |

A field is empty or nil at positions that do not carry it.

## Return value

- `nil` or an empty map: continue unchanged.
- A non-empty map: override what the run does next. The server merges it in at that point.
- An error: treated as `nil`. A broken hook never blocks the run.

That last rule makes callbacks safe for observation, and makes them a poor place for work that must
not be silently dropped. Put durable side effects in a tool; see [Tools](./tools.md).

## Next steps

Continue with [guardrails](./guardrails.md), [termination](./termination.md), or [streaming and human-in-the-loop](./streaming-hitl.md).
