package ai

import (
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// Input and output types for the golden tool fixtures. They mirror the Python
// fixture signatures in sdk/ai/testdata/agent_config/generate_fixtures.py, so the
// reflected schemas must match what Python's type hints produced.

type weatherIn struct {
	City string `json:"city"`
	Days int    `json:"days,omitempty"` // Python: days: int = 3 — has a default
}

type scalarKindsIn struct {
	S   string  `json:"s"`
	I   int     `json:"i"`
	F   float64 `json:"f"`
	B   bool    `json:"b"`
	Opt *string `json:"opt,omitempty"` // Python: Optional[str] = None
}

type containerKindsIn struct {
	Names []string       `json:"names"`
	Meta  map[string]int `json:"meta"`
}

type orderIn struct {
	OrderID string `json:"order_id"`
}

type branchIn struct {
	Branch string `json:"branch"`
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
