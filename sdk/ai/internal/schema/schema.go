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
	"unicode"
)

// Of returns the JSON Schema for t. A struct property is named by its `json`
// tag, or by Snake of the field name when there is none, in declaration order,
// which matters because `required` is a positional array. A pointer or
// omitempty field is optional, Go's stand-in for a Python default.
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
			name = Snake(f.Name)
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

// Snake converts a Go identifier to snake_case, acronyms included: AccountID
// becomes account_id, HTTPStatus becomes http_status, TempF becomes temp_f.
// It names an untagged struct field on the wire, so Go tools read like Python
// ones without a tag on every field.
func Snake(name string) string {
	runes := []rune(name)
	var b strings.Builder
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			prev := runes[i-1]
			nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextLower) {
				b.WriteByte('_')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// RekeyInput returns a copy of a task's input with the wire names Of emits for
// untagged fields renamed to the Go field names encoding/json binds, so a field
// declared AccountID receives the model's account_id. encoding/json accepts an
// untagged field's own name and case-insensitive variants of it, but not the
// underscored form, and a missed key decodes silently to a zero value; this is
// what makes untagged fields safe to use. Tagged fields are left alone because
// encoding/json already matches their tag. It descends into nested structs and
// into slices and maps of structs, and never overwrites a key already present
// under the Go name.
func RekeyInput(in map[string]any, t reflect.Type) map[string]any {
	if in == nil {
		return nil
	}
	m, ok := rekeyValue(in, t, toGo).(map[string]any)
	if !ok {
		return in
	}
	return m
}

// RekeyOutputValue is the inverse of RekeyInput for a handler's result, given
// its JSON-decoded form: encoding/json wrote each untagged field under its Go
// name, and this renames it to the snake_case name the output schema
// advertised, so the model reads the keys it was told to expect. It accepts a
// struct, or a slice or map of structs, decoded to any.
func RekeyOutputValue(v any, t reflect.Type) any {
	return rekeyValue(v, t, toWire)
}

// NeedsRekey reports whether t has, at any depth, an untagged exported field
// whose snake_case wire name differs from its Go name. When false, encoding/json
// already produces the advertised names and a result can travel untouched, in
// declaration order.
func NeedsRekey(t reflect.Type) bool {
	return needsRekey(t, map[reflect.Type]bool{})
}

func needsRekey(t reflect.Type, seen map[reflect.Type]bool) bool {
	t = deref(t)
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return needsRekey(t.Elem(), seen)
	case reflect.Struct:
	default:
		return false
	}
	if seen[t] {
		return false
	}
	seen[t] = true
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// encoding/json promotes an embedded struct's exported fields even when
		// the embedded type itself is unexported, so only skip named fields.
		if f.PkgPath != "" && !f.Anonymous {
			continue
		}
		tag, _ := parseTag(f)
		if tag == "-" {
			continue
		}
		if tag == "" && !f.Anonymous && Snake(f.Name) != f.Name {
			return true
		}
		if needsRekey(f.Type, seen) {
			return true
		}
	}
	return false
}

// direction is which way rekeyStruct renames an untagged field: toGo turns the
// wire name into the field name for decoding, toWire the reverse for encoding.
type direction bool

const (
	toGo   direction = true
	toWire direction = false
)

func rekeyStruct(m map[string]any, t reflect.Type, d direction) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" && !f.Anonymous {
			continue
		}
		tag, _ := parseTag(f)
		if tag == "-" {
			continue
		}
		if f.Anonymous && tag == "" {
			// Promoted fields bind at this level, as encoding/json treats them,
			// including those of an embedded type that is itself unexported.
			if ft := deref(f.Type); ft.Kind() == reflect.Struct {
				rekeyStruct(m, ft, d)
			}
			continue
		}
		key := tag
		if key == "" {
			wire := Snake(f.Name)
			if d == toWire {
				key = moveKey(m, f.Name, wire)
			} else {
				key = moveKey(m, wire, f.Name)
			}
		}
		if v, ok := m[key]; ok {
			m[key] = rekeyValue(v, f.Type, d)
		}
	}
}

// moveKey renames m[from] to m[to] unless the two are the same or to is
// already present, and returns to, the key the value now sits under.
func moveKey(m map[string]any, from, to string) string {
	if from == to {
		return to
	}
	if v, ok := m[from]; ok {
		if _, taken := m[to]; !taken {
			m[to] = v
			delete(m, from)
		}
	}
	return to
}

// rekeyValue copies and renames through the containers a field can be: a
// struct, or a slice or map of structs. Anything else passes through.
func rekeyValue(v any, t reflect.Type, d direction) any {
	t = deref(t)
	switch t.Kind() {
	case reflect.Struct:
		if m, ok := v.(map[string]any); ok {
			out := make(map[string]any, len(m))
			for k, e := range m {
				out[k] = e
			}
			rekeyStruct(out, t, d)
			return out
		}
	case reflect.Slice, reflect.Array:
		if list, ok := v.([]any); ok && deref(t.Elem()).Kind() == reflect.Struct {
			out := make([]any, len(list))
			for i, e := range list {
				out[i] = rekeyValue(e, t.Elem(), d)
			}
			return out
		}
	case reflect.Map:
		if m, ok := v.(map[string]any); ok && deref(t.Elem()).Kind() == reflect.Struct {
			out := make(map[string]any, len(m))
			for k, e := range m {
				out[k] = rekeyValue(e, t.Elem(), d)
			}
			return out
		}
	}
	return v
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
