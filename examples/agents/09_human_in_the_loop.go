package agents

import (
	"context"
	"encoding/json"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// The Python tool returns the float 15000.00, which serializes as 15000.0;
// a Go float64 would serialize as 15000, and the recorder tells them apart.
func checkAccountBalance(_ context.Context, in accountIn) (map[string]any, error) {
	return map[string]any{"account_id": in.AccountID, "balance": json.Number("15000.0")}, nil
}

func transferFunds(_ context.Context, in transferIn) (map[string]any, error) {
	return map[string]any{"status": "completed", "from": in.FromAcct, "to": in.ToAcct, "amount": in.Amount}, nil
}

// HumanInTheLoop is the Python SDK's examples/agents/09_human_in_the_loop.py:
// a transfer tool that pauses the run until a person approves it.
func HumanInTheLoop(model string) *ai.Agent {
	return &ai.Agent{
		Name:  "banker",
		Model: model,
		Tools: ai.Tools(
			tool.Func("check_balance", checkAccountBalance, "Check the balance of an account."),
			tool.Func("transfer_funds", transferFunds,
				"Request a funds transfer; runtime pauses for human approval before execution.", tool.RequiresApproval()),
		),
		Instructions: "You are a banking assistant. Use check_balance for balance inquiries. When asked to transfer money, first check the balance, then call transfer_funds to request the transfer. The runtime will pause for human approval before the transfer executes.",
	}
}

func runHumanInTheLoop(ctx context.Context, rt *ai.Runtime, in io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := runWithApproval(ctx, rt, HumanInTheLoop(Model()),
		"Transfer $500 from ACC-789 to ACC-456. Check the balance first.", in, out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
