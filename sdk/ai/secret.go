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

// ErrCredentialNotFound means nothing was delivered for that name. It is a
// configuration problem, not a transient one — the tool did not declare the
// credential, or the server has no value for it — so a tool that can carry on
// should report it to the model rather than failing the task.
var ErrCredentialNotFound = errors.New("credential not found")

// Secret reads a credential the server delivered with the current task.
//
// A secured server resolves the names a tool declared when the worker polls and
// attaches the values to the task on a wire-only field, which the runtime moves
// into the task context; nothing is read from the environment and this makes no
// network call. Only names declared with tool.WithCredentials are delivered,
// because Go cannot inspect a function body to discover which it uses.
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
// parent's environment. os.Setenv would be process-wide and racy across
// concurrent tool calls, since workers are goroutines in one process.
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

// deliveredNames lists what did arrive, never values, so a missing-credential
// error can say what the alternatives were.
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
