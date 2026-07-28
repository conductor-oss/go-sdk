# Observability

Metrics and logging for the Conductor Go SDK. For metrics see below; for logging
jump to [Logging Configuration](#logging-configuration).

The Conductor Go SDK can expose Prometheus metrics for worker polling, task
execution, task result updates, payload sizes, workflow starts, and HTTP API
client latency.

The SDK currently has two mutually exclusive metric surfaces:

- **Legacy metrics** are the default. They preserve the original Go SDK names
  and types, including last-value Gauges for timing and size metrics.
- **Canonical metrics** are opt-in with `WORKER_CANONICAL_METRICS=true`. They
  use the cross-SDK canonical names, labels, units, and Prometheus Histogram
  types.

Only one collector is active at a time. The SDK does not emit legacy and
canonical metrics simultaneously.

## Usage

Start the Prometheus metrics HTTP server in a goroutine. The default settings
expose port 2112 on the `/metrics` endpoint:

```go
go metrics.ProvideMetrics(settings.NewDefaultMetricsSettings())
```

`ProvideMetrics` calls `InitCollector()` internally if metrics have not already
been initialized. `InitCollector` reads the environment variables, creates the
appropriate collector, registers Prometheus metrics, and enables collection.

Running `ProvideMetrics` in a goroutine is concurrency-safe even while workers
are already polling: the collector is published atomically, so readers never
race with initialization. The only trade-off is that metrics emitted before
initialization completes are dropped. If you need first-poll coverage, call
`metrics.InitCollector()` synchronously before starting workers (see below).

To customize the port or endpoint:

```go
metricsSettings := &settings.MetricsSettings{
    ApiEndpoint: "/metrics",
    Port:        9090,
}
go metrics.ProvideMetrics(metricsSettings)
```

To initialize the collector before starting the HTTP server (useful when the
collector must be ready before the first poll):

```go
metrics.InitCollector()
go metrics.ProvideMetrics(nil)
```

## Selecting Canonical Metrics

Set `WORKER_CANONICAL_METRICS` before the worker starts:

```shell
WORKER_CANONICAL_METRICS=true ./my-worker
```

| Environment variable | Values | Effect |
|---|---|---|
| `WORKER_CANONICAL_METRICS` | `true`, `1`, or `yes` (case-insensitive, surrounding whitespace ignored) | Selects the canonical collector. |
| `WORKER_CANONICAL_METRICS` | unset, blank, `false`, `0`, `no`, or any other value | Selects the legacy collector. |

The variable is read once when the collector is created. Changing it requires a
worker restart.

`WORKER_LEGACY_METRICS` is reserved for a future default-flip phase and is not
currently read by the Go SDK factory.

### Disabling Expensive Canonical Metrics

Two optional environment variables control expensive metric categories within
canonical mode. Both default to `true` in canonical mode and have no effect in
legacy mode (which never emits these metrics).

| Environment variable | Default (canonical) | Effect when `false` |
|---|---|---|
| `WORKER_METRICS_PAYLOAD_SIZE` | `true` | Skips `json.Marshal` of task results and workflow inputs for size histograms (`task_result_size_bytes`, `workflow_input_size_bytes`). |
| `WORKER_METRICS_HTTP_REQUESTS` | `true` | Skips per-request timing in the HTTP transport layer (`http_api_client_request_seconds`). |

These variables are read once when the collector is created. Like
`WORKER_CANONICAL_METRICS`, changing them requires a worker restart.

## Canonical Metrics Catalog

Canonical timing values are seconds. Canonical size values are bytes. Label
names use camelCase. Exception labels use bounded Go type names
(`fmt.Sprintf("%T", err)`), not error messages.

### Canonical Counters

| Metric | Labels | Description |
|---|---|---|
| `task_poll_total` | `taskType` | Incremented each time the worker issues a poll request. |
| `task_execution_started_total` | `taskType` | Incremented when a polled task is dispatched to the worker function. |
| `task_poll_error_total` | `taskType`, `exception` | Incremented when a poll request fails client-side. |
| `task_execute_error_total` | `taskType`, `exception` | Incremented when the worker function returns an error. |
| `task_update_error_total` | `taskType`, `exception` | Incremented when reporting a task result back to Conductor fails. |
| `task_ack_error_total` | `taskType`, `exception` | Registered as API surface. The internal runner uses batch-poll responses as ack and does not emit this during normal polling. |
| `task_ack_failed_total` | `taskType` | Registered as API surface. The internal runner uses batch-poll responses as ack and does not emit this during normal polling. |
| `task_execution_queue_full_total` | `taskType` | Registered as API surface. Go dispatches onto goroutines with no bounded queue, so the internal runner does not emit this. |
| `task_paused_total` | `taskType` | Incremented when a worker is paused and skips acting on a poll. |
| `thread_uncaught_exceptions_total` | `exception` | Incremented when a worker goroutine terminates with an uncaught panic. |
| `external_payload_used_total` | `entityName`, `operation`, `payloadType` | Incremented when external payload storage is used for task or workflow payloads. |
| `workflow_start_error_total` | `workflowType`, `exception` | Incremented when starting a workflow fails client-side. |

### Canonical Time Histograms

All canonical time histograms use buckets (in seconds):

```text
0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10
```

| Metric | Labels | Description |
|---|---|---|
| `task_poll_time_seconds` | `taskType`, `status` | Poll request latency. `status` is `SUCCESS` or `FAILURE`. |
| `task_execute_time_seconds` | `taskType`, `status` | Worker function execution duration. `status` is `SUCCESS` or `FAILURE`. |
| `task_update_time_seconds` | `taskType`, `status` | Latency for reporting a task result back to Conductor. `status` is `SUCCESS` or `FAILURE`. |
| `http_api_client_request_seconds` | `method`, `uri`, `status` | Latency of HTTP requests made by the API client. `status` is the HTTP status code as a string, or `"0"` on network failure. |

Each histogram exposes Prometheus series such as:

```prometheus
task_execute_time_seconds_bucket{taskType="my_task",status="SUCCESS",le="0.1"} 42
task_execute_time_seconds_count{taskType="my_task",status="SUCCESS"} 50
task_execute_time_seconds_sum{taskType="my_task",status="SUCCESS"} 2.3
```

The `uri` label for `http_api_client_request_seconds` uses the path template
(e.g. `/tasks/poll/batch/{tasktype}`) rather than the fully-resolved request
path. This keeps label cardinality bounded to the number of distinct API
endpoints, regardless of how many unique workflow IDs, task IDs, or other
dynamic segments appear in actual requests.

### Canonical Size Histograms

All canonical size histograms use buckets (in bytes):

```text
100, 1000, 10000, 100000, 1000000, 10000000
```

| Metric | Labels | Description |
|---|---|---|
| `task_result_size_bytes` | `taskType` | Serialized task result output size. |
| `workflow_input_size_bytes` | `workflowType`, `version` | Serialized workflow input size. `version` is the workflow version as a string. |

### Canonical Gauges

| Metric | Labels | Description |
|---|---|---|
| `active_workers` | `taskType` | Current number of worker goroutines actively executing a task. |

## Legacy Metrics Catalog

Legacy mode is the default so existing dashboards and alerts continue to work.
Timing metrics are last-value Gauges (not Histograms). Error counters do not
carry an `exception` label.

### Legacy Counters

| Metric | Labels | Description |
|---|---|---|
| `task_poll` | `taskType` | Incremented each time polling is done. |
| `task_poll_error` | `taskType` | Incremented when a poll request fails. |
| `task_execute_error` | `taskType` | Incremented when the worker function returns an error. |
| `task_update_error` | `taskType` | Incremented when updating the task result fails. |
| `task_execution_queue_full` | `taskType` | Incremented when the execution queue is saturated. |
| `task_paused` | `taskType` | Incremented when a worker is paused and skips a poll. |
| `thread_uncaught_exceptions` | none | Uncaught exceptions in worker goroutines. |
| `external_payload_used` | `entityName`, `operation`, `payload_type` | External payload storage usage. |
| `workflow_start_error` | `workflowType` | Workflow start errors. |

### Legacy Gauges

| Metric | Labels | Unit | Description |
|---|---|---|---|
| `task_poll_time` | `taskType` | seconds | Most recent poll duration. |
| `task_execute_time` | `taskType` | milliseconds | Most recent task execution duration. |
| `task_update_time` | `taskType` | milliseconds | Most recent task result update duration. |
| `task_result_size` | `taskType` | bytes | Most recent serialized task result output size. |
| `workflow_input_size` | `workflowType`, `version` | bytes | Most recent serialized workflow input size. |

Legacy `task_execute_time` and `task_update_time` record values in milliseconds.
This is a historical unit inconsistency with `task_poll_time` which records in
seconds. The canonical collector corrects this -- all canonical time metrics use
seconds.

Legacy mode does not emit `task_execution_started_total`,
`http_api_client_request_seconds`, or `active_workers`.

## Metrics Not Applicable to Go

The cross-SDK canonical catalog defines additional metrics that are not
applicable to the Go SDK's runtime model:

| Canonical metric | Why N/A for Go |
|---|---|
| `worker_restart_total` | Python-only. Its multi-process supervisor restarts child processes. Go uses goroutines. |
| `task_ack_error_total` | Batch-poll response is the ack; there is no separate ack call. Registered as API surface for user code that layers on its own ack semantics. |
| `task_ack_failed_total` | Same reason as above. |
| `task_execution_queue_full_total` | Go dispatches onto goroutines with no bounded queue. Registered as API surface. |

Users cross-referencing the harmonization spec or documentation from other
Conductor SDKs may notice these metrics in other catalogs. Their absence in
normal Go SDK worker output is intentional.

## Labels

| Label | Used by | Values |
|---|---|---|
| `taskType` | Worker metrics (both legacy and canonical) | Task definition name. |
| `workflowType` | Workflow metrics | Workflow definition name. |
| `version` | `workflow_input_size`, `workflow_input_size_bytes` | Workflow version as a string. |
| `status` | Canonical task time histograms | `SUCCESS` or `FAILURE`. For `http_api_client_request_seconds`, the HTTP status code as a string, or `"0"` on network failure. |
| `exception` | Canonical error counters | Go type name of the error, such as `*url.Error`. |
| `entityName` | `external_payload_used` / `external_payload_used_total` | Task type or workflow name associated with the external payload. |
| `operation` | `external_payload_used` / `external_payload_used_total` | External payload operation, such as `READ` or `WRITE`. |
| `payload_type` | Legacy `external_payload_used` | Payload type, such as `TASK_INPUT`, `TASK_OUTPUT`, `WORKFLOW_INPUT`, or `WORKFLOW_OUTPUT`. |
| `payloadType` | Canonical `external_payload_used_total` | Payload type, such as `TASK_INPUT`, `TASK_OUTPUT`, `WORKFLOW_INPUT`, or `WORKFLOW_OUTPUT`. |
| `method` | HTTP metrics | HTTP verb. |
| `uri` | HTTP metrics | Request path template (e.g. `/workflow/{workflowId}`). Bounded cardinality. |

## Migration from Legacy to Canonical

Switching to canonical metrics is an explicit metrics-surface cutover. Enable
`WORKER_CANONICAL_METRICS=true` in a lower environment first, then update
dashboards, recording rules, and alerts before enabling it in production.

Key changes:

- Legacy timing metrics are last-value Gauges. Canonical timing metrics are
  Histograms with bucket boundaries. Query `_bucket` series with
  `histogram_quantile()` instead of reading Gauge values directly.
- Legacy `task_execute_time` and `task_update_time` record milliseconds.
  Canonical `task_execute_time_seconds` and `task_update_time_seconds` record
  seconds.
- Legacy size metrics are last-value Gauges. Canonical size metrics are
  Histograms.
- Canonical counters add a `_total` suffix.
- Canonical error counters add an `exception` label containing the Go type name
  of the error.
- Canonical time histograms add a `status` label (`SUCCESS` or `FAILURE`).
- Canonical mode adds metrics that legacy mode never emits:
  `task_execution_started_total`, `http_api_client_request_seconds`, and
  `active_workers`.
- Legacy uses `payload_type`; canonical uses `payloadType`.
- Canonical and legacy collectors are mutually exclusive. During a migration,
  compare scrape output by running separate worker instances or environments
  with and without `WORKER_CANONICAL_METRICS=true`.

Legacy-to-canonical replacements:

| Legacy metric | Canonical replacement |
|---|---|
| `task_poll{taskType}` | `task_poll_total{taskType}` |
| `task_poll_error{taskType}` | `task_poll_error_total{taskType,exception}` |
| `task_execute_error{taskType}` | `task_execute_error_total{taskType,exception}` |
| `task_update_error{taskType}` | `task_update_error_total{taskType,exception}` |
| `task_execution_queue_full{taskType}` | `task_execution_queue_full_total{taskType}` |
| `task_paused{taskType}` | `task_paused_total{taskType}` |
| `thread_uncaught_exceptions` | `thread_uncaught_exceptions_total{exception}` |
| `external_payload_used{entityName,operation,payload_type}` | `external_payload_used_total{entityName,operation,payloadType}` |
| `workflow_start_error{workflowType}` | `workflow_start_error_total{workflowType,exception}` |
| `task_poll_time{taskType}` (Gauge, seconds) | `task_poll_time_seconds{taskType,status}` (Histogram, seconds) |
| `task_execute_time{taskType}` (Gauge, milliseconds) | `task_execute_time_seconds{taskType,status}` (Histogram, seconds) |
| `task_update_time{taskType}` (Gauge, milliseconds) | `task_update_time_seconds{taskType,status}` (Histogram, seconds) |
| `task_result_size{taskType}` (Gauge) | `task_result_size_bytes{taskType}` (Histogram) |
| `workflow_input_size{workflowType,version}` (Gauge) | `workflow_input_size_bytes{workflowType,version}` (Histogram) |

Common PromQL replacements:

| Legacy | Canonical |
|---|---|
| `task_poll_time{taskType="my_task"}` | `histogram_quantile(0.95, sum by (le, taskType, status) (rate(task_poll_time_seconds_bucket[5m])))` |
| `task_execute_time{taskType="my_task"}` (ms) | `histogram_quantile(0.95, sum by (le, taskType, status) (rate(task_execute_time_seconds_bucket[5m])))` |
| `task_result_size{taskType="my_task"}` | `task_result_size_bytes_bucket`, `_count`, and `_sum` |
| `workflow_input_size{workflowType="my_wf"}` | `workflow_input_size_bytes_bucket`, `_count`, and `_sum` |
| `external_payload_used{payload_type="TASK_OUTPUT"}` | `external_payload_used_total{payloadType="TASK_OUTPUT"}` |

Average latency queries use `_sum` divided by `_count`. The canonical series
are cumulative histogram counters:

```promql
sum(rate(task_execute_time_seconds_sum[5m])) by (taskType)
/
sum(rate(task_execute_time_seconds_count[5m])) by (taskType)
```

## Troubleshooting

### Metrics Are Empty

- Verify that `metrics.ProvideMetrics` or `metrics.InitCollector` is called
  before workers begin polling.
- Verify workers have polled or executed tasks. Metrics are recorded when the
  relevant code path runs.
- Confirm the scrape endpoint is reachable at the expected host and port
  (default: `http://localhost:2112/metrics`).

### Missing HTTP Metrics

- `http_api_client_request_seconds` requires canonical mode. Legacy mode does
  not emit HTTP metrics. The canonical collector records HTTP latency via a
  `RoundTripper` wrapper on the API client.

### thread_uncaught_exceptions Was Always Zero

Prior to this release, the legacy `thread_uncaught_exceptions` counter never
incremented due to a label-count mismatch bug. After upgrading, the counter
will begin recording actual values. Alerts or dashboards that assumed this
counter was always zero should be updated.

### High Cardinality

- The `uri` label on `http_api_client_request_seconds` now uses path templates
  (e.g. `/workflow/{workflowId}`), so cardinality is bounded to the number of
  distinct API endpoints. If an API call site does not set a template, the
  round-tripper falls back to the raw request path.
- Prefer canonical mode for bounded `exception` labels using Go type names
  instead of error messages.
- Avoid embedding user identifiers or unbounded values in task type, workflow
  type, or external payload labels.

## Logging Configuration

The SDK uses a configurable logging system that allows you to customize the logging behavior to fit your application's needs.

### Log Levels

The default log level is **INFO**, which includes INFO, WARN, ERROR, and FATAL messages. You can change the log level by setting the `LOG_LEVEL` environment variable:

```shell
# Show all messages including debug output
LOG_LEVEL=debug go run ./your-app

# Default — show info and above (same as not setting LOG_LEVEL)
LOG_LEVEL=info go run ./your-app

# Only warnings and errors
LOG_LEVEL=warn go run ./your-app

# Only errors
LOG_LEVEL=error go run ./your-app
```

The environment variable is case-insensitive. When using a custom logger (see below), log level filtering is handled by that logger's own configuration.

### Using Custom Logger

You can configure a custom logger by calling the `SetLogger` function with any implementation of the `Logger` interface:

```go
import "github.com/conductor-sdk/conductor-go/sdk/log"

// Set your custom logger
log.SetLogger(yourCustomLogger)

// Reset to default logger
log.SetLogger(nil)
```

### Zap Logger Integration

The SDK provides built-in support for the popular [Zap logger](https://github.com/uber-go/zap). You can easily integrate Zap with the SDK:

```go
import (
    "go.uber.org/zap"
    "github.com/conductor-sdk/conductor-go/sdk/log"
)

// Create a Zap logger
zapLogger, _ := zap.NewProduction()
defer zapLogger.Sync()

// Set it as the SDK logger
log.SetLogger(log.NewZap(zapLogger))
```

The Zap adapter automatically handles structured logging and converts the SDK's logging calls to Zap's structured format.

## Detailed Technical Notes -- Unreleased

This section documents internal implementation details for developers reviewing
the metrics harmonization changes. It is not end-user-facing and will be
removed or folded into the relevant sections once the release is published.

### Architecture

The `MetricsCollector` interface (`sdk/metrics/collector.go`) defines the
contract for recording SDK metrics. Two implementations exist:

- `legacyCollector` (`sdk/metrics/legacy_collector.go`) -- emits the original
  metric names and types (Gauges for timing, no `exception` labels).
- `canonicalCollector` (`sdk/metrics/canonical_collector.go`) -- emits the
  harmonized cross-SDK catalog (Histograms in seconds/bytes, `_total` counters,
  bounded `exception` labels from Go type names).

A `noopCollector` is the default singleton before `InitCollector` is called, so
all metric calls are safe at any time.

`metrics.NewCollector()` reads `WORKER_CANONICAL_METRICS` once and returns the
appropriate implementation. `metrics.InitCollector()` is protected by
`sync.Once` for safe concurrent use.

### Capability query methods

The `MetricsCollector` interface exposes `ShouldRecordPayloadSize()` and
`ShouldRecordHTTPRequests()` so call sites can skip expensive prep work:

- `noopCollector` / `legacyCollector`: both return `false`
- `canonicalCollector`: both default to `true`, overridable to `false` via
  `WORKER_METRICS_PAYLOAD_SIZE` and `WORKER_METRICS_HTTP_REQUESTS`

`recordTaskResultPayloadSize` (in `sdk/worker/task_runner.go`) and
`recordWorkflowInputPayloadSize` (in
`sdk/workflow/executor/executor_with_context.go`) check
`ShouldRecordPayloadSize()` before calling `json.Marshal`.

### HTTP request timing

`metricsRoundTripper` (`sdk/client/metrics_roundtripper.go`) wraps the API
client's `http.Transport` to record `http_api_client_request_seconds`. It
checks `ShouldRecordHTTPRequests()` at request time and short-circuits to the
inner transport when the capability is disabled (noop, legacy, or canonical with
the env var set to `false`).

The round-tripper is installed on every `APIClient` HTTP transport in
`api_client.go`, including when legacy or noop collectors are active. When
recording is disabled the overhead is a single `ShouldRecordHTTPRequests()`
check (returns `false`) followed by a direct delegation to the inner
transport -- one additional interface dispatch per HTTP request.

### Context annotation for URI labels

Every API resource method now calls `metrics.WithPathTemplate(ctx, template)`
before building the URL path with `fmt.Sprintf`. This stores the parameterized
path template (e.g. `/workflow/{workflowId}`) in the request context so the
`metricsRoundTripper` can use it as the bounded-cardinality `uri` label.

As a fallback, `executeCall` in `sdk/client/api_client.go` calls
`metrics.WithRawPath(ctx, path)` when no template has been set. This covers
API calls that do not have path parameters (the raw path IS the template) and
any call sites not yet annotated.

Both context enrichments use `context.WithValue`, adding one small allocation
per API call regardless of whether metrics are active.

### Signature changes

- `RecordTaskPollTime`, `RecordTaskExecuteTime`, `RecordTaskUpdateTime`
  package-level functions retain their original 2-argument signatures
  `(taskType string, seconds float64)` and are marked `// Deprecated:`. They
  delegate to the collector with a nil error (always records as SUCCESS in
  canonical mode). The status-aware 3-argument signatures
  `(taskType string, seconds float64, err error)` are available on the
  `MetricsCollector` interface via `metrics.GetCollector()`. Internal SDK code
  uses the 3-argument form through the interface. The legacy collector converts
  execute and update times back to milliseconds internally.
- `IncrementUncaughtException` now takes `recovered interface{}` instead of
  `message string`. Existing callers passing string values are unaffected
  (string satisfies `interface{}`). The canonical collector derives a
  bounded-cardinality `exception` label via `fmt.Sprintf("%T", recovered)`;
  the legacy collector ignores the argument.

### New worker instrumentation call sites

- `IncrementTaskExecutionStarted` -- called in `task_runner.go` when a polled
  task is dispatched to the worker function.
- `IncrementTaskPaused` -- called in `task_runner.go` when `isPaused` is true
  (was defined but never called on `main`).
- `SetActiveWorkers` -- called on worker enter/exit in `task_runner.go`.
- `recordTaskResultPayloadSize` -- called after task execution in
  `task_runner.go`.
- `recordWorkflowInputPayloadSize` -- called before workflow start in
  `executor_with_context.go`.
- `IncrementWorkflowStartError` -- called on workflow start failure in
  `executor_with_context.go` (existed as dead code on `main`).

### Bug fixes

- `metrics.PayloadType.TASK_OUTPUT` had the value `"TASK_INPUT"` (copy-paste
  bug). Fixed to `"TASK_OUTPUT"`.
- `thread_uncaught_exceptions` was registered with zero labels but
  `IncrementUncaughtException` passed one label value, causing Prometheus to
  silently reject every increment. The legacy collector now passes `[]string{}`
  (matching the registration), so the counter increments correctly.

### Harness changes

- `harness/manifests/deployment.yaml` sets `WORKER_CANONICAL_METRICS=true`.
- `harness/main.go` calls `metrics.InitCollector()` before starting the HTTP
  server and logs which collector is active.
