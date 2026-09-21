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
// The output must match what the Python and Java SDKs put on the wire: types
// and required-ness only, no descriptions, titles, formats or constraints.
package schema

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

// Of returns the JSON Schema for t: struct properties come from `json` tags in
// declaration order, which matters because `required` is a positional array. A
// pointer or omitempty field is optional, Go's stand-in for a Python default.
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
		// Python emits an open object for a bare dict: additionalProperties
		// present but empty.
		if isAny(t.Elem()) {
			return map[string]any{"type": "object", "additionalProperties": map[string]any{}}
		}
		return map[string]any{"type": "object", "additionalProperties": Of(t.Elem())}

	case reflect.Struct:
		return structSchema(t)

	case reflect.Interface:
		if isAny(t) {
			return map[string]any{}
		}
		return map[string]any{"type": "object"}
	}

	return map[string]any{}
}

func isAny(t reflect.Type) bool {
	return t.Kind() == reflect.Interface && t.NumMethod() == 0
}

// Properties is an object schema's property set in declaration order, which a
// Go map would lose: the server copies a tool's parameter order into text the
// model reads, so it has to match what the Python SDK sends.
type Properties struct {
	keys   []string
	values map[string]any
}

// Set adds or replaces a property, keeping first-insertion order.
func (p *Properties) Set(name string, schema any) {
	if p.values == nil {
		p.values = map[string]any{}
	}
	if _, exists := p.values[name]; !exists {
		p.keys = append(p.keys, name)
	}
	p.values[name] = schema
}

// Get returns a property's schema.
func (p *Properties) Get(name string) (any, bool) {
	v, ok := p.values[name]
	return v, ok
}

// Keys lists the property names in declaration order.
func (p *Properties) Keys() []string { return append([]string(nil), p.keys...) }

// Len is the number of properties.
func (p *Properties) Len() int { return len(p.keys) }

// MarshalJSON writes the properties as a JSON object in declaration order.
func (p *Properties) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteByte('{')
	for i, k := range p.keys {
		if i > 0 {
			b.WriteByte(',')
		}
		name, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(p.values[k])
		if err != nil {
			return nil, err
		}
		b.Write(name)
		b.WriteByte(':')
		b.Write(val)
	}
	b.WriteByte('}')
	return b.Bytes(), nil
}

func structSchema(t reflect.Type) map[string]any {
	props := &Properties{}
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
			if sub, ok := structSchema(deref(f.Type))["properties"].(*Properties); ok {
				for _, k := range sub.Keys() {
					v, _ := sub.Get(k)
					props.Set(k, v)
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
		props.Set(name, Of(f.Type))
		if f.Type.Kind() != reflect.Pointer && !opts.omitempty {
			required = append(required, name)
		}
	}

	out := map[string]any{"type": "object", "properties": props}
	if len(required) > 0 {
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
