package tool

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// The media tools' wire shape is pinned against what the Python SDK's
// image_tool, audio_tool, video_tool and pdf_tool produce, captured in
// testdata/tools_media.json by sdk/ai/testdata/agent_config/generate_fixtures.py.
// The default schemas are long, and a stray field or default would compile
// to a different tool on the server, so they are compared whole.
func TestMediaToolsMatchPython(t *testing.T) {
	raw, err := os.ReadFile("testdata/tools_media.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var want []map[string]any
	if err := json.Unmarshal(raw, &want); err != nil {
		t.Fatal(err)
	}

	got := []ai.ToolDef{
		Image("ks_image", "Generate image", "openai", "dall-e-3"),
		Audio("ks_audio", "Generate audio", "openai", "tts-1"),
		Video("ks_video", "Generate video", "openai", "sora"),
		PDF("ks_pdf", "Generate PDF"),
	}
	if len(got) != len(want) {
		t.Fatalf("fixture has %d tools, test builds %d", len(want), len(got))
	}
	for i, td := range got {
		w := want[i]
		if w["name"] != td.Name {
			t.Fatalf("fixture[%d] is %v, test builds %s; keep the two in the same order", i, w["name"], td.Name)
		}
		if w["description"] != td.Description {
			t.Errorf("%s: description = %q, want %q", td.Name, td.Description, w["description"])
		}
		if w["toolType"] != string(td.ToolType) {
			t.Errorf("%s: toolType = %q, want %v", td.Name, td.ToolType, w["toolType"])
		}
		if g := roundTrip(t, td.InputSchema); !reflect.DeepEqual(g, w["inputSchema"]) {
			t.Errorf("%s: inputSchema differs from Python\n--- go ---\n%s\n--- python ---\n%s",
				td.Name, indent(g), indent(w["inputSchema"]))
		}
		if g := roundTrip(t, td.Config); !reflect.DeepEqual(g, w["config"]) {
			t.Errorf("%s: config differs from Python\n--- go ---\n%s\n--- python ---\n%s",
				td.Name, indent(g), indent(w["config"]))
		}
	}
}

// Options still apply: a custom schema replaces the default, and static
// generation parameters land in config next to the provider and model.
func TestMediaToolOptions(t *testing.T) {
	custom := map[string]any{"type": "object", "properties": map[string]any{"prompt": map[string]any{"type": "string"}}}
	td := Image("art", "Draw.", "openai", "dall-e-3", WithInputSchema(custom), WithConfig("n", 2))
	if !reflect.DeepEqual(td.InputSchema, custom) {
		t.Errorf("WithInputSchema did not replace the default schema")
	}
	if td.Config["n"] != 2 || td.Config["model"] != "dall-e-3" || td.Config["taskType"] != "GENERATE_IMAGE" {
		t.Errorf("config = %v", td.Config)
	}
}

// roundTrip normalizes a Go value through JSON so it compares against decoded
// fixture content: ints become float64, typed slices become []any.
func roundTrip(t *testing.T, v any) any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

func indent(v any) string {
	raw, _ := json.MarshalIndent(v, "", "  ")
	return string(raw)
}
