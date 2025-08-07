package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewZap(t *testing.T) {
	logger := zap.NewNop()
	zapLogger := NewZap(logger)

	assert.NotNil(t, zapLogger, "NewZap should return a non-nil logger")
	assert.IsType(t, &ZapLogger{}, zapLogger, "NewZap should return a *ZapLogger")
}

func TestZapLogger_AllLevels(t *testing.T) {
	tests := []struct {
		name   string
		level  zapcore.Level
		logFn  func(Logger)
		msg    string
		fields map[string]interface{}
	}{
		{
			name:  "Debug",
			level: zapcore.DebugLevel,
			logFn: func(l Logger) { l.Debug("debug message", "key", "value") },
			msg:   "debug message",
			fields: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name:  "Info",
			level: zapcore.InfoLevel,
			logFn: func(l Logger) { l.Info("info message", "user", "john") },
			msg:   "info message",
			fields: map[string]interface{}{
				"user": "john",
			},
		},
		{
			name:  "Warn",
			level: zapcore.WarnLevel,
			logFn: func(l Logger) { l.Warn("warning message", "status", 400) },
			msg:   "warning message",
			fields: map[string]interface{}{
				"status": int64(400),
			},
		},
		{
			name:  "Error",
			level: zapcore.ErrorLevel,
			logFn: func(l Logger) { l.Error("error message", "error", "something went wrong") },
			msg:   "error message",
			fields: map[string]interface{}{
				"error": "something went wrong",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, observed := observer.New(tt.level)
			logger := zap.New(core)
			zapLogger := NewZap(logger)

			tt.logFn(zapLogger)

			logs := observed.All()
			require.Len(t, logs, 1, "Should have exactly one log entry")

			entry := logs[0]
			assert.Equal(t, tt.level, entry.Level, "Log level should match")
			assert.Equal(t, tt.msg, entry.Message, "Message should match")
			assert.Len(t, entry.Context, len(tt.fields), "Should have expected number of fields")

			// Check all fields
			actualFields := make(map[string]interface{})
			for _, field := range entry.Context {
				switch field.Type {
				case zapcore.StringType:
					actualFields[field.Key] = field.String
				case zapcore.Int64Type:
					actualFields[field.Key] = field.Integer
				default:
					actualFields[field.Key] = field.Interface
				}
			}

			for key, expectedValue := range tt.fields {
				assert.Equal(t, expectedValue, actualFields[key], "Field %s should match", key)
			}
		})
	}
}

func TestZapLogger_Fatal(t *testing.T) {
	core, _ := observer.New(zapcore.FatalLevel)
	logger := zap.New(core)
	_ = NewZap(logger)

	msg, fields := splitArgs([]interface{}{"fatal error", "code", 500})
	assert.Equal(t, "fatal error", msg, "Message should be extracted correctly")
	assert.Len(t, fields, 1, "Should have one field")
	assert.Equal(t, "code", fields[0].Key, "Field key should match")
	assert.Equal(t, int64(500), fields[0].Integer, "Field value should match")
}

func TestZapLogger_With(t *testing.T) {
	core, observed := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)
	zapLogger := NewZap(logger)

	// Create a logger with context
	contextLogger := zapLogger.With("requestId", "123", "userId", "456")
	assert.IsType(t, &ZapLogger{}, contextLogger, "With should return a *ZapLogger")

	// Log with the context logger
	contextLogger.Info("processing request")

	logs := observed.All()
	require.Len(t, logs, 1, "Should have exactly one log entry")

	entry := logs[0]
	assert.Equal(t, "processing request", entry.Message, "Message should match")
	assert.Len(t, entry.Context, 2, "Should have two context fields")

	// Check context fields more robustly
	expectedFields := map[string]string{
		"requestId": "123",
		"userId":    "456",
	}

	actualFields := make(map[string]string)
	for _, field := range entry.Context {
		actualFields[field.Key] = field.String
	}

	assert.Equal(t, expectedFields, actualFields, "Context fields should match")
}

func TestZapLogger_ChainedWith(t *testing.T) {
	core, observed := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)
	zapLogger := NewZap(logger)

	// Chain multiple With calls
	contextLogger := zapLogger.With("service", "api").With("version", "1.0")
	contextLogger.Info("test message")

	logs := observed.All()
	require.Len(t, logs, 1, "Should have exactly one log entry")

	entry := logs[0]
	assert.Len(t, entry.Context, 2, "Should have two context fields from chained With calls")

	expectedFields := map[string]string{
		"service": "api",
		"version": "1.0",
	}

	actualFields := make(map[string]string)
	for _, field := range entry.Context {
		actualFields[field.Key] = field.String
	}

	assert.Equal(t, expectedFields, actualFields, "Chained context fields should match")
}

func TestSplitArgs_WithMessage(t *testing.T) {
	tests := []struct {
		name           string
		args           []interface{}
		expectedMsg    string
		expectedFields int
		description    string
	}{
		{
			name:           "message with key-value pairs (odd total)",
			args:           []interface{}{"hello", "key1", "value1", "key2", "value2"},
			expectedMsg:    "hello",
			expectedFields: 2,
			description:    "5 args: first is message, rest are 2 key-value pairs",
		},
		{
			name:           "only key-value pairs (even total)",
			args:           []interface{}{"key1", "value1", "key2", "value2"},
			expectedMsg:    "",
			expectedFields: 2,
			description:    "4 args: no message, 2 key-value pairs",
		},
		{
			name:           "message with even number of args",
			args:           []interface{}{"hello", "key1", "value1", "key2"},
			expectedMsg:    "",
			expectedFields: 2,
			description:    "4 args: even count means no message extraction",
		},
		{
			name:           "message with odd number of args",
			args:           []interface{}{"hello", "key1", "value1"},
			expectedMsg:    "hello",
			expectedFields: 1,
			description:    "3 args: odd count means first is message",
		},
		{
			name:           "single message only",
			args:           []interface{}{"hello"},
			expectedMsg:    "hello",
			expectedFields: 0,
			description:    "1 arg: odd count, just message",
		},
		{
			name:           "empty args",
			args:           []interface{}{},
			expectedMsg:    "",
			expectedFields: 0,
			description:    "no args at all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, fields := splitArgs(tt.args)

			assert.Equal(t, tt.expectedMsg, msg, "Message should match expected (%s)", tt.description)
			assert.Len(t, fields, tt.expectedFields, "Fields count should match expected (%s)", tt.description)

			// Verify field structure for non-empty fields
			for _, field := range fields {
				assert.NotEmpty(t, field.Key, "Field key should not be empty")
			}
		})
	}
}

func TestSplitArgs_HandlesDifferentTypes(t *testing.T) {
	testStruct := struct{ Name string }{"test"}
	args := []interface{}{"message", "string", "value", "int", 42, "bool", true, "struct", testStruct}

	msg, fields := splitArgs(args)

	assert.Equal(t, "message", msg, "Message should be extracted")
	assert.Len(t, fields, 4, "Should have 4 fields")

	expectedKeys := []string{"string", "int", "bool", "struct"}
	for i, field := range fields {
		assert.Equal(t, expectedKeys[i], field.Key, "Field key should match expected")
		assert.NotNil(t, field, "Field should not be nil")
	}
}
