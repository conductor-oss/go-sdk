# AI e2e tests

End-to-end tests for `sdk/ai`. Each test describes an agent, sends it to a
real Conductor server, hosts the agent's tool workers in the test process, and
checks the outcome. They need a server with an LLM provider configured, so they
are behind the `e2e` build tag and are not part of `go test ./...`.

## Running

Start a Conductor server with an OpenAI key in its environment. The tests use
`openai/gpt-4o-mini` unless `CONDUCTOR_AGENT_LLM_MODEL` says otherwise.

    OPENAI_API_KEY=sk-... java -jar conductor-server.jar --server.port=8080

Then point the tests at it:

    export CONDUCTOR_SERVER_URL=http://localhost:8080/api
    go test -tags e2e -count=1 -v ./test/ai_e2e/

Without `CONDUCTOR_SERVER_URL` every test skips. A test whose optional
dependency is missing skips too and says what to start, so a run is never
silently green.

| Variable | Purpose |
|---|---|
| `CONDUCTOR_SERVER_URL` | Server API base URL. Required. |
| `CONDUCTOR_AUTH_KEY`, `CONDUCTOR_AUTH_SECRET` | Credentials for a server that needs them. Not needed for OSS. |
| `CONDUCTOR_AGENT_LLM_MODEL` | Model for every agent, as `provider/model`. |
| `CONDUCTOR_E2E_MCP_URL` | An MCP server the Conductor server can reach. Default `http://localhost:3001/mcp`, for example `mcp-testkit --transport http --port 3001`. |
| `CONDUCTOR_SECRET_<NAME>` | Set on the **server**, not the test. The OSS secret store is read-only and reads secrets from the server's environment at start. Credential tests skip when the secret they declare is not provisioned. |
| `CONDUCTOR_E2E_SECRET_PROVISIONED` | Set to any value to turn a missing secret from a skip into a failure. |

Run one test with `-run`:

    go test -tags e2e -count=1 -v -run TestToolCall ./test/ai_e2e/

## What travels where

A run touches three network hops. Knowing them matters when you want to
capture traffic or replace a side with a mock.

1. **Test process to Conductor server.** Plain HTTP JSON under
   `CONDUCTOR_SERVER_URL`. The agent definition goes to `POST /agent/start`
   as the `agentConfig` document, the same document the golden files under
   `sdk/ai/testdata/agent_config` pin down. Status polls hit
   `GET /agent/{id}/status`, and the tool workers poll `GET /tasks/poll/...`
   and report with `POST /tasks`. Event streaming is a long-lived
   `GET /agent/stream/{id}` server-sent-events response.
2. **Conductor server to the LLM provider.** The server runs the agent loop
   and calls the provider itself. The SDK never talks to the LLM.
3. **Conductor server to an MCP server**, for tests that use MCP tools. The
   server discovers and calls MCP tools itself, so the test process only needs
   to reach the MCP server for its own readiness check.

Tool workers, guardrail checks, and handoff predicates run inside the test
process as goroutines. The server dispatches them as Conductor tasks named
after the tool or agent, which is why each test registers its workers before
starting the run.

## Recording and replaying with the server's LLM recorder

The Conductor server can record every LLM response it receives and later
serve those responses instead of calling a provider. That makes an e2e run
deterministic and free, and it is the only way to run these tests in CI. The
feature lives on the server's `feature/llm_mock_impl` branch (see its
`RECORD_MOCKS.md`); it is not in a released server yet.

`scripts/conductor-server.sh` starts a local server in either mode. It needs
`CONDUCTOR_SERVER_JAR` pointing at a boot jar built from that branch:

    git -C /path/to/conductor checkout feature/llm_mock_impl
    (cd /path/to/conductor && ./gradlew :conductor-server:bootJar -x test -x spotlessCheck)
    export CONDUCTOR_SERVER_JAR=/path/to/conductor/server/build/libs/conductor-server-*-boot.jar

Recordings live under `testdata/llm-recordings/<group>/`, one directory per
group of tests recorded together, and are checked in.

**Record.** The server calls the real model and writes one JSON file per
response into the directory. Run only the tests you mean to record: every
agent run against the server is saved, related or not.

    OPENAI_API_KEY=sk-... test/ai_e2e/scripts/conductor-server.sh record test/ai_e2e/testdata/llm-recordings/<group>
    CONDUCTOR_SERVER_URL=http://localhost:8080/api go test -tags e2e -count=1 -v -run <TestName> ./test/ai_e2e/
    test/ai_e2e/scripts/conductor-server.sh stop

**Replay.** The same directory, the `mock` provider, and no key. The model
name tells the server to look up each request in the recordings; a request
with no recording fails the run rather than reaching a provider.

    test/ai_e2e/scripts/conductor-server.sh replay test/ai_e2e/testdata/llm-recordings/<group>
    CONDUCTOR_SERVER_URL=http://localhost:8080/api CONDUCTOR_AGENT_LLM_MODEL=mock/mockLLM \
        go test -tags e2e -count=1 -v -run <TestName> ./test/ai_e2e/
    test/ai_e2e/scripts/conductor-server.sh stop

The script passes these server properties; set them yourself if you start the
server another way:

| Property | Record | Replay |
|---|---|---|
| `conductor.integrations.ai.enabled` | `true` | `true` |
| `conductor.ai.record-mode` | `true` | `false` |
| `conductor.ai.enable-llm-mocks` | `false` | `true` |
| `conductor.ai.recordings-directory` | the directory, absolute | the same directory |

**What has to match.** A replayed request is looked up by its full normalized
content: every message's role and text, tool calls and tool results, the tool
definitions (name, description, input schema), the output schema, and the
generation options. So between the recording run and a replay:

- The test must send the same prompt, instructions, and tools. A change to
  any of them, or to how the SDK serializes them, needs a re-record.
- Tool workers still run, and their outputs are part of the next request.
  Keep them deterministic: no timestamps, IDs, or temp paths in what a tool
  returns to the model. If a test needs proof that a worker ran, write it
  somewhere the model never sees, such as a file in the test's temp dir.
- Agent tools cannot be replayed yet. The server hands the parent model the
  sub-agent's result together with its `subWorkflowId`, a fresh UUID each run,
  and the recorder does not normalize it away. A test that uses `tool.Agent`
  should skip when `model(t) == mockModel`; fixing it needs a server change
  in `RecordedRequestNormalizer`.

## Adding a test

- Build the runtime with `newRuntime(t)` and models with `model(t)` from
  `e2e_test.go` so the environment variables above keep working.
- Assert on what the test controls: that a worker ran, with what input, and
  the run's final status. Match model prose loosely or not at all.
- Skip, with a message saying what to start, when an optional dependency is
  missing. Never fail for a missing dependency.
- Keep each run under two minutes with `context.WithTimeout`. A run that times
  out usually means a worker was never registered for a task the server
  scheduled.
