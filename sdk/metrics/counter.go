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

	"github.com/prometheus/client_golang/prometheus"
)

var counterByName = map[MetricName]*prometheus.CounterVec{}

var counterTemplates = map[MetricName]*MetricDetails{
	TASK_POLL: NewMetricDetails(
		TASK_POLL,
		TASK_POLL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTION_QUEUE_FULL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL,
		TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	THREAD_UNCAUGHT_EXCEPTION: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTION,
		THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{},
	),
	TASK_POLL_ERROR: NewMetricDetails(
		TASK_POLL_ERROR,
		TASK_POLL_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_PAUSED: NewMetricDetails(
		TASK_PAUSED,
		TASK_PAUSED_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_ERROR: NewMetricDetails(
		TASK_EXECUTE_ERROR,
		TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),

	TASK_UPDATE_ERROR: NewMetricDetails(
		TASK_UPDATE_ERROR,
		TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	EXTERNAL_PAYLOAD_USED: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED,
		EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{
			ENTITY_NAME,
			OPERATION,
			PAYLOAD_TYPE,
		},
	),
	WORKFLOW_START_ERROR: NewMetricDetails(
		WORKFLOW_START_ERROR,
		WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
		},
	),

	// Canonical _total counters. Emitted alongside the legacy counters above.
	TASK_POLL_TOTAL: NewMetricDetails(
		TASK_POLL_TOTAL,
		TASK_POLL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_POLL_ERROR_TOTAL: NewMetricDetails(
		TASK_POLL_ERROR_TOTAL,
		TASK_POLL_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_EXECUTION_STARTED_TOTAL: NewMetricDetails(
		TASK_EXECUTION_STARTED_TOTAL,
		TASK_EXECUTION_STARTED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTE_ERROR_TOTAL: NewMetricDetails(
		TASK_EXECUTE_ERROR_TOTAL,
		TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_UPDATE_ERROR_TOTAL: NewMetricDetails(
		TASK_UPDATE_ERROR_TOTAL,
		TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_ACK_ERROR_TOTAL: NewMetricDetails(
		TASK_ACK_ERROR_TOTAL,
		TASK_ACK_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_ACK_FAILED_TOTAL: NewMetricDetails(
		TASK_ACK_FAILED_TOTAL,
		TASK_ACK_FAILED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTION_QUEUE_FULL_TOTAL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL_TOTAL,
		TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_PAUSED_TOTAL: NewMetricDetails(
		TASK_PAUSED_TOTAL,
		TASK_PAUSED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	THREAD_UNCAUGHT_EXCEPTIONS_TOTAL: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTIONS_TOTAL,
		THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{EXCEPTION},
	),
	EXTERNAL_PAYLOAD_USED_TOTAL: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED_TOTAL,
		EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{ENTITY_NAME, OPERATION, PAYLOAD_TYPE},
	),
	WORKFLOW_START_ERROR_TOTAL: NewMetricDetails(
		WORKFLOW_START_ERROR_TOTAL,
		WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{WORKFLOW_TYPE, EXCEPTION},
	),
}

// ExceptionLabel returns a bounded-cardinality label value for an error. We
// emit the dynamic Go type name (e.g. "*rest.APIError") rather than the
// error message, so the "exception" label stays usable as a dashboard facet.
func ExceptionLabel(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%T", err)
}

func IncrementTaskPoll(taskType string) {
	incrementCounter(TASK_POLL, []string{taskType})
	incrementCounter(TASK_POLL_TOTAL, []string{taskType})
}

// IncrementTaskExecutionStarted records that a polled task has been dispatched
// to the user worker function. Emitted only on the canonical counter.
func IncrementTaskExecutionStarted(taskType string) {
	incrementCounter(TASK_EXECUTION_STARTED_TOTAL, []string{taskType})
}

// IncrementTaskExecutionQueueFull is provided for parity with the other SDKs
// but is not wired into any internal code path: the Go runtime dispatches
// each polled task onto its own goroutine and does not maintain a bounded
// executor queue that can overflow. Intentional N/A per the metrics
// harmonization spec (§3.1). The helper is kept so user code that spawns its
// own bounded worker pools can still emit the canonical counter.
func IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(TASK_EXECUTION_QUEUE_FULL, []string{taskType})
	incrementCounter(TASK_EXECUTION_QUEUE_FULL_TOTAL, []string{taskType})
}

// IncrementUncaughtException increments both the legacy unlabeled counter and
// the canonical thread_uncaught_exceptions_total{exception=...} counter.
// `recovered` is the value returned by recover(); we emit its Go type name
// as the exception label to keep cardinality bounded (mirroring
// ExceptionLabel for typed errors and %T for everything else).
func IncrementUncaughtException(recovered interface{}) {
	incrementCounter(THREAD_UNCAUGHT_EXCEPTION, []string{})
	incrementCounter(THREAD_UNCAUGHT_EXCEPTIONS_TOTAL, []string{recoveredLabel(recovered)})
}

// recoveredLabel returns a bounded-cardinality exception label for the value
// returned by recover(). For typed Go errors it reuses ExceptionLabel; for
// anything else (string panics, etc.) it falls back to fmt.Sprintf("%T", ...).
func recoveredLabel(recovered interface{}) string {
	if recovered == nil {
		return ""
	}
	if err, ok := recovered.(error); ok {
		return ExceptionLabel(err)
	}
	return fmt.Sprintf("%T", recovered)
}

func IncrementTaskPollError(taskType string, err error) {
	incrementCounter(TASK_POLL_ERROR, []string{taskType})
	incrementCounter(TASK_POLL_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func IncrementTaskPaused(taskType string) {
	incrementCounter(TASK_PAUSED, []string{taskType})
	incrementCounter(TASK_PAUSED_TOTAL, []string{taskType})
}

func IncrementTaskExecuteError(taskType string, err error) {
	incrementCounter(TASK_EXECUTE_ERROR, []string{taskType})
	incrementCounter(TASK_EXECUTE_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func IncrementTaskUpdateError(taskType string, err error) {
	incrementCounter(TASK_UPDATE_ERROR, []string{taskType})
	incrementCounter(TASK_UPDATE_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

// IncrementTaskAckError records that an exception was raised while
// acknowledging a polled task. Canonical-only.
//
// Intentional N/A for the internal Go runner per the metrics harmonization
// spec (§3.1): the go-sdk worker loop does not call a separate ack endpoint
// — the batch-poll response itself is the ack. Exposed for user code that
// may implement its own ack semantics.
func IncrementTaskAckError(taskType string, err error) {
	incrementCounter(TASK_ACK_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

// IncrementTaskAckFailed records that the server returned a non-success ack
// response for a polled task. Canonical-only.
//
// Intentional N/A for the internal Go runner per the metrics harmonization
// spec (§3.1); see IncrementTaskAckError above.
func IncrementTaskAckFailed(taskType string) {
	incrementCounter(TASK_ACK_FAILED_TOTAL, []string{taskType})
}

// IncrementExternalPayloadUsed is intentionally N/A for the go-sdk per the
// metrics harmonization spec (§3.1): the Go client does not currently
// integrate with the external-payload-storage branch of the conductor API.
// Kept for parity with python-sdk / java-sdk so user code can still emit
// the canonical counter if it implements its own external-payload plumbing.
func IncrementExternalPayloadUsed(entityName string, operation string, payloadType string) {
	incrementCounter(EXTERNAL_PAYLOAD_USED, []string{entityName, operation, payloadType})
	incrementCounter(EXTERNAL_PAYLOAD_USED_TOTAL, []string{entityName, operation, payloadType})
}

func IncrementWorkflowStartError(workflowType string, err error) {
	incrementCounter(WORKFLOW_START_ERROR, []string{workflowType})
	incrementCounter(WORKFLOW_START_ERROR_TOTAL, []string{workflowType, ExceptionLabel(err)})
}

func incrementCounter(metricName MetricName, labelValues []string) {
	// We skip incrementing if metrics collection is not yet enabled
	if !collectionEnabled {
		return
	}

	counter := getCounter(metricName, labelValues)
	if *counter != nil {
		(*counter).Inc()
	}
}

func newCounter(metricDetails *MetricDetails) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func getCounter(metricName MetricName, labelValues []string) *prometheus.Counter {
	counterVec, ok := counterByName[metricName]
	if !ok {
		return nil
	}
	counter, _ := counterVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &counter
}
