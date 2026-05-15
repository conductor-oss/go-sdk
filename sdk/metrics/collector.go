//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

import (
	"fmt"
	"os"
	"strings"

	"github.com/conductor-sdk/conductor-go/sdk/log"
)

// MetricsCollector defines the contract for recording SDK metrics. Two
// implementations exist: legacyCollector (original metric names and types) and
// canonicalCollector (harmonized names, histograms, exception labels). Only
// one is active at a time, selected by environment variables at startup.
type MetricsCollector interface {
	// Register creates and registers all Prometheus metrics for this collector.
	Register()

	// CollectorName returns a human-readable name for logging.
	CollectorName() string

	// --- Counters ---

	IncrementTaskPoll(taskType string)
	IncrementTaskPollError(taskType string, err error)
	IncrementTaskExecutionStarted(taskType string)
	IncrementTaskExecuteError(taskType string, err error)
	IncrementTaskUpdateError(taskType string, err error)
	IncrementTaskAckError(taskType string, err error)
	IncrementTaskAckFailed(taskType string)
	IncrementTaskExecutionQueueFull(taskType string)
	IncrementTaskPaused(taskType string)
	IncrementUncaughtException(recovered interface{})
	IncrementExternalPayloadUsed(entityName, operation, payloadType string)
	IncrementWorkflowStartError(workflowType string, err error)

	// --- Timing (callers always pass seconds) ---

	RecordTaskPollTime(taskType string, seconds float64, err error)
	RecordTaskExecuteTime(taskType string, seconds float64, err error)
	RecordTaskUpdateTime(taskType string, seconds float64, err error)
	RecordHTTPRequestTime(method, uri, status string, seconds float64)

	// --- Size ---

	RecordTaskResultPayloadSize(taskType string, bytes float64)
	RecordWorkflowInputPayloadSize(workflowType, version string, bytes float64)

	// --- Gauges ---

	SetActiveWorkers(taskType string, count float64)

	// --- Capability queries ---
	// Call sites check these before doing expensive prep work (e.g.
	// json.Marshal for payload size, timing in the HTTP round-tripper).

	ShouldRecordPayloadSize() bool
	ShouldRecordHTTPRequests() bool
}

// collector is the package-level singleton. Before ProvideMetrics is called it
// is a noopCollector so all metric calls are safe to make at any time.
var collector MetricsCollector = &noopCollector{}

// GetCollector returns the active MetricsCollector. Useful for callers that
// want to hold a typed reference rather than going through package functions.
func GetCollector() MetricsCollector {
	return collector
}

// envBoolTruthy returns true when the named environment variable is set to
// one of the common truthy values ("true", "1", "yes"), case-insensitive.
func envBoolTruthy(name string) bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv(name)))
	return v == "true" || v == "1" || v == "yes"
}

// envBoolDefault returns defaultVal when the environment variable is unset or
// empty. Otherwise it applies the same truthy check as envBoolTruthy.
func envBoolDefault(name string, defaultVal bool) bool {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return defaultVal
	}
	v := strings.ToLower(raw)
	return v == "true" || v == "1" || v == "yes"
}

// NewCollector reads the WORKER_CANONICAL_METRICS / WORKER_LEGACY_METRICS
// environment variables and returns the appropriate implementation.
//
// Selection logic:
//   - WORKER_CANONICAL_METRICS=true/1/yes  -> canonicalCollector
//   - Anything else (default)              -> legacyCollector
func NewCollector() MetricsCollector {
	if envBoolTruthy("WORKER_CANONICAL_METRICS") {
		log.Info("Metrics implementation selected: canonical")
		return &canonicalCollector{
			recordPayloadSize:  envBoolDefault("WORKER_METRICS_PAYLOAD_SIZE", true),
			recordHTTPRequests: envBoolDefault("WORKER_METRICS_HTTP_REQUESTS", true),
		}
	}
	log.Info("Metrics implementation selected: legacy")
	return &legacyCollector{}
}

// ExceptionLabel returns a bounded-cardinality label value for an error,
// using the Go type name (e.g. "*url.Error") rather than the message.
func ExceptionLabel(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%T", err)
}

// recoveredLabel returns a bounded-cardinality exception label for the value
// returned by recover(). Typed errors use ExceptionLabel; anything else falls
// back to fmt.Sprintf("%T", ...).
func recoveredLabel(recovered interface{}) string {
	if recovered == nil {
		return ""
	}
	if err, ok := recovered.(error); ok {
		return ExceptionLabel(err)
	}
	return fmt.Sprintf("%T", recovered)
}

// statusLabel returns "SUCCESS" or "FAILURE" based on whether err is nil.
func statusLabel(err error) string {
	if err == nil {
		return "SUCCESS"
	}
	return "FAILURE"
}

// ---------------------------------------------------------------------------
// Package-level convenience functions — delegate to the singleton collector.
// These preserve the existing call-site API so no import changes are needed.
// ---------------------------------------------------------------------------

