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
	"fmt"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
	taskmodel "github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
)

type crmCustomerIn struct {
	CustomerID string `json:"customer_id"`
}

type inventoryIn struct {
	ProductID string `json:"product_id"`
	Warehouse string `json:"warehouse,omitempty"`
}

type orderActionIn struct {
	OrderID string `json:"order_id"`
	Action  string `json:"action"`
}

type formatDataIn struct {
	Data map[string]any `json:"data"`
}

// External worker tools — the Python SDK's examples/agents/33_external_workers.py
// as a test.
//
// An agent mixes one local tool with three external ones: tools declared
// with a schema and no handler, whose tasks Conductor dispatches to whatever
// worker polls for them. Here those workers are plain Conductor workers
// started outside the agent runtime, standing in for the other service, and
// they return what the Python recording shows the workers returned. This is
// that flow, copied, with the print replaced by validation: the run
// completes, each external worker was dispatched with the customer, product
// and order from the request, and the local tool ran too.
//
// The recording cannot be replayed: format_response renders a dict in its
// key order, and a Go map does not keep the order the task arrived in, so
// the last request never matches; see README, "What has to match".
func TestExample33ExternalWorkers(t *testing.T) {
	runtime := newRuntime(t)
	if model(t) == mockModel {
		t.Skip("format_response renders a map, and Go cannot keep the key order the recorded result has")
	}

	// The "other service": one classic worker per external task name.
	external := worker.NewTaskRunnerWithApiClient(newAPIClient(t))
	var customerCalls, inventoryCalls, orderCalls, formatCalls atomic.Int32
	var customerID, productID, orderID, action atomic.Value
	start := func(name string, fn func(in map[string]any) map[string]any) {
		err := external.StartWorker(name, func(task *taskmodel.Task) (any, error) {
			return fn(task.InputData), nil
		}, 1, 100*time.Millisecond)
		if err != nil {
			t.Fatalf("start external worker %s: %v", name, err)
		}
		t.Cleanup(func() { external.Shutdown(name) })
	}
	start("get_customer", func(in map[string]any) map[string]any {
		customerCalls.Add(1)
		customerID.Store(fmt.Sprint(in["customer_id"]))
		return map[string]any{
			"customer_id": "C-1234",
			"name":        "Example Customer",
			"orders": []any{map[string]any{
				"order_id": "ORD-5678", "customer_id": "C-1234", "product_id": "PROD-001",
				"warehouse": "default", "status": "pending",
			}},
		}
	})
	start("check_inventory", func(in map[string]any) map[string]any {
		inventoryCalls.Add(1)
		productID.Store(fmt.Sprint(in["product_id"]))
		return map[string]any{"product_id": "PROD-001", "warehouse": "default", "in_stock": true, "quantity": 12}
	})
	start("process_order", func(in map[string]any) map[string]any {
		orderCalls.Add(1)
		orderID.Store(fmt.Sprint(in["order_id"]))
		action.Store(fmt.Sprint(in["action"]))
		return map[string]any{
			"order_id": "ORD-5678", "customer_id": "C-1234", "product_id": "PROD-001",
			"warehouse": "default", "status": "cancelled",
		}
	})

	// The local tool, as in the example.
	formatResponse := func(ctx context.Context, in formatDataIn) (string, error) {
		formatCalls.Add(1)
		keys := make([]string, 0, len(in.Data))
		for k := range in.Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		lines := make([]string, 0, len(keys))
		for _, k := range keys {
			lines = append(lines, fmt.Sprintf("  %s: %v", k, in.Data[k]))
		}
		return strings.Join(lines, "\n"), nil
	}

	supportAgent := &ai.Agent{
		Name:  "support_agent",
		Model: model(t),
		Instructions: "You are a customer support agent. Use the available tools to " +
			"look up customers, check inventory, process orders, and format " +
			"responses for the customer.",
		Tools: []ai.ToolDef{
			tool.Func("format_response", "Format a data dictionary into a human-readable string.", formatResponse),
			tool.External[crmCustomerIn, map[string]any]("get_customer", "Look up customer details from the CRM system."),
			tool.External[inventoryIn, map[string]any]("check_inventory", "Check product availability in a warehouse."),
			tool.External[orderActionIn, map[string]any]("process_order", "Process a customer order. Actions: refund, cancel, update."),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, supportAgent,
		"Customer C-1234 wants to cancel order ORD-5678. "+
			"Look up the customer, check if we have the product in stock, "+
			"and process the cancellation.")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if n := customerCalls.Load(); n != 1 {
		t.Errorf("get_customer ran %d times, want 1", n)
	}
	if id, _ := customerID.Load().(string); id != "C-1234" {
		t.Errorf("get_customer was asked about %q, want C-1234", id)
	}
	if n := inventoryCalls.Load(); n < 1 {
		t.Errorf("check_inventory ran %d times, want at least 1", n)
	}
	if id, _ := productID.Load().(string); id != "PROD-001" {
		t.Errorf("check_inventory was asked about %q, want PROD-001 from the customer's order", id)
	}
	if n := orderCalls.Load(); n != 1 {
		t.Errorf("process_order ran %d times, want 1", n)
	}
	if id, _ := orderID.Load().(string); id != "ORD-5678" {
		t.Errorf("process_order was asked about %q, want ORD-5678", id)
	}
	if a, _ := action.Load().(string); a != "cancel" {
		t.Errorf("process_order action = %q, want cancel", a)
	}
	if n := formatCalls.Load(); n < 1 {
		t.Errorf("format_response ran %d times, want at least 1", n)
	}
	if !strings.Contains(result.Output, "ORD-5678") {
		t.Errorf("output should mention the cancelled order\n--- output ---\n%.400s", result.Output)
	}
}
