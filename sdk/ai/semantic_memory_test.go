//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"reflect"
	"testing"
)

func TestSemanticMemory(t *testing.T) {
	m := &SemanticMemory{MaxResults: 2, SessionID: "s1"}
	ids := map[string]string{}
	for _, c := range []string{
		"User prefers concise answers",
		"Project uses Python 3.12 with FastAPI",
		"User's name is Alice",
	} {
		id, err := m.Add(c, map[string]any{"type": "fact"})
		if err != nil || len(id) != 16 {
			t.Fatalf("Add(%q) = %q, %v", c, id, err)
		}
		ids[c] = id
	}

	// Keyword overlap ranks the matching memory first and leaves out non-matches.
	got, err := m.Search("What does the user prefer?", 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 || got[0] != "User prefers concise answers" {
		t.Errorf("search = %q", got)
	}
	for _, g := range got {
		if g == "Project uses Python 3.12 with FastAPI" {
			t.Errorf("unrelated memory returned: %q", got)
		}
	}
	if got, _ := m.Search("quantum chromodynamics", 0); len(got) != 0 {
		t.Errorf("no-overlap query returned %q", got)
	}

	// MaxResults bounds the result; an explicit topK overrides it.
	if got, _ := m.Search("user python alice", 0); len(got) > 2 {
		t.Errorf("MaxResults not applied: %q", got)
	}
	if got, _ := m.Search("user python alice", 3); len(got) != 3 {
		t.Errorf("topK=3 returned %d", len(got))
	}

	// Session metadata is recorded.
	entries, _ := m.List()
	if len(entries) != 3 || entries[0].Metadata["session_id"] != "s1" || entries[0].Metadata["type"] != "fact" {
		t.Errorf("entries = %+v", entries)
	}

	// Context renders the Python SDK's format.
	ctx, _ := m.Context("what is the user's name")
	wantPrefix := "Relevant context from memory:\n  1. User's name is Alice"
	if !reflect.DeepEqual(ctx[:len(wantPrefix)], wantPrefix) {
		t.Errorf("context = %q", ctx)
	}
	if ctx, _ := m.Context("zzz"); ctx != "" {
		t.Errorf("context for no match = %q", ctx)
	}

	// Delete and Clear.
	if ok, _ := m.Delete(ids["User's name is Alice"]); !ok {
		t.Error("delete of an existing id reported false")
	}
	if ok, _ := m.Delete("missing"); ok {
		t.Error("delete of a missing id reported true")
	}
	if entries, _ := m.List(); len(entries) != 2 {
		t.Errorf("after delete: %d entries", len(entries))
	}
	_ = m.Clear()
	if entries, _ := m.List(); len(entries) != 0 {
		t.Errorf("after clear: %d entries", len(entries))
	}
}
