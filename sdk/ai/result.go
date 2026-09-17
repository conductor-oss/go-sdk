//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import "strings"

// Status is the state of an agent execution.
type Status string

const (
	StatusRunning    Status = "RUNNING"
	StatusCompleted  Status = "COMPLETED"
	StatusFailed     Status = "FAILED"
	StatusTerminated Status = "TERMINATED"
	StatusTimedOut   Status = "TIMED_OUT"
	StatusPaused     Status = "PAUSED"
	// StatusWaiting means the run is blocked on a human.
	StatusWaiting Status = "WAITING"
)

// Terminal reports whether no further progress will happen without intervention.
func (s Status) Terminal() bool {
	switch s {
	case StatusCompleted, StatusFailed, StatusTerminated, StatusTimedOut:
		return true
	}
	return false
}

// TokenUsage is what the run cost, when the server reports it.
type TokenUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// AgentResult is the outcome of an execution.
type AgentResult struct {
	ExecutionID string
	Status      Status
	// Output is the agent's final answer.
	Output string
	// FinishReason is why the loop stopped, e.g. "STOP" or "MAX_TURNS".
	FinishReason string
	// Error is the failure reason when Status is StatusFailed.
	Error string
	// TokenUsage is zero when the server does not report it.
	TokenUsage TokenUsage
	// Raw is the untouched status document, for fields this struct does not
	// model yet. The agent surface is still growing; this keeps callers from
	// being blocked on it.
	Raw map[string]any
}

// resultFrom reads a status document into an AgentResult.
//
// The shape is taken from a live server, not guessed:
//
//	{"executionId": "...", "status": "COMPLETED",
//	 "output": {"result": "...", "finishReason": "STOP", "rejectionReason": null}}
//
// The answer is nested under output.result, not at the top level. Fields are
// still read leniently — the payload is shared with three other SDKs and has
// accreted aliases — so an unknown field leaves a zero value rather than
// failing, and Raw keeps everything for callers that need more.
func resultFrom(executionID string, status map[string]any) *AgentResult {
	res := &AgentResult{ExecutionID: executionID, Raw: status}

	if s := firstString(status, "status", "state"); s != "" {
		res.Status = Status(strings.ToUpper(s))
	}

	if out, ok := status["output"].(map[string]any); ok {
		res.Output = firstString(out, "result", "output", "finalOutput")
		res.FinishReason = firstString(out, "finishReason", "finish_reason")
		res.Error = firstString(out, "rejectionReason", "error")
	} else {
		// Some responses carry a bare string output.
		res.Output = firstString(status, "output", "result", "finalOutput")
	}
	if res.Error == "" {
		res.Error = firstString(status, "error", "reasonForIncompletion", "failureReason")
	}

	usage, ok := status["tokenUsage"].(map[string]any)
	if !ok {
		if out, o := status["output"].(map[string]any); o {
			usage, ok = out["tokenUsage"].(map[string]any)
		}
	}
	if ok {
		res.TokenUsage = TokenUsage{
			PromptTokens:     firstInt(usage, "promptTokens", "prompt_tokens"),
			CompletionTokens: firstInt(usage, "completionTokens", "completion_tokens"),
			TotalTokens:      firstInt(usage, "totalTokens", "total_tokens"),
		}
	}
	return res
}

func firstString(m map[string]any, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k].(string); ok && v != "" {
			return v
		}
	}
	return ""
}

func firstInt(m map[string]any, keys ...string) int {
	for _, k := range keys {
		switch v := m[k].(type) {
		case float64: // JSON numbers decode as float64
			return int(v)
		case int:
			return v
		}
	}
	return 0
}
