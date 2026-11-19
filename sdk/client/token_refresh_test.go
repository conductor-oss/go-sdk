package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// After a failover the server rejects the cached token; the client must get a new one and retry once.
func TestCallAPI_RefreshesTokenAndRetriesOn401(t *testing.T) {
	var issued, calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/token" {
			fmt.Fprintf(w, `{"token":"t%d"}`, atomic.AddInt32(&issued, 1))
			return
		}
		atomic.AddInt32(&calls, 1)
		body, _ := io.ReadAll(r.Body)
		if r.Header.Get("X-Authorization") != "t2" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"INVALID_TOKEN"}`)
			return
		}
		assert.Equal(t, "payload", string(body))
		fmt.Fprint(w, `"ok"`)
	}))
	defer srv.Close()

	c := NewAPIClient(settings.NewAuthenticationSettings("key", "secret"), settings.NewHttpSettings(srv.URL))
	var out string
	resp, err := c.Post(context.Background(), "/thing", "payload", &out)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "ok", out)
	assert.Equal(t, int32(2), issued)
	assert.Equal(t, int32(2), calls)
}
