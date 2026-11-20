package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

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

// After a failover every in-flight request gets a 401 at once; together they must fetch one new token, not one each.
func TestCallAPI_ConcurrentRejectionsFetchOneNewToken(t *testing.T) {
	const n = 20
	var issued, rejected int32
	allRejected := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/token" {
			// Hold the fetch open so the other requests' Invalidate and RefreshToken calls overlap it.
			time.Sleep(50 * time.Millisecond)
			fmt.Fprintf(w, `{"token":"t%d"}`, atomic.AddInt32(&issued, 1))
			return
		}
		if r.Header.Get("X-Authorization") == "t1" {
			if atomic.AddInt32(&rejected, 1) == n {
				close(allRejected)
			}
			select {
			case <-allRejected:
			case <-time.After(5 * time.Second):
			}
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"error":"INVALID_TOKEN"}`)
			return
		}
		fmt.Fprint(w, `"ok"`)
	}))
	defer srv.Close()

	c := NewAPIClient(settings.NewAuthenticationSettings("key", "secret"), settings.NewHttpSettings(srv.URL))
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var out string
			resp, err := c.Post(context.Background(), "/thing", "payload", &out)
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, int32(n), rejected)
	assert.Equal(t, int32(2), issued, "t1 for the first attempts, t2 shared by every retry")
}
