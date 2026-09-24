//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package schema

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSnake(t *testing.T) {
	cases := map[string]string{
		"City":           "city",
		"AccountID":      "account_id",
		"OrderID":        "order_id",
		"TempF":          "temp_f",
		"ID":             "id",
		"URL":            "url",
		"HTTPStatus":     "http_status",
		"XMLHttpRequest": "xml_http_request",
		"MinChars":       "min_chars",
		"Field1":         "field1",
		"Temp1F":         "temp1_f",
		"already_snake":  "already_snake",
	}
	for in, want := range cases {
		if got := Snake(in); got != want {
			t.Errorf("Snake(%q) = %q, want %q", in, got, want)
		}
	}
}

// An untagged field is advertised to the model in snake_case; a tagged one
// keeps its tag. Both must be what the decoder binds, which TestRekeyInput
// covers from the other side.
func TestUntaggedFieldIsSnakeCased(t *testing.T) {
	type in struct {
		AccountID string
		Tagged    string `json:"renamed"`
		HTTPCode  int
	}
	raw, err := json.Marshal(Of(reflect.TypeOf(in{})))
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"properties":{"account_id":{"type":"string"},"renamed":{"type":"string"},"http_code":{"type":"integer"}},"required":["account_id","renamed","http_code"],"type":"object"}`
	if string(raw) != want {
		t.Errorf("schema JSON\n got %s\nwant %s", raw, want)
	}
}

func TestRekeyInput(t *testing.T) {
	type Address struct {
		ZipCode string
	}
	type Embedded struct {
		TraceID string
	}
	type in struct {
		Embedded
		AccountID string
		Tagged    string `json:"renamed"`
		Home      Address
		Others    []Address
		ByLabel   map[string]Address
		Skipped   string `json:"-"`
	}
	input := map[string]any{
		"account_id": "ACC-1",
		"renamed":    "kept",
		"trace_id":   "t-1",
		"home":       map[string]any{"zip_code": "94105"},
		"others":     []any{map[string]any{"zip_code": "10001"}},
		"by_label":   map[string]any{"work": map[string]any{"zip_code": "02134"}},
		"skipped":    "ignored",
	}
	got := RekeyInput(input, reflect.TypeOf(in{}))

	want := map[string]any{
		"AccountID": "ACC-1",
		"renamed":   "kept",
		"TraceID":   "t-1",
		"Home":      map[string]any{"ZipCode": "94105"},
		"Others":    []any{map[string]any{"ZipCode": "10001"}},
		"ByLabel":   map[string]any{"work": map[string]any{"ZipCode": "02134"}},
		"skipped":   "ignored",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RekeyInput =\n %#v\nwant\n %#v", got, want)
	}
	// The caller's map is untouched: guardrails read the model's original keys.
	if _, moved := input["AccountID"]; moved {
		t.Error("RekeyInput mutated its input")
	}

	// The proof that matters: the rekeyed map binds, the raw one silently does not.
	var bound in
	raw, _ := json.Marshal(got)
	if err := json.Unmarshal(raw, &bound); err != nil {
		t.Fatal(err)
	}
	if bound.AccountID != "ACC-1" || bound.TraceID != "t-1" || bound.Home.ZipCode != "94105" ||
		len(bound.Others) != 1 || bound.Others[0].ZipCode != "10001" || bound.ByLabel["work"].ZipCode != "02134" {
		t.Errorf("decoded = %+v", bound)
	}
	var unbound in
	raw, _ = json.Marshal(input)
	if err := json.Unmarshal(raw, &unbound); err != nil {
		t.Fatal(err)
	}
	if unbound.AccountID != "" {
		t.Error("expected the un-rekeyed input to miss AccountID; the rekey is unnecessary if it binds")
	}
}

func TestRekeyInputDoesNotOverwriteGoName(t *testing.T) {
	type in struct{ AccountID string }
	got := RekeyInput(map[string]any{"AccountID": "explicit", "account_id": "wire"}, reflect.TypeOf(in{}))
	if got["AccountID"] != "explicit" {
		t.Errorf("AccountID = %v, want the explicit value kept", got["AccountID"])
	}
}

func TestRekeyInputPassesNonStructsThrough(t *testing.T) {
	m := map[string]any{"a_b": 1}
	if got := RekeyInput(m, reflect.TypeOf(map[string]any{})); !reflect.DeepEqual(got, m) {
		t.Errorf("map type: %v", got)
	}
	if got := RekeyInput(nil, reflect.TypeOf(struct{ AB int }{})); got != nil {
		t.Errorf("nil input: %v", got)
	}
}

// The output direction: encoding/json writes an untagged field under its Go
// name, and the result must reach the model under the name the schema promised.
func TestRekeyOutputValue(t *testing.T) {
	type Address struct{ ZipCode string }
	type Embedded struct{ TraceID string }
	type out struct {
		Embedded
		AccountID string
		Tagged    string `json:"renamed"`
		Home      Address
		Others    []Address
	}
	// What encoding/json produces for an untagged struct, decoded back to any.
	encoded := map[string]any{
		"AccountID": "ACC-1",
		"renamed":   "kept",
		"TraceID":   "t-1",
		"Home":      map[string]any{"ZipCode": "94105"},
		"Others":    []any{map[string]any{"ZipCode": "10001"}},
	}
	got := RekeyOutputValue(encoded, reflect.TypeOf(out{}))
	want := map[string]any{
		"account_id": "ACC-1",
		"renamed":    "kept",
		"trace_id":   "t-1",
		"home":       map[string]any{"zip_code": "94105"},
		"others":     []any{map[string]any{"zip_code": "10001"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RekeyOutputValue =\n %#v\nwant\n %#v", got, want)
	}
	// A slice of structs at the top level is renamed element by element.
	list := RekeyOutputValue([]any{map[string]any{"ZipCode": "1"}}, reflect.TypeOf([]Address{}))
	if !reflect.DeepEqual(list, []any{map[string]any{"zip_code": "1"}}) {
		t.Errorf("slice = %#v", list)
	}
}

func TestNeedsRekey(t *testing.T) {
	type tagged struct {
		City string `json:"city"`
	}
	type untagged struct{ TempF int }
	type nested struct {
		Inner untagged `json:"inner"`
	}
	type single struct{ City string } // Snake("City") == "city", but the Go name differs
	type recursive struct {
		Next *recursive `json:"next"`
	}
	cases := []struct {
		name string
		typ  reflect.Type
		want bool
	}{
		{"all tagged", reflect.TypeOf(tagged{}), false},
		{"untagged", reflect.TypeOf(untagged{}), true},
		{"pointer to untagged", reflect.TypeOf(&untagged{}), true},
		{"slice of untagged", reflect.TypeOf([]untagged{}), true},
		{"map of untagged", reflect.TypeOf(map[string]untagged{}), true},
		{"untagged nested under a tag", reflect.TypeOf(nested{}), true},
		{"untagged whose snake form differs only by case", reflect.TypeOf(single{}), true},
		{"recursive type terminates", reflect.TypeOf(recursive{}), false},
		{"scalar", reflect.TypeOf(""), false},
		{"free-form map", reflect.TypeOf(map[string]any{}), false},
	}
	for _, c := range cases {
		if got := NeedsRekey(c.typ); got != c.want {
			t.Errorf("NeedsRekey(%s) = %v, want %v", c.name, got, c.want)
		}
	}
}

// An embedded struct whose type is unexported still has its exported fields
// promoted by encoding/json, so the rekey must look inside it in both
// directions. This is the shape that surfaced in suite 20's enriched record.
func TestRekeyThroughUnexportedEmbeddedType(t *testing.T) {
	type record struct { // unexported type name, exported fields
		RecordID string
		Value    int
	}
	type enriched struct {
		record
		ValueSquared int
	}
	typ := reflect.TypeOf(enriched{})
	if !NeedsRekey(typ) {
		t.Fatal("NeedsRekey must see through an embedded unexported struct type")
	}

	// Output: what encoding/json writes, renamed to what the schema promised.
	raw, err := json.Marshal(enriched{record{"r-1", 42}, 1764})
	if err != nil {
		t.Fatal(err)
	}
	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	got := RekeyOutputValue(decoded, typ)
	want := map[string]any{"record_id": "r-1", "value": float64(42), "value_squared": float64(1764)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("output rekey = %#v, want %#v", got, want)
	}

	// Input: the model's snake keys must bind onto the promoted fields.
	in := RekeyInput(map[string]any{"record_id": "r-2", "value": 7, "value_squared": 49}, typ)
	var bound enriched
	raw, _ = json.Marshal(in)
	if err := json.Unmarshal(raw, &bound); err != nil {
		t.Fatal(err)
	}
	if bound.RecordID != "r-2" || bound.Value != 7 || bound.ValueSquared != 49 {
		t.Errorf("bound = %+v", bound)
	}
}
