# AI integration tests

End-to-end tests for `sdk/ai`. Each test describes an agent, sends it to a
real Conductor server, hosts the agent's tool workers in the test process, and
checks the outcome. They need a server with an LLM provider configured, so they
are behind the `integration` build tag and are not part of `go test ./...`.

## Running

Start a Conductor server with an OpenAI key in its environment. The tests use
`openai/gpt-4o-mini` unless `CONDUCTOR_AGENT_LLM_MODEL` says otherwise.

    OPENAI_API_KEY=sk-... java -jar conductor-server.jar --server.port=8080

Then point the tests at it:

    export CONDUCTOR_SERVER_URL=http://localhost:8080/api
    go test -tags integration -count=1 -v ./test/integration_tests/ai/

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

    go test -tags integration -count=1 -v -run TestToolCall ./test/integration_tests/ai/

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
serve those responses instead of calling a provider. That makes a run
deterministic and free, and it is the only way to run these tests in CI. The
feature lives on the server's `feature/llm_mock_impl` branch (see its
`RECORD_MOCKS.md`); it is not in a released server yet.

`scripts/conductor-server.sh` starts a local server in either mode. It needs
`CONDUCTOR_SERVER_JAR` pointing at a boot jar built from that branch:

    git -C /path/to/conductor checkout feature/llm_mock_impl
    (cd /path/to/conductor && ./gradlew :conductor-server:bootJar -x test -x spotlessCheck)
    export CONDUCTOR_SERVER_JAR=/path/to/conductor/server/build/libs/conductor-server-*-boot.jar

Recordings live under `testdata/llm-recordings/<test file>/`, one directory
per test file named after it without `_test.go`, the layout the shared
recordings in the conductor repository use, and are checked in. The whole package was recorded in one pass with
`openai/gpt-4o-mini`; the server loads the directory recursively, so one
replay run covers every test. The credential tests were recorded with the
secret provisioned, so the server needs it in both modes and the tests need
to be told it is there:

    CONDUCTOR_SECRET_GH_TOKEN=ghp_fake_e2e_token_value_12345 \
    CONDUCTOR_SECRET_MCP_AUTH_KEY=e2e-test-secret-key-12345 \
    CONDUCTOR_SECRET_HTTP_AUTH_KEY=e2e-http-test-secret-key-67890 \
        test/integration_tests/ai/scripts/conductor-server.sh replay test/integration_tests/ai/testdata/llm-recordings
    CONDUCTOR_SERVER_URL=http://localhost:8080/api CONDUCTOR_AGENT_LLM_MODEL=mock/mockLLM \
        CONDUCTOR_E2E_SECRET_PROVISIONED=1 go test -tags integration -count=1 -v ./test/integration_tests/ai/

Without the secrets those tests skip, and their recordings go unused.

**Record.** The server calls the real model and writes one JSON file per
response. The recorder names files by UUID and knows nothing about tests, so
record with `scripts/record-suite.sh`: it runs each test on its own and files
that test's recordings under `<dir>/<test file>/`, then writes `<dir>/INDEX.md`
listing every recording with its test, the agent's tools and the first user
message. Run only the tests you mean to record; every agent run against the
server is saved, related or not.

    OPENAI_API_KEY=sk-... test/integration_tests/ai/scripts/conductor-server.sh record test/integration_tests/ai/testdata/llm-recordings
    CONDUCTOR_SERVER_URL=http://localhost:8080/api test/integration_tests/ai/scripts/record-suite.sh test/integration_tests/ai/testdata/llm-recordings
    test/integration_tests/ai/scripts/conductor-server.sh stop

Pass a Go test regex as the second argument to re-record a subset, for
example `'^TestSkill'`; delete that test's directory first so stale files do
not linger.

**Replay.** The same directory, the `mock` provider, and no key. The model
name tells the server to look up each request in the recordings; a request
with no recording fails the run rather than reaching a provider.

    test/integration_tests/ai/scripts/conductor-server.sh replay test/integration_tests/ai/testdata/llm-recordings
    CONDUCTOR_SERVER_URL=http://localhost:8080/api CONDUCTOR_AGENT_LLM_MODEL=mock/mockLLM \
        go test -tags integration -count=1 -v -run <TestName> ./test/integration_tests/ai/
    test/integration_tests/ai/scripts/conductor-server.sh stop

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
- Numbers in tool results must serialize the same way. Python writes an
  integral float as `15000.0`; Go's `encoding/json` writes the same `float64`
  as `15000`. The recorder compares tool results as JSON nodes, so the two do
  not match. `TestExample09HumanInTheLoop` skips for this reason until the
  recorder compares numbers by value; a tool returning integral floats should
  avoid the issue by returning an int where the Python tool does.
