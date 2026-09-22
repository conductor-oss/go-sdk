# Stateful agents

How `Agent.Stateful` and `tool.Stateful()` pin one run's tool calls to the process that started it.

## Declaring it

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

`Agent.Stateful` has no wire key of its own: the serializer stamps `"stateful": true` onto each
of that agent's tools. One stateful tool anywhere in the tree — a sub-agent, a router, a planner,
a fallback, an agent-as-tool — turns on domain routing for the whole run.

## Per-execution worker domains

A Conductor task queue can be partitioned by *domain*. Stateful routing uses one domain per run:

1. `Runtime.Start` and `Runtime.Run` mint a run id when anything in the tree is stateful — 32
   lowercase hex characters, the same format as the Python SDK's `uuid4().hex` — and send it as
   `runId` in the start request. Nothing stateful means no run id and no domain.
2. The run id doubles as a queue name. The server records the mapping on the workflow, as its
   task-to-domain map, and schedules that run's tool tasks into those queues.
3. Once the run exists the runtime reads the mapping back off the execution and starts a worker
   per task name on the queue named there. So the process that started the run is the one polling
   that run's queue, and two concurrent runs never take each other's tasks.

The same tool in two stateful runs is two pollers, one per (task name, domain) pair.

A task name the server leaves out of the mapping is scheduled on the shared queue and its worker
polls undomained; that is how a nested skill's own tasks behave, since the server routes only the
agent's declared tools.

### Ordering

Workers for a stateful run can only start after the run does, because the domain does not exist
until then — `Start` and `Run` do it in that order, before returning. A run with no stateful tools
registers its workers *before* the start call instead, so no tool call is enqueued unheard.

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
