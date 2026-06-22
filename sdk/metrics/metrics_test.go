package metrics

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setCollector publishes c as the active collector. Test-only helper mirroring
// the atomic publication done by InitCollector.
func setCollector(c MetricsCollector) {
	collectorPtr.Store(&c)
}

// resetState isolates each test by clearing package-level maps, disabling
// collection, restoring the noop collector, and installing a fresh Prometheus
// registry so MustRegister never collides with a previous test.
func resetState(t *testing.T) {
	t.Helper()

	origCollector := GetCollector()
	origEnabled := collectionEnabled.Load()
	origRegisterer := prometheus.DefaultRegisterer
	origGatherer := prometheus.DefaultGatherer

	counterByName = map[MetricName]*prometheus.CounterVec{}
	histogramByName = map[MetricName]*prometheus.HistogramVec{}
	gaugeByName = map[MetricName]*prometheus.GaugeVec{}
	collectionEnabled.Store(false)
	setCollector(&noopCollector{})
	resetInitOnce()

	reg := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = reg
	prometheus.DefaultGatherer = reg

	t.Cleanup(func() {
		setCollector(origCollector)
		collectionEnabled.Store(origEnabled)
		resetInitOnce()
		counterByName = map[MetricName]*prometheus.CounterVec{}
		histogramByName = map[MetricName]*prometheus.HistogramVec{}
		gaugeByName = map[MetricName]*prometheus.GaugeVec{}
		prometheus.DefaultRegisterer = origRegisterer
		prometheus.DefaultGatherer = origGatherer
	})
}

// ---------------------------------------------------------------------------
// Helper function tests
// ---------------------------------------------------------------------------

func TestEnvBoolTruthy(t *testing.T) {
	const envKey = "TEST_ENVBOOLTRUTHY"
	t.Cleanup(func() { os.Unsetenv(envKey) })

	for _, tc := range []struct {
		val  string
		want bool
	}{
		{"true", true},
		{"TRUE", true},
		{"True", true},
		{"1", true},
		{"yes", true},
		{"YES", true},
		{" yes ", true},
		{"false", false},
		{"0", false},
		{"no", false},
		{"", false},
		{"maybe", false},
	} {
		t.Run(fmt.Sprintf("val=%q", tc.val), func(t *testing.T) {
			os.Setenv(envKey, tc.val)
			assert.Equal(t, tc.want, envBoolTruthy(envKey))
		})
	}

	t.Run("unset", func(t *testing.T) {
		os.Unsetenv(envKey)
		assert.False(t, envBoolTruthy(envKey))
	})
}

func TestExceptionLabel(t *testing.T) {
	assert.Equal(t, "", ExceptionLabel(nil))
	assert.Equal(t, "*errors.errorString", ExceptionLabel(errors.New("boom")))
}

func TestRecoveredLabel(t *testing.T) {
	assert.Equal(t, "", recoveredLabel(nil))
	assert.Equal(t, "*errors.errorString", recoveredLabel(errors.New("boom")))
	assert.Equal(t, "string", recoveredLabel("not an error"))
	assert.Equal(t, "int", recoveredLabel(42))
}

func TestStatusLabel(t *testing.T) {
	assert.Equal(t, "SUCCESS", statusLabel(nil))
	assert.Equal(t, "FAILURE", statusLabel(errors.New("fail")))
}

func TestEnvBoolDefault(t *testing.T) {
	const envKey = "TEST_ENVBOOLDEFAULT"
	t.Cleanup(func() { os.Unsetenv(envKey) })

	for _, tc := range []struct {
		val        string
		set        bool
		defaultVal bool
		want       bool
	}{
		{"true", true, false, true},
		{"1", true, false, true},
		{"yes", true, false, true},
		{"YES", true, false, true},
		{" true ", true, false, true},
		{"false", true, true, false},
		{"0", true, true, false},
		{"no", true, true, false},
		{"maybe", true, true, false},
		{"", true, true, true},
		{"", false, true, true},
		{"", false, false, false},
	} {
		label := fmt.Sprintf("val=%q/set=%v/default=%v", tc.val, tc.set, tc.defaultVal)
		t.Run(label, func(t *testing.T) {
			if tc.set {
				os.Setenv(envKey, tc.val)
			} else {
				os.Unsetenv(envKey)
			}
			assert.Equal(t, tc.want, envBoolDefault(envKey, tc.defaultVal))
		})
	}
}

