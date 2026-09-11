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

## Step 2: the operational runtime

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

## Step 3: pause, resume, and resume from an instance

`Runtime.Pause`, `Resume`, and `Resume(ctx, executionID)` for an existing
execution, with the handle able to reattach. Unlocks suite 23 (from instance
and event human-in-the-loop, 23 tests), the largest Python suite.

## Step 4: callbacks

`before_agent`, `after_agent`, `before_model`, `after_model` and the
`CallbackHandler` form. They run as workers, the pattern guardrails and
handoffs already use in `runtime.go`. Proof: suite 13 (callbacks, 5 tests).

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
