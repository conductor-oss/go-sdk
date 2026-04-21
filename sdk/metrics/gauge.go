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
	"github.com/prometheus/client_golang/prometheus"
)

var gaugeByName = map[MetricName]*prometheus.GaugeVec{}

var gaugeTemplates = map[MetricName]*MetricDetails{
	WORKFLOW_INPUT_SIZE: NewMetricDetails(
		WORKFLOW_INPUT_SIZE,
		WORKFLOW_INPUT_SIZE_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
			WORKFLOW_VERSION,
		},
	),
	WORKFLOW_INPUT_SIZE_BYTES: NewMetricDetails(
		WORKFLOW_INPUT_SIZE_BYTES,
		WORKFLOW_INPUT_SIZE_BYTES_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
			WORKFLOW_VERSION,
		},
	),
	TASK_RESULT_SIZE: NewMetricDetails(
		TASK_RESULT_SIZE,
		TASK_RESULT_SIZE_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_RESULT_SIZE_BYTES: NewMetricDetails(
		TASK_RESULT_SIZE_BYTES,
		TASK_RESULT_SIZE_BYTES_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_POLL_TIME: NewMetricDetails(
		TASK_POLL_TIME,
		TASK_POLL_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_TIME: NewMetricDetails(
		TASK_EXECUTE_TIME,
		TASK_EXECUTE_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_UPDATE_TIME: NewMetricDetails(
		TASK_UPDATE_TIME,
		TASK_UPDATE_TIME_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
}

func RecordWorkflowInputPayloadSize(workflowType string, version string, payloadSize float64) {
	setGauge(WORKFLOW_INPUT_SIZE, []string{workflowType, version}, payloadSize)
	setGauge(WORKFLOW_INPUT_SIZE_BYTES, []string{workflowType, version}, payloadSize)
}

func RecordTaskResultPayloadSize(taskType string, payloadSize float64) {
	setGauge(TASK_RESULT_SIZE, []string{taskType}, payloadSize)
	setGauge(TASK_RESULT_SIZE_BYTES, []string{taskType}, payloadSize)
}

// RecordTaskPollTime records a task-poll duration (in seconds) on the legacy
// task_poll_time Gauge. For the canonical task_poll_time_seconds Histogram,
// use ObserveTaskPollTimeSeconds — it requires a status label and should be
// called after the poll completes so SUCCESS/FAILURE can be distinguished.
func RecordTaskPollTime(taskType string, timeSpent float64) {
	setGauge(TASK_POLL_TIME, []string{taskType}, timeSpent)
}

// RecordTaskUpdateTime records a task-update duration.
// NOTE: for backward compatibility the legacy task_update_time Gauge is
// updated using the caller-supplied value (which the existing worker call
// site passes in milliseconds). The canonical task_update_time_seconds
// Histogram is NOT observed here: callers that have the correct seconds
// value should call ObserveTaskUpdateTimeSeconds directly.
func RecordTaskUpdateTime(taskType string, timeSpent float64) {
	setGauge(TASK_UPDATE_TIME, []string{taskType}, timeSpent)
}

// RecordTaskExecuteTime records a task-execution duration.
// NOTE: for backward compatibility the legacy task_execute_time Gauge is
// updated using the caller-supplied value (which the existing worker call
// site passes in milliseconds). The canonical task_execute_time_seconds
// Histogram is NOT observed here: callers that have the correct seconds
// value should call ObserveTaskExecuteTimeSeconds directly.
func RecordTaskExecuteTime(taskType string, timeSpent float64) {
	setGauge(TASK_EXECUTE_TIME, []string{taskType}, timeSpent)
}

func newGauge(metricDetails *MetricDetails) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func setGauge(metricName MetricName, labelValues []string, value float64) {
	// We skip setting gauge if Metrics collection is not enabled
	if !collectionEnabled {
		return
	}

	gauge := getGauge(metricName, labelValues)
	if *gauge != nil {
		(*gauge).Set(value)
	}
}

func getGauge(metricName MetricName, labelValues []string) *prometheus.Gauge {
	gaugeVec, ok := gaugeByName[metricName]
	if !ok {
		return nil
	}
	gauge, _ := gaugeVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &gauge
}
