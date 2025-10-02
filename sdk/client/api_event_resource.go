// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// EventResourceApiService is the service for the event resource.
type EventResourceApiService struct {
	*APIClient
}

// AddEventHandler adds a new event handler.
func (a *EventResourceApiService) AddEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	req := a.http_orkes.EventResourceAPI.AddEventHandler(ctx)

	// Convert domain model to orkes model using mapper
	orkesHandler := toGeneratedEventHandlerForOrkes(body)

	req = req.EventHandler([]orkes.EventHandler{orkesHandler})
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// DeleteQueueConfig deletes queue configuration.
func (a *EventResourceApiService) DeleteQueueConfig(ctx context.Context, queueType string, queueName string) (*http.Response, error) {
	req := a.http_orkes.EventResourceAPI.DeleteQueueConfig(ctx, queueType, queueName)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// GetEventHandlers gets all event handlers.
func (a *EventResourceApiService) GetEventHandlers(ctx context.Context) ([]model.EventHandler, *http.Response, error) {
	req := a.http_orkes.EventResourceAPI.GetEventHandlers(ctx)
	orkesHandlers, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert orkes models to domain models using mapper
	result := toDomainEventHandlersFromOrkes(orkesHandlers)
	return result, resp, nil
}

// EventResourceApiGetEventHandlersForEventOpts is the optional parameters for the GetEventHandlersForEvent method.
type EventResourceApiGetEventHandlersForEventOpts struct {
	ActiveOnly optional.Bool
}

// GetEventHandlersForEvent gets event handlers for a specific event.
func (a *EventResourceApiService) GetEventHandlersForEvent(ctx context.Context, event string, opts *EventResourceApiGetEventHandlersForEventOpts) ([]model.EventHandler, *http.Response, error) {
	req := a.http_orkes.EventResourceAPI.GetEventHandlersForEvent(ctx, event)

	// Apply optional parameters
	if opts != nil && opts.ActiveOnly.IsSet() {
		req = req.ActiveOnly(opts.ActiveOnly.Value())
	}

	orkesHandlers, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	// Convert orkes models to domain models using mapper
	result := toDomainEventHandlersFromOrkes(orkesHandlers)
	return result, resp, nil
}

// GetQueueConfig gets queue configuration.
func (a *EventResourceApiService) GetQueueConfig(ctx context.Context, queueType string, queueName string) (map[string]interface{}, *http.Response, error) {
	// Use orkes generated client
	req := a.http_orkes.EventResourceAPI.GetQueueConfig(ctx, queueType, queueName)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	flatResult := make(map[string]interface{})
	for k, v := range result {
		flatResult[k] = v
	}

	return flatResult, resp, nil
}

// GetQueueNames gets all queue names.
func (a *EventResourceApiService) GetQueueNames(ctx context.Context) (map[string]string, *http.Response, error) {
	req := a.http_orkes.EventResourceAPI.GetQueueNames(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	return result, resp, nil
}

// PutQueueConfig puts queue configuration.
func (a *EventResourceApiService) PutQueueConfig(ctx context.Context, body string, queueType string, queueName string) (*http.Response, error) {
	req := a.http_orkes.EventResourceAPI.PutQueueConfig(ctx, queueType, queueName)
	req = req.Body(body)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// RemoveEventHandlerStatus removes event handler.
func (a *EventResourceApiService) RemoveEventHandler(ctx context.Context, name string) (*http.Response, error) {
	req := a.http_orkes.EventResourceAPI.RemoveEventHandlerStatus(ctx, name)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}

// UpdateEventHandler updates event handler.
func (a *EventResourceApiService) UpdateEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	req := a.http_orkes.EventResourceAPI.UpdateEventHandler(ctx)

	// Convert domain model to orkes model using mapper
	orkesHandler := toGeneratedEventHandlerForOrkes(body)

	req = req.EventHandler(orkesHandler)
	resp, err := req.Execute()
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}
	return resp, nil
}