- A tool that renders a JSON object must not depend on key order. Python's
  `format_response` in example 33 prints a dict in the order the task input
  arrived, and that order is what the recording holds; a Go `map` has no
  order, so the Go port sorts the keys and its result never matches. The
  matcher itself is order-free for JSON objects; by the time it sees this
  value the dict has become one string, and the order is inside the string.
  `TestExample33ExternalWorkers` runs live only and skips in playback.
  Proposed fix: have the Python example iterate `sorted(data.items())`, a
  one-line change to the example, then re-record 33. The Go port already
  sorts, so the recording would then match and the skip can go.
- An example that calls a live third-party API is not a replay fixture.
  Server-side HTTP tools call the real endpoint even in playback, and the
  whole response is part of the next model request. Example 16e calls GitHub
  with a personal token: the response carries headers tied to that token
  (`X-OAuth-Scopes`, `x-oauth-client-id`), and its body is live data. Everyone
  has a different token, so only the recording author can replay it, and only
  until the data changes. `TestExample16eCredentialsHTTPTool` therefore runs
  live only, with `CONDUCTOR_SECRET_GITHUB_TOKEN` set on the server, and
  skips in playback.
- A tool's parameters must be in the same order as the Python tool declares
  them. The server copies the order into text the model reads, such as a
  planner's tool catalog, so the SDK's schema builder keeps struct field
  order rather than the sorted order `encoding/json` gives a map.
- Paths in prompts must be stable. `TestCLICommand` lists a fixed relative
  directory rather than `t.TempDir()`, whose random name would make every
  run's prompt unique.