// ---------------------------------------------------------------------------
// NewCollector / GetCollector
// ---------------------------------------------------------------------------

func TestNewCollector_Default(t *testing.T) {
	os.Unsetenv("WORKER_CANONICAL_METRICS")
	c := NewCollector()
	assert.Equal(t, "legacy", c.CollectorName())
}

func TestNewCollector_Canonical(t *testing.T) {
	os.Setenv("WORKER_CANONICAL_METRICS", "true")
	t.Cleanup(func() { os.Unsetenv("WORKER_CANONICAL_METRICS") })
	c := NewCollector()
	assert.Equal(t, "canonical", c.CollectorName())
}

func TestGetCollector(t *testing.T) {
	resetState(t)
	assert.Equal(t, "noop", GetCollector().CollectorName())
}

// ---------------------------------------------------------------------------
// noopCollector — exercise every method for coverage
// ---------------------------------------------------------------------------

func TestNoopCollector_Capabilities(t *testing.T) {
	n := &noopCollector{}
	assert.False(t, n.ShouldRecordPayloadSize())
	assert.False(t, n.ShouldRecordHTTPRequests())
}

func TestNoopCollector(t *testing.T) {
	n := &noopCollector{}
	assert.Equal(t, "noop", n.CollectorName())

	assert.NotPanics(t, func() {
		n.Register()
		n.IncrementTaskPoll("t")
		n.IncrementTaskPollError("t", errors.New("e"))
		n.IncrementTaskExecutionStarted("t")
		n.IncrementTaskExecuteError("t", errors.New("e"))
		n.IncrementTaskUpdateError("t", errors.New("e"))
		n.IncrementTaskAckError("t", errors.New("e"))
		n.IncrementTaskAckFailed("t")
		n.IncrementTaskExecutionQueueFull("t")
		n.IncrementTaskPaused("t")
		n.IncrementUncaughtException("panic!")
		n.IncrementExternalPayloadUsed("e", "o", "p")
		n.IncrementWorkflowStartError("w", errors.New("e"))
		n.RecordTaskPollTime("t", 1.0, nil)
		n.RecordTaskExecuteTime("t", 1.0, nil)
		n.RecordTaskUpdateTime("t", 1.0, nil)
		n.RecordHTTPRequestTime("GET", "/api", "200", 0.5)
		n.RecordTaskResultPayloadSize("t", 100)
		n.RecordWorkflowInputPayloadSize("w", "1", 200)
		n.SetActiveWorkers("t", 5)
	})
}

// ---------------------------------------------------------------------------
// counter.go / gauge.go / histogram.go — low-level helpers
// ---------------------------------------------------------------------------

func TestIncrementCounter_CollectionDisabled(t *testing.T) {
	resetState(t)
	assert.NotPanics(t, func() {
		incrementCounter(TASK_POLL, []string{"t"})
	})
}

func TestSetGauge_CollectionDisabled(t *testing.T) {
	resetState(t)
	assert.NotPanics(t, func() {
		setGauge(TASK_POLL_TIME, []string{"t"}, 1.0)
	})
}

func TestObserveHistogram_CollectionDisabled(t *testing.T) {
	resetState(t)
	assert.NotPanics(t, func() {
		observeHistogram(TASK_POLL_TIME_SECONDS, []string{"t", "SUCCESS"}, 0.5)
	})
}

func TestGetCounter_NotRegistered(t *testing.T) {
	resetState(t)
	c := getCounter("nonexistent", []string{})
	assert.Nil(t, c)
}

func TestGetGauge_NotRegistered(t *testing.T) {
	resetState(t)
	g := getGauge("nonexistent", []string{})
	assert.Nil(t, g)
}

func TestGetHistogram_NotRegistered(t *testing.T) {
	resetState(t)
	h := getHistogram("nonexistent", []string{})
	assert.Nil(t, h)
}

func TestNewCounter(t *testing.T) {
	d := NewMetricDetails(TASK_POLL, TASK_POLL_DOC, []MetricLabel{TASK_TYPE})
	cv := newCounter(d)
	require.NotNil(t, cv)
}

