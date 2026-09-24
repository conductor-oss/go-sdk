# Deploy, serve, run

The five `ai.Runtime` lifecycle calls, what each one actually does, and how to split them
across processes in production.

Conductor has always kept three concerns apart: registering a definition on the server,
hosting the workers that answer its tasks, and starting an execution. Each call covers a
different subset.

| Call | Registers on server | Starts a run | Hosts tool workers | Blocks | Returns |
|---|---|---|---|---|---|
| `Plan(ctx, agent)` | no | no | no | no | `map[string]any` |
| `Deploy(ctx, agent)` | yes | no | no | no | workflow name |
| `Serve(ctx, agents...)` | yes | no | yes, until `ctx` is cancelled | yes | `ctx.Err()` |
| `Run(ctx, agent, prompt, opts...)` | yes, as a side effect | yes | yes | yes, until the run ends | `*ai.AgentResult` |
| `Start(ctx, agent, prompt, opts...)` | yes, as a side effect | yes | yes | no | `*ai.AgentHandle` |

All five validate the agent tree first; see `Agent.Validate`.

```go
rt := ai.NewRuntime(ai.Config{})
defer rt.Shutdown()
```

`ai.NewRuntime` reads `CONDUCTOR_SERVER_URL` and, on Orkes Conductor, `CONDUCTOR_AUTH_KEY`
and `CONDUCTOR_AUTH_SECRET`. `ai.NewRuntimeWithClient` takes an existing `*client.APIClient`
instead. `ai.Config` tunes `WorkerPollInterval`, `WorkerBatchSize` and `StatusPollInterval`.

## Plan — compile only

Asks the server to compile the agent and returns the result without registering or starting
anything: the workflow under `"workflowDef"` and the worker task names it needs under
`"requiredWorkers"`. Use it in CI to inspect or assert the compiled shape.

```go
compiled, err := rt.Plan(ctx, agent)
```

## Deploy — register, don't run

Compiles and stores the definition under the agent's name, and returns that name. Nothing
runs and no workers start. Safe to call repeatedly.

```go
name, err := rt.Deploy(ctx, agent)
```

## Serve — host the workers

Deploys each agent, registers its task definitions, starts its tool workers, then blocks
until `ctx` is cancelled, at which point it calls `Shutdown` and returns `ctx.Err()`. Use it
in a long-lived worker process for agents whose runs are started elsewhere.

```go
err := rt.Serve(ctx, researcher, writer)
```

Add replicas to scale; Conductor spreads tasks across whoever is polling. Restarting one
loses nothing: an in-flight run waits on its tool task until a worker polls again.

## Run and Start — register, start, and optionally wait

`Run` starts the agent and polls until a terminal state. `Start` is the same without the
wait: it returns an `*ai.AgentHandle` once the workers are polling.

```go
res, err := rt.Run(ctx, agent, "What is the capital of France?")
res.PrintResult()

h, err := rt.Start(ctx, agent, "Summarize today's incidents")
res, err := h.Result(ctx)
```

Both take `ai.RunOption`s: `ai.WithPlan`, `ai.WithMedia`, `ai.WithRunSettings`.

**`Run` and `Start` register the definition as a side effect of starting.** They send the
full agent definition with every start, and the server registers it under the agent's name,
so they overwrite whatever an earlier `Deploy` stored under that name. That is what makes
them the development path — one binary does everything, and changing the agent means
redeploying that binary — and why a production release should not run them against a name
that `Deploy` owns.

## The production split

Three roles, each owned, deployed and scaled on its own.

| Role | Call | Runs where | Notes |
|---|---|---|---|
| Publish | `Deploy(ctx, agent)` | A release step, once per version | Compiles and stores the definition. Nothing runs yet. |
| Serve | `Serve(ctx, agents...)` | A long-lived worker process | Hosts the tool workers and blocks. |
| Trigger | Start by name | Wherever the trigger lives | A [schedule](./scheduling.md), a webhook, a service in another language, or `POST /agent/start` with the agent's name. The caller needs no knowledge of the tools. |

`rt.AgentClient()` exposes the control plane for triggering by name from Go without
re-sending a definition.

## Driving a run in flight

`Start` returns an `*ai.AgentHandle` for the execution it began.

| Method | Effect |
|---|---|
| `h.ExecutionID` | The execution's ID |
| `h.Events(ctx)` | A channel of `ai.Event` until the run ends or `ctx` is cancelled |
| `h.Status(ctx)` | Current `*ai.AgentResult` without waiting |
| `h.Result(ctx)` | Blocks until terminal |
| `h.Waiting(ctx)` | Whether the run is blocked on a human |
| `h.For(ev)` | A handle to the execution an event came from, so a nested agent is answered on its own execution |
| `h.Respond(ctx, output)`, `h.Approve(ctx)`, `h.Reject(ctx, reason)` | Answer a waiting run or a tool call awaiting approval |
| `h.Stop(ctx)` | Terminate the run |
| `h.Signal`, `h.SendMessage`, `h.Pause`, `h.Resume` | The `Runtime` methods below, bound to this execution |

The same operations exist on the runtime, keyed by execution ID, so they reach a run this
process did not start:

| Method | Effect |
|---|---|
| `rt.Signal(ctx, executionID, message)` | Injects a persistent signal the agent prepends to its next LLM turn. It persists until overwritten; an empty message clears it. Works on any agent. |
| `rt.SendMessage(ctx, executionID, message)` | Pushes a message into the execution's queue, for an agent waiting on `tool.WaitForMessage`. A non-map value is wrapped as `{"message": value}`. |
| `rt.Pause(ctx, executionID)` | Suspends the execution, keeping its state. |
| `rt.Resume(ctx, executionID)` | Continues a paused execution. |
| `rt.Shutdown()` | Stops every worker this runtime started. Local only; it does not touch server state. |

`Resume` is the inverse of `Pause`. It is not the Python SDK's
`runtime.resume(execution_id, agent)`, which re-registers workers for a run another process
started; `Serve` covers that case.

See [streaming and human-in-the-loop](./streaming-hitl.md) for events and approvals.

## This is Conductor's standard lifecycle

None of it is specific to agents. The same lifecycle applies to every Conductor workflow;
the agent runtime only bundles calls the workflow API leaves separate.

| Agent runtime | Plain workflow |
|---|---|
| `Plan` | `ConductorWorkflow.ToWorkflowDef()` — built locally, whereas `Plan` compiles on the server |
| `Deploy` | `ConductorWorkflow.Register(overwrite)`, plus registering task definitions through the metadata client |
| `Serve` | `TaskRunner.StartWorker(...)` per task, then `WaitWorkers()` |
| `Run` | `WorkflowExecutor.ExecuteWorkflow(request, waitUntilTask)` — which blocks only to the server's synchronous timeout, while `Run` polls to a terminal state |
| `Start` | `WorkflowExecutor.StartWorkflow(request)`, then `MonitorExecution(id)` — a channel rather than a handle |

Two differences matter when moving between the two. A plain workflow must be registered
before it is started, whereas an agent `Run` registers as a side effect. And a plain
workflow does not start workers for you: the runtime can, because it derives each worker's
task name from the agent definition — one per tool with a Go handler, plus guardrails,
handoff predicates, gates, callbacks, router, termination and stop-when.

## Next steps

Continue with [scheduling](./scheduling.md), [stateful agents](./stateful.md), or the [runtime reference](../reference/runtime.md).
