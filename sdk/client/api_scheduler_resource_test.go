// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

// PauseSchedule and ResumeSchedule have to work against three server generations
// that accept different verbs. No single live server exercises all of them, so they
// are modelled here with stub handlers that record the verbs they receive — letting
// the tests assert not only that the call succeeds but that the client prefers PUT
// and falls back only when refused.

type verbSpy struct {
	methods []string
	paths   []string
}

// handler answers with the status configured for the received method, or 405 when
// the method is not in the map.
func (v *verbSpy) handler(allow map[string]int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v.methods = append(v.methods, r.Method)
		v.paths = append(v.paths, r.URL.EscapedPath())
		if code, ok := allow[r.Method]; ok {
			w.WriteHeader(code)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func newSchedulerService(t *testing.T, h http.Handler) (*SchedulerResourceApiService, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	api := NewAPIClient(
		settings.NewAuthenticationSettings("", ""),
		settings.NewHttpSettings(srv.URL+"/api"),
	)
	return &SchedulerResourceApiService{APIClient: api}, srv
}

func assertMethods(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("methods = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("methods = %v, want %v", got, want)
		}
	}
}

func TestPauseScheduleVerbNegotiation(t *testing.T) {
	tests := []struct {
		name        string
		allow       map[string]int
		wantMethods []string
		wantErr     bool
	}{
		{
			// Upstream OSS Conductor declares these endpoints @PutMapping.
			name:        "server accepts put only",
			allow:       map[string]int{http.MethodPut: http.StatusOK},
			wantMethods: []string{http.MethodPut},
		},
		{
			// Orkes Conductor from 2026-07-14 accepts both; PUT must be preferred so
			// callers converge on it and the fallback can eventually be removed.
			name: "server accepts both, put preferred",
			allow: map[string]int{
				http.MethodPut: http.StatusOK,
				http.MethodGet: http.StatusOK,
			},
			wantMethods: []string{http.MethodPut},
		},
		{
			// Orkes Conductor before 2026-07-14 accepts GET only.
			name:        "legacy server falls back to get",
			allow:       map[string]int{http.MethodGet: http.StatusOK},
			wantMethods: []string{http.MethodPut, http.MethodGet},
		},
		{
			// Neither verb accepted: report the failure rather than claim success.
			name:        "neither verb accepted returns error",
			allow:       map[string]int{},
			wantMethods: []string{http.MethodPut, http.MethodGet},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spy := &verbSpy{}
			svc, _ := newSchedulerService(t, spy.handler(tt.allow))

			_, _, err := svc.PauseSchedule(context.Background(), "probe")
			if tt.wantErr && err == nil {
				t.Fatal("expected an error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertMethods(t, spy.methods, tt.wantMethods)
		})
	}
}

func TestResumeScheduleVerbNegotiation(t *testing.T) {
	// Resume shares the helper with pause; this guards against only one of the two
	// being wired up.
	spy := &verbSpy{}
	svc, _ := newSchedulerService(t, spy.handler(map[string]int{http.MethodGet: http.StatusOK}))

	if _, _, err := svc.ResumeSchedule(context.Background(), "probe"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertMethods(t, spy.methods, []string{http.MethodPut, http.MethodGet})
	for _, p := range spy.paths {
		if p != "/api/scheduler/schedules/probe/resume" {
			t.Fatalf("path = %q, want the resume endpoint", p)
		}
	}
}

// A 5xx must not trigger the fallback: retrying a server fault with a different verb
// would hide it behind a misleading "method not supported".
func TestPauseScheduleDoesNotFallBackOnServerError(t *testing.T) {
	spy := &verbSpy{}
	svc, _ := newSchedulerService(t, spy.handler(map[string]int{
		http.MethodPut: http.StatusInternalServerError,
		http.MethodGet: http.StatusOK,
	}))

	if _, _, err := svc.PauseSchedule(context.Background(), "probe"); err == nil {
		t.Fatal("expected the 500 to surface, got nil")
	}
	assertMethods(t, spy.methods, []string{http.MethodPut})
}
