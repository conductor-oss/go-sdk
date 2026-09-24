package agents

import (
	"errors"
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
)

// StartExternalWorkers starts the services 33_external_workers calls, plain
// Conductor workers for get_customer, check_inventory and process_order, and
// returns a function that stops them. They run in another process in the
// Python example; the command in ./cmd starts them that way.
func StartExternalWorkers(apiClient *client.APIClient) (stop func(), err error) {
	order := func(orderID string) map[string]any {
		return map[string]any{"order_id": orderID, "customer_id": "C-1234", "product_id": "PROD-001", "warehouse": "default"}
	}
	workers := map[string]func(*model.Task) (any, error){
		"get_customer": func(task *model.Task) (any, error) {
			pending := order("ORD-5678")
			pending["status"] = "pending"
			return map[string]any{
				"customer_id": task.InputData["customer_id"],
				"name":        "Example Customer",
				"orders":      []any{pending},
			}, nil
		},
		"check_inventory": func(task *model.Task) (any, error) {
			warehouse, ok := task.InputData["warehouse"]
			if !ok || warehouse == "" {
				warehouse = "default"
			}
			return map[string]any{
				"product_id": task.InputData["product_id"],
				"warehouse":  warehouse,
				"in_stock":   true,
				"quantity":   12,
			}, nil
		},
		"process_order": func(task *model.Task) (any, error) {
			if task.InputData["action"] != "cancel" {
				return nil, errors.New("this demo supports cancellation only")
			}
			cancelled := order(fmt.Sprint(task.InputData["order_id"]))
			cancelled["status"] = "cancelled"
			return cancelled, nil
		},
	}
	runner := worker.NewTaskRunnerWithApiClient(apiClient)
	for name, fn := range workers {
		if err := runner.StartWorker(name, fn, 1, 100*time.Millisecond); err != nil {
			return nil, fmt.Errorf("start worker %s: %w", name, err)
		}
	}
	return func() {
		for name := range workers {
			runner.Shutdown(name)
		}
		runner.WaitWorkers()
	}, nil
}
