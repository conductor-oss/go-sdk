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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

// Agent Skills (agentskills.io) loaded as agents.
//
// A skill directory becomes its own agent: SKILL.md is the orchestrator's
// instructions, each *-agent.md file is a sub-agent it can call, each file in
// scripts/ is a worker tool, and everything else is a resource the agent reads
// on demand through a read_skill_file tool. The skill runs as a separate
// workflow — used as a tool, a parent agent receives its result, never its
// instructions. That is the delegation model the Agent Skills integration
// guide describes as optional, and it is the one every Conductor SDK uses.
//
// The SDK does no compiling. It reads the directory into the raw document
// below and the server's SkillNormalizer turns that into an agentConfig. The
// document is the cross-SDK contract, so its shape and quirks follow the
// Python SDK's skill.py exactly; the golden fixture 18_skill pins it.

const (
	skillFramework = "skill"

	// sectionSplitThreshold is the SKILL.md body length, in characters, above
	// which the body is split into ## sections the agent loads on demand.
	// Python measures len(str), so this counts code points, not bytes.
	sectionSplitThreshold = 50000

	skillSectionPrefix = "skill_section:"
	skillScriptTimeout = 300 * time.Second
)

var (
	// frontmatterRe matches the YAML block between --- fences and captures the
	// body after it. It is the union of Python's two patterns, with the same
	// \s* tolerance around the fences.
	frontmatterRe = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*\n(.*)`)

	// crossSkillRe finds prose like "invoke the writing-plans skill". The name
	// is looked up as a sibling directory, so a match that resolves to nothing
	// is ignored rather than reported.
	crossSkillRe = regexp.MustCompile(`(?i)(?:invoke|use|call)\s+(?:the\s+)?([a-z][a-z0-9-]*)\s+skill`)

	slugDropRe  = regexp.MustCompile(`[^a-z0-9\s-]`)
	slugSpaceRe = regexp.MustCompile(`\s+`)
	slugDashRe  = regexp.MustCompile(`-+`)
)

// scriptLanguages maps a script's extension to the language the worker runs
// it with; scriptInterpreters maps that language to the interpreter binary.
var scriptLanguages = map[string]string{
	".py": "python", ".sh": "bash", ".js": "node", ".mjs": "node", ".ts": "node", ".rb": "ruby",
}

// shebangLanguages is checked in order against the first line of a script
// with no recognised extension. The order matters: "sh" would otherwise
// claim "bash".
var shebangLanguages = []struct{ key, lang string }{
	{"python", "python"}, {"python3", "python"}, {"bash", "bash"},
	{"sh", "bash"}, {"node", "node"}, {"ruby", "ruby"},
}

var scriptInterpreters = map[string]string{
	"python": "python3", "bash": "bash", "node": "node", "ruby": "ruby",
}

// SkillOption adjusts how LoadSkill and LoadSkills read a skill directory.
type SkillOption func(*skillOptions)

type skillOptions struct {
	model          string
	agentModels    map[string]string
	perSkillModels map[string]map[string]string
	searchPath     []string
	params         map[string]any
}

// WithSkillModel sets the model of the skill's orchestrator agent. Sub-agents
// declared in *-agent.md files use it too unless WithAgentModels overrides
// them.
func WithSkillModel(model string) SkillOption {
	return func(o *skillOptions) { o.model = model }
}

// WithAgentModels overrides the model per sub-agent, keyed by the name derived
// from its file: "gilfoyle-agent.md" is "gilfoyle".
func WithAgentModels(models map[string]string) SkillOption {
	return func(o *skillOptions) { o.agentModels = models }
}

// WithSkillAgentModels is WithAgentModels for LoadSkills. The outer key is the
// skill's directory name, so each skill gets its own overrides; it mirrors the
// agent_models argument of Python's load_skills.
func WithSkillAgentModels(perSkill map[string]map[string]string) SkillOption {
	return func(o *skillOptions) { o.perSkillModels = perSkill }
}

// WithSkillSearchPath adds directories searched for skills that this skill's
// instructions reference. They are searched after the skill's own parent
// directory and the .agents/skills directories under the working directory
// and the home directory.
func WithSkillSearchPath(dirs ...string) SkillOption {
	return func(o *skillOptions) { o.searchPath = append(o.searchPath, dirs...) }
}

// WithSkillParams overrides the defaults of the params the SKILL.md
// frontmatter declares, and may add params it does not declare. The merged
// values are appended to the instructions as a [Skill Parameters] block.
func WithSkillParams(params map[string]any) SkillOption {
	return func(o *skillOptions) { o.params = params }
}

// LoadSkill reads an Agent Skills directory and returns it as an Agent.
//
// The directory must contain a SKILL.md with a name in its frontmatter. The
// agent can be run, exposed as a tool with tool.Agent, or placed in another
// agent's Agents list like any other; its Name is the skill's name.
//
//	dg, err := ai.LoadSkill("~/.claude/skills/dg",
//	    ai.WithSkillModel("openai/gpt-4o"),
//	    ai.WithAgentModels(map[string]string{"gilfoyle": "anthropic/claude-sonnet-4-6"}))
func LoadSkill(path string, opts ...SkillOption) (*Agent, error) {
	var o skillOptions
	for _, opt := range opts {
		opt(&o)
	}
	dir := resolveSkillPath(path)
	agentModels := o.agentModels
	if agentModels == nil {
		agentModels = o.perSkillModels[filepath.Base(dir)]
	}
	return loadSkill(dir, o.model, agentModels, o.searchPath, o.params)
}

// LoadSkills loads every skill directory directly under dir, keyed by
// directory name. Options apply to each skill; use WithSkillAgentModels for
// per-skill sub-agent models.
func LoadSkills(dir string, opts ...SkillOption) (map[string]*Agent, error) {
	root := resolveSkillPath(dir)
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("read skills directory %s: %w", root, err)
	}
	skills := map[string]*Agent{}
	for _, e := range entries {
		sub := filepath.Join(root, e.Name())
		if !isDir(sub) || !fileExists(filepath.Join(sub, "SKILL.md")) {
			continue
		}
		agent, err := LoadSkill(sub, opts...)
		if err != nil {
			return nil, err
		}
		skills[e.Name()] = agent
	}
	return skills, nil
}

// skillConfig is what LoadSkill read, kept on the Agent so the serializer can
// emit the raw document and the runtime can register the skill's workers.
type skillConfig struct {
	dir           string
	model         string
	agentModels   map[string]string
	skillMd       string // full SKILL.md, with the [Skill Parameters] block appended
	agentFiles    map[string]string
	scripts       []skillScript
	resourceFiles []string
	crossRefs     map[string]any
	defaultParams map[string]any
	params        map[string]any
	sections      map[string]string
}

type skillScript struct {
	name     string // file stem; the tool is "<skill>__<name>"
	filename string
	language string
	path     string
}

func loadSkill(dir, model string, agentModels map[string]string, searchPath []string, overrides map[string]any) (*Agent, error) {
	skillMdPath := filepath.Join(dir, "SKILL.md")
	raw, err := os.ReadFile(skillMdPath) //nolint:gosec // path is inside the skill directory
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("directory %s is not a valid skill: SKILL.md not found", dir)
		}
		return nil, fmt.Errorf("read %s: %w", skillMdPath, err)
	}
	skillMd := string(raw)

	fm, err := parseSkillFrontmatter(skillMd)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", skillMdPath, err)
	}
	if fm.name == "" {
		return nil, fmt.Errorf("%s: SKILL.md missing required 'name' field in frontmatter", skillMdPath)
	}

	defaults, order := fm.paramDefaults()
	params, paramOrder := mergeSkillParams(defaults, order, overrides)

	agentFiles, err := readAgentFiles(dir)
	if err != nil {
		return nil, err
	}
	scripts, err := discoverScripts(dir)
	if err != nil {
		return nil, err
	}
	resources, err := listResourceFiles(dir, true)
	if err != nil {
		return nil, err
	}
	crossRefs, err := resolveCrossSkills(skillMd, dir, searchPath, nil)
	if err != nil {
		return nil, err
	}

	// A body too long to hold in context is replaced server side by a table
	// of contents; the sections become virtual files the agent reads on
	// demand, listed alongside the real resources.
	sections := map[string]string{}
	if utf8.RuneCountInString(fm.body) > sectionSplitThreshold {
		var names []string
		names, sections = splitIntoSections(fm.body)
		for _, n := range names {
			resources = append(resources, skillSectionPrefix+n)
		}
	}

	// The parameters are appended to SKILL.md itself so the orchestrator sees
	// them in its system prompt however the skill is invoked.
	if len(params) > 0 {
		skillMd = skillMd + "\n\n" + formatSkillParams(params, paramOrder) + "\n"
	}

	if agentModels == nil {
		agentModels = map[string]string{}
	}
	cfg := &skillConfig{
		dir:           dir,
		model:         model,
		agentModels:   agentModels,
		skillMd:       skillMd,
		agentFiles:    agentFiles,
		scripts:       scripts,
		resourceFiles: resources,
		crossRefs:     crossRefs,
		defaultParams: defaults,
		params:        params,
		sections:      sections,
	}
	return &Agent{Name: fm.name, Model: model, skill: cfg}, nil
}

// rawConfig is the document the server's SkillNormalizer consumes. Every key
// is always present, with empty maps and lists rather than nulls, because
// that is what the Python SDK sends.
func (s *skillConfig) rawConfig() map[string]any {
	return map[string]any{
		"model":          s.model,
		"agentModels":    s.agentModels,
		"skillMd":        s.skillMd,
		"agentFiles":     s.agentFiles,
		"scripts":        scriptsWire(s.scripts),
		"resourceFiles":  s.resourceFiles,
		"crossSkillRefs": s.crossRefs,
		"defaultParams":  s.defaultParams,
		"params":         s.params,
	}
}

// wireConfig is rawConfig in the position of an agentConfig: as the document
// for /agent/start, or nested under an agent tool. The _framework marker is
// what tells the server to normalize it rather than read it as agentConfig.
func (s *skillConfig) wireConfig(name string) map[string]any {
	cfg := s.rawConfig()
	cfg["name"] = name
	cfg["_framework"] = skillFramework
	return cfg
}

func scriptsWire(scripts []skillScript) map[string]any {
	out := map[string]any{}
	for _, sc := range scripts {
		out[sc.name] = map[string]any{"filename": sc.filename, "language": sc.language}
	}
	return out
}

// ── SKILL.md parsing ────────────────────────────────────────────────

type skillFrontmatter struct {
	name   string
	params *yaml.Node // the params mapping, in file order; nil when absent
	body   string     // markdown after the frontmatter, trimmed
}

// parseSkillFrontmatter splits SKILL.md into frontmatter and body. A file
// with no frontmatter is not an error here — a referenced skill may lack
// one — but frontmatter without a name is, as in Python.
func parseSkillFrontmatter(content string) (skillFrontmatter, error) {
	m := frontmatterRe.FindStringSubmatch(content)
	if m == nil {
		return skillFrontmatter{body: strings.TrimSpace(content)}, nil
	}
	fm := skillFrontmatter{body: strings.TrimSpace(m[2])}

	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(m[1]), &doc); err != nil {
		return fm, fmt.Errorf("SKILL.md frontmatter: %w", err)
	}
	mapping := documentMapping(&doc)
	if mapping == nil {
		return fm, errors.New("SKILL.md missing required 'name' field in frontmatter")
	}
	for i := 0; i+1 < len(mapping.Content); i += 2 {
		key, val := mapping.Content[i], mapping.Content[i+1]
		switch key.Value {
		case "name":
			if val.Kind == yaml.ScalarNode {
				fm.name = val.Value
			}
		case "params":
			if val.Kind == yaml.MappingNode {
				fm.params = val
			}
		}
	}
	if fm.name == "" {
		return fm, errors.New("SKILL.md missing required 'name' field in frontmatter")
	}
	return fm, nil
}

func documentMapping(doc *yaml.Node) *yaml.Node {
	if doc.Kind != yaml.DocumentNode || len(doc.Content) == 0 {
		return nil
	}
	if doc.Content[0].Kind != yaml.MappingNode {
		return nil
	}
	return doc.Content[0]
}

// paramDefaults reads the frontmatter's params. Each entry is either a
// mapping with a "default" key or a bare value that is its own default. The
// order is the file's, because it decides the order of the [Skill
// Parameters] lines.
func (fm skillFrontmatter) paramDefaults() (map[string]any, []string) {
	defaults := map[string]any{}
	var order []string
	if fm.params == nil {
		return defaults, order
	}
	for i := 0; i+1 < len(fm.params.Content); i += 2 {
		key := fm.params.Content[i].Value
		var val any
		if err := fm.params.Content[i+1].Decode(&val); err != nil {
			continue
		}
		if m, ok := val.(map[string]any); ok {
			if d, has := m["default"]; has {
				val = d
			}
		}
		if _, dup := defaults[key]; !dup {
			order = append(order, key)
		}
		defaults[key] = val
	}
	return defaults, order
}

// mergeSkillParams lays overrides over the defaults. Declared params keep
// their file order; params only the override names follow, sorted, since a
// Go map has no order to preserve.
func mergeSkillParams(defaults map[string]any, order []string, overrides map[string]any) (map[string]any, []string) {
	merged := map[string]any{}
	mergedOrder := append([]string(nil), order...)
	for _, k := range order {
		merged[k] = defaults[k]
	}
	var extra []string
	for k, v := range overrides {
		if _, declared := merged[k]; !declared {
			extra = append(extra, k)
		}
		merged[k] = v
	}
	sort.Strings(extra)
	return merged, append(mergedOrder, extra...)
}

// formatSkillParams renders the block appended to SKILL.md.
func formatSkillParams(params map[string]any, order []string) string {
	if len(params) == 0 {
		return ""
	}
	lines := make([]string, 0, len(order))
	for _, k := range order {
		lines = append(lines, k+": "+pyString(params[k]))
	}
	return "[Skill Parameters]\n" + strings.Join(lines, "\n")
}

// pyString formats a value the way Python's str() does, so the prompt text
// matches across SDKs: True rather than true, 1.0 rather than 1. Lists and
// maps fall back to Go formatting; skills declare scalar params.
func pyString(v any) string {
	switch x := v.(type) {
	case nil:
		return "None"
	case bool:
		if x {
			return "True"
		}
		return "False"
	case string:
		return x
	case float32:
		return pyString(float64(x))
	case float64:
		if x == math.Trunc(x) && math.Abs(x) < 1e16 {
			return strconv.FormatFloat(x, 'f', 1, 64)
		}
		return strconv.FormatFloat(x, 'g', -1, 64)
	default:
		return fmt.Sprint(x)
	}
}

// ── directory conventions ───────────────────────────────────────────

// readAgentFiles collects *-agent.md files as sub-agent instructions, keyed
// by the name before the suffix.
func readAgentFiles(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read skill directory %s: %w", dir, err)
	}
	out := map[string]string{}
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, "-agent.md") || !isRegularFile(filepath.Join(dir, name)) {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, name)) //nolint:gosec // path is inside the skill directory
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", name, err)
		}
		out[strings.TrimSuffix(name, "-agent.md")] = string(data)
	}
	return out, nil
}

// discoverScripts lists scripts/ in name order. Two scripts with the same
// stem collapse to one entry, the later name winning, as a Python dict would.
func discoverScripts(dir string) ([]skillScript, error) {
	scriptsDir := filepath.Join(dir, "scripts")
	entries, err := os.ReadDir(scriptsDir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("read %s: %w", scriptsDir, err)
	}
	out := make([]skillScript, 0, len(entries))
	index := map[string]int{}
	for _, e := range entries {
		p := filepath.Join(scriptsDir, e.Name())
		if !isRegularFile(p) {
			continue
		}
		sc := skillScript{
			name:     pyStem(e.Name()),
			filename: e.Name(),
			language: detectScriptLanguage(p),
			path:     p,
		}
		if i, dup := index[sc.name]; dup {
			out[i] = sc
			continue
		}
		index[sc.name] = len(out)
		out = append(out, sc)
	}
	return out, nil
}

// detectScriptLanguage decides how a script is run: by extension, else by
// shebang, else as bash.
func detectScriptLanguage(path string) string {
	if lang, ok := scriptLanguages[strings.ToLower(pySuffix(filepath.Base(path)))]; ok {
		return lang
	}
	if data, err := os.ReadFile(path); err == nil { //nolint:gosec // a script under the skill directory
		first, _, _ := strings.Cut(string(data), "\n")
		if strings.HasPrefix(first, "#!") {
			for _, sb := range shebangLanguages {
				if strings.Contains(first, sb.key) {
					return sb.lang
				}
			}
		}
	}
	return "bash"
}

// listResourceFiles lists what read_skill_file may serve: everything under
// references/, examples/ and assets/, each subtree sorted, then — for the
// skill itself, not for a referenced one — the loose files in the root other
// than SKILL.md, the agent files and the skill manifests.
func listResourceFiles(dir string, includeRoot bool) ([]string, error) {
	out := []string{}
	for _, sub := range []string{"references", "examples", "assets"} {
		root := filepath.Join(dir, sub)
		if !isDir(root) {
			continue
		}
		found, err := walkResourceDir(dir, root)
		if err != nil {
			return nil, err
		}
		out = append(out, found...)
	}
	if !includeRoot {
		return out, nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read skill directory %s: %w", dir, err)
	}
	for _, e := range entries {
		name := e.Name()
		if name == "SKILL.md" || strings.HasSuffix(name, "-agent.md") ||
			name == "skill.yaml" || name == "skill.toml" {
			continue
		}
		if isRegularFile(filepath.Join(dir, name)) {
			out = append(out, name)
		}
	}
	return out, nil
}

// walkResourceDir lists the regular files under root, as sorted paths
// relative to dir with forward slashes.
func walkResourceDir(dir, root string) ([]string, error) {
	var found []string
	err := filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !isRegularFile(p) {
			return nil
		}
		rel, err := filepath.Rel(dir, p)
		if err != nil {
			return err
		}
		found = append(found, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list %s: %w", root, err)
	}
	sort.Strings(found)
	return found, nil
}

// ── cross-skill references ──────────────────────────────────────────

// resolveCrossSkills finds skills the body names and loads each as a nested
// raw document, so the server can compile it into an agent tool. A name that
// resolves nowhere is skipped. A skill that, through any chain of references,
// names one of its ancestors is a cycle and an error.
func resolveCrossSkills(skillMd, skillDir string, searchPath []string, ancestors map[string]bool) (map[string]any, error) {
	body := skillBody(skillMd)
	refs := map[string]any{}
	names := map[string]bool{}
	for _, m := range crossSkillRe.FindAllStringSubmatch(body, -1) {
		names[m[1]] = true
	}
	if len(names) == 0 {
		return refs, nil
	}
	sorted := make([]string, 0, len(names))
	for n := range names {
		sorted = append(sorted, n)
	}
	sort.Strings(sorted)

	self := resolvePath(skillDir)
	seen := map[string]bool{self: true}
	for k := range ancestors {
		seen[k] = true
	}
	dirs := skillSearchDirs(skillDir, searchPath)
	for _, name := range sorted {
		for _, d := range dirs {
			refDir := filepath.Join(d, name)
			if !fileExists(filepath.Join(refDir, "SKILL.md")) {
				continue
			}
			refResolved := resolvePath(refDir)
			if refResolved == self {
				continue
			}
			if seen[refResolved] {
				return nil, fmt.Errorf("circular skill reference detected: %s", name)
			}
			ref, err := loadCrossSkillRef(refResolved, searchPath, seen)
			if err != nil {
				return nil, err
			}
			refs[name] = ref
			break
		}
	}
	return refs, nil
}

// loadCrossSkillRef reads a referenced skill into the nested document shape.
// It differs from the top-level document in three ways Python has and this
// keeps: params carry no overrides, the root's loose files are not listed as
// resources, and skillSections is included.
func loadCrossSkillRef(refDir string, searchPath []string, ancestors map[string]bool) (map[string]any, error) {
	refMdPath := filepath.Join(refDir, "SKILL.md")
	raw, err := os.ReadFile(refMdPath) //nolint:gosec // path is inside a skill directory
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", refMdPath, err)
	}
	refMd := string(raw)
	fm, err := parseSkillFrontmatter(refMd)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", refMdPath, err)
	}
	defaults, _ := fm.paramDefaults()
	agentFiles, err := readAgentFiles(refDir)
	if err != nil {
		return nil, err
	}
	scripts, err := discoverScripts(refDir)
	if err != nil {
		return nil, err
	}
	resources, err := listResourceFiles(refDir, false)
	if err != nil {
		return nil, err
	}
	sections := map[string]string{}
	if utf8.RuneCountInString(fm.body) > sectionSplitThreshold {
		var names []string
		names, sections = splitIntoSections(fm.body)
		for _, n := range names {
			resources = append(resources, skillSectionPrefix+n)
		}
	}
	nextAncestors := map[string]bool{refDir: true}
	for k := range ancestors {
		nextAncestors[k] = true
	}
	nested, err := resolveCrossSkills(refMd, refDir, searchPath, nextAncestors)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"skillMd":        refMd,
		"agentFiles":     agentFiles,
		"scripts":        scriptsWire(scripts),
		"resourceFiles":  resources,
		"crossSkillRefs": nested,
		"defaultParams":  defaults,
		"params":         defaults,
		"skillSections":  sections,
	}, nil
}

// skillSearchDirs lists where a referenced skill may live, in lookup order.
func skillSearchDirs(skillDir string, searchPath []string) []string {
	dirs := make([]string, 0, len(searchPath)+3)
	if parent := filepath.Dir(skillDir); fileExists(parent) {
		dirs = append(dirs, parent)
	}
	if cwd, err := os.Getwd(); err == nil {
		dirs = append(dirs, filepath.Join(cwd, ".agents", "skills"))
	}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, filepath.Join(home, ".agents", "skills"))
	}
	for _, p := range searchPath {
		dirs = append(dirs, resolveSkillPath(p))
	}
	return dirs
}

// ── large SKILL.md bodies ───────────────────────────────────────────

// splitIntoSections cuts the body at each line that starts a ## heading and
// keys the pieces by the heading's slug, in order. Text before the first
// heading is the preamble and is not a section.
func splitIntoSections(body string) ([]string, map[string]string) {
	var parts []string
	start := 0
	for i := 0; i < len(body); {
		if i != start && strings.HasPrefix(body[i:], "## ") {
			parts = append(parts, body[start:i])
			start = i
		}
		nl := strings.IndexByte(body[i:], '\n')
		if nl < 0 {
			break
		}
		i += nl + 1
	}
	parts = append(parts, body[start:])

	var order []string
	sections := map[string]string{}
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if !strings.HasPrefix(trimmed, "## ") {
			continue
		}
		first, _, _ := strings.Cut(trimmed, "\n")
		slug := slugify(strings.TrimSpace(first[3:]))
		if slug == "" {
			continue
		}
		if _, dup := sections[slug]; !dup {
			order = append(order, slug)
		}
		sections[slug] = trimmed
	}
	return order, sections
}

// slugify turns a heading into a section name: lower case, spaces to
// hyphens, everything else dropped.
func slugify(text string) string {
	slug := slugDropRe.ReplaceAllString(strings.ToLower(text), "")
	slug = slugSpaceRe.ReplaceAllString(strings.TrimSpace(slug), "-")
	slug = slugDashRe.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

// ── workers ─────────────────────────────────────────────────────────

// skillScriptIn is what the model sends a script tool: one command line of
// arguments, split the way a shell would.
type skillScriptIn struct {
	Command string `json:"command"`
}

// skillFileIn names a resource, relative to the skill directory, or a
// skill_section:<name> virtual file.
type skillFileIn struct {
	Path string `json:"path"`
}

// workers returns the tools the runtime must serve for this skill: one per
// script and, when there is anything to read, read_skill_file. The names are
// what the server's SkillNormalizer emits, so the runtime registers workers
// under exactly these.
//
// Failures a model can act on — a script that exits non-zero, a path that is
// not in the skill — come back as ERROR strings rather than task failures,
// matching the Python workers.
func (s *skillConfig) workers(skillName string) []ToolDef {
	out := make([]ToolDef, 0, len(s.scripts)+1)
	for _, sc := range s.scripts {
		out = append(out, ToolDef{
			Name:        skillName + "__" + sc.name,
			Description: fmt.Sprintf("Run %s script from %s skill", sc.name, skillName),
			Handler:     sc.run,
		})
	}
	if len(s.resourceFiles) > 0 {
		out = append(out, ToolDef{
			Name:        skillName + "__read_skill_file",
			Description: fmt.Sprintf("Read resource files from %s skill", skillName),
			Handler:     s.readFile,
		})
	}
	return out
}

// run executes the script with the interpreter for its language.
func (sc skillScript) run(ctx context.Context, in skillScriptIn) (string, error) {
	var args []string
	if in.Command != "" {
		var err error
		if args, err = shellSplit(in.Command); err != nil {
			return "ERROR: " + err.Error(), nil
		}
	}
	interpreter, ok := scriptInterpreters[sc.language]
	if !ok {
		interpreter = scriptInterpreters["bash"]
	}

	ctx, cancel := context.WithTimeout(ctx, skillScriptTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, interpreter, append([]string{sc.path}, args...)...) //nolint:gosec // running the skill's own script is the tool
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return fmt.Sprintf("ERROR: Script execution timed out (%.0fs)", skillScriptTimeout.Seconds()), nil
	}
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return fmt.Sprintf("ERROR (exit %d):\n%s", exitErr.ExitCode(), stderr.String()), nil
		}
		return "ERROR: " + err.Error(), nil
	}
	return stdout.String(), nil
}

// readFile serves one resource. Only listed files are readable, and a path
// that escapes the skill directory is refused even if it is listed.
func (s *skillConfig) readFile(_ context.Context, in skillFileIn) (string, error) {
	allowed := false
	for _, f := range s.resourceFiles {
		if f == in.Path {
			allowed = true
			break
		}
	}
	if !allowed {
		return fmt.Sprintf("ERROR: '%s' not found. Available: %s", in.Path, pyList(s.resourceFiles)), nil
	}
	if name, ok := strings.CutPrefix(in.Path, skillSectionPrefix); ok {
		if content, ok := s.sections[name]; ok {
			return content, nil
		}
		return fmt.Sprintf("ERROR: section '%s' not found", name), nil
	}

	root := resolvePath(s.dir)
	target := resolvePath(filepath.Join(s.dir, filepath.FromSlash(in.Path)))
	if target != root && !strings.HasPrefix(target, root+string(filepath.Separator)) {
		return fmt.Sprintf("ERROR: '%s' is outside the skill directory", in.Path), nil
	}
	data, err := os.ReadFile(target) //nolint:gosec // target was checked to be inside the skill directory
	if err != nil {
		return fmt.Sprintf("ERROR reading '%s': %s", in.Path, err), nil
	}
	return string(data), nil
}

// pyList renders names as Python's sorted(list) repr, which is what the
// Python worker puts in its error message.
func pyList(items []string) string {
	sorted := append([]string(nil), items...)
	sort.Strings(sorted)
	quoted := make([]string, len(sorted))
	for i, s := range sorted {
		quoted[i] = "'" + s + "'"
	}
	return "[" + strings.Join(quoted, ", ") + "]"
}

// ── path helpers ────────────────────────────────────────────────────

// skillBody is the markdown after the frontmatter, or the whole file when
// there is none.
func skillBody(skillMd string) string {
	if m := frontmatterRe.FindStringSubmatch(skillMd); m != nil {
		return strings.TrimSpace(m[2])
	}
	return skillMd
}

// resolveSkillPath expands a leading ~ and resolves the path, following
// symlinks when it exists, like Python's expanduser().resolve().
func resolveSkillPath(p string) string {
	if p == "~" || strings.HasPrefix(p, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			p = filepath.Join(home, p[1:])
		}
	}
	return resolvePath(p)
}

func resolvePath(p string) string {
	abs, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	if real, err := filepath.EvalSymlinks(abs); err == nil {
		return real
	}
	return abs
}

// pySuffix and pyStem follow pathlib: a leading dot or a trailing dot is not
// a suffix, so ".bashrc" and "notes." keep their whole name as the stem.
func pySuffix(name string) string {
	i := strings.LastIndex(name, ".")
	if i > 0 && i < len(name)-1 {
		return name[i:]
	}
	return ""
}

func pyStem(name string) string {
	return strings.TrimSuffix(name, pySuffix(name))
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func isDir(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

// isRegularFile follows symlinks, as pathlib's is_file does.
func isRegularFile(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.Mode().IsRegular()
}
