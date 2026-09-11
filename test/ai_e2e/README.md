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

Recordings live under `testdata/llm-recordings/<group>/<TestName>/`, one
directory per test inside a group recorded together, and are checked in. The
server loads the group directory recursively, so one replay run covers the
whole group. `suite/` holds the
whole package recorded in one pass with `openai/gpt-4o-mini`; replay it with
no `-run` filter. The credential tests were recorded with the secret
provisioned, so the server needs it in both modes and the tests need to be
told it is there:

    CONDUCTOR_SECRET_GH_TOKEN=ghp_fake_e2e_token_value_12345 \
        test/ai_e2e/scripts/conductor-server.sh replay test/ai_e2e/testdata/llm-recordings/suite
    CONDUCTOR_SERVER_URL=http://localhost:8080/api CONDUCTOR_AGENT_LLM_MODEL=mock/mockLLM \
        CONDUCTOR_E2E_SECRET_PROVISIONED=1 go test -tags e2e -count=1 -v ./test/ai_e2e/

Without the secret those two tests skip, and their recordings go unused.

**Record.** The server calls the real model and writes one JSON file per
response. The recorder names files by UUID and knows nothing about tests, so
record with `scripts/record-suite.sh`: it runs each test on its own and files
that test's recordings under `<dir>/<TestName>/`, then writes `<dir>/INDEX.md`
listing every recording with its test, the agent's tools and the first user
message. Run only the tests you mean to record; every agent run against the
server is saved, related or not.

    OPENAI_API_KEY=sk-... test/ai_e2e/scripts/conductor-server.sh record test/ai_e2e/testdata/llm-recordings/suite
    CONDUCTOR_SERVER_URL=http://localhost:8080/api test/ai_e2e/scripts/record-suite.sh test/ai_e2e/testdata/llm-recordings/suite
    test/ai_e2e/scripts/conductor-server.sh stop

Pass a Go test regex as the second argument to re-record a subset, for
example `'^TestSkill'`; delete that test's directory first so stale files do
not linger.

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
- MCP tool loops need three server fixes that are not yet on
  `feature/llm_mock_impl`: name a tool result
  after the call it answers rather than the task type (`LLMHelper`), treat a
  history the recorder cannot describe as a miss rather than an error
  (`MockLLM`), and render tool results under the dispatched tool's name rather
  than a reference built from the model's call ID (`ToolCompiler` and
  `JavaScriptBuilder`). The `suite/` recordings were made with those fixes, so
  `TestMCPToolResultReachesTheAnswer` fails in replay against a server
  without them. Ordinary worker tools replay on the branch as is.
- Paths in prompts must be stable. `TestCLICommand` lists a fixed relative
  directory rather than `t.TempDir()`, whose random name would make every
  run's prompt unique.
- Tests recorded together must not send identical requests. The model may
  answer the same prompt two ways, and playback refuses to start when one
  request has two answers ("Conflicting recorded responses for the same
  request"). Give each test its own prompt; the two guardrail tests ask
  different questions for this reason.

## Compile parity with the Python SDK

`suite1_basic_validation_test.go` is the Python SDK's
`e2e/test_suite1_basic_validation.py`, test for test and under the same
names, `test_x` as `TestX`. Those tests never run an agent: each compiles one
with `Runtime.Plan`, the counterpart of Python's `runtime.plan()`, and asserts
on the returned workflow. No model is called and no recording is involved.

Each test then compares the whole compiled workflow with the one the server
returned for the Python agent, and the `LLM_CHAT_COMPLETE` task inside it on
its own, since that task is the request the server will send the model when
the agent runs. `testdata/compiled/<agent>.json` holds the Python results,
captured by `testdata/compiled/generate_compiled.py`, which builds the same
agents with the Python SDK. Compilation is deterministic, so a difference is a
difference in what the two SDKs sent. One thing is compared loosely: the
JavaScript the compiler generates embeds our configs as JSON text, and Python
writes object keys in declaration order while Go sorts them, so those scripts
are compared without regard to key order.

To refresh the fixtures, from a python-sdk checkout with a server running and
the MCP test server on port 3001 for the kitchen sink:

    PYTHONPATH=src python3 /path/to/go-sdk/test/ai_e2e/testdata/compiled/generate_compiled.py \
        --out /path/to/go-sdk/test/ai_e2e/testdata/compiled

Nine of the suite's ten tests are ported. The tests live on the lowest branch
of the stack that has what they need: tools, credentials, sub-agents and
`base_url` on the foundation, guardrails and the kitchen sink where guardrails
and handoffs arrive. The tenth, the LLM-judge test, grades the compiled JSON by
calling a provider directly from pytest rather than through Conductor, so it
is not an SDK behaviour to port.

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
