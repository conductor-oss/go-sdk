package agents

import (
	"context"
	"io"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

// CredentialsHTTPTool is the Python SDK's examples/agents/16e_credentials_http_tool.py:
// an HTTP tool whose Authorization header names a credential the server
// resolves. GITHUB_REPOS_URL points the tool at a fixture instead of GitHub.
func CredentialsHTTPTool(model string) *ai.Agent {
	url := os.Getenv("GITHUB_REPOS_URL")
	if url == "" {
		url = "https://api.github.com/users/Conductor/repos?per_page=5&sort=updated"
	}
	listRepos := tool.HTTP("list_github_repos", url,
		"List public GitHub repositories for a user. Returns JSON array with name, url, and stars.",
		tool.WithHeaders(map[string]string{
			"Authorization": "Bearer ${GITHUB_TOKEN}",
			"Accept":        "application/vnd.github.v3+json",
		}),
		tool.WithCredentials("GITHUB_TOKEN"),
	)
	return &ai.Agent{
		Name:         "github_http_agent",
		Model:        model,
		Tools:        ai.Tools(listRepos),
		Instructions: "You list GitHub repos using the list_github_repos tool. Summarize the results.",
	}
}

func runCredentialsHTTPTool(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, CredentialsHTTPTool(Model()), "List the repos for Conductor", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
