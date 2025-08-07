package log

type nopLogger struct{}

func (nopLogger) Debug(...interface{})       {}
func (nopLogger) Info(...interface{})        {}
func (nopLogger) Warn(...interface{})        {}
func (nopLogger) Error(...interface{})       {}
func (nopLogger) Fatal(...interface{})       {}
func (nopLogger) With(...interface{}) Logger { return nopLogger{} }

// NewNop creates a new no-op logger implementation.
func NewNop() Logger {
	return nopLogger{}
}
