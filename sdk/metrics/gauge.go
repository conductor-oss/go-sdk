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

func setGauge(metricName MetricName, labelValues []string, value float64) {
	if !collectionEnabled.Load() {
		return
	}

	gauge := getGauge(metricName, labelValues)
	if *gauge != nil {
		(*gauge).Set(value)
	}
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
