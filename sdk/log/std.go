package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type stdLogger struct {
	l     *log.Logger
	level int
}

const (
	lvlDebug = iota
	lvlInfo
	lvlWarn
	lvlError
)

func levelStr(l int) string {
	switch l {
	case lvlDebug:
		return "[DEBUG]"
	case lvlInfo:
		return "[INFO]"
	case lvlWarn:
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

func (s *stdLogger) SetLevel(level int) { s.level = level }

func (s *stdLogger) logf(lvl int, args ...interface{}) {
	if lvl < s.level {
		return
	}
	s.l.Printf("%s %s", levelStr(lvl), formatArgs(args...))
}

func (s *stdLogger) Debug(a ...interface{}) { s.logf(lvlDebug, a...) }
func (s *stdLogger) Info(a ...interface{})  { s.logf(lvlInfo, a...) }
func (s *stdLogger) Warn(a ...interface{})  { s.logf(lvlWarn, a...) }
func (s *stdLogger) Error(a ...interface{}) { s.logf(lvlError, a...) }

func (s *stdLogger) Fatal(a ...interface{}) {
	s.l.Fatal(formatArgs(a...))
}

func (s *stdLogger) With(vals ...interface{}) Logger {
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
	return &stdLogger{l: child, level: s.level}
}

// NewStd creates a new Logger that wraps a log.Logger.
func NewStd(l *log.Logger) Logger {
	if l == nil {
		l = log.New(os.Stdout, "", log.LstdFlags)
	}
	return &stdLogger{l: l}
}
