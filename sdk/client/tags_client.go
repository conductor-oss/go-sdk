// Package client provides API client functionality for Conductor
//
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

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// TagsClient provides methods for managing tags on tasks and workflows
type TagsClient interface {
	AddTaskTag(ctx context.Context, body model.TagObject, taskName string) (interface{}, *http.Response, error)
	DeleteTaskTag(ctx context.Context, body model.TagString, taskName string) (interface{}, *http.Response, error)
	GetTaskTags(ctx context.Context, taskName string) ([]model.TagObject, *http.Response, error)
	SetTaskTags(ctx context.Context, body []model.TagObject, taskName string) (interface{}, *http.Response, error)
	AddWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error)
	DeleteWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error)
	GetWorkflowTags(ctx context.Context, name string) ([]model.TagObject, *http.Response, error)
	SetWorkflowTags(ctx context.Context, body []model.TagObject, name string) (interface{}, *http.Response, error)
	GetTags(ctx context.Context) ([]model.TagObject, *http.Response, error)
}

// NewTagsClient creates a new TagsClient instance
func NewTagsClient(client *APIClient) TagsClient {
	return &TagsApiService{client}
}
