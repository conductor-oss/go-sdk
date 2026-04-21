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

func IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(TASK_EXECUTION_QUEUE_FULL, []string{taskType})
	incrementCounter(TASK_EXECUTION_QUEUE_FULL_TOTAL, []string{taskType})
}

// IncrementUncaughtException increments both the legacy unlabeled counter and
// the canonical thread_uncaught_exceptions_total{exception=...} counter. The
// `message` argument, previously silently dropped, is now retained on the
// canonical counter as the exception label (bounded-cardinality: callers
// should pass a short identifier, not a full stack trace).
func IncrementUncaughtException(message string) {
	incrementCounter(THREAD_UNCAUGHT_EXCEPTION, []string{})
	incrementCounter(THREAD_UNCAUGHT_EXCEPTIONS_TOTAL, []string{message})
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
func IncrementTaskAckError(taskType string, err error) {
	incrementCounter(TASK_ACK_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

// IncrementTaskAckFailed records that the server returned a non-success ack
// response for a polled task. Canonical-only.
func IncrementTaskAckFailed(taskType string) {
	incrementCounter(TASK_ACK_FAILED_TOTAL, []string{taskType})
}

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