- Tests recorded together must not send identical requests. The model may
  answer the same prompt two ways, and playback refuses to start when one
  request has two answers ("Conflicting recorded responses for the same
  request"). Give each test its own prompt; the two guardrail tests ask
  different questions for this reason.

## Tests that skip in playback

Every skip is a deliberate `t.Skip` with its reason in the message. In a
playback run of the whole package these are the skips, and why:

| Test | Reason |
|---|---|
| `TestExample09HumanInTheLoop` | The Python recording holds `15000.0` where Go sends `15000`; the recorder compared numbers by type. Fixed by the first commit of conductor PR #1633; skips until it lands. |
| `TestSkillAsAgentTool` | An agent used as a tool puts the sub-workflow's fresh id into the next request, so no recording matches. Fixed by the second commit of PR #1633; skips until it lands. |
| `TestExample16eCredentialsHTTPTool` | Calls GitHub with a personal token; the response differs per token and over time. Live only. |
| `TestExample33ExternalWorkers` | Its tool renders a dict into text whose key order a Go map cannot keep, so the last request never matches. Live only until the Python example sorts the keys and is re-recorded. |
| `TestCredentialLifecycle` | Replays steps 1–3, then skips at the first secret write: the OSS store is read-only (HTTP 501). The Python test skips at the same point. |
| `TestCliCredentialLifecycle` | `cli_mktemp` returns a fresh temp path the model reads, and `cli_gh` needs a real `GITHUB_TOKEN`. Live only; on OSS it also skips at the store write. |
| `TestPdfGenerationAndRoundtrip` | The tool hands the model the URL of a freshly generated file. Live only; the text round trip also needs `pdftotext`. |
| `TestImageOpenai` | Live only, same reason as the PDF test. Live, it skips on the failure Python marks xfail: the server sends a `style` parameter the image API rejects. |
| `TestImageGemini` | Live only; needs `GOOGLE_AI_API_KEY`. |
| `TestAudioOpenai` | Live only, same reason as the PDF test; needs `OPENAI_API_KEY`. Passes live. |
| `TestJupyterStateful` | Asked to run `print(x * 73)` exactly as provided, the model defined its own `x` in 12 of 13 live attempts. Live only rather than a recording of the one lucky run. |
| `TestStatefulSwarmHandoffCompletes` | Disabled in the Python suite too: a stateful swarm handoff does not reliably complete in its domain. |
| The rest of suite 14, and `TestAgentToolSkillWorkersWithDomain` | A stateful run's domain, and a nested agent tool's sub-workflow id, are new on every run, so nothing about them can be replayed. Live only. |

## Suite 1 compile tests

`suite1_basic_validation_test.go` is the Python SDK's
`e2e/test_suite1_basic_validation.py`, test for test and under the same
names, `test_x` as `TestX`. Those tests never run an agent: each compiles one
with `Runtime.Plan`, the counterpart of Python's `runtime.plan()`, and asserts
on the returned workflow the way the Python test does. No model is called and
no recording is involved, so they run live in every mode.

Nine of the suite's ten tests are ported. The tenth, the LLM-judge test,
grades the compiled JSON by calling a provider directly from pytest rather
than through Conductor, so it is not an SDK behaviour to port.

## Suite 2–15 tests

`suite2_…` through `suite15_…_test.go` are the Python SDK's
`e2e/test_suite2_tool_calling.py` through `test_suite15_skills.py`, test for
test and under the same names. Unlike suite 1 these run agents, so each test
that can replay has its recordings under `testdata/llm-recordings/<test file>/`,
and the rest run live only and skip in playback with the reason in the skip
message. `suite_helpers_test.go` holds what the Python suites repeat at module
level: fetching the workflow, finding a tool's task, the secret store calls,
starting mcp-testkit, and the checks on a result.

| Python suite | Go file | Tests | Mode | Needs |
|---|---|---|---|---|
| 2 tool calling | `suite2_tool_calling_test.go` | 1 | replays steps 1–3, then skips | a writable secret store for steps 4–5; the OSS store is read-only, so it skips there as the Python test does |
| 3 CLI tools | `suite3_cli_tools_test.go` | 1 | live only | `GITHUB_TOKEN` in the environment and `gh`; skips at the store write on OSS |
| 4 MCP tools | `suite4_mcp_tools_test.go` | 1 of 2 | replays | `mcp-testkit` on PATH; the test starts it on port 3002. `test_mcp_result_reaches_the_answer` is `TestMCPToolResultReachesTheAnswer` in `mcp_test.go` |
| 5 HTTP tools | `suite5_http_tools_test.go` | 2 | replays | `mcp-testkit` on port 3003; `developer.orkescloud.com` reachable |
| 6 PDF tools | `suite6_pdf_tools_test.go` | 1 | live only | `pdftotext` for the round trip, else it skips after checking the file |
| 7 media tools | `suite7_media_tools_test.go` | 3 | live only | `OPENAI_API_KEY`; `GOOGLE_AI_API_KEY` for Gemini |
| 8 guardrails | `suite8_guardrails_test.go` | 7 | 3 plan-only, 4 replay | — |
| 9 handoffs | `suite9_handoffs_test.go` | 8 | 2 plan-only, 6 replay | — |
| 10 code execution | `suite10_code_execution_test.go` | 9 | 3 plan-only, 5 replay, Jupyter live only | Docker; `python3` with `jupyter_client` and `ipykernel` |
| 11 langgraph | not ported | — | — | Exercises the Python LangGraph adapter (`langgraph`, `langchain`). The Go SDK has no framework adapter, so there is nothing to port. |
| 12 termination and gates | `suite12_termination_gates_test.go` | 5 | 2 plan-only, 3 replay | — |
| 13 callbacks | `suite13_callbacks_test.go` | 5 | 2 plan-only, 3 replay | — |
| 14 stateful domain | `suite14_stateful_domain_test.go` | 6 | live only | — |
| 15 skills | `suite15_skills_test.go` | 4 | 3 plan-only, 1 live only | `bash` |

The server needs two more secrets for suites 4 and 5, with the values the
Python suites use, since the tests start mcp-testkit in auth mode with the
value the store holds: `CONDUCTOR_SECRET_MCP_AUTH_KEY=e2e-test-secret-key-12345`
and `CONDUCTOR_SECRET_HTTP_AUTH_KEY=e2e-http-test-secret-key-67890`.

Suite 15's other tests load and serialize a skill without a server; those are
unit tests of the SDK itself in `sdk/ai/skill_test.go`, and its two standalone
run tests are `TestSkillScriptRunsAsWorker` and `TestSkillAsAgentTool` in
`skill_test.go` here. Suite 13's three run tests were ported earlier; this
round added its two compile-only ones.

Why some run live only:

- Suite 3's `cli_mktemp` returns a fresh temp path every run, which the model
  then reads, so no recording can match; and `cli_gh` needs a real token.
- Suites 6 and 7 hand the model the URL of a freshly generated file.
- Suite 14 reads the run's task-to-domain map, and a stateful run's domain is
  new every time, so nothing about it can be replayed.
- Suite 15's stateful skill test nests an agent tool, whose result carries a
  per-run sub-workflow id.
- pytest reruns every e2e test up to twice. Go does not, and
  `TestJupyterStateful` shows why that mattered: asked to run `print(x * 73)`
  exactly as provided, the model defined its own `x` first in 12 of 13 live
  attempts. It runs live only rather than replay a recording of the one run
  where it did not.
- Suite 5's `TestExternalOpenapiSpec` does replay, but the public Orkes
  document is part of the request; if it changes, the run fails, which the
  test tolerates the way the Python one does.

Where Go differs from Python and how the port handles it:

- Python checks its CLI allow-list validator directly in suite 3; the Go
  validator is covered by the SDK's own `cli_runner_test.go`.
- Python's `markitdown` reads the PDF back; the Go test uses `pdftotext` and
  skips the round trip without it. The OSS server reports the file as a
  `file://` path on its own disk, which the test reads directly.
- Python's `math >> text` pipeline operator has no Go form;
  `TestPipeOperatorSequential` and suite 12's gate tests build the same parent
  by hand.
- Python detects which callback hooks a handler overrides; Go sets them as
  fields on `ai.Callbacks`, which the compile tests read back out of the plan.
- Python marks `test_image_openai` xfail; `TestImageOpenai` skips on the
  known failure and passes if it ever works.
- A tool's own guardrails run in the Go worker around the handler, as the
  Python worker's `run_tool_task` does: input guardrails over the arguments,
  output guardrails over the result, with fix, raise and the blocked marker.
  The server gates only a tool's input, so without this a tool-level output
  guardrail did nothing; `TestToolOutputRegexRetry` is what showed it.
- The unauthenticated and authenticated phases of suites 4 and 5, and steps 2
  and 3 of suite 2, send identical model requests. The recorder refuses two
  different answers for one request, so only the first phase's recordings are
  kept; the second phase replays the same answers and follows the same tool
  path.

## Example tests

The Go ports of the Python SDK's `examples/agents`, run as tests against the
Python SDK's own recordings, so the two SDKs are checked against the same
model traffic and the model is deterministic.

The recordings live in the conductor repository under
`llm-recordings/<example>/`, one folder per example, recorded from the Python
examples with the server's LLM recorder (see the README there). Each
test here runs the agent from one example with the mock model and asserts on
the result the way the other tests here do: the run completes and the final message
is the recorded one. The mock only answers a request that matches a
recording, so completing with the recorded answer also shows the Go example
sent the same request the Python one did.

### Running the example tests

Build the server from the conductor branch that carries the recordings and
start it in playback on the `llm-recordings` directory:

    CONDUCTOR_SERVER_JAR=/path/to/conductor-server-*-boot.jar \
      test/integration_tests/ai/scripts/conductor-server.sh replay \
      /path/to/conductor/llm-recordings

Then run the tests with the same directory in `CONDUCTOR_RECORDINGS_DIR`:

    CONDUCTOR_SERVER_URL=http://localhost:8080/api \
    CONDUCTOR_RECORDINGS_DIR=/path/to/conductor/llm-recordings \
      go test -tags integration -count=1 -v ./test/integration_tests/ai/

The tests skip when either variable is unset.

### Adding an example

Copy the flow of `examples/agents/<name>.py` into `example_<name>_test.go`
here — the same agent, instructions and prompt, character for character — and
replace its printing with assertions, using `recordedAnswers(t, "<name>")`
for what the Python example printed. The recordings only match an identical
request, so any drift from the Python example shows up as a failed run. The
test stands on its own; it does not read `examples/agents/<name>.go`.

## Adding a test

- Build the runtime with `newRuntime(t)` and models with `model(t)` from
  `agent_test.go` so the environment variables above keep working.
- Assert on what the test controls: that a worker ran, with what input, and
  the run's final status. Match model prose loosely or not at all.
- Skip, with a message saying what to start, when an optional dependency is
  missing. Never fail for a missing dependency.
- Keep each run under two minutes with `context.WithTimeout`. A run that times
  out usually means a worker was never registered for a task the server
  scheduled.
