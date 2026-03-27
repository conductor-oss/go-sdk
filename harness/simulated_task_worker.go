package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

const alphanumericChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var instanceID = func() string {
	if h := os.Getenv("HOSTNAME"); h != "" {
		return h
	}
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "00000000"
	}
	return hex.EncodeToString(b)
}()

type SimulatedTaskWorker struct {
	taskName     string
	codename     string
	defaultDelay time.Duration
	batchSize    int
	pollInterval time.Duration
	workerID     string
	rng          *mrand.Rand //nolint:gosec // intentionally non-crypto: fast seedable RNG for simulation delays and output
}

func NewSimulatedTaskWorker(taskName, codename string, sleepSeconds, batchSize, pollIntervalMs int) *SimulatedTaskWorker {
	w := &SimulatedTaskWorker{
		taskName:     taskName,
		codename:     codename,
		defaultDelay: time.Duration(sleepSeconds) * time.Second,
		batchSize:    batchSize,
		pollInterval: time.Duration(pollIntervalMs) * time.Millisecond,
		workerID:     fmt.Sprintf("%s-%s", taskName, instanceID),
		rng:          mrand.New(mrand.NewSource(time.Now().UnixNano())), //nolint:gosec // see struct field comment
	}

	fmt.Printf("[%s] Initialized worker [workerId=%s, codename=%s, batchSize=%d, pollInterval=%dms]\n",
		w.taskName, w.workerID, w.codename, w.batchSize, pollIntervalMs)

	return w
}

func (w *SimulatedTaskWorker) Execute(task *model.Task) (interface{}, error) {
	input := task.InputData
	if input == nil {
		input = make(map[string]interface{})
	}
	taskID := task.TaskId
	taskIndex := getIntOrDefault(input, "taskIndex", -1)

	fmt.Printf("[%s] Starting simulated task [id=%s, index=%d, codename=%s]\n",
		w.taskName, taskID, taskIndex, w.codename)

	startTime := time.Now()

	delayType := getStringOrDefault(input, "delayType", "fixed")
	minDelay := getIntOrDefault(input, "minDelay", int(w.defaultDelay.Milliseconds()))
	maxDelay := getIntOrDefault(input, "maxDelay", minDelay+100)
	meanDelay := getIntOrDefault(input, "meanDelay", (minDelay+maxDelay)/2)
	stdDeviation := getIntOrDefault(input, "stdDeviation", 30)
	successRate := getFloat64OrDefault(input, "successRate", 1.0)
	failureMode := getStringOrDefault(input, "failureMode", "random")
	outputSize := getIntOrDefault(input, "outputSize", 1024)

	var delayMs int64
	if !strings.EqualFold(delayType, "wait") {
		delayMs = w.calculateDelay(delayType, minDelay, maxDelay, meanDelay, stdDeviation)

		fmt.Printf("[%s] Simulated task [id=%s, index=%d] sleeping for %d ms\n",
			w.taskName, taskID, taskIndex, delayMs)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}

	if !w.shouldTaskSucceed(successRate, failureMode, input) {
		fmt.Printf("[%s] Simulated task [id=%s, index=%d] failed as configured\n",
			w.taskName, taskID, taskIndex)
		return nil, fmt.Errorf("simulated task failure based on configuration")
	}

	elapsed := time.Since(startTime).Milliseconds()
	output := w.generateOutput(input, taskID, taskIndex, delayMs, elapsed, outputSize)
	return output, nil
}

func (w *SimulatedTaskWorker) calculateDelay(delayType string, minDelay, maxDelay, meanDelay, stdDeviation int) int64 {
	switch strings.ToLower(delayType) {
	case "fixed":
		return int64(minDelay)

	case "random":
		spread := maxDelay - minDelay + 1
		if spread < 1 {
			spread = 1
		}
		return int64(minDelay) + int64(w.rng.Intn(spread))

	case "normal":
		gaussian := w.nextGaussian()
		delay := math.Round(float64(meanDelay) + gaussian*float64(stdDeviation))
		if delay < 1 {
			return 1
		}
		return int64(delay)

	case "exponential":
		exp := -float64(meanDelay) * math.Log(1-w.rng.Float64())
		result := int64(exp)
		if result < int64(minDelay) {
			return int64(minDelay)
		}
		if result > int64(maxDelay) {
			return int64(maxDelay)
		}
		return result

	default:
		return int64(minDelay)
	}
}

func (w *SimulatedTaskWorker) nextGaussian() float64 {
	u1 := 1.0 - w.rng.Float64()
	u2 := w.rng.Float64()
	return math.Sqrt(-2.0*math.Log(u1)) * math.Sin(2.0*math.Pi*u2)
}

