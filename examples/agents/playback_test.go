package agents

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// TestExamplesPlayback runs every example against a Conductor server in
// playback mode, the mock model answering from the recordings shared in the
// conductor repository. It checks that each run reached a terminal state;
// whether the outcome was right is the shared check-playback action's job.
// scripts/run-agents-playback.sh sets it up and starts the external workers.
func TestExamplesPlayback(t *testing.T) {
	if os.Getenv("CONDUCTOR_AGENTS_PLAYBACK") != "true" {
		t.Skip("set CONDUCTOR_AGENTS_PLAYBACK=true and start the playback server")
	}
	for _, ex := range Catalog {
		t.Run(ex.Name, func(t *testing.T) {
			rt := ai.NewRuntime(ai.Config{})
			defer rt.Shutdown()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()
			var out bytes.Buffer
			results, err := ex.Run(ctx, rt, strings.NewReader("y\ny\n"), &out)
			t.Log(out.String())
			if err != nil {
				t.Fatal(err)
			}
			if len(results) == 0 {
				t.Fatal("no executions")
			}
			for _, res := range results {
				if !res.Status.Terminal() {
					t.Errorf("%s ended as %s", res.ExecutionID, res.Status)
				}
			}
		})
	}
}
