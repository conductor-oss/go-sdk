//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

// Package schema derives JSON Schema from Go types for tool input and output.
//
// The output must match what the Python and Java SDKs put on the wire, which is
// narrower than JSON Schema allows: types and required-ness only, no
// descriptions, titles, formats or constraints. Python's schema_from_function
// carries the same restriction — the comment there claims docstrings supply
// parameter descriptions, but the code never attaches them.
package schema

import (
	"reflect"
	"strings"
)

// Of returns the JSON Schema for t.
//
// Structs become objects whose properties come from `json` tags, in declaration
// order. A field is required unless it is a pointer or carries omitempty; Go has
// no parameter defaults, so those two stand in for Python's "has a default".
// Declaration order is load-bearing: `required` is a JSON array, so the golden
// comparison is positional.
func Of(t reflect.Type) map[string]any {
	for t != nil && t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t == nil {
		return map[string]any{}
	}

	switch t.Kind() {
	case reflect.String:
		return map[string]any{"type": "string"}
	case reflect.Bool:
		return map[string]any{"type": "boolean"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return map[string]any{"type": "integer"}
	case reflect.Float32, reflect.Float64:
		return map[string]any{"type": "number"}

	case reflect.Slice, reflect.Array:
		// []byte is not an array of integers on the wire.
		if t.Elem().Kind() == reflect.Uint8 {
			return map[string]any{"type": "string"}
		}
		return map[string]any{"type": "array", "items": Of(t.Elem())}

	case reflect.Map:
		// map[string]any is the open object Python emits for a bare dict:
		// additionalProperties is present but empty, not absent.
		if isAny(t.Elem()) {
			return map[string]any{"type": "object", "additionalProperties": map[string]any{}}
		}
		return map[string]any{"type": "object", "additionalProperties": Of(t.Elem())}

	case reflect.Struct:
		return structSchema(t)

	case reflect.Interface:
		if isAny(t) {
			// An untyped value places no constraint at all.
			return map[string]any{}
		}
		return map[string]any{"type": "object"}
	}

	return map[string]any{}
}

func isAny(t reflect.Type) bool {
	return t.Kind() == reflect.Interface && t.NumMethod() == 0
}

func structSchema(t reflect.Type) map[string]any {
	props := map[string]any{}
	var required []string

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" { // unexported
			continue
		}
		name, opts := parseTag(f)
		if name == "-" {
			continue
		}
		if f.Anonymous && name == "" {
			// Embedded struct: lift its properties, as encoding/json does.
			if sub, ok := structSchema(deref(f.Type))["properties"].(map[string]any); ok {
				for k, v := range sub {
					props[k] = v
				}
			}
			if sub, ok := structSchema(deref(f.Type))["required"].([]string); ok {
				required = append(required, sub...)
			}
			continue
		}
		if name == "" {
			name = f.Name
		}
		props[name] = Of(f.Type)
		if f.Type.Kind() != reflect.Pointer && !opts.omitempty {
			required = append(required, name)
		}
	}

	out := map[string]any{"type": "object", "properties": props}
	if len(required) > 0 {
		// []string rather than []any: the golden comparison round-trips through
		// JSON, so either encodes identically.
		out["required"] = required
	}
	return out
}

func deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

type tagOpts struct{ omitempty bool }

func parseTag(f reflect.StructField) (string, tagOpts) {
	tag := f.Tag.Get("json")
	if tag == "" {
		return "", tagOpts{}
	}
	parts := strings.Split(tag, ",")
	var o tagOpts
	for _, p := range parts[1:] {
		if p == "omitempty" {
			o.omitempty = true
		}
	}
	return parts[0], o
}
