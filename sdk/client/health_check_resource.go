//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// HealthCheckResourceApiService is the service for checking the health of the Conductor server
type HealthCheckResourceApiService struct {
	*APIClient
}

// DoCheck checks the health of the Conductor server
func (a *HealthCheckResourceApiService) DoCheck(ctx context.Context) (model.HealthCheckStatus, *http.Response, error) {
	req := a.http_orkes.HealthCheckResourceAPI.DoCheck(ctx)
	result, resp, err := req.Execute()
	if err != nil {
		return model.HealthCheckStatus{}, resp, err
	}

	return toDomainHealthCheckStatusFromOrkes(result), resp, nil
}
