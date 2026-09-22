package ai

import (
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// Input and output types for the golden tool fixtures. They mirror the Python
// fixture signatures in sdk/ai/testdata/agent_config/generate_fixtures.py, so the
// reflected schemas must match what Python's type hints produced.

type weatherIn struct {
	City string
	Days int `json:"days,omitempty"` // Python: days: int = 3 — has a default
}

type scalarKindsIn struct {
	S   string
	I   int
	F   float64
	B   bool
	Opt *string `json:"opt,omitempty"` // Python: Optional[str] = None
}

type containerKindsIn struct {
	Names []string
	Meta  map[string]int
}

type orderIn struct {
	OrderID string
}

type branchIn struct {
	Branch string
}

// mkTool builds a worker ToolDef with schemas reflected from the given input
// and output values. The tool package's generic constructors need a function
// literal per case; this keeps the fixture table readable.
func mkTool(name, description string, in, out any) ToolDef {
	return ToolDef{
		Name:         name,
		Description:  description,
		InputSchema:  schema.Of(reflect.TypeOf(in)),
		OutputSchema: schema.Of(reflect.TypeOf(out)),
		ToolType:     ToolTypeWorker,
	}
}
