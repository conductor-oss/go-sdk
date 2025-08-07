package log

var defaultLogger Logger

func init() {
	defaultLogger = NewStd(nil)
}

// Logger is an interface for logging.
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	With(value ...interface{}) Logger
}

// SetLogger sets a custom logger implementation. If nil is passed, uses the standard logger.
func SetLogger(l Logger) {
	if l == nil {
		defaultLogger = NewStd(nil)
		return
	}
	defaultLogger = l
}

// Info logs an info level message.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Debug logs a debug level message.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Warn logs a warning level message.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Error logs an error level message.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Fatal logs a formatted error level message.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// With creates a new logger with the given key-value pair.
func With(value ...interface{}) Logger {
	return defaultLogger.With(value...)
}
