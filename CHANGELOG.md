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
- `RecordTaskPollTime` / `RecordTaskExecuteTime` / `RecordTaskUpdateTime` package-level functions retain their original 2-argument signatures (now deprecated). The new status-aware 3-argument signatures are available on the `MetricsCollector` interface via `metrics.GetCollector()` -- see [docs/metrics.md](docs/metrics.md#detailed-technical-notes----unreleased)
- `IncrementUncaughtException` parameter type changed from `string` to `interface{}`; existing callers passing string values are unaffected (string satisfies `interface{}`)

### Fixed

- `metrics.PayloadType.TASK_OUTPUT` constant value (was `"TASK_INPUT"`)
- `IncrementTaskPaused` was previously never called
- `IncrementWorkflowStartError` was previously never called; the `workflow_start_error` counter will now increment on workflow start failures in both legacy and canonical mode
- `thread_uncaught_exceptions` counter now increments correctly in legacy mode (was silently failing due to a label-count mismatch)
- `JumpToTask` path had a literal `{taskReferenceName}` in the URL string (now removed; the parameter is correctly passed as a query param)

### Deprecated

- Legacy metric names remain the default. Migration guidance is in [docs/metrics.md](docs/metrics.md#migration-from-legacy-to-canonical).
- `RecordTaskPollTime(taskType, seconds)`, `RecordTaskExecuteTime(taskType, seconds)`, `RecordTaskUpdateTime(taskType, seconds)` package-level functions. Use `metrics.GetCollector().RecordTaskPollTime(taskType, seconds, err)` (and equivalents) for status-aware recording.
