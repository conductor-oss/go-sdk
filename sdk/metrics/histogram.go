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

// CanonicalDurationBuckets is the shared bucket set for all timing histograms.
// Matches the canonical catalog in the SDK metrics harmonization spec.
var CanonicalDurationBuckets = []float64{
	0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10,
}

// CanonicalSizeBuckets is the shared bucket set for all size histograms.
var CanonicalSizeBuckets = []float64{
	100, 1000, 10_000, 100_000, 1_000_000, 10_000_000,
}

var histogramByName = map[MetricName]*prometheus.HistogramVec{}

func observeHistogram(metricName MetricName, labelValues []string, value float64) {
	if !collectionEnabled.Load() {
		return
	}
	hist := getHistogram(metricName, labelValues)
	if hist == nil {
		return
	}
	hist.Observe(value)
}

func newHistogram(metricDetails *MetricDetails, buckets []float64) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricDetails.Name,
			Help:    metricDetails.Description,
			Buckets: buckets,
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
