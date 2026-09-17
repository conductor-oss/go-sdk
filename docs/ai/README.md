# Agents — Design Docs (Go SDK)

Status: **implemented on `feat/agent-golden-fixtures`, under review.** These documents describe
the port of the python-sdk "Agents" feature (`conductor.ai.agents`) to the Go SDK as `sdk/ai`.
They follow the same set of documents the Rust SDK used to plan its port
([conductor-oss/rust-sdk#10](https://github.com/conductor-oss/rust-sdk/pull/10)), so the SDKs
can be compared side by side — but where the Rust docs propose, these record what was built and
what was verified against a live server.

Source of truth for behaviour is the python-sdk. java-sdk (`conductor-client-ai`) is the
secondary typed-language reference and was diffed against directly where it mattered.

| Doc | Purpose |
|---|---|
| [`go-sdk-design.md`](go-sdk-design.md) | The Go types, package layout, how they map onto the existing worker framework, and where Go deliberately differs. |
| [`secrets-and-credentials.md`](secrets-and-credentials.md) | The server→worker credential delivery contract and the Go surface over it. Flagged up front because it is **not** a 1:1 port of Python's env-var mechanics. |
| [`framework-support.md`](framework-support.md) | Which external agent frameworks Go supports, which it does not, and why. |
| [`examples.md`](examples.md) | Three worked examples of the shipped API, each with the e2e test that exercises it. |
| [`parity-plan.md`](parity-plan.md) | One-page summary (classes / examples / secrets / frameworks) in the cross-SDK comparison format. |

## How to read these

Start with `go-sdk-design.md`. The other docs go deep on the parts that needed the most
scrutiny — credentials, frameworks — and on what is actually exercised end to end. The Python
SDK remains the behavioural source of truth, but it is documented in its own repo, not here.

## How it was verified

Two bars, both enforced by tests in this repo:

- **Wire conformance.** `sdk/ai/testdata/agent_config/` holds `agentConfig` documents captured
  verbatim from the Python serializer. `TestGoldenAgentConfig` builds the equivalent Go agent for
  each and requires the same document (compared semantically, so key order is free). Every
  fixture is implemented; `TestGoldenCoverage` reports the current count.
- **Behaviour.** `test/ai_e2e/` runs its scenarios against a live Conductor server — tools,
  streaming with human approval, credentials, code execution, CLI, guardrails, handoffs, MCP,
  plan-execute — and each asserts the effect on the worker side (a counter incremented, a value
  bound), not just that the run finished. Run with `-tags e2e` and `CONDUCTOR_SERVER_URL` set.

## Open items surfaced while building this

These are server- or sibling-SDK-side, found because Go's e2e suite checks things the other suites
do not. The first is covered in `go-sdk-design.md`; the rest are recorded here.

1. **`on_condition` handoffs stall in python-sdk and java-sdk against current servers.** The server
   moved to one worker per condition (`{agent}_handoff_{target}`) in July 2026; both SDKs still
   register the retired `{agent}_handoff_check`. Go registers the per-condition worker. See
   `go-sdk-design.md` § Derived worker names.
2. **Plan-execute could not pass a typed value between steps.** Root-caused to the server's planner
   prompt and compiler; a fix is on a conductor branch. Go's `TestPlanExecute` is the test that
   catches it, because it chains an `int` between tools where the Python suite chains only strings.
3. **`Position` on guardrails is not acted on by current servers** at either level. Kept for wire
   parity; documented honestly on the constant.
4. **`RequiredTools` is kept for parity but is not usable on current servers.**
