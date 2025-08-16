//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package worker

import (
	"encoding/json"
	"fmt"
)

// InputBinder performs conversion of Conductor task input (map[string]any) into a typed destination value.
type InputBinder interface {
	Bind(dst any, src map[string]any) error
}

// JSONBinder implements InputBinder via JSON marshal/unmarshal.
type JSONBinder struct{}

// Bind converts the provided task input map into the destination typed value using JSON marshal/unmarshal.
// The dst parameter must be a non-nil pointer to the destination type.
func (JSONBinder) Bind(dst any, src map[string]any) error {
	if dst == nil {
		return fmt.Errorf("destination pointer is nil - cannot bind task input")
	}
	raw, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("failed to marshal input data: %w", err)
	}
	return json.Unmarshal(raw, dst)
}
