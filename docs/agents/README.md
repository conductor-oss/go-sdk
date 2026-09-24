# Conductor Go Agent SDK

Build durable Go AI agents on Conductor. An agent is described as a Go struct, compiled by the
server into a workflow, and executed like any other workflow, so every model turn, tool call and
handoff is a task that is persisted, retryable and visible in the Conductor UI. Tools are ordinary
Go functions the SDK registers as workers.

**New here?** Follow [Getting Started](getting-started.md) to point the SDK at a server and run an
agent.

> [!note]
> Agents need a Conductor server with an LLM provider configured **on the server**. Your Go process
> never calls the model provider, so no provider API key is needed locally.

## Install

Requirements: Go 1.23+ and a Conductor server.

```shell
go get github.com/conductor-sdk/conductor-go
```

## Start here

- **[Getting Started](getting-started.md)** — configure a server and run your first agent.
- **[Deploy · Serve · Run](concepts/deploy-serve-run.md)** — choose the right runtime mode.
- **[Scheduling](concepts/scheduling.md)** — run a deployed agent on a cron cadence.

## Build agents

- **[Agents](concepts/agents.md)** — the `Agent` struct and its fields.
- **[Tools](concepts/tools.md)** — Go functions, HTTP and MCP tools, human approval, media, and credentials.
- **[Multi-Agent](concepts/multi-agent.md)** — handoff, sequential, parallel, router, swarm and plan-execute.
- **[Guardrails](concepts/guardrails.md)**, **[Termination](concepts/termination.md)**, **[Callbacks](concepts/callbacks.md)**, **[Stateful Agents](concepts/stateful.md)**, **[Streaming & Human-in-the-Loop](concepts/streaming-hitl.md)**, and **[Structured Output](concepts/structured-output.md)**.

## Reference

- **[Runtime](reference/runtime.md)** — every `Runtime` and `AgentHandle` method, `Config`, and run options.
- **[Agent definition](reference/agent-definition.md)** — every `Agent` and `ToolDef` field, with defaults and wire keys.
- **[Client](reference/client.md)** — the lower-level control plane, including starting a deployed agent by name.

## Framework bridges

The Go SDK ships none, deliberately. MCP servers are supported directly through `tool.MCP`, since
MCP is a protocol the server speaks rather than a dependency. See
**[Framework support](framework-support.md)** for the policy and what to do instead.

## More

- **[Worked examples](examples.md)** — annotated examples with the tests that exercise them.
- **[Secrets and credentials](secrets-and-credentials.md)** — how the server delivers secrets to your tools.
- **[Runnable examples](../../examples/agents/)** — the maintained example programs.
- **[Design notes](design/README.md)** — how the port was built and verified, and what is still open.
