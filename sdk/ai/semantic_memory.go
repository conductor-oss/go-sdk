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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// Semantic memory: the counterpart of the Python SDK's SemanticMemory,
// MemoryStore and InMemoryStore. Neither SDK wires it into the agent runtime
// yet, so use it from application code, for example with SemanticMemory.Context.

// MemoryEntry is one remembered item.
type MemoryEntry struct {
	// ID identifies the entry; the store assigns one when it is empty.
	ID string
	// Content is the remembered text.
	Content string
	// Metadata is free-form: type, source, importance, session.
	Metadata map[string]any
	// Embedding is an optional vector for stores that search by embedding.
	Embedding []float64
	// CreatedAt is when the entry was stored.
	CreatedAt time.Time
}

// A MemoryStore keeps entries and finds the ones relevant to a query; implement
// it to back memory with a vector database. InMemoryStore is built in.
type MemoryStore interface {
	// Add stores an entry and returns its ID.
	Add(entry MemoryEntry) (string, error)
	// Search returns up to topK entries relevant to the query, most relevant first.
	Search(query string, topK int) ([]MemoryEntry, error)
	// Delete removes an entry, reporting whether it existed.
	Delete(id string) (bool, error)
	// Clear removes every entry.
	Clear() error
	// List returns every entry.
	List() ([]MemoryEntry, error)
}

// InMemoryStore keeps entries in memory, non-persistently, ranked by Jaccard
// similarity between the words of the query and of each entry. It is a fallback
// for development and tests; production memory belongs in a real store.
type InMemoryStore struct {
	mu      sync.Mutex
	entries map[string]MemoryEntry
	order   []string // insertion order, so List and ties are stable
}

func (s *InMemoryStore) Add(entry MemoryEntry) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.entries == nil {
		s.entries = map[string]MemoryEntry{}
	}
	if entry.CreatedAt.IsZero() {
		entry.CreatedAt = time.Now()
	}
	if entry.ID == "" {
		sum := sha256.Sum256([]byte(fmt.Sprintf("%s%d", entry.Content, entry.CreatedAt.UnixNano())))
		entry.ID = hex.EncodeToString(sum[:])[:16]
	}
	if _, exists := s.entries[entry.ID]; !exists {
		s.order = append(s.order, entry.ID)
	}
	s.entries[entry.ID] = entry
	return entry.ID, nil
}

func (s *InMemoryStore) Search(query string, topK int) ([]MemoryEntry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	queryWords := wordSet(query)
	type scored struct {
		score float64
		entry MemoryEntry
	}
	ranked := make([]scored, 0, len(s.order))
	for _, id := range s.order {
		entry := s.entries[id]
		ranked = append(ranked, scored{jaccard(queryWords, wordSet(entry.Content)), entry})
	}
	sort.SliceStable(ranked, func(i, j int) bool { return ranked[i].score > ranked[j].score })
	out := make([]MemoryEntry, 0, min(topK, len(ranked)))
	for _, r := range ranked {
		if r.score <= 0 || len(out) >= topK {
			break
		}
		out = append(out, r.entry)
	}
	return out, nil
}

func (s *InMemoryStore) Delete(id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.entries[id]; !ok {
		return false, nil
	}
	delete(s.entries, id)
	for i, k := range s.order {
		if k == id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}
	return true, nil
}

func (s *InMemoryStore) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries, s.order = map[string]MemoryEntry{}, nil
	return nil
}

func (s *InMemoryStore) List() ([]MemoryEntry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]MemoryEntry, 0, len(s.order))
	for _, id := range s.order {
		out = append(out, s.entries[id])
	}
	return out, nil
}

func wordSet(text string) map[string]struct{} {
	set := map[string]struct{}{}
	for _, w := range strings.Fields(strings.ToLower(text)) {
		set[w] = struct{}{}
	}
	return set
}

func jaccard(a, b map[string]struct{}) float64 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	common := 0
	for w := range a {
		if _, ok := b[w]; ok {
			common++
		}
	}
	return float64(common) / float64(len(a)+len(b)-common)
}

// SemanticMemory adds and retrieves memories through a MemoryStore and renders
// the relevant ones as context for a prompt.
type SemanticMemory struct {
	// Store holds the entries; nil means a fresh InMemoryStore.
	Store MemoryStore
	// MaxResults bounds a search; zero means 5.
	MaxResults int
	// SessionID, when set, is recorded in every entry's metadata.
	SessionID string

	once  sync.Once
	store MemoryStore
}

func (m *SemanticMemory) backing() MemoryStore {
	m.once.Do(func() {
		m.store = m.Store
		if m.store == nil {
			m.store = &InMemoryStore{}
		}
	})
	return m.store
}

func (m *SemanticMemory) limit(topK int) int {
	switch {
	case topK > 0:
		return topK
	case m.MaxResults > 0:
		return m.MaxResults
	}
	return 5
}

// Add remembers content with optional metadata and returns the entry's ID.
func (m *SemanticMemory) Add(content string, metadata map[string]any) (string, error) {
	meta := map[string]any{}
	for k, v := range metadata {
		meta[k] = v
	}
	if m.SessionID != "" {
		meta["session_id"] = m.SessionID
	}
	return m.backing().Add(MemoryEntry{Content: content, Metadata: meta})
}

// Search returns the content of the memories relevant to the query, most
// relevant first. topK of zero means MaxResults.
func (m *SemanticMemory) Search(query string, topK int) ([]string, error) {
	entries, err := m.SearchEntries(query, topK)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(entries))
	for _, e := range entries {
		out = append(out, e.Content)
	}
	return out, nil
}

// SearchEntries is Search with the full entries.
func (m *SemanticMemory) SearchEntries(query string, topK int) ([]MemoryEntry, error) {
	return m.backing().Search(query, m.limit(topK))
}

// Delete removes one memory, reporting whether it existed.
func (m *SemanticMemory) Delete(id string) (bool, error) { return m.backing().Delete(id) }

// Clear forgets everything.
func (m *SemanticMemory) Clear() error { return m.backing().Clear() }

// List returns every memory.
func (m *SemanticMemory) List() ([]MemoryEntry, error) { return m.backing().List() }

// Context renders the memories relevant to a query, in the Python SDK's format,
// as a block to prepend to an agent's instructions; "" when nothing is relevant.
func (m *SemanticMemory) Context(query string) (string, error) {
	memories, err := m.Search(query, 0)
	if err != nil || len(memories) == 0 {
		return "", err
	}
	lines := []string{"Relevant context from memory:"}
	for i, mem := range memories {
		lines = append(lines, fmt.Sprintf("  %d. %s", i+1, mem))
	}
	return strings.Join(lines, "\n"), nil
}
