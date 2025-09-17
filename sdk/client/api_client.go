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
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/generated/http/conductor"
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes"
	"github.com/conductor-sdk/conductor-go/sdk/generated/http/orkes_v4"
	"github.com/conductor-sdk/conductor-go/sdk/log"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const (
	CONDUCTOR_AUTH_KEY            = "CONDUCTOR_AUTH_KEY"
	CONDUCTOR_AUTH_SECRET         = "CONDUCTOR_AUTH_SECRET"
	CONDUCTOR_SERVER_URL          = "CONDUCTOR_SERVER_URL"
	CONDUCTOR_CLIENT_HTTP_TIMEOUT = "CONDUCTOR_CLIENT_HTTP_TIMEOUT"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

type APIClient struct {
	http_orkes     *orkes.APIClient
	http_conductor *conductor.APIClient
	http_orkes_v4  *orkes_v4.APIClient
	httpClient     *http.Client
	tokenManager   authentication.TokenManager
	httpSettings   *settings.HttpSettings
}

func NewAPIClient(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		nil,
		nil,
	)
}
func NewAPIClientFromEnv() *APIClient {
	return NewAPIClient(NewAuthenticationSettingsFromEnv(), NewHttpSettingsFromEnv())
}

func NewAuthenticationSettingsFromEnv() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		os.Getenv(CONDUCTOR_AUTH_KEY),
		os.Getenv(CONDUCTOR_AUTH_SECRET),
	)
}

func NewHttpSettingsFromEnv() *settings.HttpSettings {
	url := os.Getenv(CONDUCTOR_SERVER_URL)
	if url == "" {
		log.Error("Environment variable CONDUCTOR_SERVER_URL is not set")
	}

	return settings.NewHttpSettings(url)
}

func NewAPIClientWithTokenExpiration(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	tokenExpiration *authentication.TokenExpiration,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		tokenExpiration,
		nil,
	)
}

func NewAPIClientWithTokenManager(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
	tokenExpiration *authentication.TokenExpiration,
	tokenManager authentication.TokenManager,
) *APIClient {
	return newAPIClient(
		authenticationSettings,
		httpSettings,
		tokenExpiration,
		tokenManager,
	)
}

func newAPIClient(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings, tokenExpiration *authentication.TokenExpiration, tokenManager authentication.TokenManager) *APIClient {
	if httpSettings == nil {
		httpSettings = settings.NewHttpDefaultSettings()
	}
	var httpTimeout = 30 * time.Second // Set default value once

	timeoutStr := os.Getenv(CONDUCTOR_CLIENT_HTTP_TIMEOUT)
	if timeoutStr != "" {
		// Only try to parse if the environment variable is actually set
		if timeoutInt, err := strconv.Atoi(timeoutStr); err == nil {
			httpTimeout = time.Duration(timeoutInt) * time.Second
		}
		// If parsing fails, we'll keep the default value
	}

	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	baseTransport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         baseDialer.DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		DisableCompression:  false, // This ensures automatic gzip handling
	}

	// Wrap base transport with gzip handling
	gzipTransport := &GzipTransport{
		Transport: baseTransport,
	}

	// Create ONE HTTP client for token refresh (without auth to avoid recursion)
	tokenRefreshClient := &http.Client{
		Transport: baseTransport,
		Timeout:   httpTimeout,
	}

	// Initialize token manager if not provided
	if authenticationSettings != nil && !authenticationSettings.IsEmpty() {
		if tokenManager == nil {
			tokenManager = authentication.NewTokenManager(*authenticationSettings, tokenExpiration)
		}
	}

	var finalTransport http.RoundTripper
	// Wrap transport with auth if we have token manager
	if tokenManager != nil {
		finalTransport = &AuthTransport{
			TokenManager:  tokenManager,
			HttpSettings:  httpSettings,
			HttpClient:    tokenRefreshClient, // Client for token refresh only
			BaseTransport: gzipTransport,      // Use gzip transport as base
		}
	}

	// Create THE SINGLE HTTP client with the final transport
	httpClient := &http.Client{
		Transport: finalTransport,
		Timeout:   httpTimeout,
	}

	// Create configuration for orkes client
	config := orkes.NewConfiguration()

	// Set base URL properly
	if httpSettings != nil && httpSettings.BaseUrl != "" {
		baseUrl := httpSettings.BaseUrl
		// Use the full URL as the server URL
		config.Servers = orkes.ServerConfigurations{
			{
				URL:         baseUrl,
				Description: "Configured server",
			},
		}
	}

	// Copy headers
	config.DefaultHeader = make(map[string]string)
	for key, value := range httpSettings.Headers {
		config.DefaultHeader[key] = value
	}

	// Use THE SAME HTTP client for orkes API
	config.HTTPClient = httpClient

	// Create orkes API client
	http_orkes := orkes.NewAPIClient(config)

	// Create configuration for conductor client
	conductorConfig := conductor.NewConfiguration()
	// Convert server configurations
	conductorServers := make(conductor.ServerConfigurations, len(config.Servers))
	for i, server := range config.Servers {
		conductorServers[i] = conductor.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
		}
	}
	conductorConfig.Servers = conductorServers
	conductorConfig.DefaultHeader = config.DefaultHeader
	conductorConfig.HTTPClient = httpClient
	http_conductor := conductor.NewAPIClient(conductorConfig)

	// Create configuration for orkes_v4 client
	orkesV4Config := orkes_v4.NewConfiguration()
	// Convert server configurations
	orkesV4Servers := make(orkes_v4.ServerConfigurations, len(config.Servers))
	for i, server := range config.Servers {
		orkesV4Servers[i] = orkes_v4.ServerConfiguration{
			URL:         server.URL,
			Description: server.Description,
		}
	}
	orkesV4Config.Servers = orkesV4Servers
	orkesV4Config.DefaultHeader = config.DefaultHeader
	orkesV4Config.HTTPClient = httpClient
	http_orkes_v4 := orkes_v4.NewAPIClient(orkesV4Config)

	// Save all necessary components in APIClient
	return &APIClient{
		http_orkes:     http_orkes,
		http_conductor: http_conductor,
		http_orkes_v4:  http_orkes_v4,
		httpClient:     httpClient,
		tokenManager:   tokenManager,
		httpSettings:   httpSettings,
	}
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}

	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		if err = json.Unmarshal(b, v); err != nil {
			// Hacky - if json unmarshalling fails, return a string.
			// it's because the backend might reply with content-type: application/json and a string.
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.String {
				rv.Elem().SetString(string(b))
				return nil
			}
			return err
		}
		return nil
	} else if strings.Contains(contentType, "text/plain") {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return errors.New("undefined response type")
		}
		rv.Elem().SetString(string(b))
		return nil
	}

	return errors.New("undefined response type")
}