func TestNewGauge(t *testing.T) {
	d := NewMetricDetails(TASK_POLL_TIME, TASK_POLL_TIME_DOC, []MetricLabel{TASK_TYPE})
	gv := newGauge(d)
	require.NotNil(t, gv)
}

func TestNewHistogram(t *testing.T) {
	d := NewMetricDetails(TASK_POLL_TIME_SECONDS, TASK_POLL_TIME_SECONDS_DOC, []MetricLabel{TASK_TYPE, STATUS})
	hv := newHistogram(d, CanonicalDurationBuckets)
	require.NotNil(t, hv)
}

// ---------------------------------------------------------------------------
// canonicalCollector
// ---------------------------------------------------------------------------

func TestCanonicalCollector_RegisterAndMethods(t *testing.T) {
	resetState(t)
	c := &canonicalCollector{recordPayloadSize: true, recordHTTPRequests: true}
	c.Register()
	collectionEnabled.Store(true)

	assert.Equal(t, "canonical", c.CollectorName())

	testErr := errors.New("test-error")

	assert.NotPanics(t, func() {
		c.IncrementTaskPoll("myTask")
		c.IncrementTaskPollError("myTask", testErr)
		c.IncrementTaskExecutionStarted("myTask")
		c.IncrementTaskExecuteError("myTask", testErr)
		c.IncrementTaskUpdateError("myTask", testErr)
		c.IncrementTaskAckError("myTask", testErr)
		c.IncrementTaskAckFailed("myTask")
		c.IncrementTaskExecutionQueueFull("myTask")
		c.IncrementTaskPaused("myTask")
		c.IncrementUncaughtException(testErr)
		c.IncrementExternalPayloadUsed("entity", "op", "payload")
		c.IncrementWorkflowStartError("myWorkflow", testErr)
		c.RecordTaskPollTime("myTask", 0.5, nil)
		c.RecordTaskPollTime("myTask", 0.1, testErr)
		c.RecordTaskExecuteTime("myTask", 1.0, nil)
		c.RecordTaskUpdateTime("myTask", 0.2, nil)
		c.RecordHTTPRequestTime("GET", "/api/tasks", "200", 0.05)
		c.RecordTaskResultPayloadSize("myTask", 1024)
		c.RecordWorkflowInputPayloadSize("myWorkflow", "1", 2048)
		c.SetActiveWorkers("myTask", 3)
	})

	// Verify a counter was actually incremented.
	cv, ok := counterByName[TASK_POLL_TOTAL]
	require.True(t, ok)
	m := &dto.Metric{}
	counter, _ := cv.GetMetricWithLabelValues("myTask")
	require.NoError(t, counter.Write(m))
	assert.Equal(t, float64(1), m.GetCounter().GetValue())

	// Verify a gauge was set.
	gv, ok := gaugeByName[ACTIVE_WORKERS]
	require.True(t, ok)
	gauge, _ := gv.GetMetricWithLabelValues("myTask")
	m = &dto.Metric{}
	require.NoError(t, gauge.Write(m))
	assert.Equal(t, float64(3), m.GetGauge().GetValue())
}

func TestCanonicalCollector_Capabilities_Defaults(t *testing.T) {
	os.Unsetenv("WORKER_METRICS_PAYLOAD_SIZE")
	os.Unsetenv("WORKER_METRICS_HTTP_REQUESTS")
	c := &canonicalCollector{
		recordPayloadSize:  envBoolDefault("WORKER_METRICS_PAYLOAD_SIZE", true),
		recordHTTPRequests: envBoolDefault("WORKER_METRICS_HTTP_REQUESTS", true),
	}
	assert.True(t, c.ShouldRecordPayloadSize())
	assert.True(t, c.ShouldRecordHTTPRequests())
}

func TestCanonicalCollector_Capabilities_EnvOverride(t *testing.T) {
	os.Setenv("WORKER_METRICS_PAYLOAD_SIZE", "false")
	os.Unsetenv("WORKER_METRICS_HTTP_REQUESTS")
	t.Cleanup(func() {
		os.Unsetenv("WORKER_METRICS_PAYLOAD_SIZE")
		os.Unsetenv("WORKER_METRICS_HTTP_REQUESTS")
	})
	c := &canonicalCollector{
		recordPayloadSize:  envBoolDefault("WORKER_METRICS_PAYLOAD_SIZE", true),
		recordHTTPRequests: envBoolDefault("WORKER_METRICS_HTTP_REQUESTS", true),
	}
	assert.False(t, c.ShouldRecordPayloadSize())
	assert.True(t, c.ShouldRecordHTTPRequests())
}

