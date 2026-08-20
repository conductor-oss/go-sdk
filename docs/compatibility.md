# Compatibility matrix

**Audience:** teams selecting runtime baselines and deciding which capabilities
this SDK can provide.
**Last verified:** repository `go.mod`, `Dockerfile`, and CI workflow configuration.

| Area | Supported baseline | Notes |
|---|---|---|
| Go SDK | Go 1.23+ | `go.mod` declares `go 1.23`; CI builds and tests on 1.23. |
| OSS Conductor | Supported server deployment | Test your target server during an upgrade; the CI matrix is the exercised baseline. |
| Orkes | Supported tenant API | Set the tenant API endpoint and credentials; availability of enterprise features is tenant-specific. |
| Exercised server matrix | Orkes SM clusters `v4` and `v5` | `.github/workflows/integration-tests-sm.yml` runs integration and backward-compatibility suites against both. |
| Backward compatibility | Against the last released tag | `test/backward_compatibility/` compiles the current tree and a released version side by side. |

## Published module

| Module | Purpose | Reference |
|---|---|---|
| `github.com/conductor-sdk/conductor-go` | Single module: workflow, task, metadata, scheduler, schema, secret, RBAC, human-task, prompt, and integration clients, plus the worker runtime and metrics | [pkg.go.dev](https://pkg.go.dev/github.com/conductor-sdk/conductor-go) |

Install with `go get github.com/conductor-sdk/conductor-go` rather than pinning a
version literal from a guide. Generated pkg.go.dev documentation is the signature
reference; linked source in this repository remains the fallback.

## Capability differences from the Java and Python SDKs

| Capability | Java / Python | Go |
|---|---|---|
| Workflow and worker clients | Supported | Supported |
| Schema, scheduler, secret, RBAC, human-task clients | Supported | Present in `sdk/client`, not yet documented |
| LLM prompt templates and provider integrations | Supported | `PromptClient` and `IntegrationClient` present, not yet documented |
| **Conductor agent runtime** | Supported | **Not provided** — no agent runtime, definitions, tools, or framework bridges |
| Workflow-scoped `FileClient` | Java only | Not exposed |
| Spring / Spring Boot integration | Java only | Not applicable; use Go application hosting patterns |

The agent-runtime row is the one that most often causes confusion: the absence of
`docs/agents/` in this repository reflects a missing *runtime*, not merely missing
documentation. Prompt and integration support does exist. See
[documentation parity](documentation-parity.md).

Next: [connection and authentication](connection-authentication.md),
[core quickstart](core-quickstart.md), and [upgrading](upgrading.md).
