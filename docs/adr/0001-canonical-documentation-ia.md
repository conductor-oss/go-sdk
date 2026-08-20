# 1. Adopt the canonical cross-SDK documentation information architecture

Date: 2026-07-28

## Status

Accepted

## Context

The Conductor SDKs share a documentation information architecture that originated
in the [Java SDK](https://github.com/conductor-oss/java-sdk/tree/main/docs):
canonical page names at `docs/*.md` (`workers.md`, `workflows.md`,
`observability.md`, `security.md`, `connection-authentication.md`, `upgrading.md`,
…) plus a `docs/README.md` hub. The
[Python SDK adopted it](https://github.com/conductor-oss/python-sdk/pull/441) in
July 2026.

This repository had none of it. Its eight guides used repo-local names
(`workers_sdk.md`, `workflow_sdk.md`, `migration_guide.md`, `metrics.md`,
`logger_sdk.md`, and a `docs/api_client/` subtree) with no hub, so a reader or tool
following the convention from another SDK found nothing.

Two premises behind the original request turned out to be wrong, and are recorded
here so they are not re-litigated:

1. **The Python PR was not the source of the structure.** Java was. Python was the
   follower. This repository is a second follower, not a recipient of a Python port.
2. **The `Agentspan` → `Conductor` rebrand does not apply here.** That PR renamed
   `AGENTSPAN_*` environment variables to `CONDUCTOR_AGENT_*` and stripped
   `Agentspan` copyright headers. This repository contains **zero** occurrences of
   `agentspan` in any casing, and `sdk/settings/env.go` is already uniformly
   `CONDUCTOR_*`. Note also that the Python PR description claims it kept legacy
   environment-variable aliases for backward compatibility; the diff shows it
   deleted them. Neither the rename nor the alias question applies to Go.

What remained was a genuine information-architecture gap — but one that could not
be closed by copying files, because roughly half of the canonical page set
documents a Conductor agent runtime this SDK does not provide.

## Decision

Adopt the canonical information architecture for the content that already exists,
and make every remaining gap explicit rather than filling it with skeletons.

1. **Rename to canonical filenames.** Four pure renames; two merges of pages that
   already split one subject across two files (`metrics.md` + `logger_sdk.md` →
   `observability.md`; `api_client/README.md` + `api_client/proxy_configuration.md`
   → `connection-authentication.md`).
2. **Keep redirect stubs at the two externally-linked paths.** `conductor-oss/conductor`
   hardlinks `docs/workers_sdk.md` and `docs/api_client/README.md`. Deleting them
   would 404 the published OSS documentation. Stubs are pointer-only and are
   removed once the upstream reference is updated.
3. **Do not create `docs/agents/`.** The absence is a capability gap, stated in
   `compatibility.md`, `documentation-parity.md`, and the hub. An empty or
   "unsupported" agent tree would document a non-feature and create a namespace
   that must later be filled or deleted.
4. **Adopt `documentation-standard.md` forward-looking.** It binds new and edited
   pages. Pre-existing pages do not conform; that is declared in the standard
   itself and tracked, not left silent.
5. **Author nothing that does not already exist.** The twelve missing canonical
   pages are filed as issues #266–#272 rather than written as stubs.
6. **Validate in the native idiom.** Structural assertions live in a Go test that
   contributors run with `go test ./...`; link liveness uses the same lychee action
   as the sibling repositories rather than a hand-rolled checker.

## Consequences

- A reader or tool following the Java/Python convention now finds the same page
  names here, for the subjects Go covers.
- The repository ships a documentation standard its own older pages do not meet.
  This is a deliberate, declared trade: fusing a mechanical rename with editorial
  rewrites would have produced an unreviewable diff.
- Two redirect stubs exist that must be cleaned up after an upstream change. A test
  asserts they stay pointer-only so they cannot silently grow back into guides.
- `security.md` claims a canonical name while covering TLS only. Java's equivalent
  also covers secrets and RBAC. Tracked in #272.
- The parity page carries a coverage statement — 17 of 20 client interfaces
  undocumented — so the state of the documentation is a recorded fact rather than
  something a reader has to infer.
- Anchors that `CHANGELOG.md` points into were preserved verbatim through the
  `metrics.md` → `observability.md` rename. Future merges must do the same or
  update the referring links in the same change.