// ---------------------------------------------------------------------------
// legacyCollector
// ---------------------------------------------------------------------------

func TestLegacyCollector_Capabilities(t *testing.T) {
	l := &legacyCollector{}
	assert.False(t, l.ShouldRecordPayloadSize())
	assert.False(t, l.ShouldRecordHTTPRequests())
}

func TestLegacyCollector_RegisterAndMethods(t *testing.T) {
	resetState(t)
	l := &legacyCollector{}
	l.Register()
	collectionEnabled.Store(true)

	assert.Equal(t, "legacy", l.CollectorName())

	testErr := errors.New("test-error")

	assert.NotPanics(t, func() {
		l.IncrementTaskPoll("myTask")
		l.IncrementTaskPollError("myTask", testErr)
		l.IncrementTaskExecuteError("myTask", testErr)
		l.IncrementTaskUpdateError("myTask", testErr)
		l.IncrementTaskExecutionQueueFull("myTask")
		l.IncrementTaskPaused("myTask")
		l.IncrementUncaughtException(testErr)
		l.IncrementExternalPayloadUsed("entity", "op", "payload")
		l.IncrementWorkflowStartError("myWorkflow", testErr)
		l.RecordTaskPollTime("myTask", 0.5, nil)
		l.RecordTaskExecuteTime("myTask", 1.0, nil)
		l.RecordTaskUpdateTime("myTask", 0.2, nil)
		l.RecordTaskResultPayloadSize("myTask", 512)
		l.RecordWorkflowInputPayloadSize("myWorkflow", "1", 1024)

		// Noop stubs for canonical-only methods.
		l.IncrementTaskExecutionStarted("myTask")
		l.IncrementTaskAckError("myTask", testErr)
		l.IncrementTaskAckFailed("myTask")
		l.RecordHTTPRequestTime("GET", "/api", "200", 0.5)
		l.SetActiveWorkers("myTask", 5)
	})

	// Verify a counter was incremented.
	cv, ok := counterByName[TASK_POLL]
	require.True(t, ok)
	m := &dto.Metric{}
	counter, _ := cv.GetMetricWithLabelValues("myTask")
	require.NoError(t, counter.Write(m))
	assert.Equal(t, float64(1), m.GetCounter().GetValue())

	// Verify a gauge was set (execute time stored as ms).
	gv, ok := gaugeByName[TASK_EXECUTE_TIME]
	require.True(t, ok)
	gauge, _ := gv.GetMetricWithLabelValues("myTask")
	m = &dto.Metric{}
	require.NoError(t, gauge.Write(m))
	assert.Equal(t, float64(1000), m.GetGauge().GetValue())
}

// ---------------------------------------------------------------------------
// Package-level convenience functions
// ---------------------------------------------------------------------------

func TestPackageLevelFunctions(t *testing.T) {
	resetState(t)

	os.Unsetenv("WORKER_CANONICAL_METRICS")
	InitCollector()

	assert.NotPanics(t, func() {
		IncrementTaskPoll("t")
		IncrementTaskPollError("t", errors.New("e"))
		IncrementTaskExecutionStarted("t")
		IncrementTaskExecuteError("t", errors.New("e"))
		IncrementTaskUpdateError("t", errors.New("e"))
		IncrementTaskAckError("t", errors.New("e"))
		IncrementTaskAckFailed("t")
		IncrementTaskExecutionQueueFull("t")
		IncrementTaskPaused("t")
		IncrementUncaughtException("panic!")
		IncrementExternalPayloadUsed("e", "o", "p")
		IncrementWorkflowStartError("w", errors.New("e"))
		RecordTaskPollTime("t", 1.0)
		RecordTaskExecuteTime("t", 1.0)
		RecordTaskUpdateTime("t", 1.0)
		RecordHTTPRequestTime("GET", "/api", "200", 0.5)
		RecordTaskResultPayloadSize("t", 100)
		RecordWorkflowInputPayloadSize("w", "1", 200)
		SetActiveWorkers("t", 5)
	})
}

