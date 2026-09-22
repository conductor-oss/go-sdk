# Streaming and human-in-the-loop

Following a run as it happens with `AgentHandle.Events`, and answering it when it pauses for a
person.

## Starting a run you can watch

`Runtime.Start` returns once the run exists, with the agent's tool workers already polling.

```go
runtime := ai.NewRuntime(ai.Config{})
defer runtime.Shutdown()

handle, err := runtime.Start(ctx, agent, "Deploy v2.1 to production")
if err != nil {
    return err
}
fmt.Println("execution:", handle.ExecutionID)
```

`Runtime.Run` is the blocking form: the same start, followed by polling for the result.

## Events

`handle.Events(ctx)` opens the server-sent event stream for this execution and returns a channel
that is closed when the stream ends or `ctx` is cancelled.

```go
events, err := handle.Events(ctx)
if err != nil {
    return err
}
for ev := range events {
    switch ev.Type {
    case ai.EventThinking:
        fmt.Println("thinking:", ev.Text)
    case ai.EventToolCall:
        fmt.Println("tool call:", ev.Data)
    case ai.EventToolResult:
        fmt.Println("tool result:", ev.Data)
    case ai.EventDone:
        fmt.Println("done:", ev.Text)
    }
}
```

An already finished run yields no events at all, so take the outcome from `handle.Result`, not
from the stream.

### `Event`

| Field | Meaning |
|---|---|
| `Type` | `EventType`, the normalized kind |
| `Name` | the server's raw event name |
| `Text` | streamed output or message text, when the event has one |
| `Data` | the decoded payload, for what `Type` and `Text` do not cover |
| `ExecutionID` | the execution this event came from; a nested agent's is a sub-execution |

`EventType` values: `EventThinking`, `EventToolCall`, `EventToolResult`, `EventHandoff`,
`EventWaiting`, `EventMessage`, `EventError`, `EventDone`, `EventGuardrailPass`,
`EventGuardrailFail`.

`Type` is a direct conversion of the wire name, so an event name a server build adds is not one
of the constants but still arrives intact, in `Name` and `Type`.

## Polling instead

| Call | Returns |
|---|---|
| `handle.Status(ctx)` | `*AgentResult` for the current state, without waiting |
| `handle.Waiting(ctx)` | whether the run is blocked on a human |
| `handle.Result(ctx)` | blocks until the run reaches a terminal status |

## Human-in-the-loop

Three ways a run pauses for a person:

- `tool.RequiresApproval()` on a tool — the run pauses before that call is dispatched.
- `tool.Human(name, description, opts...)` — a tool the model calls to ask someone. Its default
  schema is one `question` string; `tool.WithInputSchema` collects more structure.
- A guardrail with `OnFail: ai.OnFailHuman`, on the output position.

```go
Tools: ai.Tools(
    tool.Func("check_balance", checkBalance, "Check an account balance."),
    tool.Func("transfer_funds", transferFunds,
        "Transfer funds between accounts.", tool.RequiresApproval()),
    tool.Human("ask_customer", "Ask the customer to confirm the destination account."),
),
```

### Answering

| Call | What it posts |
|---|---|
| `handle.Approve(ctx)` | `{"approved": true}` |
| `handle.Reject(ctx, reason)` | `{"approved": false, "reason": reason}` |
| `handle.Respond(ctx, output)` | `output`, for a human tool's own fields |

```go
for ev := range events {
    if ev.Type != ai.EventWaiting {
        continue
    }
    target, err := handle.For(ev)
    if err != nil {
        return err
    }
    if approved {
        err = target.Approve(ctx)
    } else {
        err = target.Reject(ctx, "declined by operator")
    }
    if err != nil {
        return err
    }
}
```

`handle.For(ev)` returns the handle for the execution the event came from, so a sub-agent's
request for input is answered on its own execution rather than on the run this handle started.
It returns the same handle when the ids match, and an error when the event names no execution:
answering the wrong one approves it unseen.

### Why `Respond` waits

`Approve` and `Reject` are `Respond` with a fixed payload, and `Respond` first polls the server
until it reports the run as waiting. The stream's waiting event announces the pause slightly
before the server will accept a response, so a response posted on the event alone is dropped and
the run hangs there; retrying instead injects spurious turns that re-run the pending tool. The
wait is built in, so answering straight from an `EventWaiting` event is correct — `Respond`
blocks until the response will land.

`Respond` returns an error if the execution is already terminal, and stops early if `ctx` does.

## Other controls

| Call | Effect |
|---|---|
| `handle.Signal(ctx, message)` | sets a persistent signal the agent prepends to its next LLM turn; an empty message clears it |
| `handle.SendMessage(ctx, message)` | pushes a message into the execution's message queue, for an agent waiting on `tool.WaitForMessage` |
| `handle.Pause(ctx)` / `handle.Resume(ctx)` | suspend and continue the execution |
| `handle.Stop(ctx)` | terminate the run |

## Next steps

Continue with [structured output](./structured-output.md), [guardrails](./guardrails.md), or the [runtime reference](../reference/runtime.md).
