package log

import (
	"bytes"
	stdlog "log"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyLogger struct{ called atomic.Bool }

func (d *dummyLogger) Debug(a ...interface{})           { d.called.Store(true) }
func (d *dummyLogger) Info(a ...interface{})            { d.called.Store(true) }
func (d *dummyLogger) Warn(a ...interface{})            { d.called.Store(true) }
func (d *dummyLogger) Error(a ...interface{})           { d.called.Store(true) }
func (d *dummyLogger) Fatal(a ...interface{})           { d.called.Store(true) }
func (d *dummyLogger) With(value ...interface{}) Logger { return d }

func TestSetLogger_DefaultFallback(t *testing.T) {
	defer SetLogger(nil)
	d := &dummyLogger{}
	SetLogger(d)
	Info("hi")
	assert.True(t, d.called.Load(), "custom logger must be used after SetLogger")

	SetLogger(nil) // reset to default std logger
	d.called.Store(false)
	Info("again")

	assert.False(t, d.called.Load(), "dummy logger must NOT be called after SetLogger(nil)")
}

func TestSetLogger_WithCustomLogger(t *testing.T) {
	// Save original to restore later
	originalBuf := &bytes.Buffer{}
	originalLogger := NewStd(stdlog.New(originalBuf, "", 0))
	defer SetLogger(originalLogger)

	buf := &bytes.Buffer{}
	customLogger := NewStd(stdlog.New(buf, "", 0))
	SetLogger(customLogger)

	Info("test message")

	assert.Contains(t, buf.String(), "test message", "SetLogger should use the provided custom logger")
}
