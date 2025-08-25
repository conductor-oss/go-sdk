package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// StdLogger is a logger that writes to the standard output
type StdLogger struct {
	l     *log.Logger
	level int
}

const (
	// LvlDebug is the debug level
	LvlDebug = iota
	// LvlInfo is the info level
	LvlInfo
	// LvlWarn is the warn level
	LvlWarn
	// LvlError is the error level
	LvlError
)

func levelStr(l int) string {
	switch l {
	case LvlDebug:
		return "[DEBUG]"
	case LvlInfo:
		return "[INFO]"
	case LvlWarn:
		return "[WARN]"
	default:
		return "[ERROR]"
	}
}

func formatArgs(args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}

	var b strings.Builder
	start := 0

	if len(args)%2 == 1 {
		fmt.Fprint(&b, args[0])
		start = 1
	}

	for i := start; i < len(args); i += 2 {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		key := args[i]
		var val interface{} = "<nil>"
		if i+1 < len(args) {
			val = args[i+1]
		}
		fmt.Fprintf(&b, "%v=%v", key, val)
	}

	return b.String()
}

// SetLevel sets the level of the logger
func (s *StdLogger) SetLevel(level int) { s.level = level }

func (s *StdLogger) logf(lvl int, args ...interface{}) {
	if lvl < s.level {
		return
	}
	s.l.Printf("%s %s", levelStr(lvl), formatArgs(args...))
}

// Debug logs a debug level message
func (s *StdLogger) Debug(a ...interface{}) { s.logf(LvlDebug, a...) }

// Info logs an info level message
func (s *StdLogger) Info(a ...interface{}) { s.logf(LvlInfo, a...) }

// Warn logs a warning level message
func (s *StdLogger) Warn(a ...interface{}) { s.logf(LvlWarn, a...) }

// Error logs an error level message
func (s *StdLogger) Error(a ...interface{}) { s.logf(LvlError, a...) }

// Fatal logs a fatal level message
func (s *StdLogger) Fatal(a ...interface{}) {
	s.l.Fatal(formatArgs(a...))
}

// With returns a new Logger with the given values
func (s *StdLogger) With(vals ...interface{}) Logger {
	prefix := s.l.Prefix()

	var b strings.Builder
	b.WriteString(prefix)
	if len(vals) > 0 {
		b.WriteByte('[')

		for i := 0; i < len(vals); i += 2 {
			if i > 0 {
				b.WriteByte(' ')
			}

			key := fmt.Sprintf("%v", vals[i])
			var val interface{} = "<nil>"
			if i+1 < len(vals) {
				val = vals[i+1]
			}
			fmt.Fprintf(&b, "%s=%v", key, val)
		}
		b.WriteString("] ")
	}

	child := log.New(s.l.Writer(), b.String(), s.l.Flags())
	return &StdLogger{l: child, level: s.level}
}

// NewStd creates a new StdLogger that wraps a log.Logger.
func NewStd(l *log.Logger) *StdLogger {
	if l == nil {
		l = log.New(os.Stdout, "", log.LstdFlags)
	}
	return &StdLogger{l: l}
}
