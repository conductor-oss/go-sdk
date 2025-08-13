//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

// TestDataIn represents a simple typed input payload for typed workers.
type TestDataIn struct {
	A string `json:"a"`
	B int    `json:"b"`
}

// TestDataOut represents a simple typed output payload for typed workers.
type TestDataOut struct {
	Message string `json:"message"`
	Sum     int    `json:"sum"`
	TaskId  string `json:"taskId"`
}
