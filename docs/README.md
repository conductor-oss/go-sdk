# Go SDK documentation

Build durable workflow workers and orchestration clients with Conductor. This
repository documents both **OSS** and **Orkes** in place; pages call out
capabilities that need Orkes.

## Start here

| Goal | Guide | Expected result |
|---|---|---|
| Connect to a server | [Connection and authentication](connection-authentication.md) | A local or remote API endpoint accepts SDK requests. |
| Build a workflow and Go worker | [Core quickstart](core-quickstart.md) | Hello World prints a greeting from the workflow output. |
| Check version and capability baselines | [Compatibility](compatibility.md) | You know which Go and server versions are exercised. |

## Build

- [Workflows](workflows.md) — authoring workflow definitions in Go
- [Workers](workers.md) — writing and running task workers
- [Core quickstart](core-quickstart.md) — the end-to-end Hello World

## Operate

- [Security](security.md) — TLS, mTLS, and certificate configuration
- [Observability](observability.md) — Prometheus metrics and logging
- [Upgrading](upgrading.md) — deprecated methods and their replacements

## About this documentation

- [Compatibility](compatibility.md) — supported baselines and the published module
- [Documentation standard](documentation-standard.md) — what every guide must contain
- [Documentation parity](documentation-parity.md) — how this set maps to the Java and Python SDKs
- [Architecture decisions](adr/) — recorded decisions about this repository

## AI agents

The Go SDK does **not** provide the Conductor agent runtime that the
[Java](https://github.com/conductor-oss/java-sdk/tree/main/docs/agents) and
[Python](https://github.com/conductor-oss/python-sdk/tree/main/docs/agents) SDKs
document, so there is no `docs/agents/` section here. Use the Java or Python SDK
to build agents.

Go does expose the AI-orchestration primitives — `PromptClient` for LLM prompt
templates and `IntegrationClient` for provider configuration. Both are currently
undocumented; see [documentation parity](documentation-parity.md) and
[#268](https://github.com/conductor-oss/go-sdk/issues/268).
