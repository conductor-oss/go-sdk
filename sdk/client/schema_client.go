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

type SchemaClient interface {
	// GetAll gets all schemas.
	GetAll(ctx context.Context, opts *SchemaResourceApiGetAllOpts) ([]model.SchemaDefinition, *http.Response, error)

	// CreateSchema creates one or more schemas.
	CreateSchema(ctx context.Context, schemas []model.SchemaDefinition, opts *SchemaResourceApiCreateSchemaOpts) (*http.Response, error)

	// DeleteSchema deletes all versions of a schema by name.
	DeleteSchema(ctx context.Context, name string) (*http.Response, error)

	// GetSchema gets a schema by name with the latest version.
	GetSchema(ctx context.Context, name string) (*model.SchemaDefinition, *http.Response, error)

	// DeleteSchemaVersion deletes a specific version of a schema.
	DeleteSchemaVersion(ctx context.Context, name string, version int32) (*http.Response, error)

	// GetSchemaVersion gets a schema by name and specific version.
	GetSchemaVersion(ctx context.Context, name string, version int32) (*model.SchemaDefinition, *http.Response, error)
}

// NewSchemaClient creates a new SchemaClient.
func NewSchemaClient(apiClient *APIClient) SchemaClient {
	return &SchemaResourceApiService{apiClient}
}
