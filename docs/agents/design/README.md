# Agents — design notes (Go SDK)

How `sdk/ai` was built and verified. These are build records, not user documentation; start from
the [agent docs index](../README.md) if you are using the SDK rather than changing it.

The python-sdk is the source of truth for behaviour. The java-sdk (`conductor-client-ai`) is the
secondary typed-language reference and was diffed against directly where it mattered.

| Doc | Purpose |
|---|---|
| [`go-sdk-design.md`](go-sdk-design.md) | The Go types, package layout, how they map onto the existing worker framework, and where Go deliberately differs. |
| [`parity-plan.md`](parity-plan.md) | One-page summary in the cross-SDK comparison format. |
| [`remaining-work-plan.md`](remaining-work-plan.md) | What was built, in order, and what was deferred. |

## How it was verified

Two bars, both enforced by tests in this repo:

- **Wire conformance.** `sdk/ai/testdata/agent_config/` holds `agentConfig` documents captured
  verbatim from the Python serializer. The golden test builds the equivalent Go agent for each and
  requires the same document, compared semantically so key order is free.
- **Behaviour.** `examples/agents/` are exact ports of the Python examples, and CI replays every
  one against the recordings shared in conductor-oss/conductor with the mock model. Completing
  against Python's recording is what shows the two SDKs send the same requests; the conductor
  repository's `check-playback` action judges the outcomes, so no example carries its own
  pass or fail rule.

## Open items surfaced while building this

Server- or sibling-SDK-side, found because Go's suite checks things the other suites do not.

1. **`on_condition` handoffs stall in python-sdk and java-sdk against current servers.** The server
   moved to one worker per condition (`{agent}_handoff_{target}`); both SDKs still register the
   retired `{agent}_handoff_check`. Go registers the per-condition worker. See
   [`go-sdk-design.md`](go-sdk-design.md).
2. **Plan-execute could not pass a typed value between steps.** Root-caused to the server's planner
   prompt and compiler. Go's plan-execute test catches it because it chains an `int` between tools
   where the Python suite chains only strings.
3. **`Position` on guardrails is not acted on by current servers** at either level. Kept for wire
   parity and documented on the constant.
4. **`RequiredTools` is kept for parity but is not usable on current servers.**
