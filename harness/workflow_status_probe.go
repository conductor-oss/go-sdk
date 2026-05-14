package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
)

const maxTrackedIDs = 256

// WorkflowStatusProbe exercises UUID-bearing workflow lookup endpoints so
// http_api_client_request_seconds picks up entries with
// uri=/workflow/{workflowId} and uri=/workflow/{workflowId}/status.
//
// Default harness traffic only hits bounded, no-path-param URLs (poll/update),
// making the high-cardinality concern on the uri label invisible without this
// probe.
//
// Default off. Runs only when HARNESS_PROBE_RATE_PER_SEC > 0.
// Side-effect-free: only issues read calls (GetExecutionStatus, GetWorkflowState).
// Self-bounded: fixed-size FIFO of workflow IDs.
type WorkflowStatusProbe struct {
	workflowClient client.WorkflowClient
	callsPerSecond int

	mu        sync.Mutex
	recentIDs []string
	rng       *rand.Rand
}

func NewWorkflowStatusProbe(workflowClient client.WorkflowClient, callsPerSecond int) *WorkflowStatusProbe {
	return &WorkflowStatusProbe{
		workflowClient: workflowClient,
		callsPerSecond: callsPerSecond,
		rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Offer captures a workflow ID for later probing. Thread-safe.
func (p *WorkflowStatusProbe) Offer(workflowID string) {
	if workflowID == "" {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.recentIDs = append(p.recentIDs, workflowID)
	if len(p.recentIDs) > maxTrackedIDs {
		p.recentIDs = p.recentIDs[len(p.recentIDs)-maxTrackedIDs:]
	}
}

// Run starts the probe loop. Blocks until ctx is cancelled.
func (p *WorkflowStatusProbe) Run(ctx context.Context) {
	if p.callsPerSecond <= 0 {
		fmt.Println("WorkflowStatusProbe disabled (HARNESS_PROBE_RATE_PER_SEC<=0)")
		return
	}
	fmt.Printf("WorkflowStatusProbe started: rate=%d/sec, retainedIds<=%d\n",
		p.callsPerSecond, maxTrackedIDs)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("WorkflowStatusProbe stopped")
			return
		case <-ticker.C:
			p.tick(ctx)
		}
	}
}

func (p *WorkflowStatusProbe) tick(ctx context.Context) {
	p.mu.Lock()
	budget := p.callsPerSecond
	if budget > len(p.recentIDs) {
		budget = len(p.recentIDs)
	}
	if budget == 0 {
		p.mu.Unlock()
		return
	}

	ids := make([]string, budget)
	for i := range ids {
		ids[i] = p.recentIDs[p.rng.Intn(len(p.recentIDs))]
	}
	p.mu.Unlock()

	for _, id := range ids {
		if p.rng.Float64() < 0.5 {
			_, _, _ = p.workflowClient.GetExecutionStatus(ctx, id, nil)
		} else {
			_, _, _ = p.workflowClient.GetWorkflowState(ctx, id, false, false)
		}
	}
}
