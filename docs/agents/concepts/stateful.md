# Stateful agents

A run is independent by default: the agent remembers nothing from last time, and any worker
process may serve any tool call. Four separate mechanisms relax that, and you can mix them.

| Want | Use |
|---|---|
| Several runs to form one conversation | `ai.WithSession(id)` |
| One run's tool calls to reach one process | `Agent.Stateful`, `tool.Stateful()` |
| To supply history the run starts from | `Agent.Memory` |
| Tools to hand data to each other mid-run | `ai.SetState` and `ai.State` |

## Sessions: linking runs into a conversation

Pass the same session id to several runs and the server treats them as turns of one
conversation.

```go
for _, turn := range []string{"My name is Alice.", "What is my name?"} {
    res, err := rt.Run(ctx, agent, turn, ai.WithSession("user-42"))
    ...
}
```

The id is yours to choose; a user id or a chat id is the usual thing. Leave it off and the
server keys continuity to the execution instead, so the run stands alone.

This is a property of the run, not of the agent, which is why it is an option on `Run` and
`Start` rather than a field on `Agent`. One deployed agent can serve every user that way.

## Stateful mode: declaring it

`Agent.Stateful` marks every tool on that agent; `tool.Stateful()` marks one tool.

```go
agent := &ai.Agent{
    Name:         "hr_assistant",
    Model:        "openai/gpt-4o",
    Instructions: "You are an HR assistant.",
    Stateful:     true,
    Tools: ai.Tools(
        tool.Func("open_case", openCase, "Open an HR case."),
    ),
}
```

```go
Tools: ai.Tools(
    tool.Func("counter", counter, "Increment this run's counter.", tool.Stateful()),
)
```

`Agent.Stateful` is shorthand for marking each of that agent's tools; the two are equivalent.
Marking one tool anywhere in the tree — on a sub-agent, a router, a planner, a fallback, an
agent-as-tool — makes the whole run stateful.

## What it gives you

Without it, any process serving the agent can pick up any tool call. Run three copies of your
worker and the calls spread across them, which is what you want most of the time.

```
                  ┌── process A   ← any call, any run
tool "open_case" ─┼── process B
                  └── process C
```

That falls apart when a tool holds something in memory between calls. A run opens a case in
process A, its next call lands in process B, and the case is not there.

Stateful pins each run to one process. Every call from a run goes back to the process that
started it, for the life of the run, so anything a tool keeps in memory is still there next call.

```
tool "open_case" ─┬── process A   ← run 1's calls, always
                  └── process B   ← run 2's calls, always
```

## When to use it

Use it when a tool holds per-run state that is expensive or impossible to rebuild: an open
database transaction, a logged-in browser session, a loaded model, a file handle, a scratch
directory. One stateful tool anywhere in the tree pins the whole run.

Skip it otherwise. Pinning costs you the load spreading above: a busy run's calls all queue
behind one process instead of fanning out.

Where a value is small and serializable, prefer agent state to keeping it in memory. The server
carries that across tool calls whether or not the agent is stateful; see below.

## What to expect

- Two concurrent runs never take each other's tool calls, even for the same tool.
- Run a single process and nothing looks different; the benefit shows up once you scale out.
- `Run` and `Start` set the pinning up themselves. There is nothing extra to call.
- A nested skill's own tasks are not pinned, only the tools your agent declares.
- A run cannot be picked up by a different process later. `Serve` re-hosts an agent after a
  restart, but a run in flight when the process died does not resume; see the note at the end
  of this page.

## Seeding prior turns

`Agent.Memory` supplies conversation history up front and caps what is retained.

```go
agent.Memory = &ai.ConversationMemory{
    Messages: []map[string]any{
        {"role": "user", "message": "My name is Alice."},
        {"role": "assistant", "message": "Nice to meet you, Alice."},
    },
    MaxMessages: 20,
}
```

Each message is a `role`/`message` map, the shape the Python SDK's `ConversationMemory` helpers
append. Both fields are omitted from the wire when empty, so an empty `ConversationMemory` still
sends `"memory": {}`: non-nil memory means stateful, which is distinct from having no memory.

`ai.SemanticMemory` over a `ai.MemoryStore` is separate and in-process. Neither the Go nor the
Python runtime wires it into a run; call `SemanticMemory.Context` from application code and put
the result in the prompt.

## Sharing state between tools

Within one run, tools can hand data to each other through a map the server carries: `ai.State`,
`ai.StateValue` and `ai.SetState`, the counterpart of `ToolContext.state` in Python and
`ToolContext.getState()` in Java. See [Tools → Agent state](./tools.md#agent-state).

This is independent of `Stateful`. The server carries the map either way, so reach for it first
and mark the agent stateful only when a tool holds something the server cannot carry.

## Not implemented: resume from an instance

The Go SDK has no counterpart of Python's `runtime.resume(id, agent)`, which re-registers a
run's tool workers in a new process. `Runtime.Resume` and `AgentHandle.Resume` are the inverse of
`Pause` — they un-pause an execution — and nothing more.

Go workers are goroutines, so the case that loses them is a process restart, and a standing
`Runtime.Serve(ctx, agents...)` covers that for the whole fleet: it deploys each agent and polls
the shared queue for every agent it serves, without a run of its own. Re-attaching to one
execution's domain is the only thing it does not cover.

## Next steps

Continue with [deploy, serve, run](./deploy-serve-run.md), [streaming and human-in-the-loop](./streaming-hitl.md), or the [runtime reference](../reference/runtime.md).