func TestGetCollector_DirectCallWithError(t *testing.T) {
	resetState(t)

	os.Setenv("WORKER_CANONICAL_METRICS", "true")
	t.Cleanup(func() { os.Unsetenv("WORKER_CANONICAL_METRICS") })
	InitCollector()

	testErr := errors.New("poll-timeout")
	c := GetCollector()

	assert.NotPanics(t, func() {
		c.RecordTaskPollTime("myTask", 0.5, nil)
		c.RecordTaskPollTime("myTask", 1.2, testErr)
		c.RecordTaskExecuteTime("myTask", 0.3, nil)
		c.RecordTaskExecuteTime("myTask", 5.0, testErr)
		c.RecordTaskUpdateTime("myTask", 0.1, nil)
		c.RecordTaskUpdateTime("myTask", 2.0, testErr)
	})

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	foundSuccess := false
	foundFailure := false
	for _, fam := range families {
		if fam.GetName() != "task_poll_time_seconds" {
			continue
		}
		for _, m := range fam.GetMetric() {
			for _, lp := range m.GetLabel() {
				if lp.GetName() == "status" && lp.GetValue() == "SUCCESS" {
					foundSuccess = true
				}
				if lp.GetName() == "status" && lp.GetValue() == "FAILURE" {
					foundFailure = true
				}
			}
		}
	}
	assert.True(t, foundSuccess, "expected SUCCESS status label from nil error")
	assert.True(t, foundFailure, "expected FAILURE status label from non-nil error")
}

// ---------------------------------------------------------------------------
// InitCollector / handlePanicError
// ---------------------------------------------------------------------------

func TestInitCollector(t *testing.T) {
	resetState(t)
	os.Unsetenv("WORKER_CANONICAL_METRICS")

	InitCollector()

	assert.True(t, collectionEnabled.Load())
	assert.Equal(t, "legacy", GetCollector().CollectorName())
}

func TestInitCollector_Canonical(t *testing.T) {
	resetState(t)
	os.Setenv("WORKER_CANONICAL_METRICS", "1")
	t.Cleanup(func() { os.Unsetenv("WORKER_CANONICAL_METRICS") })

	InitCollector()

	assert.True(t, collectionEnabled.Load())
	assert.Equal(t, "canonical", GetCollector().CollectorName())
}

func TestInitCollector_ConcurrentSafe(t *testing.T) {
	resetState(t)
	os.Unsetenv("WORKER_CANONICAL_METRICS")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			InitCollector()
		}()
	}
	wg.Wait()

	assert.True(t, collectionEnabled.Load())
	assert.Equal(t, "legacy", GetCollector().CollectorName())
}

// TestInitCollector_RaceWithReaders forces the interleaving that the singleton
// publication has to be safe against: worker/round-tripper hot-path readers
// hitting the package-level collector while InitCollector publishes a new one.
// Run with `go test -race` to catch any unsynchronized access to the collector,
// the collectionEnabled flag, or the metric maps populated by Register().
func TestInitCollector_RaceWithReaders(t *testing.T) {
	resetState(t)
	os.Unsetenv("WORKER_CANONICAL_METRICS")

	const readers = 8
	start := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 2000; j++ {
				IncrementTaskPoll("race_task")
				GetCollector().RecordTaskPollTime("race_task", 0.01, nil)
				_ = ShouldRecordHTTPRequests()
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-start
		InitCollector()
	}()

	close(start)
	wg.Wait()

	assert.True(t, collectionEnabled.Load())
	assert.Equal(t, "legacy", GetCollector().CollectorName())
}

func TestHandlePanicError_NoPanic(t *testing.T) {
	resetState(t)
	assert.NotPanics(t, func() {
		defer handlePanicError("test")
	})
}

func TestHandlePanicError_WithPanic(t *testing.T) {
	resetState(t)
	assert.NotPanics(t, func() {
		defer handlePanicError("test")
		panic("kaboom")
	})
}

// ---------------------------------------------------------------------------
// MetricDetails
// ---------------------------------------------------------------------------

func TestNewMetricDetails(t *testing.T) {
	d := NewMetricDetails(TASK_POLL_TOTAL, TASK_POLL_DOC, []MetricLabel{TASK_TYPE, EXCEPTION})
	assert.Equal(t, "task_poll_total", d.Name)
	assert.Equal(t, string(TASK_POLL_DOC), d.Description)
	assert.Equal(t, []string{"taskType", "exception"}, d.Labels)
}

