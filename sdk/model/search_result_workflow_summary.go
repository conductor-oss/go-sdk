//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

// SearchResultWorkflowSummary is the response from the search workflow endpoint.
type SearchResultWorkflowSummary struct {
	// TotalHits the total number of workflows that match the search criteria.
	TotalHits int64 `json:"totalHits,omitempty"`
	// Results the list of workflows that match the search criteria.
	Results []WorkflowSummary `json:"results,omitempty"`
}
