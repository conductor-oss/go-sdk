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

// canonicalCollector emits the harmonized metric names defined in the SDK
// metrics harmonization spec. Counters use _total suffix with exception
// labels, timing uses Histograms in seconds, and sizes use Histograms in
// bytes.
type canonicalCollector struct {
	recordPayloadSize  bool
	recordHTTPRequests bool
}

var canonicalCounterTemplates = map[MetricName]*MetricDetails{
	TASK_POLL_TOTAL: NewMetricDetails(
		TASK_POLL_TOTAL, TASK_POLL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_POLL_ERROR_TOTAL: NewMetricDetails(
		TASK_POLL_ERROR_TOTAL, TASK_POLL_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_EXECUTION_STARTED_TOTAL: NewMetricDetails(
		TASK_EXECUTION_STARTED_TOTAL, TASK_EXECUTION_STARTED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTE_ERROR_TOTAL: NewMetricDetails(
		TASK_EXECUTE_ERROR_TOTAL, TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_UPDATE_ERROR_TOTAL: NewMetricDetails(
		TASK_UPDATE_ERROR_TOTAL, TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_ACK_ERROR_TOTAL: NewMetricDetails(
		TASK_ACK_ERROR_TOTAL, TASK_ACK_ERROR_DOC,
		[]MetricLabel{TASK_TYPE, EXCEPTION},
	),
	TASK_ACK_FAILED_TOTAL: NewMetricDetails(
		TASK_ACK_FAILED_TOTAL, TASK_ACK_FAILED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_EXECUTION_QUEUE_FULL_TOTAL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL_TOTAL, TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	TASK_PAUSED_TOTAL: NewMetricDetails(
		TASK_PAUSED_TOTAL, TASK_PAUSED_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	THREAD_UNCAUGHT_EXCEPTIONS_TOTAL: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTIONS_TOTAL, THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{EXCEPTION},
	),
	EXTERNAL_PAYLOAD_USED_TOTAL: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED_TOTAL, EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{ENTITY_NAME, OPERATION, PAYLOAD_TYPE_CAMEL},
	),
	WORKFLOW_START_ERROR_TOTAL: NewMetricDetails(
		WORKFLOW_START_ERROR_TOTAL, WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{WORKFLOW_TYPE, EXCEPTION},
	),
}

var canonicalDurationHistogramTemplates = map[MetricName]*MetricDetails{
	TASK_POLL_TIME_SECONDS: NewMetricDetails(
		TASK_POLL_TIME_SECONDS, TASK_POLL_TIME_SECONDS_DOC,
		[]MetricLabel{TASK_TYPE, STATUS},
	),
	TASK_EXECUTE_TIME_SECONDS: NewMetricDetails(
		TASK_EXECUTE_TIME_SECONDS, TASK_EXECUTE_TIME_SECONDS_DOC,
		[]MetricLabel{TASK_TYPE, STATUS},
	),
	TASK_UPDATE_TIME_SECONDS: NewMetricDetails(
		TASK_UPDATE_TIME_SECONDS, TASK_UPDATE_TIME_SECONDS_DOC,
		[]MetricLabel{TASK_TYPE, STATUS},
	),
	HTTP_API_CLIENT_REQUEST_SECONDS: NewMetricDetails(
		HTTP_API_CLIENT_REQUEST_SECONDS, HTTP_API_CLIENT_REQUEST_SECONDS_DOC,
		[]MetricLabel{METHOD, URI, STATUS},
	),
}

var canonicalSizeHistogramTemplates = map[MetricName]*MetricDetails{
	TASK_RESULT_SIZE_BYTES: NewMetricDetails(
		TASK_RESULT_SIZE_BYTES, TASK_RESULT_SIZE_BYTES_DOC,
		[]MetricLabel{TASK_TYPE},
	),
	WORKFLOW_INPUT_SIZE_BYTES: NewMetricDetails(
		WORKFLOW_INPUT_SIZE_BYTES, WORKFLOW_INPUT_SIZE_BYTES_DOC,
		[]MetricLabel{WORKFLOW_TYPE, WORKFLOW_VERSION},
	),
}

var canonicalGaugeTemplates = map[MetricName]*MetricDetails{
	ACTIVE_WORKERS: NewMetricDetails(
		ACTIVE_WORKERS, ACTIVE_WORKERS_DOC,
		[]MetricLabel{TASK_TYPE},
	),
}

func (c *canonicalCollector) Register() {
	for name, details := range canonicalCounterTemplates {
		counterByName[name] = newCounter(details)
		prometheus.MustRegister(counterByName[name])
	}
	for name, details := range canonicalDurationHistogramTemplates {
		histogramByName[name] = newHistogram(details, CanonicalDurationBuckets)
		prometheus.MustRegister(histogramByName[name])
	}
	for name, details := range canonicalSizeHistogramTemplates {
		histogramByName[name] = newHistogram(details, CanonicalSizeBuckets)
		prometheus.MustRegister(histogramByName[name])
	}
	for name, details := range canonicalGaugeTemplates {
		gaugeByName[name] = newGauge(details)
		prometheus.MustRegister(gaugeByName[name])
	}
}

func (c *canonicalCollector) CollectorName() string { return "canonical" }

// --- Counters ---

func (c *canonicalCollector) IncrementTaskPoll(taskType string) {
	incrementCounter(TASK_POLL_TOTAL, []string{taskType})
}

func (c *canonicalCollector) IncrementTaskPollError(taskType string, err error) {
	incrementCounter(TASK_POLL_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func (c *canonicalCollector) IncrementTaskExecutionStarted(taskType string) {
	incrementCounter(TASK_EXECUTION_STARTED_TOTAL, []string{taskType})
}

func (c *canonicalCollector) IncrementTaskExecuteError(taskType string, err error) {
	incrementCounter(TASK_EXECUTE_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func (c *canonicalCollector) IncrementTaskUpdateError(taskType string, err error) {
	incrementCounter(TASK_UPDATE_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func (c *canonicalCollector) IncrementTaskAckError(taskType string, err error) {
	incrementCounter(TASK_ACK_ERROR_TOTAL, []string{taskType, ExceptionLabel(err)})
}

func (c *canonicalCollector) IncrementTaskAckFailed(taskType string) {
	incrementCounter(TASK_ACK_FAILED_TOTAL, []string{taskType})
}

func (c *canonicalCollector) IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(TASK_EXECUTION_QUEUE_FULL_TOTAL, []string{taskType})
}

func (c *canonicalCollector) IncrementTaskPaused(taskType string) {
	incrementCounter(TASK_PAUSED_TOTAL, []string{taskType})
}

func (c *canonicalCollector) IncrementUncaughtException(recovered interface{}) {
	incrementCounter(THREAD_UNCAUGHT_EXCEPTIONS_TOTAL, []string{recoveredLabel(recovered)})
}

func (c *canonicalCollector) IncrementExternalPayloadUsed(entityName, operation, payloadType string) {
	incrementCounter(EXTERNAL_PAYLOAD_USED_TOTAL, []string{entityName, operation, payloadType})
}

func (c *canonicalCollector) IncrementWorkflowStartError(workflowType string, err error) {
	incrementCounter(WORKFLOW_START_ERROR_TOTAL, []string{workflowType, ExceptionLabel(err)})
}

// --- Timing (Histograms in seconds) ---

func (c *canonicalCollector) RecordTaskPollTime(taskType string, seconds float64, err error) {
	observeHistogram(TASK_POLL_TIME_SECONDS, []string{taskType, statusLabel(err)}, seconds)
}

func (c *canonicalCollector) RecordTaskExecuteTime(taskType string, seconds float64, err error) {
	observeHistogram(TASK_EXECUTE_TIME_SECONDS, []string{taskType, statusLabel(err)}, seconds)
}

func (c *canonicalCollector) RecordTaskUpdateTime(taskType string, seconds float64, err error) {
	observeHistogram(TASK_UPDATE_TIME_SECONDS, []string{taskType, statusLabel(err)}, seconds)
}

func (c *canonicalCollector) RecordHTTPRequestTime(method, uri, status string, seconds float64) {
	observeHistogram(HTTP_API_CLIENT_REQUEST_SECONDS, []string{method, uri, status}, seconds)
}

// --- Size (Histograms in bytes) ---

func (c *canonicalCollector) RecordTaskResultPayloadSize(taskType string, bytes float64) {
	observeHistogram(TASK_RESULT_SIZE_BYTES, []string{taskType}, bytes)
}

func (c *canonicalCollector) RecordWorkflowInputPayloadSize(workflowType, version string, bytes float64) {
	observeHistogram(WORKFLOW_INPUT_SIZE_BYTES, []string{workflowType, version}, bytes)
}

// --- Gauges ---

func (c *canonicalCollector) SetActiveWorkers(taskType string, count float64) {
	setGauge(ACTIVE_WORKERS, []string{taskType}, count)
}

// --- Capability queries ---

func (c *canonicalCollector) ShouldRecordPayloadSize() bool  { return c.recordPayloadSize }
func (c *canonicalCollector) ShouldRecordHTTPRequests() bool { return c.recordHTTPRequests }