func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte,
) (localVarRequest *http.Request, err error) {
	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	urlStr, err := url.Parse(c.httpSettings.BaseUrl + path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := urlStr.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	urlStr.RawQuery = query.Encode()

	if body != nil {
		localVarRequest, err = http.NewRequestWithContext(ctx, method, urlStr.String(), body)
	} else {
		localVarRequest, err = http.NewRequestWithContext(ctx, method, urlStr.String(), nil)
	}

	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add default headers from httpSettings
	for header, value := range c.httpSettings.Headers {
		localVarRequest.Header.Add(header, value)
	}

	// Note: Token is already handled by AuthTransport, no need to add it here

	return localVarRequest, nil
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

func getDecompressedBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	var reader io.ReadCloser
	var err error
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			log.Error("Unable to decompress the response", "error", err)
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
	default:
		reader = response.Body
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

func isSuccessfulStatus(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

// executeCall performs an HTTP request with centralized error handling
// Supports all CRUD operations through a common interface
func (c *APIClient) executeCall(ctx context.Context, method, path string, queryParams url.Values, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	// Create headers
	headers := make(map[string]string)

	// Set content type if body is provided
	if body != nil {
		cType := "application/json"
		if len(contentType) > 0 && contentType != "" {
			cType = contentType
		}
		headers["Content-Type"] = cType
	}

	// Set accept header for all requests
	headers["Accept"] = "application/json"

	// Prepare the request
	req, err := c.prepareRequest(ctx, path, method, body, headers, queryParams, nil, "", nil)
	if err != nil {
		return nil, err
	}

	// Call the API
	resp, err := c.callAPI(req)
	if err != nil || resp == nil {
		return resp, wrapGeneratedError(err, resp)
	}

	// Get response body
	respBody, err := getDecompressedBody(resp)
	if err != nil {
		return resp, wrapGeneratedError(err, resp)
	}

	// Handle successful response
	if isSuccessfulStatus(resp.StatusCode) {
		if result != nil && len(respBody) > 0 {
			err = c.decode(result, respBody, resp.Header.Get("Content-Type"))
		}
		return resp, wrapGeneratedError(err, resp)
	}

	// Handle error response - create GenericSwaggerError with status code
	newErr := NewGenericSwaggerError(respBody, string(respBody), nil, resp.StatusCode)
	return resp, newErr
}

// Deprecated: use generated resource clients instead; this method will be removed.
// Get performs a GET request
func (c *APIClient) Get(ctx context.Context, path string, queryParams url.Values, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "GET", path, queryParams, nil, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// Post performs a POST request
func (c *APIClient) Post(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, nil, body, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// PostWithContentType performs post with given content type
func (c *APIClient) PostWithContentType(ctx context.Context, path string, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, nil, body, contentType, result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// PostWithParams performs a POST request with query parameters
func (c *APIClient) PostWithParams(ctx context.Context, path string, queryParams url.Values, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "POST", path, queryParams, body, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// Put performs a PUT request
func (c *APIClient) Put(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, nil, body, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// PutWithContentType performs a PUT request
func (c *APIClient) PutWithContentType(ctx context.Context, path string, body interface{}, contentType string, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, nil, body, contentType, result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// PutWithParams performs a PUT request with query parameters
func (c *APIClient) PutWithParams(ctx context.Context, path string, queryParams url.Values, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PUT", path, queryParams, body, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// Delete performs a DELETE request without a body
func (c *APIClient) Delete(ctx context.Context, path string, queryParams url.Values, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "DELETE", path, queryParams, nil, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// DeleteWithBody performs a DELETE request with a body
func (c *APIClient) DeleteWithBody(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "DELETE", path, nil, body, "", result)
}

// Deprecated: use generated resource clients instead; this method will be removed.
// Patch performs a PATCH request
func (c *APIClient) Patch(ctx context.Context, path string, body interface{}, result interface{}) (*http.Response, error) {
	return c.executeCall(ctx, "PATCH", path, nil, body, "", result)
}