func (w *SimulatedTaskWorker) shouldTaskSucceed(successRate float64, failureMode string, input map[string]interface{}) bool {
	if v, ok := input["forceSuccess"]; ok && v != nil {
		if b, ok := toBool(v); ok {
			return b
		}
	}
	if v, ok := input["forceFail"]; ok && v != nil {
		if b, ok := toBool(v); ok {
			return !b
		}
	}

	switch strings.ToLower(failureMode) {
	case "random":
		return w.rng.Float64() < successRate
	case "conditional":
		return w.shouldConditionalSucceed(successRate, input)
	case "sequential":
		attempt := getIntOrDefault(input, "attempt", 1)
		failUntilAttempt := getIntOrDefault(input, "failUntilAttempt", 2)
		return attempt >= failUntilAttempt
	default:
		return w.rng.Float64() < successRate
	}
}

func (w *SimulatedTaskWorker) shouldConditionalSucceed(successRate float64, input map[string]interface{}) bool {
	taskIndex := getIntOrDefault(input, "taskIndex", -1)
	if taskIndex >= 0 {
		if failIndexes, ok := input["failIndexes"]; ok {
			if arr, ok := failIndexes.([]interface{}); ok {
				for _, idx := range arr {
					if toInt(idx) == taskIndex {
						return false
					}
				}
			}
		}
		failEvery := getIntOrDefault(input, "failEvery", 0)
		if failEvery > 0 && taskIndex%failEvery == 0 {
			return false
		}
	}
	return w.rng.Float64() < successRate
}

func (w *SimulatedTaskWorker) generateOutput(
	input map[string]interface{}, taskID string, taskIndex int,
	delayMs, elapsedTimeMs int64, outputSize int,
) map[string]interface{} {
	output := map[string]interface{}{
		"taskId":                taskID,
		"taskIndex":             taskIndex,
		"codename":              w.codename,
		"status":                "completed",
		"configuredDelayMs":     delayMs,
		"actualExecutionTimeMs": elapsedTimeMs,
		"a_or_b": func() string {
			if w.rng.Intn(100) > 20 {
				return "a"
			}
			return "b"
		}(),
		"c_or_d": func() string {
			if w.rng.Intn(100) > 33 {
				return "c"
			}
			return "d"
		}(),
	}

	if getBoolOrDefault(input, "includeInput", false) {
		output["input"] = input
	}

	if prev, ok := input["previousTaskOutput"]; ok && prev != nil {
		output["previousTaskData"] = prev
	}

	if outputSize > 0 {
		output["data"] = generateRandomData(w.rng, outputSize)
	}

	if tmpl, ok := input["outputTemplate"]; ok {
		if m, ok := tmpl.(map[string]interface{}); ok {
			for k, v := range m {
				output[k] = v
			}
		}
	}

	return output
}

func generateRandomData(rng *mrand.Rand, size int) string {
	if size <= 0 {
		return ""
	}
	b := make([]byte, size)
	for i := range b {
		b[i] = alphanumericChars[rng.Intn(len(alphanumericChars))]
	}
	return string(b)
}

func getIntOrDefault(input map[string]interface{}, key string, defaultVal int) int {
	v, ok := input[key]
	if !ok || v == nil {
		return defaultVal
	}
	return toInt(v)
}

func getFloat64OrDefault(input map[string]interface{}, key string, defaultVal float64) float64 {
	v, ok := input[key]
	if !ok || v == nil {
		return defaultVal
	}
	return toFloat64(v)
}

func getStringOrDefault(input map[string]interface{}, key string, defaultVal string) string {
	v, ok := input[key]
	if !ok || v == nil {
		return defaultVal
	}
	if s, ok := v.(string); ok {
		return s
	}
	return defaultVal
}

func getBoolOrDefault(input map[string]interface{}, key string, defaultVal bool) bool {
	v, ok := input[key]
	if !ok || v == nil {
		return defaultVal
	}
	if b, ok := toBool(v); ok {
		return b
	}
	return defaultVal
}

func toInt(v interface{}) int {
	switch n := v.(type) {
	case int:
		return n
	case int32:
		return int(n)
	case int64:
		return int(n)
	case float32:
		return int(n)
	case float64:
		return int(n)
	case string:
		if i, err := strconv.Atoi(n); err == nil {
			return i
		}
		return 0
	default:
		return 0
	}
}

func toFloat64(v interface{}) float64 {
	switch n := v.(type) {
	case float64:
		return n
	case float32:
		return float64(n)
	case int:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	case string:
		if f, err := strconv.ParseFloat(n, 64); err == nil {
			return f
		}
		return 0
	default:
		return 0
	}
}

func toBool(v interface{}) (bool, bool) {
	switch b := v.(type) {
	case bool:
		return b, true
	case string:
		switch strings.ToLower(b) {
		case "true", "1":
			return true, true
		case "false", "0":
			return false, true
		}
	case float64:
		return b != 0, true
	case int:
		return b != 0, true
	}
	return false, false
}
