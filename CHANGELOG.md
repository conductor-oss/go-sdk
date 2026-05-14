# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Canonical metrics: opt-in harmonized metric surface via `WORKER_CANONICAL_METRICS=true` -- see [docs/metrics.md](docs/metrics.md) for the full catalog, configuration, and migration guide
- `WORKER_METRICS_PAYLOAD_SIZE` and `WORKER_METRICS_HTTP_REQUESTS` env vars to selectively disable expensive canonical metrics
- Bounded `uri` label on `http_api_client_request_seconds`: the `uri` label now uses path templates (e.g. `/workflow/{workflowId}`) instead of fully-resolved paths, preventing metric cardinality explosion from dynamic IDs
- `WorkflowStatusProbe` in harness: opt-in probe (via `HARNESS_PROBE_RATE_PER_SEC`) that exercises UUID-bearing endpoints to validate template URI metrics

### Changed

- Legacy metrics emit unchanged by default; no action required for existing deployments
- `RecordTaskPollTime` / `RecordTaskExecuteTime` / `RecordTaskUpdateTime` signatures now accept `(seconds, error)` -- see [docs/metrics.md](docs/metrics.md#detailed-technical-notes----unreleased)

### Fixed

- `metrics.PayloadType.TASK_OUTPUT` constant value (was `"TASK_INPUT"`)
- `IncrementTaskPaused` was previously never called
- `thread_uncaught_exceptions` counter now increments correctly in legacy mode (was silently failing due to a label-count mismatch)
- `JumpToTask` path had a literal `{taskReferenceName}` in the URL string (now removed; the parameter is correctly passed as a query param)

### Deprecated

- Legacy metric names remain the default. Migration guidance is in [docs/metrics.md](docs/metrics.md#migration-from-legacy-to-canonical).
