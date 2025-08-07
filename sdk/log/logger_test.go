package log

import (
	"bytes"
	stdlog "log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyLogger struct {
	mu     sync.Mutex
	called bool
}

func (d *dummyLogger) setCalled(val bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.called = val
}

func (d *dummyLogger) getCalled() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.called
}

func (d *dummyLogger) Debug(a ...interface{})           { d.setCalled(true) }
func (d *dummyLogger) Info(a ...interface{})            { d.setCalled(true) }
func (d *dummyLogger) Warn(a ...interface{})            { d.setCalled(true) }
func (d *dummyLogger) Error(a ...interface{})           { d.setCalled(true) }
func (d *dummyLogger) Fatal(a ...interface{})           { d.setCalled(true) }
func (d *dummyLogger) With(value ...interface{}) Logger { return d }

func TestSetLogger_DefaultFallback(t *testing.T) {
	defer SetLogger(nil)
	d := &dummyLogger{}
	SetLogger(d)
	Info("hi")
	assert.True(t, d.getCalled(), "custom logger must be used after SetLogger")

	SetLogger(nil) // reset to default std logger
	d.setCalled(false)
	Info("again")

	assert.False(t, d.getCalled(), "dummy logger must NOT be called after SetLogger(nil)")
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