func TestNewMetricDetails_NoLabels(t *testing.T) {
	d := NewMetricDetails(THREAD_UNCAUGHT_EXCEPTION, THREAD_UNCAUGHT_EXCEPTION_DOC, []MetricLabel{})
	assert.Equal(t, "thread_uncaught_exceptions", d.Name)
	assert.Nil(t, d.Labels)
}

// ---------------------------------------------------------------------------
// Prometheus output tests — verify template URIs in scraped metrics
// ---------------------------------------------------------------------------

func TestPrometheusOutput_HttpMetric_TemplateURI(t *testing.T) {
	resetState(t)
	c := &canonicalCollector{recordPayloadSize: true, recordHTTPRequests: true}
	c.Register()
	collectionEnabled.Store(true)
	setCollector(c)

	RecordHTTPRequestTime("GET", "/workflow/{workflowId}", "200", 0.042)

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	found := false
	for _, fam := range families {
		if fam.GetName() != "http_api_client_request_seconds" {
			continue
		}
		for _, m := range fam.GetMetric() {
			for _, lp := range m.GetLabel() {
				if lp.GetName() == "uri" && lp.GetValue() == "/workflow/{workflowId}" {
					found = true
				}
			}
		}
	}
	assert.True(t, found, "expected uri=/workflow/{workflowId} in Prometheus output")
}

func TestPrometheusOutput_HttpMetric_BoundedCardinality(t *testing.T) {
	resetState(t)
	c := &canonicalCollector{recordPayloadSize: true, recordHTTPRequests: true}
	c.Register()
	collectionEnabled.Store(true)
	setCollector(c)

	for i := 0; i < 100; i++ {
		RecordHTTPRequestTime("GET", "/workflow/{workflowId}", "200", 0.01)
	}

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	uris := map[string]bool{}
	for _, fam := range families {
		if fam.GetName() != "http_api_client_request_seconds" {
			continue
		}
		for _, m := range fam.GetMetric() {
			for _, lp := range m.GetLabel() {
				if lp.GetName() == "uri" {
					uris[lp.GetValue()] = true
				}
			}
		}
	}
	assert.Len(t, uris, 1, "expected exactly 1 distinct uri series")
	assert.True(t, uris["/workflow/{workflowId}"])
}

func TestPrometheusOutput_HttpMetric_MultipleEndpoints(t *testing.T) {
	resetState(t)
	c := &canonicalCollector{recordPayloadSize: true, recordHTTPRequests: true}
	c.Register()
	collectionEnabled.Store(true)
	setCollector(c)

	RecordHTTPRequestTime("GET", "/workflow/{workflowId}", "200", 0.01)
	RecordHTTPRequestTime("POST", "/tasks/{taskId}/log", "200", 0.02)
	RecordHTTPRequestTime("PUT", "/tasks", "500", 0.5)

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	uris := map[string]bool{}
	for _, fam := range families {
		if fam.GetName() != "http_api_client_request_seconds" {
			continue
		}
		for _, m := range fam.GetMetric() {
			for _, lp := range m.GetLabel() {
				if lp.GetName() == "uri" {
					uris[lp.GetValue()] = true
				}
			}
		}
	}
	assert.True(t, uris["/workflow/{workflowId}"])
	assert.True(t, uris["/tasks/{taskId}/log"])
	assert.True(t, uris["/tasks"])
}

func TestPrometheusOutput_HttpMetric_ErrorStatus(t *testing.T) {
	resetState(t)
	c := &canonicalCollector{recordPayloadSize: true, recordHTTPRequests: true}
	c.Register()
	collectionEnabled.Store(true)
	setCollector(c)

	RecordHTTPRequestTime("GET", "/workflow/{workflowId}", "0", 0.1)

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	foundStatus := false
	for _, fam := range families {
		if fam.GetName() != "http_api_client_request_seconds" {
			continue
		}
		for _, m := range fam.GetMetric() {
			for _, lp := range m.GetLabel() {
				if lp.GetName() == "status" && lp.GetValue() == "0" {
					foundStatus = true
				}
			}
		}
	}
	assert.True(t, foundStatus, "expected status=0 for transport errors")
}
