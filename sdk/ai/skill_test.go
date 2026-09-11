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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// The fixture skills double as golden inputs; see serializer_golden_test.go.
const skillFixtures = "testdata/agent_config/skills"

func fixtureSkill(t *testing.T, name string, opts ...SkillOption) *Agent {
	t.Helper()
	agent, err := LoadSkill(filepath.Join(skillFixtures, name), opts...)
	if err != nil {
		t.Fatalf("LoadSkill(%s): %v", name, err)
	}
	return agent
}

// writeSkill lays out a skill directory from a map of relative path to content.
func writeSkill(t *testing.T, dir string, files map[string]string) string {
	t.Helper()
	for rel, content := range files {
		p := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestLoadSkillDiscovery(t *testing.T) {
	agent := fixtureSkill(t, "review-skill",
		WithSkillModel(testModel),
		WithAgentModels(map[string]string{"critic": "openai/gpt-4o-mini"}),
		WithSkillParams(map[string]any{"rounds": 1, "style": "terse"}))

	if agent.Name != "review-skill" || agent.Model != testModel {
		t.Fatalf("agent identity = %q/%q", agent.Name, agent.Model)
	}
	if err := agent.Validate(); err != nil {
		t.Fatalf("skill agent does not validate: %v", err)
	}
	s := agent.skill

	if got := sortedKeys(s.agentFiles); !reflect.DeepEqual(got, []string{"critic", "defender"}) {
		t.Errorf("agent files = %v", got)
	}
	if !strings.Contains(s.agentFiles["critic"], "You Are the Critic") {
		t.Errorf("critic instructions not read: %q", s.agentFiles["critic"])
	}

	// Scripts in name order, languages by extension then shebang.
	wantScripts := []skillScript{
		{name: "count_lines", filename: "count_lines.sh", language: "bash"},
		{name: "echo_args", filename: "echo_args.py", language: "python"},
		{name: "render", filename: "render", language: "node"},
	}
	if len(s.scripts) != len(wantScripts) {
		t.Fatalf("scripts = %+v", s.scripts)
	}
	for i, want := range wantScripts {
		got := s.scripts[i]
		if got.name != want.name || got.filename != want.filename || got.language != want.language {
			t.Errorf("script %d = %+v, want %+v", i, got, want)
		}
		if !strings.HasSuffix(got.path, filepath.Join("scripts", want.filename)) {
			t.Errorf("script %d path = %q", i, got.path)
		}
	}

	// Subtrees first, sorted, then loose root files; agent files and SKILL.md
	// are not resources.
	wantResources := []string{"references/guide.md", "comic-template.html"}
	if !reflect.DeepEqual(s.resourceFiles, wantResources) {
		t.Errorf("resources = %v, want %v", s.resourceFiles, wantResources)
	}

	// The cross reference resolved from the sibling directory, in the nested
	// shape: no root resources, params equal to defaults, skillSections present.
	ref, ok := s.crossRefs["cleanup-skill"].(map[string]any)
	if !ok {
		t.Fatalf("cross refs = %v", s.crossRefs)
	}
	if !strings.Contains(ref["skillMd"].(string), "# Cleanup Skill") {
		t.Errorf("ref skillMd not read")
	}
	if got := ref["resourceFiles"]; !reflect.DeepEqual(got, []string{"references/notes.md"}) {
		t.Errorf("ref resources = %v; root files must not be listed for a reference", got)
	}
	if got := ref["defaultParams"]; !reflect.DeepEqual(got, map[string]any{"dry_run": false}) {
		t.Errorf("ref defaults = %v", got)
	}
	if !reflect.DeepEqual(ref["params"], ref["defaultParams"]) {
		t.Errorf("ref params must equal its defaults, got %v", ref["params"])
	}
	if got, ok := ref["skillSections"].(map[string]string); !ok || len(got) != 0 {
		t.Errorf("ref skillSections = %v", ref["skillSections"])
	}

	// Params: defaults from frontmatter in file order, override applied, the
	// added key last, and the block appended to SKILL.md.
	if !reflect.DeepEqual(s.defaultParams, map[string]any{"rounds": 3, "verbose": true}) {
		t.Errorf("defaults = %v", s.defaultParams)
	}
	if !reflect.DeepEqual(s.params, map[string]any{"rounds": 1, "verbose": true, "style": "terse"}) {
		t.Errorf("params = %v", s.params)
	}
	wantTail := "\n\n[Skill Parameters]\nrounds: 1\nverbose: True\nstyle: terse\n"
	if !strings.HasSuffix(s.skillMd, wantTail) {
		t.Errorf("skillMd tail = %q", s.skillMd[len(s.skillMd)-min(len(s.skillMd), 80):])
	}
	if !strings.HasPrefix(s.skillMd, "---\nname: review-skill\n") {
		t.Errorf("skillMd must be the verbatim file, got prefix %q", s.skillMd[:30])
	}
}

func TestLoadSkillErrors(t *testing.T) {
	t.Run("missing SKILL.md", func(t *testing.T) {
		_, err := LoadSkill(t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "SKILL.md not found") {
			t.Fatalf("err = %v", err)
		}
	})
	t.Run("frontmatter without name", func(t *testing.T) {
		dir := writeSkill(t, t.TempDir(), map[string]string{
			"SKILL.md": "---\ndescription: no name\n---\n# Body\n",
		})
		_, err := LoadSkill(dir)
		if err == nil || !strings.Contains(err.Error(), "missing required 'name'") {
			t.Fatalf("err = %v", err)
		}
	})
	t.Run("no frontmatter", func(t *testing.T) {
		dir := writeSkill(t, t.TempDir(), map[string]string{"SKILL.md": "# Just a body\n"})
		_, err := LoadSkill(dir)
		if err == nil || !strings.Contains(err.Error(), "missing required 'name'") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestLoadSkills(t *testing.T) {
	skills, err := LoadSkills(skillFixtures,
		WithSkillModel(testModel),
		WithSkillAgentModels(map[string]map[string]string{
			"review-skill": {"critic": "openai/gpt-4o-mini"},
		}))
	if err != nil {
		t.Fatal(err)
	}
	if got := sortedKeys(skills); !reflect.DeepEqual(got, []string{"cleanup-skill", "review-skill"}) {
		t.Fatalf("skills = %v", got)
	}
	if got := skills["review-skill"].skill.agentModels["critic"]; got != "openai/gpt-4o-mini" {
		t.Errorf("per-skill agent model not applied: %q", got)
	}
	if got := skills["cleanup-skill"].skill.agentModels; len(got) != 0 {
		t.Errorf("cleanup-skill must get no overrides, got %v", got)
	}
}

func TestSkillToConfig(t *testing.T) {
	agent := fixtureSkill(t, "cleanup-skill", WithSkillModel(testModel))
	cfg := normalize(t, agent.toConfig())

	if cfg["_framework"] != skillFramework || cfg["name"] != "cleanup-skill" || cfg["model"] != testModel {
		t.Fatalf("identity fields = %v %v %v", cfg["_framework"], cfg["name"], cfg["model"])
	}
	// Empty collections are {} and [], never null: the server reads them
	// with getOrDefault and a null would take the default's place badly.
	for _, key := range []string{"agentModels", "agentFiles", "crossSkillRefs"} {
		if _, ok := cfg[key].(map[string]any); !ok {
			t.Errorf("%s = %v (%T), want an object", key, cfg[key], cfg[key])
		}
	}
	if _, ok := cfg["resourceFiles"].([]any); !ok {
		t.Errorf("resourceFiles = %v (%T), want an array", cfg["resourceFiles"], cfg["resourceFiles"])
	}
	// Scripts on the wire carry filename and language only, never the path.
	tidy := cfg["scripts"].(map[string]any)["tidy"].(map[string]any)
	if !reflect.DeepEqual(tidy, map[string]any{"filename": "tidy.sh", "language": "bash"}) {
		t.Errorf("scripts.tidy = %v", tidy)
	}
	// Frontmatter defaults alone still produce the parameters block.
	if !strings.HasSuffix(cfg["skillMd"].(string), "[Skill Parameters]\ndry_run: False\n") {
		t.Errorf("skillMd tail = %q", cfg["skillMd"])
	}
}

func TestDetectScriptLanguage(t *testing.T) {
	dir := t.TempDir()
	cases := []struct{ file, content, want string }{
		{"a.py", "print(1)", "python"},
		{"b.sh", "echo hi", "bash"},
		{"c.js", "console.log(1)", "node"},
		{"d.rb", "puts 1", "ruby"},
		{"e.PY", "print(1)", "python"},
		{"plain", "echo hi", "bash"},
		{"shebang_py", "#!/usr/bin/env python3\nprint(1)", "python"},
		{"shebang_bash", "#!/bin/bash\necho hi", "bash"},
		{"shebang_node", "#!/usr/bin/env node\n", "node"},
		{"f.txt", "#!/usr/bin/env ruby\n", "ruby"}, // unknown extension falls through to the shebang
	}
	for _, c := range cases {
		p := filepath.Join(dir, c.file)
		if err := os.WriteFile(p, []byte(c.content), 0o755); err != nil {
			t.Fatal(err)
		}
		if got := detectScriptLanguage(p); got != c.want {
			t.Errorf("%s: language = %q, want %q", c.file, got, c.want)
		}
	}
}

func TestPyStemAndSuffix(t *testing.T) {
	cases := []struct{ name, stem, suffix string }{
		{"echo_args.py", "echo_args", ".py"},
		{"archive.tar.gz", "archive.tar", ".gz"},
		{"render", "render", ""},
		{".bashrc", ".bashrc", ""},
		{"notes.", "notes.", ""},
	}
	for _, c := range cases {
		if got := pyStem(c.name); got != c.stem {
			t.Errorf("pyStem(%q) = %q, want %q", c.name, got, c.stem)
		}
		if got := pySuffix(c.name); got != c.suffix {
			t.Errorf("pySuffix(%q) = %q, want %q", c.name, got, c.suffix)
		}
	}
}

func TestSlugify(t *testing.T) {
	cases := map[string]string{
		"Workflow Definitions":   "workflow-definitions",
		"  API  Reference (v2) ": "api-reference-v2",
		"Error--Handling":        "error-handling",
		"-- --":                  "",
		"Configuration Guide!":   "configuration-guide",
	}
	for in, want := range cases {
		if got := slugify(in); got != want {
			t.Errorf("slugify(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSplitIntoSections(t *testing.T) {
	body := "# Title\n\nPreamble text.\n\n## First Part\nalpha\n\n## Second Part\nbeta\n### Sub\n## First Part\ngamma\n"
	order, sections := splitIntoSections(body)
	if !reflect.DeepEqual(order, []string{"first-part", "second-part"}) {
		t.Fatalf("order = %v", order)
	}
	// A repeated heading keeps its position and takes the later content.
	if sections["first-part"] != "## First Part\ngamma" {
		t.Errorf("first-part = %q", sections["first-part"])
	}
	if sections["second-part"] != "## Second Part\nbeta\n### Sub" {
		t.Errorf("second-part = %q", sections["second-part"])
	}
	// A body that opens with a heading has no preamble and no empty part.
	order, _ = splitIntoSections("## Only\nx")
	if !reflect.DeepEqual(order, []string{"only"}) {
		t.Errorf("order = %v", order)
	}
}

// largeSkill writes a skill whose body exceeds the split threshold, with five
// sections and one reference file, mirroring the Python test fixture.
func largeSkill(t *testing.T) string {
	t.Helper()
	var b strings.Builder
	b.WriteString("---\nname: large-skill\ndescription: A large skill.\n---\n")
	b.WriteString("# Large Skill\n\nYou are the orchestrator. Follow these core rules:\n1. Always validate inputs\n2. Never skip error handling\n\n")
	for _, name := range []string{"Workflow Definitions", "Running Workflows", "Error Handling", "API Reference", "Configuration Guide"} {
		fmt.Fprintf(&b, "## %s\n\nThis section covers %s.\n\n", name, strings.ToLower(name))
		for i := 1; i <= 100; i++ {
			fmt.Fprintf(&b, "### Rule %d for %s\n\nWhen handling %s scenario %d, validate inputs, check permissions, execute operation, verify result.\n\n",
				i, name, strings.ToLower(name), i)
		}
	}
	if b.Len() < sectionSplitThreshold {
		t.Fatalf("fixture body too short: %d", b.Len())
	}
	return writeSkill(t, filepath.Join(t.TempDir(), "large-skill"), map[string]string{
		"SKILL.md":            b.String(),
		"references/guide.md": "# Guide\nSome content.",
	})
}

func TestLargeSkillSplitsIntoSections(t *testing.T) {
	agent, err := LoadSkill(largeSkill(t), WithSkillModel(testModel))
	if err != nil {
		t.Fatal(err)
	}
	s := agent.skill
	wantSections := []string{"workflow-definitions", "running-workflows", "error-handling", "api-reference", "configuration-guide"}
	if got := sortedKeys(s.sections); len(got) != len(wantSections) {
		t.Fatalf("sections = %v", got)
	}
	// Virtual section files follow the real resources, in heading order.
	want := append([]string{"references/guide.md"}, prefixed(skillSectionPrefix, wantSections)...)
	if !reflect.DeepEqual(s.resourceFiles, want) {
		t.Errorf("resources = %v, want %v", s.resourceFiles, want)
	}
	if !strings.Contains(s.sections["workflow-definitions"], "Rule 1 for Workflow Definitions") {
		t.Errorf("section content missing")
	}

	// The read worker serves sections and files alike, and refuses the rest.
	read := findWorker(t, s.workers(agent.Name), "large-skill__read_skill_file").Handler.(func(context.Context, skillFileIn) (string, error))
	ctx := context.Background()
	if out, _ := read(ctx, skillFileIn{Path: "skill_section:workflow-definitions"}); !strings.HasPrefix(out, "## Workflow Definitions") {
		t.Errorf("section read = %q", out[:min(len(out), 60)])
	}
	if out, _ := read(ctx, skillFileIn{Path: "references/guide.md"}); !strings.Contains(out, "# Guide") {
		t.Errorf("file read = %q", out)
	}
	if out, _ := read(ctx, skillFileIn{Path: "skill_section:nonexistent"}); !strings.HasPrefix(out, "ERROR") {
		t.Errorf("unknown section = %q", out)
	}
}

func TestSmallSkillHasNoSections(t *testing.T) {
	agent := fixtureSkill(t, "review-skill")
	if len(agent.skill.sections) != 0 {
		t.Errorf("sections = %v", agent.skill.sections)
	}
	for _, r := range agent.skill.resourceFiles {
		if strings.HasPrefix(r, skillSectionPrefix) {
			t.Errorf("unexpected virtual resource %q", r)
		}
	}
}

func TestCrossSkillReferences(t *testing.T) {
	t.Run("unresolved reference is skipped", func(t *testing.T) {
		dir := writeSkill(t, filepath.Join(t.TempDir(), "lonely-skill"), map[string]string{
			"SKILL.md": "---\nname: lonely-skill\ndescription: test\n---\nInvoke the nonexistent-skill-xyz skill.\n",
		})
		agent, err := LoadSkill(dir)
		if err != nil {
			t.Fatal(err)
		}
		if len(agent.skill.crossRefs) != 0 {
			t.Errorf("cross refs = %v", agent.skill.crossRefs)
		}
	})

	t.Run("nested references resolve transitively", func(t *testing.T) {
		root := t.TempDir()
		writeSkill(t, root, map[string]string{
			"parent-skill/SKILL.md":     "---\nname: parent-skill\n---\n# Parent\nUse the child-skill skill.\n",
			"child-skill/SKILL.md":      "---\nname: child-skill\n---\n# Child\nUse the grandchild-skill skill.\n",
			"grandchild-skill/SKILL.md": "---\nname: grandchild-skill\n---\n# Grandchild\n",
		})
		agent, err := LoadSkill(filepath.Join(root, "parent-skill"))
		if err != nil {
			t.Fatal(err)
		}
		child := agent.skill.crossRefs["child-skill"].(map[string]any)
		nested := child["crossSkillRefs"].(map[string]any)
		if _, ok := nested["grandchild-skill"]; !ok {
			t.Errorf("grandchild not resolved through child: %v", nested)
		}
	})

	t.Run("cycle is an error", func(t *testing.T) {
		root := t.TempDir()
		writeSkill(t, root, map[string]string{
			"a-skill/SKILL.md": "---\nname: a-skill\n---\nCall the b-skill skill.\n",
			"b-skill/SKILL.md": "---\nname: b-skill\n---\nCall the a-skill skill.\n",
		})
		_, err := LoadSkill(filepath.Join(root, "a-skill"))
		if err == nil || !strings.Contains(err.Error(), "circular skill reference") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("self reference is ignored", func(t *testing.T) {
		root := t.TempDir()
		writeSkill(t, root, map[string]string{
			"solo-skill/SKILL.md": "---\nname: solo-skill\n---\nUse the solo-skill skill again.\n",
		})
		agent, err := LoadSkill(filepath.Join(root, "solo-skill"))
		if err != nil {
			t.Fatal(err)
		}
		if len(agent.skill.crossRefs) != 0 {
			t.Errorf("cross refs = %v", agent.skill.crossRefs)
		}
	})

	t.Run("search path is consulted after siblings", func(t *testing.T) {
		root := t.TempDir()
		elsewhere := t.TempDir()
		writeSkill(t, root, map[string]string{
			"caller-skill/SKILL.md": "---\nname: caller-skill\n---\nInvoke the remote-skill skill.\n",
		})
		writeSkill(t, elsewhere, map[string]string{
			"remote-skill/SKILL.md": "---\nname: remote-skill\n---\n# Remote\n",
		})
		agent, err := LoadSkill(filepath.Join(root, "caller-skill"), WithSkillSearchPath(elsewhere))
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := agent.skill.crossRefs["remote-skill"]; !ok {
			t.Errorf("remote-skill not found via search path: %v", agent.skill.crossRefs)
		}
	})
}

func TestSkillParams(t *testing.T) {
	dir := writeSkill(t, filepath.Join(t.TempDir(), "param-skill"), map[string]string{
		"SKILL.md": "---\nname: param-skill\ndescription: test\n" +
			"params:\n  rounds:\n    type: integer\n    default: 3\n    description: Number of rounds\n" +
			"  style:\n    type: string\n    default: concise\n  ratio: 0.5\n  weight: 2.0\n---\n# Body",
	})

	t.Run("defaults only", func(t *testing.T) {
		agent, err := LoadSkill(dir)
		if err != nil {
			t.Fatal(err)
		}
		want := map[string]any{"rounds": 3, "style": "concise", "ratio": 0.5, "weight": 2.0}
		if !reflect.DeepEqual(agent.skill.defaultParams, want) {
			t.Errorf("defaults = %v", agent.skill.defaultParams)
		}
		if !reflect.DeepEqual(agent.skill.params, want) {
			t.Errorf("params = %v", agent.skill.params)
		}
		// Floats print as Python does: 2.0 stays 2.0.
		wantBlock := "[Skill Parameters]\nrounds: 3\nstyle: concise\nratio: 0.5\nweight: 2.0\n"
		if !strings.HasSuffix(agent.skill.skillMd, wantBlock) {
			t.Errorf("skillMd tail = %q", agent.skill.skillMd[len(agent.skill.skillMd)-len(wantBlock):])
		}
	})

	t.Run("overrides and additions", func(t *testing.T) {
		agent, err := LoadSkill(dir, WithSkillParams(map[string]any{"rounds": 7, "verbose": true, "audience": "ops"}))
		if err != nil {
			t.Fatal(err)
		}
		if agent.skill.params["rounds"] != 7 || agent.skill.params["style"] != "concise" {
			t.Errorf("params = %v", agent.skill.params)
		}
		// Declared params keep file order; additions follow, sorted.
		wantBlock := "[Skill Parameters]\nrounds: 7\nstyle: concise\nratio: 0.5\nweight: 2.0\naudience: ops\nverbose: True\n"
		if !strings.HasSuffix(agent.skill.skillMd, wantBlock) {
			t.Errorf("skillMd tail = %q", agent.skill.skillMd[len(agent.skill.skillMd)-len(wantBlock):])
		}
	})

	t.Run("no params means no block", func(t *testing.T) {
		agent := fixtureSkill(t, "cleanup-skill")
		plain := writeSkill(t, filepath.Join(t.TempDir(), "plain-skill"), map[string]string{
			"SKILL.md": "---\nname: plain-skill\n---\n# Plain\n",
		})
		agentPlain, err := LoadSkill(plain)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(agentPlain.skill.skillMd, "[Skill Parameters]") {
			t.Errorf("block appended without params")
		}
		if formatSkillParams(nil, nil) != "" {
			t.Errorf("empty params must format to empty string")
		}
		_ = agent
	})
}

func TestPyString(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{true, "True"}, {false, "False"}, {nil, "None"}, {"x", "x"},
		{3, "3"}, {int64(4), "4"}, {2.0, "2.0"}, {0.5, "0.5"}, {1e20, "1e+20"},
	}
	for _, c := range cases {
		if got := pyString(c.in); got != c.want {
			t.Errorf("pyString(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSkillWorkers(t *testing.T) {
	agent := fixtureSkill(t, "review-skill")
	workers := agent.skill.workers(agent.Name)

	var names []string
	for _, w := range workers {
		names = append(names, w.Name)
		if w.Handler == nil {
			t.Errorf("worker %s has no handler", w.Name)
		}
		if _, err := toolExecutor(w); err != nil {
			t.Errorf("worker %s: %v", w.Name, err)
		}
	}
	want := []string{
		"review-skill__count_lines", "review-skill__echo_args", "review-skill__render",
		"review-skill__read_skill_file",
	}
	if !reflect.DeepEqual(names, want) {
		t.Fatalf("workers = %v, want %v", names, want)
	}

	ctx := context.Background()
	read := findWorker(t, workers, "review-skill__read_skill_file").Handler.(func(context.Context, skillFileIn) (string, error))
	if out, _ := read(ctx, skillFileIn{Path: "comic-template.html"}); !strings.Contains(out, "{{PANELS}}") {
		t.Errorf("read = %q", out)
	}
	if out, _ := read(ctx, skillFileIn{Path: "../../etc/passwd"}); !strings.HasPrefix(out, "ERROR: '../../etc/passwd' not found. Available: ['comic-template.html', 'references/guide.md']") {
		t.Errorf("unknown path = %q", out)
	}
	// A listed name that resolves outside the directory is still refused.
	agent.skill.resourceFiles = append(agent.skill.resourceFiles, "../cleanup-skill/SKILL.md")
	if out, _ := read(ctx, skillFileIn{Path: "../cleanup-skill/SKILL.md"}); !strings.Contains(out, "outside the skill directory") {
		t.Errorf("escape = %q", out)
	}

	if _, err := exec.LookPath("python3"); err != nil {
		t.Skip("python3 not on PATH")
	}
	echo := findWorker(t, workers, "review-skill__echo_args").Handler.(func(context.Context, skillScriptIn) (string, error))
	if out, _ := echo(ctx, skillScriptIn{Command: `Conductor "Agents Go"`}); strings.TrimSpace(out) != "ECHO_ARGS_RESULT:Conductor Agents Go" {
		t.Errorf("echo = %q", out)
	}
	if out, _ := echo(ctx, skillScriptIn{}); strings.TrimSpace(out) != "ECHO_ARGS_RESULT:no-args" {
		t.Errorf("echo no args = %q", out)
	}
	count := findWorker(t, workers, "review-skill__count_lines").Handler.(func(context.Context, skillScriptIn) (string, error))
	if out, _ := count(ctx, skillScriptIn{Command: "a b c"}); strings.TrimSpace(out) != "3" {
		t.Errorf("count = %q", out)
	}
}

func TestSkillScriptFailureIsReported(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not on PATH")
	}
	dir := writeSkill(t, filepath.Join(t.TempDir(), "fail-skill"), map[string]string{
		"SKILL.md":        "---\nname: fail-skill\n---\n# Fails\n",
		"scripts/fail.sh": "#!/bin/bash\necho boom >&2\nexit 3\n",
	})
	agent, err := LoadSkill(dir)
	if err != nil {
		t.Fatal(err)
	}
	run := findWorker(t, agent.skill.workers(agent.Name), "fail-skill__fail").Handler.(func(context.Context, skillScriptIn) (string, error))
	out, err := run(context.Background(), skillScriptIn{})
	if err != nil {
		t.Fatalf("a failing script is a result, not a task failure: %v", err)
	}
	if out != "ERROR (exit 3):\nboom\n" {
		t.Errorf("out = %q", out)
	}
	// No resources: no read worker.
	if len(agent.skill.workers(agent.Name)) != 1 {
		t.Errorf("workers = %d, want only the script", len(agent.skill.workers(agent.Name)))
	}
}

func TestSkillStartPayload(t *testing.T) {
	rt := NewRuntimeWithClient(nil, Config{})
	// The runtime would poll for the skill's workers; stub them as started so
	// startPayload can be exercised without a server.
	agent := fixtureSkill(t, "cleanup-skill", WithSkillModel(testModel))
	for _, w := range agent.skill.workers(agent.Name) {
		rt.started[w.Name] = true
	}
	payload, err := rt.startPayload(agent, "tidy up", nil)
	if err != nil {
		t.Fatal(err)
	}
	if payload["framework"] != skillFramework {
		t.Errorf("framework = %v", payload["framework"])
	}
	if _, ok := payload["agentConfig"]; ok {
		t.Errorf("a skill must not be sent as agentConfig")
	}
	raw, ok := payload["rawConfig"].(map[string]any)
	if !ok || raw["model"] != testModel {
		t.Fatalf("rawConfig = %v", payload["rawConfig"])
	}
	if _, has := raw["_framework"]; has {
		t.Errorf("rawConfig carries the framework in the request, not as a key")
	}
	for _, key := range []string{"prompt", "sessionId", "media", "context"} {
		if _, ok := payload[key]; !ok {
			t.Errorf("payload missing %s", key)
		}
	}
}

// ── helpers ─────────────────────────────────────────────────────────

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sortStrings(keys)
	return keys
}

func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func prefixed(prefix string, items []string) []string {
	out := make([]string, len(items))
	for i, s := range items {
		out[i] = prefix + s
	}
	return out
}

func findWorker(t *testing.T, workers []ToolDef, name string) ToolDef {
	t.Helper()
	for _, w := range workers {
		if w.Name == name {
			return w
		}
	}
	t.Fatalf("no worker %q in %d workers", name, len(workers))
	return ToolDef{}
}
