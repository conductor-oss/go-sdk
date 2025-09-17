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

type SecretResourceApiService struct {
	*APIClient
}

// NewSecretResourceApiService creates a new SecretResourceApiService instance
func NewSecretResourceApiService(client *APIClient) *SecretResourceApiService {
	return &SecretResourceApiService{client}
}

// ClearLocalCache Clear local cache
func (a *SecretResourceApiService) ClearLocalCache(ctx context.Context) (map[string]string, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.ClearLocalCache(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ClearRedisCache Clear redis cache
func (a *SecretResourceApiService) ClearRedisCache(ctx context.Context) (map[string]string, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.ClearRedisCache(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteSecret Delete a secret value by key
func (a *SecretResourceApiService) DeleteSecret(ctx context.Context, key string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.DeleteSecret(ctx, key).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// DeleteTagForSecret Delete tags of the secret
func (a *SecretResourceApiService) DeleteTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error) {
	genTags := toGeneratedTags(body)
	resp, err := a.http_orkes.SecretResourceAPI.DeleteTagForSecret(ctx, key).Tag(genTags).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// GetSecret Get secret value by key
func (a *SecretResourceApiService) GetSecret(ctx context.Context, key string) (string, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.GetSecret(ctx, key).Execute()
	if err != nil {
		return "", resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// GetTags Get tags by secret
func (a *SecretResourceApiService) GetTags(ctx context.Context, key string) ([]model.Tag, *http.Response, error) {
	genResult, resp, err := a.http_orkes.SecretResourceAPI.GetTags(ctx, key).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainTagsFromGenerated(genResult)
	return result, resp, nil
}

// ListAllSecretNames List all secret names
func (a *SecretResourceApiService) ListAllSecretNames(ctx context.Context) ([]string, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.ListAllSecretNames(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ListSecretsThatUserCanGrantAccessTo List all secret names user can grant access to
func (a *SecretResourceApiService) ListSecretsThatUserCanGrantAccessTo(ctx context.Context) ([]string, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.ListSecretsThatUserCanGrantAccessTo(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// ListSecretsWithTagsThatUserCanGrantAccessTo List all secret names along with tags user can grant access to
func (a *SecretResourceApiService) ListSecretsWithTagsThatUserCanGrantAccessTo(ctx context.Context) ([]model.Secret, *http.Response, error) {
	genResult, resp, err := a.http_orkes.SecretResourceAPI.ListSecretsWithTagsThatUserCanGrantAccessTo(ctx).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}

	result := toDomainSecretsFromGenerated(genResult)
	return result, resp, nil
}

// PutSecret Put a secret value by key
func (a *SecretResourceApiService) PutSecret(ctx context.Context, body string, key string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.PutSecret(ctx, key).Body(body).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}

// PutTagForSecret Tag a secret
func (a *SecretResourceApiService) PutTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error) {
	genTags := toGeneratedTags(body)
	resp, err := a.http_orkes.SecretResourceAPI.PutTagForSecret(ctx, key).Tag(genTags).Execute()
	return resp, wrapGeneratedError(err, resp)
}

// SecretExists Check if secret exists
func (a *SecretResourceApiService) SecretExists(ctx context.Context, key string) (interface{}, *http.Response, error) {
	result, resp, err := a.http_orkes.SecretResourceAPI.SecretExists(ctx, key).Execute()
	if err != nil {
		return nil, resp, wrapGeneratedError(err, resp)
	}
	return result, resp, nil
}
