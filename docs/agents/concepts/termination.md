# Termination

Termination conditions stop an agent loop before it reaches `Agent.MaxTurns`. This page covers the
four conditions, how they compose, and how `Agent.Termination` differs from `Agent.StopWhen`.

```go
agent := &ai.Agent{
    Name:        "researcher",
    Model:       "openai/gpt-4o",
    MaxTurns:    20,
    Termination: ai.TextMentionTermination{Text: "TASK_COMPLETE"},
}
```

## Conditions

`TerminationCondition` is a sealed interface — the server knows a fixed set of types, so a
caller-defined one could not be serialized. The four implementations are value types:

| Condition | Fields | Stops when |
|---|---|---|
| `ai.TextMentionTermination` | `Text string`, `CaseSensitive bool` | `Text` appears in the output. |
| `ai.StopMessageTermination` | `StopMessage string` | The output is exactly `StopMessage`. Empty means `"TERMINATE"`. |
| `ai.MaxMessageTermination` | `MaxMessages int` | The conversation reaches `MaxMessages`. |
| `ai.TokenUsageTermination` | `MaxTotalTokens`, `MaxPromptTokens`, `MaxCompletionTokens` (all `*int`) | A token budget is exhausted. |

The token limits are pointers because each is independent and optional; an unset one is omitted
rather than sent as a zero the server would read as "no tokens allowed". Use `ai.Ptr`:

```go
ai.TokenUsageTermination{MaxTotalTokens: ai.Ptr(50_000)}
```

## Composing

Go cannot overload Python's `&` and `|`, so composites are variadic constructors:

```go
// Stop when either holds.
cond := ai.OrTermination(
    ai.MaxMessageTermination{MaxMessages: 20},
    ai.StopMessageTermination{StopMessage: "DONE"},
)

// Stop only when both hold.
cond = ai.AndTermination(
    ai.MaxMessageTermination{MaxMessages: 20},
    ai.TextMentionTermination{Text: "DONE"},
)
```

Both return a `TerminationCondition` and nest freely.

## Termination versus StopWhen

The two fields are independent and may both be set.

| | `Agent.Termination` | `Agent.StopWhen` |
|---|---|---|
| Type | `TerminationCondition` | `StopWhenFunc` |
| Expressed as | A condition tree the server understands | Arbitrary Go logic |
| Serialized as | The tree itself, under `termination` | A worker reference, under `stopWhen` |
| Worker task | `<agent>_termination` | `<agent>_stop_when` |

`StopWhenFunc` returns `true` to stop:

```go
type StopWhenFunc func(ctx context.Context, state ai.StopWhenState) (bool, error)
```

```go
agent.StopWhen = func(ctx context.Context, state ai.StopWhenState) (bool, error) {
    return strings.Contains(state.Result, "confidence: high"), nil
}
```

`StopWhenState`:

| Field | Type | Meaning |
|---|---|---|
| `Result` | `string` | The agent's output so far. |
| `Messages` | `[]map[string]any` | The conversation history at this point. |
| `Iteration` | `int` | Loop passes, starting at 0. |

An error from `StopWhenFunc` continues the loop.

Prefer `Termination` for checks the server can express: it needs no worker round trip. Use
`StopWhen` when the decision needs Go code or state from your process.

## Workers

The runtime registers a worker for each field that is set, and the server compiles a task for it.
Both must be served, or the run stalls on a task nothing polls.

The `<agent>_termination` worker evaluates the same condition tree in Go and answers
`should_continue`. `TokenUsageTermination` never fires there — token counts are absent from that
task's input — so token budgets are enforced by the server alone.

## Validation

`Agent.Validate` walks the tree and rejects `MaxMessages` below 1, a `TokenUsageTermination` with
no limit set, and an `AndTermination` or `OrTermination` with no conditions.

## Next steps

Continue with [callbacks](./callbacks.md), [guardrails](./guardrails.md), or [multi-agent](./multi-agent.md).
