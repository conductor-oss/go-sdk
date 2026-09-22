//go:build integration

//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// Credentials in an HTTP tool — the Python SDK's
// examples/agents/16e_credentials_http_tool.py as a test.
//
// An HTTP tool carries "Bearer ${GITHUB_TOKEN}" in a header and declares the
// credential; the server resolves it and makes the call, so no worker runs
// here. This is that flow, copied, with the print replaced by validation: the
// run completes and the model summarized what the call returned.
//
// The server needs the credential: start it with CONDUCTOR_SECRET_GITHUB_TOKEN
// set. Without it the run does not complete and the test skips, unless
// CONDUCTOR_E2E_SECRET_PROVISIONED says the secret should be there.
func TestExample16eCredentialsHTTPTool(t *testing.T) {
	runtime := newRuntime(t)
	if model(t) == mockModel {
		// The HTTP call is real even in playback, and the whole GitHub
		// response is part of the next model request. It carries headers
		// tied to the caller's token and live data in its body, so only
		// the recording author's token can match, and only until the data
		// changes. This example is not a replay fixture; see README, "What
		// has to match".
		t.Skip("calls a live API with a personal token; the response differs per token and over time, so it cannot replay")
	}

	listRepos := tool.HTTP("list_github_repos",
		"https://api.github.com/users/Conductor/repos?per_page=5&sort=updated",
		"List public GitHub repositories for a user. Returns JSON array with name, url, and stars.",
		tool.WithHeaders(map[string]string{
			"Authorization": "Bearer ${GITHUB_TOKEN}",
			"Accept":        "application/vnd.github.v3+json",
		}),
		tool.WithCredentials("GITHUB_TOKEN"),
	)
	agent := &ai.Agent{
		Name:         "github_http_agent",
		Model:        model(t),
		Tools:        ai.Tools(listRepos),
		Instructions: "You list GitHub repos using the list_github_repos tool. Summarize the results.",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, agent, "List the repos for Conductor")
	if err != nil || result.Status != ai.StatusCompleted {
		msg := fmt.Sprintf("err=%v", err)
		if result != nil {
			msg = fmt.Sprintf("status=%q error=%q", result.Status, result.Error)
		}
		if os.Getenv("CONDUCTOR_E2E_SECRET_PROVISIONED") == "" {
			t.Skipf("run did not complete (%s); the tool needs GITHUB_TOKEN on the server — "+
				"start it with CONDUCTOR_SECRET_GITHUB_TOKEN set, "+
				"and set CONDUCTOR_E2E_SECRET_PROVISIONED to make this a failure", msg)
		}
		t.Fatalf("run did not complete: %s", msg)
	}

	// Validation, in place of the example's result.print_result(). The
	// summary is model prose, so only require that there is one.
	if strings.TrimSpace(result.Output) == "" {
		t.Error("the run completed with an empty summary")
	}
}
