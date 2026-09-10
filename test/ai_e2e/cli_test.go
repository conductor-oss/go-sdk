//go:build e2e

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai_e2e

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Attaching CLIConfig gives the agent a derived "<agent>_run_command" tool.
// The golden fixture proves the config serializes; this asks whether a run
// using it works — the same question that found code execution hanging.
//
// The command lists a directory this test creates, so the assertion is on a
// file name only this run could have produced, not on model wording.
func TestCLICommand(t *testing.T) {
	rt := newRuntime(t)
	defer rt.Shutdown()

	// A fixed, relative directory rather than t.TempDir(): the path is part
	// of the prompt, and a replayed recording has to see the same text on
	// every machine. It resolves against the worker's working directory,
	// which is this package.
	dir := filepath.Join("testdata", "cli_listing")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	marker := "conductor_e2e_marker_" + strings.ReplaceAll(t.Name(), "/", "_")
	if err := os.WriteFile(filepath.Join(dir, marker), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	agent := &ai.Agent{
		Name:  "go_e2e_cli",
		Model: model(t),
		Instructions: "You can run shell commands with the run_command tool. Use it " +
			"rather than guessing, and report the tool's exact output.",
		CLI: &ai.CLIConfig{
			// A tight allow-list: the model needs only ls, and anything else
			// it tries should be refused rather than run.
			AllowedCommands: []string{"ls"},
			TimeoutSeconds:  30,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	res, err := rt.Run(ctx, agent,
		"List the files in the directory "+dir+" using run_command with command=\"ls\" "+
			"and args=[\""+dir+"\"]. Report the file names you see.")
	if err != nil {
		t.Fatalf("Run: %v", err)
	}

	t.Logf("status=%s output=%.200q", res.Status, res.Output)
	if res.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q", res.Status, ai.StatusCompleted)
	}
	// Only an actual ls of that directory could have surfaced this name.
	if !strings.Contains(res.Output, marker) {
		t.Errorf("output does not mention %s, so run_command did not list the "+
			"directory: %q", marker, res.Output)
	}
}
