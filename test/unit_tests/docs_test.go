package unit_tests

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests guard the documentation information architecture this repository
// adopted from the Java and Python SDKs. See
// docs/adr/0001-canonical-documentation-ia.md for why the structure looks the way
// it does.
//
// Link *liveness* is not checked here -- the lychee action in
// .github/workflows/docs.yml does that, and does it better than a hand-rolled
// walker. These tests assert structure: which pages exist, what stays a stub,
// what is reachable, and what must not appear at all.

const (
	repoRoot = "../.."
	docsDir  = repoRoot + "/docs"
	docsHub  = docsDir + "/README.md"
)

// canonicalPages are the cross-SDK page names this repository has adopted. A
// reader or tool arriving from another Conductor SDK expects to find these names.
var canonicalPages = []string{
	"README.md",
	"compatibility.md",
	"connection-authentication.md",
	"core-quickstart.md",
	"documentation-parity.md",
	"documentation-standard.md",
	"observability.md",
	"security.md",
	"upgrading.md",
	"workers.md",
	"workflows.md",
}

// redirectStubs were retired by the move to canonical filenames but are kept as
// pointer-only files, because conductor-oss/conductor links to every one of them
// from docs/documentation/clientsdks/go-sdk.md and deleting them would 404 the
// published OSS documentation. They are deliberately unreachable from the hub,
// and must not grow back into guides while they wait to be removed.
//
// All six are deleted together once conductor-oss/conductor#1417 repoints those
// links. Do not add to this list for any other reason: a stub is a bridge for an
// external referrer, not a way to keep an internal link alive.
var redirectStubs = []string{
	"workers_sdk.md",
	"workflow_sdk.md",
	"migration_guide.md",
	"logger_sdk.md",
	"api_client/README.md",
	"api_client/tls_configuration.md",
}

// retiredDocPaths must never be the target of a Markdown link.
//
// This is asserted against link targets only, never against raw file contents:
// docs/adr/0001-canonical-documentation-ia.md names these files in prose while
// describing the renames, which is correct and must not fail the build.
var retiredDocPaths = []string{
	"docs/workers_sdk.md",
	"docs/workflow_sdk.md",
	"docs/migration_guide.md",
	"docs/metrics.md",
	"docs/logger_sdk.md",
	"docs/api_client/README.md",
	"docs/api_client/proxy_configuration.md",
	"docs/api_client/tls_configuration.md",
}

var markdownLink = regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`)

// markdownFiles returns every Markdown file in the repository.
func markdownFiles(t *testing.T) []string {
	t.Helper()
	var found []string
	err := filepath.Walk(repoRoot, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(info.Name(), ".md") {
			found = append(found, path)
		}
		return nil
	})
	require.NoError(t, err)
	require.NotEmpty(t, found, "found no Markdown files; is docs/ present in the test image?")
	return found
}

// linkTargets returns the repo-relative target of every local Markdown link in
// file. External URLs, in-page anchors and mailto: links are skipped, and any
// trailing #fragment is trimmed.
func linkTargets(t *testing.T, file string) []string {
	t.Helper()
	body, err := os.ReadFile(file)
	require.NoErrorf(t, err, "could not read %s", file)

	var targets []string
	for _, match := range markdownLink.FindAllStringSubmatch(string(body), -1) {
		target := strings.TrimSpace(match[1])
		if target == "" || strings.Contains(target, "://") ||
			strings.HasPrefix(target, "#") || strings.HasPrefix(target, "mailto:") {
			continue
		}
		if hash := strings.IndexByte(target, '#'); hash >= 0 {
			target = target[:hash]
		}
		if target == "" {
			continue
		}
		rel, relErr := filepath.Rel(repoRoot, filepath.Join(filepath.Dir(file), target))
		require.NoError(t, relErr)
		targets = append(targets, filepath.ToSlash(rel))
	}
	return targets
}

func TestCanonicalDocumentationPagesExist(t *testing.T) {
	for _, page := range canonicalPages {
		_, err := os.Stat(filepath.Join(docsDir, page))
		assert.NoErrorf(t, err,
			"canonical page docs/%s is missing; see docs/adr/0001-canonical-documentation-ia.md", page)
	}
}

func TestRedirectStubsRemainPointersOnly(t *testing.T) {
	const maxStubLines = 20

	for _, stub := range redirectStubs {
		path := filepath.Join(docsDir, stub)

		body, err := os.ReadFile(path)
		require.NoErrorf(t, err,
			"redirect stub docs/%s is missing; conductor-oss/conductor links to that path", stub)
		content := string(body)

		lines := strings.Count(strings.TrimSpace(content), "\n") + 1
		assert.LessOrEqualf(t, lines, maxStubLines,
			"redirect stub docs/%s has grown to %d lines; it must stay a pointer, not become a guide", stub, lines)
		assert.Containsf(t, content, "redirect stub",
			"docs/%s must say it is a redirect stub, so the next reader knows not to add content", stub)
		assert.NotContainsf(t, content, "```",
			"redirect stub docs/%s contains a code block, which means it is turning back into a guide", stub)
		assert.NotEmptyf(t, linkTargets(t, path),
			"redirect stub docs/%s must link to its replacement", stub)
	}
}

func TestNoLinksToRetiredDocumentationPaths(t *testing.T) {
	retired := make(map[string]bool, len(retiredDocPaths))
	for _, path := range retiredDocPaths {
		retired[path] = true
	}

	for _, file := range markdownFiles(t) {
		for _, target := range linkTargets(t, file) {
			assert.Falsef(t, retired[target],
				"%s links to retired path %s; use the canonical page instead", file, target)
		}
	}
}

func TestEveryDocumentationPageIsReachableFromHub(t *testing.T) {
	linked := make(map[string]bool)
	for _, target := range linkTargets(t, docsHub) {
		linked[target] = true
	}

	isStub := make(map[string]bool, len(redirectStubs))
	for _, stub := range redirectStubs {
		isStub[filepath.ToSlash(filepath.Join("docs", stub))] = true
	}

	pages, err := filepath.Glob(filepath.Join(docsDir, "*.md"))
	require.NoError(t, err)

	for _, page := range pages {
		rel, relErr := filepath.Rel(repoRoot, page)
		require.NoError(t, relErr)
		rel = filepath.ToSlash(rel)

		if rel == "docs/README.md" || isStub[rel] {
			continue
		}
		assert.Truef(t, linked[rel],
			"%s is not linked from docs/README.md; an orphan page is invisible, which is how the "+
				"documentation drifted out of parity in the first place", rel)
	}
}

func TestNoAgentDocumentationTree(t *testing.T) {
	_, err := os.Stat(filepath.Join(docsDir, "agents"))
	assert.Truef(t, os.IsNotExist(err),
		"docs/agents/ exists, but this SDK provides no Conductor agent runtime. The absence is recorded "+
			"in docs/compatibility.md and docs/documentation-parity.md rather than as an empty page tree.")

	for _, file := range markdownFiles(t) {
		for _, target := range linkTargets(t, file) {
			assert.NotContainsf(t, target, "docs/agents/",
				"%s links into docs/agents/, which does not exist in this SDK", file)
		}
	}
}
