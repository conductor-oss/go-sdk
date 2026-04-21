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

// CanonicalDurationBuckets is the shared bucket set for all timing histograms
// emitted by the SDK. It matches the set defined in
// longrunning-wfstest/sdk-metrics-harmonization.md and is kept in sync with
// the Python / Java / Ruby / Rust SDKs so cross-SDK dashboards line up.
var CanonicalDurationBuckets = []float64{
	0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
}

var histogramByName = map[MetricName]*prometheus.HistogramVec{}

var histogramTemplates = map[MetricName]*MetricDetails{
	TASK_POLL_TIME_SECONDS: NewMetricDetails(
		TASK_POLL_TIME_SECONDS,
		TASK_POLL_TIME_SECONDS_DOC,
		[]MetricLabel{
			TASK_TYPE,
			STATUS,
		},
	),
	TASK_EXECUTE_TIME_SECONDS: NewMetricDetails(
		TASK_EXECUTE_TIME_SECONDS,
		TASK_EXECUTE_TIME_SECONDS_DOC,
		[]MetricLabel{
			TASK_TYPE,
			STATUS,
		},
	),
	TASK_UPDATE_TIME_SECONDS: NewMetricDetails(
		TASK_UPDATE_TIME_SECONDS,
		TASK_UPDATE_TIME_SECONDS_DOC,
		[]MetricLabel{
			TASK_TYPE,
			STATUS,
		},
	),
	HTTP_API_CLIENT_REQUEST_SECONDS: NewMetricDetails(
		HTTP_API_CLIENT_REQUEST_SECONDS,
		HTTP_API_CLIENT_REQUEST_SECONDS_DOC,
		[]MetricLabel{
			METHOD,
			URI,
			STATUS,
		},
	),
}

// ObserveTaskPollTimeSeconds records a task-poll duration (in seconds) into
// the canonical task_poll_time_seconds histogram.
func ObserveTaskPollTimeSeconds(taskType string, status string, seconds float64) {
	observeHistogram(
		TASK_POLL_TIME_SECONDS,
		[]string{taskType, status},
		seconds,
	)
}

// ObserveTaskExecuteTimeSeconds records a task-execution duration (in seconds)
// into the canonical task_execute_time_seconds histogram.
func ObserveTaskExecuteTimeSeconds(taskType string, status string, seconds float64) {
	observeHistogram(
		TASK_EXECUTE_TIME_SECONDS,
		[]string{taskType, status},
		seconds,
	)
}

// ObserveTaskUpdateTimeSeconds records a task-update (result-report) duration
// (in seconds) into the canonical task_update_time_seconds histogram.
func ObserveTaskUpdateTimeSeconds(taskType string, status string, seconds float64) {
	observeHistogram(
		TASK_UPDATE_TIME_SECONDS,
		[]string{taskType, status},
		seconds,
	)
}

// ObserveHTTPAPIClientRequestSeconds records an HTTP API client request
// duration (in seconds) into the canonical http_api_client_request_seconds
// histogram. The uri label should be a path template (e.g.
// "/tasks/poll/batch/{taskType}") rather than an interpolated path, to keep
// label cardinality bounded.
func ObserveHTTPAPIClientRequestSeconds(method, uri, status string, seconds float64) {
	observeHistogram(
		HTTP_API_CLIENT_REQUEST_SECONDS,
		[]string{method, uri, status},
		seconds,
	)
}

func observeHistogram(metricName MetricName, labelValues []string, value float64) {
	if !collectionEnabled {
		return
	}
	hist := getHistogram(metricName, labelValues)
	if hist == nil {
		return
	}
	hist.Observe(value)
}

func newHistogram(metricDetails *MetricDetails) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricDetails.Name,
			Help:    metricDetails.Description,
			Buckets: CanonicalDurationBuckets,
		},
		metricDetails.Labels,
	)
}

func getHistogram(metricName MetricName, labelValues []string) prometheus.Observer {
	vec, ok := histogramByName[metricName]
	if !ok {
		return nil
	}
	obs, err := vec.GetMetricWithLabelValues(labelValues...)
	if err != nil {
		return nil
	}
	return obs
}
