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

// Properties are written in declaration order, as the Python SDK writes them,
// not in the sorted order encoding/json gives a map. The server copies that
// order into text the model reads, so it is part of prompt parity.
func TestStructSchemaKeepsDeclarationOrder(t *testing.T) {
	type Embedded struct {
		Zeta string `json:"zeta"`
	}
	type in struct {
		Embedded
		Text     string `json:"text"`
		MinChars int    `json:"min_chars"`
		Alpha    *bool  `json:"alpha,omitempty"`
	}
	raw, err := json.Marshal(Of(reflect.TypeOf(in{})))
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"properties":{"zeta":{"type":"string"},"text":{"type":"string"},"min_chars":{"type":"integer"},"alpha":{"type":"boolean"}},"required":["zeta","text","min_chars"],"type":"object"}`
	if string(raw) != want {
		t.Errorf("schema JSON\n got %s\nwant %s", raw, want)
	}
	props := Of(reflect.TypeOf(in{}))["properties"].(*Properties)
	if got := props.Keys(); !reflect.DeepEqual(got, []string{"zeta", "text", "min_chars", "alpha"}) {
		t.Errorf("Keys() = %v", got)
	}
}
