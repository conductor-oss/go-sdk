package main

import (
	"context"
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

type WorkflowGovernor struct {
	workflowExecutor   *executor.WorkflowExecutor
	workflowName       string
	workflowsPerSecond int
}

func NewWorkflowGovernor(
	workflowExecutor *executor.WorkflowExecutor,
	workflowName string,
	workflowsPerSecond int,
) *WorkflowGovernor {
	return &WorkflowGovernor{
		workflowExecutor:   workflowExecutor,
		workflowName:       workflowName,
		workflowsPerSecond: workflowsPerSecond,
	}
}

func (g *WorkflowGovernor) Run(ctx context.Context) {
	fmt.Printf("WorkflowGovernor started: workflow=%s, rate=%d/sec\n",
		g.workflowName, g.workflowsPerSecond)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("WorkflowGovernor stopped")
			return
		case <-ticker.C:
			g.startBatch()
		}
	}
}

func (g *WorkflowGovernor) startBatch() {
	for i := 0; i < g.workflowsPerSecond; i++ {
		_, err := g.workflowExecutor.StartWorkflow(
			&model.StartWorkflowRequest{
				Name:    g.workflowName,
				Version: 1,
			},
		)
		if err != nil {
			fmt.Printf("Governor: error starting workflows: %v\n", err)
			return
		}
	}
	fmt.Printf("Governor: started %d workflow(s)\n", g.workflowsPerSecond)
}
