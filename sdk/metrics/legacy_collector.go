//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

import "github.com/prometheus/client_golang/prometheus"

// legacyCollector emits the original metric names that existed on the main
// branch before the metrics harmonization effort. Timing metrics are Gauges
// (not Histograms) and counters lack an exception label.
type legacyCollector struct{}

var legacyCounterTemplates = map[MetricName]*MetricDetails{
	TASK_POLL: NewMetricDetails(
		TASK_POLL, TASK_POLL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTION_QUEUE_FULL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL, TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	THREAD_UNCAUGHT_EXCEPTION: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTION, THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{},
	),
	TASK_POLL_ERROR: NewMetricDetails(
		TASK_POLL_ERROR, TASK_POLL_ERROR_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_PAUSED: NewMetricDetails(
		TASK_PAUSED, TASK_PAUSED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTE_ERROR: NewMetricDetails(
		TASK_EXECUTE_ERROR, TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_UPDATE_ERROR: NewMetricDetails(
		TASK_UPDATE_ERROR, TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	EXTERNAL_PAYLOAD_USED: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED, EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{ENTITY_NAME, OPERATION, PAYLOAD_TYPE},
	),
	WORKFLOW_START_ERROR: NewMetricDetails(
		WORKFLOW_START_ERROR, WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{WORKFLOW_TYPE},
	),
}

var legacyGaugeTemplates = map[MetricName]*MetricDetails{
	WORKFLOW_INPUT_SIZE: NewMetricDetails(
		WORKFLOW_INPUT_SIZE, WORKFLOW_INPUT_SIZE_DOC,
		[]MetricLabel{WORKFLOW_TYPE, WORKFLOW_VERSION},
	),
	TASK_RESULT_SIZE: NewMetricDetails(
		TASK_RESULT_SIZE, TASK_RESULT_SIZE_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_POLL_TIME: NewMetricDetails(
		TASK_POLL_TIME, TASK_POLL_TIME_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTE_TIME: NewMetricDetails(
		TASK_EXECUTE_TIME, TASK_EXECUTE_TIME_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_UPDATE_TIME: NewMetricDetails(
		TASK_UPDATE_TIME, TASK_UPDATE_TIME_DOC,
		[]MetricLabel{TASK_TYPE},
	),
}

func (l *legacyCollector) Register() {
	for name, details := range legacyCounterTemplates {
		counterByName[name] = newCounter(details)
		prometheus.MustRegister(counterByName[name])
	}
	for name, details := range legacyGaugeTemplates {
		gaugeByName[name] = newGauge(details)
		prometheus.MustRegister(gaugeByName[name])
	}
}

func (l *legacyCollector) CollectorName() string { return "legacy" }

// --- Counters (real implementations for legacy metrics) ---

func (l *legacyCollector) IncrementTaskPoll(taskType string) {
	incrementCounter(TASK_POLL, []string{taskType})
}

func (l *legacyCollector) IncrementTaskPollError(taskType string, _ error) {
	incrementCounter(TASK_POLL_ERROR, []string{taskType})
}

func (l *legacyCollector) IncrementTaskExecuteError(taskType string, _ error) {
	incrementCounter(TASK_EXECUTE_ERROR, []string{taskType})
}

func (l *legacyCollector) IncrementTaskUpdateError(taskType string, _ error) {
	incrementCounter(TASK_UPDATE_ERROR, []string{taskType})
}

func (l *legacyCollector) IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(TASK_EXECUTION_QUEUE_FULL, []string{taskType})
}

func (l *legacyCollector) IncrementTaskPaused(taskType string) {
	incrementCounter(TASK_PAUSED, []string{taskType})
}

func (l *legacyCollector) IncrementUncaughtException(_ interface{}) {
	incrementCounter(THREAD_UNCAUGHT_EXCEPTION, []string{})
}

func (l *legacyCollector) IncrementExternalPayloadUsed(entityName, operation, payloadType string) {
	incrementCounter(EXTERNAL_PAYLOAD_USED, []string{entityName, operation, payloadType})
}

func (l *legacyCollector) IncrementWorkflowStartError(workflowType string, _ error) {
	incrementCounter(WORKFLOW_START_ERROR, []string{workflowType})
}

// --- Timing (Gauges; execute and update preserve the legacy millisecond unit) ---

func (l *legacyCollector) RecordTaskPollTime(taskType string, seconds float64, _ error) {
	setGauge(TASK_POLL_TIME, []string{taskType}, seconds)
}

func (l *legacyCollector) RecordTaskExecuteTime(taskType string, seconds float64, _ error) {
	setGauge(TASK_EXECUTE_TIME, []string{taskType}, seconds*1000)
}

func (l *legacyCollector) RecordTaskUpdateTime(taskType string, seconds float64, _ error) {
	setGauge(TASK_UPDATE_TIME, []string{taskType}, seconds*1000)
}

// --- Size (Gauges) ---

func (l *legacyCollector) RecordTaskResultPayloadSize(taskType string, bytes float64) {
	setGauge(TASK_RESULT_SIZE, []string{taskType}, bytes)
}

func (l *legacyCollector) RecordWorkflowInputPayloadSize(workflowType, version string, bytes float64) {
	setGauge(WORKFLOW_INPUT_SIZE, []string{workflowType, version}, bytes)
}

// --- Noop stubs for canonical-only methods ---

func (l *legacyCollector) IncrementTaskExecutionStarted(string)                  {}
func (l *legacyCollector) IncrementTaskAckError(string, error)                   {}
func (l *legacyCollector) IncrementTaskAckFailed(string)                         {}
func (l *legacyCollector) RecordHTTPRequestTime(string, string, string, float64) {}
func (l *legacyCollector) SetActiveWorkers(string, float64)                      {}
func (l *legacyCollector) ShouldRecordPayloadSize() bool                         { return false }
func (l *legacyCollector) ShouldRecordHTTPRequests() bool                        { return false }
