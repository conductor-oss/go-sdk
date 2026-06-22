//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

import "context"

type pathTemplateKey struct{}
type rawPathKey struct{}

// WithPathTemplate returns a derived context carrying the URI path template
// (e.g. "/workflow/{workflowId}"). The metricsRoundTripper uses this value
// as the "uri" label instead of the fully-resolved request path, keeping
// metric cardinality bounded.
func WithPathTemplate(ctx context.Context, template string) context.Context {
	return context.WithValue(ctx, pathTemplateKey{}, template)
}

// PathTemplateFromContext extracts the URI path template stored by
// WithPathTemplate. Returns "" when the context carries no template.
func PathTemplateFromContext(ctx context.Context) string {
	v, ok := ctx.Value(pathTemplateKey{}).(string)
	if !ok {
		return ""
	}
	return v
}

// WithRawPath returns a derived context carrying the pre-concatenation API
// path (e.g. "/tasks"). This is set automatically by the HTTP client layer so
// that the metricsRoundTripper can use the clean path when no explicit
// template has been provided, avoiding the base-URL prefix in the uri label.
func WithRawPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, rawPathKey{}, path)
}

// RawPathFromContext extracts the raw API path stored by WithRawPath.
// Returns "" when the context carries no raw path.
func RawPathFromContext(ctx context.Context) string {
	v, ok := ctx.Value(rawPathKey{}).(string)
	if !ok {
		return ""
	}
	return v
}
