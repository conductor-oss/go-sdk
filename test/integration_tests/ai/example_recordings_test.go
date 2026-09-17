//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Support for the example tests (example_*_test.go).
//
// Those tests are the Go ports of the Python SDK's examples/agents, run as
// tests: the Python SDK's example runs are recorded by the server's LLM
// recorder and kept in the conductor repository under llm-recordings/, one
// directory per example. The server plays them back alongside this
// directory's own recordings; CONDUCTOR_RECORDINGS_DIR points at the
// conductor copy so the tests can read what the Python example printed.
// Each example test runs the example's agent with the mock model and asserts
// on the result the way the other tests here do; the recorded answer is what
// the run must print. Completing against Python's recording is what shows
// the two SDKs send the same requests.

// recordedAnswers returns the model answers recorded for one example, in
// file order, read from CONDUCTOR_RECORDINGS_DIR/<name>/. The last one is
// what the example printed.
func recordedAnswers(t *testing.T, name string) []string {
	t.Helper()
	root := os.Getenv("CONDUCTOR_RECORDINGS_DIR")
	if root == "" {
		t.Skip("CONDUCTOR_RECORDINGS_DIR is not set; point it at the conductor llm-recordings directory the playback server was started on")
	}
	files, err := filepath.Glob(filepath.Join(root, name, "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no recordings for %s under %s", name, root)
	}
	// Files are named "<n>_<uuid>.json" in call order; sort by n, not as
	// strings, or the tenth call sorts before the second.
	sort.Slice(files, func(i, j int) bool { return recordingSeq(files[i]) < recordingSeq(files[j]) })
	var answers []string
	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		var doc struct {
			Response struct {
				Results []struct {
					Output struct {
						Text string `json:"text"`
					} `json:"output"`
				} `json:"results"`
			} `json:"response"`
		}
		if err := json.Unmarshal(raw, &doc); err != nil {
			t.Fatalf("parse %s: %v", filepath.Base(f), err)
		}
		for _, r := range doc.Response.Results {
			answers = append(answers, r.Output.Text)
		}
	}
	return answers
}

// recordingSeq is the numeric prefix of a recording file name, or a large
// number for a file without one so it sorts last.
func recordingSeq(path string) int {
	base := filepath.Base(path)
	if i := strings.IndexByte(base, '_'); i > 0 {
		if n, err := strconv.Atoi(base[:i]); err == nil {
			return n
		}
	}
	return 1 << 30
}
