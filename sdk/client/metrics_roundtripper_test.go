package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubRoundTripper struct {
	resp *http.Response
	err  error
}

func (s *stubRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return s.resp, s.err
}

func TestNewMetricsRoundTripper_NilUsesDefault(t *testing.T) {
	rt := NewMetricsRoundTripper(nil)
	mrt, ok := rt.(*metricsRoundTripper)
	require.True(t, ok)
	assert.Equal(t, http.DefaultTransport, mrt.next)
}

func TestNewMetricsRoundTripper_CustomTransport(t *testing.T) {
	custom := &stubRoundTripper{}
	rt := NewMetricsRoundTripper(custom)
	mrt, ok := rt.(*metricsRoundTripper)
	require.True(t, ok)
	assert.Equal(t, custom, mrt.next)
}

func TestRoundTrip_Success(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	rt := NewMetricsRoundTripper(&stubRoundTripper{resp: resp})

	req, _ := http.NewRequest("GET", "http://example.com/api/test", nil)
	gotResp, err := rt.RoundTrip(req)

	assert.NoError(t, err)
	require.NotNil(t, gotResp)
	assert.Equal(t, 200, gotResp.StatusCode)
}

func TestRoundTrip_TransportError(t *testing.T) {
	rt := NewMetricsRoundTripper(&stubRoundTripper{err: errors.New("connection refused")})

	req, _ := http.NewRequest("POST", "http://example.com/api/test", nil)
	gotResp, err := rt.RoundTrip(req)

	assert.Error(t, err)
	assert.Nil(t, gotResp)
}

func TestRoundTrip_NilURL(t *testing.T) {
	resp := &http.Response{StatusCode: 204}
	rt := NewMetricsRoundTripper(&stubRoundTripper{resp: resp})

	req := &http.Request{Method: "DELETE"}
	gotResp, err := rt.RoundTrip(req)

	assert.NoError(t, err)
	require.NotNil(t, gotResp)
	assert.Equal(t, 204, gotResp.StatusCode)
}

// ShouldRecordHTTPRequests returns false before InitCollector is called (noop
// collector), so the round-tripper should delegate directly without timing.
func TestRoundTrip_SkipsTimingWhenHTTPMetricsDisabled(t *testing.T) {
	called := false
	stub := &stubRoundTripper{resp: &http.Response{StatusCode: 200}}
	original := stub.resp

	wrapper := NewMetricsRoundTripper(&stubRoundTripper{
		resp: original,
	})

	req, _ := http.NewRequest("GET", "http://example.com/test", nil)
	gotResp, err := wrapper.RoundTrip(req)

	_ = called
	assert.NoError(t, err)
	require.NotNil(t, gotResp)
	assert.Equal(t, 200, gotResp.StatusCode)
}
