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
	"context"
	"errors"
	"sync"
)

// Agent state is a map the server keeps for one execution and hands to every
// tool call in it, so one tool can leave data for a later tool without routing
// it through the model. The server sends it on the task as agentStateKey and
// reads back what changed under stateUpdatesKey; Join propagates those updates
// across parallel branches. It is the counterpart of ToolContext.state in the
// Python SDK and ToolContext.getState() in the Java SDK.
const (
	agentStateKey   = "_agent_state"
	stateUpdatesKey = "_state_updates"
)

// ErrNoTaskContext means a state call was made outside a tool handler, where
// there is no task to read from or write to.
var ErrNoTaskContext = errors.New("no task in context")

type stateContextKey struct{}

// taskState holds one task's view of the agent state: what arrived, and what
// the handler set. Tools may use goroutines, so it is guarded.
type taskState struct {
	mu      sync.Mutex
	initial map[string]any
	updates map[string]any
}

func (s *taskState) get(key string) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v, ok := s.updates[key]; ok {
		return v, true
	}
	v, ok := s.initial[key]
	return v, ok
}

func (s *taskState) all() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make(map[string]any, len(s.initial)+len(s.updates))
	for k, v := range s.initial {
		out[k] = v
	}
	for k, v := range s.updates {
		out[k] = v
	}
	return out
}

func (s *taskState) set(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.updates == nil {
		s.updates = map[string]any{}
	}
	s.updates[key] = value
}

// changed reports what the handler set, or nil if it set nothing. Only changes
// travel back, so a tool that reads state adds nothing to its output.
func (s *taskState) changed() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.updates) == 0 {
		return nil
	}
	out := make(map[string]any, len(s.updates))
	for k, v := range s.updates {
		out[k] = v
	}
	return out
}

func withTaskState(ctx context.Context, s *taskState) context.Context {
	return context.WithValue(ctx, stateContextKey{}, s)
}

func stateFromContext(ctx context.Context) (*taskState, bool) {
	s, ok := ctx.Value(stateContextKey{}).(*taskState)
	return s, ok
}

// StateValue reads one key of the agent state the server delivered with this
// task, including anything SetState recorded earlier in the same call.
//
//	tenant, ok := ai.StateValue(ctx, "tenant")
//
// The second result is false when the key is absent, and also when ctx is not a
// tool handler's context.
func StateValue(ctx context.Context, key string) (any, bool) {
	s, ok := stateFromContext(ctx)
	if !ok {
		return nil, false
	}
	return s.get(key)
}

// State returns a copy of the whole agent state for this task. Writing to the
// returned map changes nothing; use SetState.
func State(ctx context.Context) map[string]any {
	s, ok := stateFromContext(ctx)
	if !ok {
		return map[string]any{}
	}
	return s.all()
}

// SetState records a value to persist into the agent's execution state, where
// later tool calls in the same run can read it. The value must serialize to
// JSON, since it travels back on the task output.
//
//	if err := ai.SetState(ctx, "ticket_id", id); err != nil {
//	    return "", err
//	}
//
// It returns ErrNoTaskContext outside a tool handler rather than discarding the
// write, so a misplaced call is visible instead of silently doing nothing.
func SetState(ctx context.Context, key string, value any) error {
	s, ok := stateFromContext(ctx)
	if !ok {
		return ErrNoTaskContext
	}
	s.set(key, value)
	return nil
}
