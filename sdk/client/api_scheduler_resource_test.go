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

// Stub servers stand in for the server generations that accept different verbs,
// recording what they receive so the tests can assert PUT is preferred.

type verbSpy struct {
	methods []string
	paths   []string
}

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
			name:        "server accepts put only",
			allow:       map[string]int{http.MethodPut: http.StatusOK},
			wantMethods: []string{http.MethodPut},
		},
		{
			// PUT must win so callers converge on it.
			name: "server accepts both, put preferred",
			allow: map[string]int{
				http.MethodPut: http.StatusOK,
				http.MethodGet: http.StatusOK,
			},
			wantMethods: []string{http.MethodPut},
		},
		{
			name:        "legacy server falls back to get",
			allow:       map[string]int{http.MethodGet: http.StatusOK},
			wantMethods: []string{http.MethodPut, http.MethodGet},
		},
		{
			name:        "neither verb accepted returns error",
			allow:       map[string]int{},
			wantMethods: []string{http.MethodPut, http.MethodGet},
			wantErr:     true,
		},
		{
			// 404 means the schedule or scheduler module is absent, not a bad verb.
			name:        "404 does not trigger the fallback",
			allow:       map[string]int{http.MethodPut: http.StatusNotFound},
			wantMethods: []string{http.MethodPut},
			wantErr:     true,
		},
		{
			name:        "401 does not trigger the fallback",
			allow:       map[string]int{http.MethodPut: http.StatusUnauthorized},
			wantMethods: []string{http.MethodPut},
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
	// Guards against only one of the two methods being wired up.
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

// A 5xx must not trigger the fallback; it would hide the fault behind a verb error.
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
