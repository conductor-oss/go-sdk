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
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// The Python SDK's e2e/test_suite25_media_input.py as Go tests, one per Python
// test and under the same names.
//
// An image given to a run reaches the model. The proof is a word that exists
// nowhere but in the image's pixels, so the model can only report it by having
// been shown the picture; the second test runs the same prompt with no image
// and requires that the word does not appear, which is what stops the first
// from passing on a lucky guess.
//
// Live only: the server reads the file itself, and a vision model's wording
// varies, so there is nothing stable to replay.

const (
	s25Secret       = "MELON7391"
	s25Instructions = "You are an OCR assistant. Read text from images precisely."
	s25Prompt       = "Transcribe the exact text shown in the image. Reply with only that text and nothing else."
)

// s25Normalize keeps only letters and digits, upper-cased, so the check does
// not turn on the model's punctuation or spacing.
func s25Normalize(s string) string {
	var b strings.Builder
	for _, r := range strings.ToUpper(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// s25Image writes the test image where the server is allowed to read it and
// returns that path. The server opens the file itself, so it has to be on the
// server's own disk, under the directory it permits.
func s25Image(t *testing.T) string {
	t.Helper()
	source := filepath.Join("testdata", "melon7391.png")
	png, err := os.ReadFile(source)
	if err != nil {
		t.Skipf("no test image at %s: %v", source, err)
	}
	dir := os.Getenv("CONDUCTOR_AGENT_MEDIA_DIR")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			t.Skipf("cannot find the server's media directory: %v", err)
		}
		dir = filepath.Join(home, "worker-payload")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Skipf("cannot create the server media directory %s: %v", dir, err)
	}
	path := filepath.Join(dir, "e2e_s25_media_input.png")
	if err := os.WriteFile(path, png, 0o644); err != nil {
		t.Skipf("cannot write the image to %s: %v", path, err)
	}
	t.Cleanup(func() { os.Remove(path) })
	return path
}

func s25Agent(t *testing.T, name string) *ai.Agent {
	t.Helper()
	return &ai.Agent{Name: name, Model: model(t), Instructions: s25Instructions}
}

// The model reads the word out of the picture it was given.
func TestVisionReadsTextFromImage(t *testing.T) {
	skipInPlayback(t, "a vision model's wording varies and the server reads the image from disk")
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set — provider unavailable")
	}
	rt := newRuntime(t)
	path := s25Image(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res, err := rt.Run(ctx, s25Agent(t, "e2e_s25_vision"), s25Prompt, ai.WithMedia(path))
	if err != nil || res.Status != ai.StatusCompleted {
		t.Fatalf("run did not complete: status=%v err=%v", statusOf(res), err)
	}
	if !strings.Contains(s25Normalize(res.Output), s25Secret) {
		t.Errorf("the model did not read %s out of the image\n--- answer ---\n%.400s", s25Secret, res.Output)
	}
}

// Without the image the word does not appear, so the check above is really
// reading the picture rather than guessing.
func TestWithoutMediaTokenIsAbsent(t *testing.T) {
	skipInPlayback(t, "the counterfactual of the vision test, which runs live only")
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY not set — provider unavailable")
	}
	rt := newRuntime(t)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	res, err := rt.Run(ctx, s25Agent(t, "e2e_s25_no_media"), s25Prompt)
	if err != nil || res.Status != ai.StatusCompleted {
		t.Fatalf("run did not complete: status=%v err=%v", statusOf(res), err)
	}
	if strings.Contains(s25Normalize(res.Output), s25Secret) {
		t.Errorf("%s appeared with no image sent, so the vision test proves nothing\n--- answer ---\n%.400s",
			s25Secret, res.Output)
	}
}
