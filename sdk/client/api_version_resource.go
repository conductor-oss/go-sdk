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
)

// VersionResourceAPIService is a service for getting the server's version
type VersionResourceAPIService struct {
	*APIClient
}

// NewVersionResourceAPIService creates a new VersionResourceAPIService
func NewVersionResourceAPIService(apiClient *APIClient) *VersionResourceAPIService {
	return &VersionResourceAPIService{APIClient: apiClient}
}

// GetVersion gets the server's version
func (a *VersionResourceAPIService) GetVersion(ctx context.Context) (string, *http.Response, error) {
	var result string
	path := "/version"
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}
