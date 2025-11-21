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
	"fmt"
	"net/http"
	"net/url"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/log"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/validation"
)

// SchemaResourceApiService is the service for the schema resource.
type SchemaResourceApiService struct {
	*APIClient
}

// SchemaResourceApiGetAllOpts contains optional parameters for GetAll
type SchemaResourceApiGetAllOpts struct {
	// Short if true, returns a short version of the schemas (default: false)
	Short optional.Bool
}

// GetAll gets all schemas.
func (s *SchemaResourceApiService) GetAll(ctx context.Context, opts *SchemaResourceApiGetAllOpts) ([]model.SchemaDefinition, *http.Response, error) {
	var result []model.SchemaDefinition
	path := "/schema"
	queryParams := url.Values{}
	if opts != nil && opts.Short.IsSet() {
		queryParams.Add("short", parameterToString(opts.Short.Value(), ""))
	}
	resp, err := s.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// SchemaResourceApiCreateSchemaOpts contains optional parameters for CreateSchema
type SchemaResourceApiCreateSchemaOpts struct {
	// NewVersion if true, creates a new version of the schema (default: false)
	NewVersion optional.Bool
}

// CreateSchema creates one or more schemas.
func (s *SchemaResourceApiService) CreateSchema(ctx context.Context, schemas []model.SchemaDefinition, opts *SchemaResourceApiCreateSchemaOpts) (*http.Response, error) {
	// Validate parameters
	if err := validation.NewValidator().
		SliceNotEmpty("schemas", schemas).
		Error(); err != nil {
		log.Error("error validating parameters", "error", err)
		return nil, err
	}

	path := "/schema"
	queryParams := url.Values{}
	if opts != nil && opts.NewVersion.IsSet() {
		queryParams.Add("newVersion", parameterToString(opts.NewVersion.Value(), ""))
	}
	resp, err := s.PostWithParams(ctx, path, queryParams, schemas, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteSchema deletes all versions of a schema by name.
func (s *SchemaResourceApiService) DeleteSchema(ctx context.Context, name string) (*http.Response, error) {
	// Validate parameters
	if err := validation.NewValidator().RequiredString("name", name).Error(); err != nil {
		log.Error("error validating parameters", "error", err)
		return nil, err
	}

	path := fmt.Sprintf("/schema/%s", name)
	resp, err := s.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetSchema gets a schema by name with the latest version.
func (s *SchemaResourceApiService) GetSchema(ctx context.Context, name string) (*model.SchemaDefinition, *http.Response, error) {
	// Validate parameters
	if err := validation.NewValidator().RequiredString("name", name).Error(); err != nil {
		log.Error("error validating parameters", "error", err)
		return nil, nil, err
	}

	var result model.SchemaDefinition
	path := fmt.Sprintf("/schema/%s", name)
	resp, err := s.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteSchemaVersion deletes a specific version of a schema.
func (s *SchemaResourceApiService) DeleteSchemaVersion(ctx context.Context, name string, version int32) (*http.Response, error) {
	// Validate parameters
	if err := validation.NewValidator().RequiredString("name", name).Error(); err != nil {
		log.Error("error validating parameters", "error", err)
		return nil, err
	}

	path := fmt.Sprintf("/schema/%s/%d", name, version)
	resp, err := s.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetSchemaVersion gets a schema by name and specific version.
func (s *SchemaResourceApiService) GetSchemaVersion(ctx context.Context, name string, version int32) (*model.SchemaDefinition, *http.Response, error) {
	// Validate parameters
	if err := validation.NewValidator().RequiredString("name", name).Error(); err != nil {
		log.Error("error validating parameters", "error", err)
		return nil, nil, err
	}

	var result model.SchemaDefinition
	path := fmt.Sprintf("/schema/%s/%d", name, version)
	resp, err := s.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
