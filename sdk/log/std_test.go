package log

import (
	"bytes"
	stdlog "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func capture() (*bytes.Buffer, *stdlog.Logger) {
	buf := &bytes.Buffer{}
	return buf, stdlog.New(buf, "", 0)
}

func TestStd_NewStd(t *testing.T) {
	l := NewStd(nil)
	assert.NotNil(t, l, "NewStd(nil) returns working logger")

	buf, goStd := capture()
	NewStd(goStd).Info("ping")
	assert.Contains(t, buf.String(), "ping")
}

func TestStd_LevelsAndFilter(t *testing.T) {
	cases := []struct {
		name        string
		logFn       func(Logger)
		setLevel    int
		expectPrint bool
		wantPrefix  string
	}{
		{"Debug_pass", func(l Logger) { l.Debug("d") }, lvlDebug, true, "[DEBUG]"},
		{"Debug_filtered", func(l Logger) { l.Debug("d") }, lvlInfo, false, ""},
		{"Info_pass", func(l Logger) { l.Info("i") }, lvlInfo, true, "[INFO]"},
		{"Warn_pass", func(l Logger) { l.Warn("w") }, lvlInfo, true, "[WARN]"},
		{"Error_pass", func(l Logger) { l.Error("e") }, lvlInfo, true, "[ERROR]"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			buf, goStd := capture()
			l := NewStd(goStd).(*stdLogger)
			l.SetLevel(tc.setLevel)

			tc.logFn(l)

			out := buf.String()
			if tc.expectPrint {
				assert.Contains(t, out, tc.wantPrefix)
			} else {
				assert.Empty(t, out)
			}
		})
	}
}

func TestStd_With(t *testing.T) {
	buf := &bytes.Buffer{}
	root := NewStd(stdlog.New(buf, "", 0))

	t.Run("single With adds kv pair", func(t *testing.T) {
		buf.Reset()
		root.With("req", 99).Info("ping")

		out := buf.String()
		assert.Contains(t, out, "[req=99]", "prefix must contain key=value in brackets")
		assert.Contains(t, out, "ping")
	})

	t.Run("chained With accumulates multiple brackets", func(t *testing.T) {
		buf.Reset()
		root.With("req", 42).With("user", "bob").Info("processing")

		out := buf.String()
		assert.Contains(t, out, "[req=42]", "first kv bracket missing")
		assert.Contains(t, out, "[user=bob]", "second kv bracket missing")
		assert.Contains(t, out, "processing")
	})
}

func TestStd_FormatArgs(t *testing.T) {
	buf, goStd := capture()
	l := NewStd(goStd)

	tests := []struct {
		name           string
		logFn          func()
		wantContains   []string
		wantNotContain []string
	}{
		{
			name: "message plus kv",
			logFn: func() {
				l.Info("process started", "id", 1, "state", "init")
			},
			wantContains:   []string{"[INFO]", "process started", "id=1", "state=init"},
			wantNotContain: []string{"process startedid="},
		},
		{
			name: "only kv pairs",
			logFn: func() {
				l.Info("id", 2, "state", "running")
			},
			wantContains:   []string{"[INFO]", "id=2", "state=running"},
			wantNotContain: []string{"runningid"},
		},
		{
			name: "only message",
			logFn: func() {
				l.Info("message")
			},
			wantContains: []string{"[INFO]", "message"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFn()
			out := buf.String()

			for _, substr := range tc.wantContains {
				assert.Contains(t, out, substr)
			}
			for _, substr := range tc.wantNotContain {
				assert.NotContains(t, out, substr)
			}
		})
	}
}
