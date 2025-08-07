package log

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger wraps zap.Logger
type ZapLogger struct {
	l *zap.Logger
}

// NewZap creates a new ZapLogger that wraps a zap.Logger.
func NewZap(l *zap.Logger) Logger {
	return &ZapLogger{l: l}
}

// Debug logs a debug level message
func (z *ZapLogger) Debug(args ...interface{}) {
	msg, f := splitArgs(args)
	z.l.Debug(msg, f...)
}

// Info logs an info level message
func (z *ZapLogger) Info(args ...interface{}) {
	msg, f := splitArgs(args)
	z.l.Info(msg, f...)
}

// Warn logs a warning level message
func (z *ZapLogger) Warn(args ...interface{}) {
	msg, f := splitArgs(args)
	z.l.Warn(msg, f...)
}

// Error logs an error level message
func (z *ZapLogger) Error(args ...interface{}) {
	msg, f := splitArgs(args)
	z.l.Error(msg, f...)
}

// Fatal logs a fatal level message
func (z *ZapLogger) Fatal(args ...interface{}) {
	msg, f := splitArgs(args)
	z.l.Fatal(msg, f...)
}

// With creates a new logger with the given key-value pair
func (z *ZapLogger) With(value ...interface{}) Logger {
	_, f := splitArgs(value)
	return &ZapLogger{l: z.l.With(f...)}
}

func splitArgs(args []interface{}) (string, []zap.Field) {
	var msg string
	start := 0

	if len(args) > 0 && len(args)%2 == 1 {
		msg = fmt.Sprintf("%v", args[0])
		start = 1
	}

	fields := make([]zap.Field, 0, (len(args)-start+1)/2)
	for i := start; i < len(args); i += 2 {
		key := fmt.Sprintf("%v", args[i])

		if i+1 < len(args) {
			fields = append(fields, zap.Any(key, args[i+1]))
		}
	}
	return msg, fields
}
