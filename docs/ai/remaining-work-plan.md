# Go Agent SDK: remaining work, in order

Status as of 2026-09-10, measured against the Python SDK on
`feat/agent-golden-fixtures` (`d1a512a`). The definition layer is at parity
except for the fields listed in step 1; the tool layer is complete for every
type the compiler runs; the runtime covers run, start, stream, approvals and
plan, and lacks the operational layer around it.

Two harnesses decide the order. Anything that is purely serialization is
verified in seconds by a golden fixture generated from Python
(`sdk/ai/testdata/agent_config/generate_fixtures.py`). Anything that runs is
verified by porting the matching Python e2e suite under the same test names
and replaying it from a recording the Python SDK made
(`test/ai_e2e/README.md`, "Recording and replaying"). Each item below names
the harness that proves it.

Keep one squashed commit per branch and add a branch per step rather than
growing `feat/agent-sdk-3-workers`.

## Step 1: close the serialization gaps — DONE 2026-09-10

Landed on `feat/agent-golden-fixtures`. Two corrections to the original
list: tool retries are not on the `agentConfig` wire, Python applies them to
the task definition it registers for each worker, so the Go runtime now
registers task definitions the same way and `WithRetry` feeds them; and
`dependencies` never leaves the process, so it is not ported.

| Item | Python | Go change | Proof |
|---|---|---|---|
| Tool retries | `@tool(retry_count, retry_delay_seconds, retry_policy)` | `tool.WithRetry(count, delaySeconds, policy)`; `retryCount`, `retryDelaySeconds`, `retryPolicy` on the tool document | golden fixture |
| API tool | `api_tool` | `tool.API(...)`; `ToolTypeAPI` exists | golden fixture |
| RAG tools | `index_tool`, `search_tool` | `tool.Index`, `tool.Search`; `ToolTypeRAGIndex`, `ToolTypeRAGSearch` exist | golden fixture |
| Workflow messages | `wait_for_message_tool` | `tool.WaitForMessage`; `ToolTypePullWorkflowMessages` exists | golden fixture; server needs `conductor.workflow-message-queue.enabled=true` |
| Gate | `gate` | `Agent.Gate GateFunc`, worker under `<agent>_gate` (suffix already in `worker_ref.go`) | golden fixture; suite 12 |
| Planning fields | `prefill_tools`, `plan_source`, `synthesize`, `planner_context` | fields on `Agent`, emitted only when set | golden fixture; suite 20 |
| Dependencies | `dependencies` | `Agent.Dependencies` | golden fixture |

Then port Python suites 12 (termination gates, 5 tests) and 20 (plan execute,
9 tests) and record them with Python. Not yet done.

## Step 2: the operational runtime — PARTLY DONE 2026-09-11

Done on `feat/agent-golden-fixtures` (uncommitted at time of writing):
`Runtime.Deploy`, `Runtime.Serve`, `Runtime.Signal`, `Runtime.SendMessage`
(the last two also on `AgentHandle`), and per-run `WithMedia` and
`WithRunSettings`. `SendMessage` posts to `/workflow/{id}/messages`, added
to the agent client. Unit tests cover all of it against a fake server;
suite 24's two run tests are ported and replay green. Still open in this
step: `Runtime.Prepare` (macOS fork batching, Python-specific, likely
unneeded in Go), and a live e2e for `Serve` and for `SendMessage` (needs a
WMQ-enabled server). Suite 25 (media input) is not e2e-ported yet: it needs
a vision-model recording and the server's allowed media directory.

In this order; the first three are wiring over calls the client already has.

1. `AgentHandle.Signal(ctx, message)` and `SendMessage`: `client.Signal` exists.
2. `Runtime.Deploy(ctx, agent)`: `client.Deploy` exists. Return the registered name.
3. Per-run `RunSettings` (model, temperature and other overrides merged into
   `agentConfig` before start) and media input on `Run` and `Start`. Small
   payload changes; suite 25 (media input) becomes portable.
4. `Runtime.Serve(ctx, agents...)`: a long-lived process hosting workers for
   deployed agents without starting a run. The largest item here and the one
   production users need most: today a Go process can only serve an agent it
   started itself.

Proof: suites 24 (agent client, 5 tests) and 25 (media input, 2 tests), plus
a Go-only e2e for Serve that deploys, then starts the run from `AgentClient`.

## Step 3: pause and resume — DONE 2026-09-11 (re-attach deferred)

Done on `feat/agent-golden-fixtures` (uncommitted at time of writing).
`Runtime.Pause` and `Runtime.Resume` suspend and un-pause an execution
(PUT `/workflow/{id}/pause` and `/resume`), and are also on `AgentHandle`.
Wire behavior is covered by fake-server unit tests; live pause/resume is
timing-dependent (the replay model returns before a pause could land), so it
is not e2e-recorded.

**Resume from an instance (re-attach) is deferred**, not implemented. In
Python `runtime.resume(id, agent)` re-registers a run's tool workers in the
process, keyed to the execution's per-execution worker domain. In Go two
things make it low-value today: workers are in-process goroutines, so the
only case they vanish is a process restart, which a standing `Serve(agent)`
already covers for the whole fleet; and the Go runtime does not yet route
stateful agents to per-execution domains, so re-attach's one advantage over
`Serve` — reconnecting to a specific execution's domain — cannot be
exercised or tested. Add it together with stateful-domain routing, when it
can be verified end to end, rather than ship a method whose main use is
untestable. Suite 23's event HITL flows still remain to port.

## Step 4: callbacks — DONE 2026-09-11

Done on `feat/agent-golden-fixtures` (uncommitted at time of writing).
`Agent.Callbacks` is a struct of six optional `CallbackFunc` fields, the
counterpart of the Python SDK's `CallbackHandler` and its six overridable
methods (before/after agent, model, tool). Each set hook serializes as
`{position, taskName}` and runs as a `<agent>_<position>` worker, the same
worker pattern guardrails and gates use. A hook returns a non-empty map to
override what the run does next, or nil to continue; an error is treated as
nil so a broken hook never blocks the run.

The worker input matches the server's `buildCallbackTask`, not the Python
worker's older signature: `callback_position`, `agent_name`, `llm_result`
(the model output, shape varies by turn) and `tool_calls`, surfaced on
`CallbackInput` with the varying fields typed `any`. Golden fixture
`25_callbacks` matches Python; suite 13's three executable tests are ported
under their names and replay green, each proving its hook fired in-process.

## Step 5: schedules

A `Schedule` type and a schedules client: create, update, list, pause,
resume, delete, run once. Mostly control plane. Proof: suite 21 (scheduling,
11 tests).

## Step 6: when a user asks

- Code executors beyond the local one: Docker, Jupyter, serverless.
- Memory stores and semantic memory; the server side is still moving.

## Not planned

The framework bridges, ClaudeCode, GPTAssistantAgent, LangGraph and ADK, wrap
Python libraries. There is nothing in Go to bridge and a Go user would not
expect them.

## Known loose ends to carry along

- The two skill e2e tests have no recordings on the branch; re-record them
  with the patched server.
- Replay of the MCP test depends on `conductor-llm-recorder-mcp.patch`, which
  is not yet on the conductor `feature/llm_mock_impl` branch. Get it merged.
- Python suites 2 and 3 skip on conductor-oss because they write to the
  secret store; they need a server with a writable store to port.
- The parity plan table lacks a row for the media tool constructors.
- `RequiredTools` is unusable on current servers, same as Python.
