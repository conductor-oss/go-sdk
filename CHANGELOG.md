# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **Metrics harmonization** - canonical metric surface aligned with the cross-SDK catalog, opt-in via `WORKER_CANONICAL_METRICS=true`
  - New `canonicalCollector` emits the harmonized cross-SDK catalog: `_total`-suffixed counters, duration histograms in seconds (`task_poll_time_seconds`, `task_execute_time_seconds`, `task_update_time_seconds`, `http_api_client_request_seconds{method,uri,status}`) with buckets `0.001…10s`, size histograms in bytes (`task_result_size_bytes`, `workflow_input_size_bytes{workflowType,version}`) with buckets `100…10_000_000`, and an `active_workers{taskType}` gauge. Labels are camelCase; `status` is `SUCCESS`/`FAILURE` for task histograms and the HTTP status code (or `"0"` on transport failure) for the HTTP histogram.
  - `metrics.NewCollector()` factory and `metrics.InitCollector()` (decoupled from the HTTP server) select between `legacyCollector` (default) and `canonicalCollector` based on `WORKER_CANONICAL_METRICS` (truthy: `true`, `1`, `yes`, case-insensitive, whitespace-trimmed). `WORKER_LEGACY_METRICS` is reserved for a future default-flip phase and is not currently read.
  - New `metricsRoundTripper` wraps the API client transport to record `http_api_client_request_seconds` (with `status="0"` for network errors).
  - Worker instrumentation: `IncrementTaskExecutionStarted`, `IncrementTaskPaused`, `SetActiveWorkers` on enter/exit, `recordTaskResultPayloadSize`, `recordWorkflowInputPayloadSize`, `IncrementWorkflowStartError`.
  - Harness manifest sets `WORKER_CANONICAL_METRICS=true`; `harness/main.go` logs which collector is active.

### Changed

- **Metrics harmonization** - defaults preserved; legacy metrics emit unchanged when `WORKER_CANONICAL_METRICS` is unset
  - `RecordTaskPollTime` / `RecordTaskExecuteTime` / `RecordTaskUpdateTime` signatures now take `(seconds float64, err error)`. Legacy gauge units are preserved at runtime (seconds for poll; milliseconds for execute/update).
  - `IncrementUncaughtException` now takes the recovered value so the canonical collector can derive a bounded-cardinality `exception` label.
  - API client transport is wrapped with `metricsRoundTripper` regardless of mode; the legacy collector treats the HTTP request observation as a no-op.
  - Default behavior is unchanged: with no env var set, the legacy metric names (no `_total` suffix, no `exception` label, snake_case `payload_type`) and shapes shipped in v1.9.0 are preserved.
  - New `docs/metrics.md` covering usage, env-var selection, full canonical and legacy catalogs, label semantics, metrics not applicable to Go (`worker_restart_total`, `task_ack_*_total`, `task_execution_queue_full_total`), legacy → canonical migration with PromQL recipes, and troubleshooting.
  - `docs/workers_sdk.md` now points at `metrics.md` instead of inlining the catalog.

### Fixed

- `metrics.PayloadType.TASK_OUTPUT` constant value (was `"TASK_INPUT"`).
- `IncrementTaskPaused` was previously never called.

### Deprecated

- Legacy gauges and counters remain the default and continue to emit unchanged. Migration to the canonical surface is documented in `docs/metrics.md`.
