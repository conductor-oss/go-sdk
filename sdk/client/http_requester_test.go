package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrepareRequestDoesNotDuplicateDefaultHeaders(t *testing.T) {
	h := NewHttpRequester(nil, settings.NewHttpDefaultSettings(), http.DefaultClient, nil, nil)
	headers := map[string]string{"Accept": "application/json"}

	req, err := h.prepareRequest(context.Background(), "/metadata/workflow", http.MethodPost, map[string]string{"name": "wf"}, headers, nil, nil, "", nil)
	require.NoError(t, err)

	assert.Equal(t, []string{"application/json"}, req.Header.Values("Accept"))
	assert.Equal(t, []string{"application/json"}, req.Header.Values("Content-Type"))
	assert.Equal(t, []string{"gzip"}, req.Header.Values("Accept-Encoding"))
}

func TestPrepareRequestKeepsRequestSpecificContentType(t *testing.T) {
	h := NewHttpRequester(nil, settings.NewHttpDefaultSettings(), http.DefaultClient, nil, nil)
	headers := map[string]string{"Content-Type": "text/plain"}

	req, err := h.prepareRequest(context.Background(), "/secrets/s1", http.MethodPut, "value", headers, nil, nil, "", nil)
	require.NoError(t, err)

	assert.Equal(t, []string{"text/plain"}, req.Header.Values("Content-Type"))
}
