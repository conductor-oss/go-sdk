//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package ai

import (
	"context"
	"errors"
	"fmt"
	"sort"
)

// ErrCredentialNotFound means nothing was delivered for that name.
//
// It is a configuration problem, not a transient one: either the tool did not
// declare the credential, or the server's store has no value for it. Retrying
// will not help, so a tool that can carry on should report it to the model
// rather than returning the error and failing the task.
var ErrCredentialNotFound = errors.New("credential not found")

// Secret reads a credential the server delivered with the current task.
//
// This is a lookup, not a fetch. A secured server resolves the names a tool
// declared when the worker polls, and attaches the values to the task on a
// wire-only field; the runtime moves them into the task context before calling
// the handler. So there is no network call here and no error but absence.
//
// Only names declared with tool.WithCredentials are delivered — that
// declaration is what the server resolves against, which is why it exists
// separately from this read. Go cannot inspect a function body to discover
// which credentials it uses.
//
//	func openPR(ctx context.Context, in PRIn) (PRResult, error) {
//	    token, err := ai.Secret(ctx, "GH_TOKEN")
//	    if err != nil {
//	        return PRResult{Error: err.Error()}, nil
//	    }
//	    ...
//	}
func Secret(ctx context.Context, name string) (string, error) {
	task, ok := taskFromContext(ctx)
	if !ok {
		return "", fmt.Errorf("%w: %q — Secret is only usable inside a tool handler",
			ErrCredentialNotFound, name)
	}
	if v, ok := task.RuntimeMetadata[name]; ok && v != "" {
		return v, nil
	}
	return "", fmt.Errorf(
		"%w: %q was not delivered with this task. Declare it with "+
			"tool.WithCredentials(%q) and store it on the server; delivered names: %v",
		ErrCredentialNotFound, name, name, deliveredNames(task.RuntimeMetadata))
}

// SecretsEnv returns the named credentials as KEY=VALUE strings for a child
// process, so a tool that shells out can pass them without touching the
// parent's environment.
//
//	cmd := exec.CommandContext(ctx, "gh", "issue", "create", "--title", title)
//	env, err := ai.SecretsEnv(ctx, "GH_TOKEN")
//	cmd.Env = append(os.Environ(), env...)
//
// os.Setenv would be process-wide and racy across concurrent tool calls, since
// workers are goroutines in one process. Scoping the values to cmd.Env is the
// safe equivalent of what the Python SDK does with a subprocess.
func SecretsEnv(ctx context.Context, names ...string) ([]string, error) {
	env := make([]string, 0, len(names))
	for _, n := range names {
		v, err := Secret(ctx, n)
		if err != nil {
			return nil, err
		}
		env = append(env, n+"="+v)
	}
	return env, nil
}

// deliveredNames lists what did arrive, so a missing-credential error says what
// the alternative was. Values are never included.
func deliveredNames(m map[string]string) []string {
	if len(m) == 0 {
		return []string{}
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