func IncrementTaskPoll(taskType string) { collector.IncrementTaskPoll(taskType) }
func IncrementTaskPollError(taskType string, err error) {
	collector.IncrementTaskPollError(taskType, err)
}
func IncrementTaskExecutionStarted(taskType string) {
	collector.IncrementTaskExecutionStarted(taskType)
}
func IncrementTaskExecuteError(taskType string, err error) {
	collector.IncrementTaskExecuteError(taskType, err)
}
func IncrementTaskUpdateError(taskType string, err error) {
	collector.IncrementTaskUpdateError(taskType, err)
}
func IncrementTaskAckError(taskType string, err error) {
	collector.IncrementTaskAckError(taskType, err)
}
func IncrementTaskAckFailed(taskType string) { collector.IncrementTaskAckFailed(taskType) }
func IncrementTaskExecutionQueueFull(taskType string) {
	collector.IncrementTaskExecutionQueueFull(taskType)
}
func IncrementTaskPaused(taskType string) { collector.IncrementTaskPaused(taskType) }
func IncrementUncaughtException(recovered interface{}) {
	collector.IncrementUncaughtException(recovered)
}
func IncrementExternalPayloadUsed(entityName, operation, payloadType string) {
	collector.IncrementExternalPayloadUsed(entityName, operation, payloadType)
}
func IncrementWorkflowStartError(workflowType string, err error) {
	collector.IncrementWorkflowStartError(workflowType, err)
}

// Deprecated: does not propagate success/failure status to the collector.
// Use GetCollector().RecordTaskPollTime(taskType, seconds, err) for
// status-aware recording.
func RecordTaskPollTime(taskType string, seconds float64) {
	collector.RecordTaskPollTime(taskType, seconds, nil)
}

// Deprecated: does not propagate success/failure status to the collector.
// Use GetCollector().RecordTaskExecuteTime(taskType, seconds, err) for
// status-aware recording.
func RecordTaskExecuteTime(taskType string, seconds float64) {
	collector.RecordTaskExecuteTime(taskType, seconds, nil)
}

// Deprecated: does not propagate success/failure status to the collector.
// Use GetCollector().RecordTaskUpdateTime(taskType, seconds, err) for
// status-aware recording.
func RecordTaskUpdateTime(taskType string, seconds float64) {
	collector.RecordTaskUpdateTime(taskType, seconds, nil)
}
func RecordHTTPRequestTime(method, uri, status string, seconds float64) {
	collector.RecordHTTPRequestTime(method, uri, status, seconds)
}
func RecordTaskResultPayloadSize(taskType string, bytes float64) {
	collector.RecordTaskResultPayloadSize(taskType, bytes)
}
func RecordWorkflowInputPayloadSize(workflowType, version string, bytes float64) {
	collector.RecordWorkflowInputPayloadSize(workflowType, version, bytes)
}
func SetActiveWorkers(taskType string, count float64) { collector.SetActiveWorkers(taskType, count) }
func ShouldRecordPayloadSize() bool                   { return collector.ShouldRecordPayloadSize() }
func ShouldRecordHTTPRequests() bool                  { return collector.ShouldRecordHTTPRequests() }

// ---------------------------------------------------------------------------
// noopCollector — all methods are no-ops. Used as the default before
// ProvideMetrics is called.
// ---------------------------------------------------------------------------

type noopCollector struct{}

func (n *noopCollector) Register()                                              {}
func (n *noopCollector) CollectorName() string                                  { return "noop" }
func (n *noopCollector) IncrementTaskPoll(string)                               {}
func (n *noopCollector) IncrementTaskPollError(string, error)                   {}
func (n *noopCollector) IncrementTaskExecutionStarted(string)                   {}
func (n *noopCollector) IncrementTaskExecuteError(string, error)                {}
func (n *noopCollector) IncrementTaskUpdateError(string, error)                 {}
func (n *noopCollector) IncrementTaskAckError(string, error)                    {}
func (n *noopCollector) IncrementTaskAckFailed(string)                          {}
func (n *noopCollector) IncrementTaskExecutionQueueFull(string)                 {}
func (n *noopCollector) IncrementTaskPaused(string)                             {}
func (n *noopCollector) IncrementUncaughtException(interface{})                 {}
func (n *noopCollector) IncrementExternalPayloadUsed(string, string, string)    {}
func (n *noopCollector) IncrementWorkflowStartError(string, error)              {}
func (n *noopCollector) RecordTaskPollTime(string, float64, error)              {}
func (n *noopCollector) RecordTaskExecuteTime(string, float64, error)           {}
func (n *noopCollector) RecordTaskUpdateTime(string, float64, error)            {}
func (n *noopCollector) RecordHTTPRequestTime(string, string, string, float64)  {}
func (n *noopCollector) RecordTaskResultPayloadSize(string, float64)            {}
func (n *noopCollector) RecordWorkflowInputPayloadSize(string, string, float64) {}
func (n *noopCollector) SetActiveWorkers(string, float64)                       {}
func (n *noopCollector) ShouldRecordPayloadSize() bool                          { return false }
func (n *noopCollector) ShouldRecordHTTPRequests() bool                         { return false }
