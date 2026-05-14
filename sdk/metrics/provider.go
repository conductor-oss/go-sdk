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
	"net/http"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

var collectionEnabled bool = false

var initOnce = &sync.Once{}

// resetInitOnce replaces the sync.Once so InitCollector can fire again.
// Only intended for use in tests.
func resetInitOnce() {
	initOnce = &sync.Once{}
}

// InitCollector selects and registers a MetricsCollector based on env vars.
// Safe to call from multiple goroutines; only the first call takes effect.
func InitCollector() {
	initOnce.Do(func() {
		collector = NewCollector()
		collector.Register()
		collectionEnabled = true
	})
}

// ProvideMetrics initializes the metrics collector (if not already done) and
// starts the HTTP metrics server. This call blocks on ListenAndServe, so it
// should typically be launched in a goroutine.
func ProvideMetrics(metricsSettings *settings.MetricsSettings) {
	defer handlePanicError("provide_metrics")
	if metricsSettings == nil {
		metricsSettings = settings.NewDefaultMetricsSettings()
	}

	InitCollector()

	http.Handle(
		metricsSettings.ApiEndpoint,
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		),
	)
	portString := strconv.Itoa(metricsSettings.Port)
	http.ListenAndServe(":"+portString, nil)
}

func handlePanicError(message string) {
	err := recover()
	if err == nil {
		return
	}
	IncrementUncaughtException(err)
	log.Warn(
		"Uncaught panic",
		"message", message,
		"error", err,
	)
}
