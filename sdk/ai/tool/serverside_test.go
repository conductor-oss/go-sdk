package tool

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Pinned against the Python SDK's api_tool, index_tool, search_tool and
// wait_for_message_tool, captured in testdata/tools_serverside.json by
// sdk/ai/testdata/agent_config/generate_fixtures.py, in this order.
func TestServerSideToolsMatchPython(t *testing.T) {
	raw, err := os.ReadFile("testdata/tools_serverside.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var want []map[string]any
	if err := json.Unmarshal(raw, &want); err != nil {
		t.Fatal(err)
	}

	got := []ai.ToolDef{
		API("", "https://api.example.test/openapi.json", ""),
		API("stripe", "https://api.stripe.test/openapi.json", "Stripe API.",
			WithHeaders(map[string]string{"Authorization": "Bearer ${STRIPE_KEY}"}),
			WithToolNames("GetCharge", "ListCharges"), WithMaxTools(20), WithCredentials("STRIPE_KEY")),
		Index("index_document", "pgvectordb",
			"product_docs", "openai", "text-embedding-3-small", "Add a document to the knowledge base."),
		Index("index_chunked", "pineconedb",
			"notes", "openai", "text-embedding-3-large", "Index with chunking.",
			WithNamespace("team_a"), WithChunking(512, 64), WithDimensions(3072)),
		Search("search_knowledge_base", "pgvectordb",
			"product_docs", "openai", "text-embedding-3-small", "Search the product documentation."),
		Search("search_notes", "pineconedb",
			"notes", "openai", "text-embedding-3-large", "Search notes.",
			WithNamespace("team_a"), WithMaxResults(3), WithDimensions(3072)),
		WaitForMessage("wait_for_message", "Wait until a message is sent to this agent."),
		WaitForMessage("poll_messages", "Take up to five queued messages without waiting.",
			WithBatchSize(5), NonBlocking()),
	}
	if len(got) != len(want) {
		t.Fatalf("fixture has %d tools, test builds %d", len(want), len(got))
	}
	for i, td := range got {
		w := want[i]
		if w["name"] != td.Name {
			t.Fatalf("fixture[%d] is %v, test builds %s; keep the two in the same order", i, w["name"], td.Name)
		}
		if err := td.Validate(); err != nil {
			t.Errorf("%s: %v", td.Name, err)
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
		// Credentials land under config on the wire; fold them in the same way
		// the serializer does before comparing config.
		cfg := map[string]any{}
		for k, v := range td.Config {
			cfg[k] = v
		}
		if len(td.Credentials) > 0 {
			cfg["credentials"] = td.Credentials
		}
		if g := roundTrip(t, cfg); !reflect.DeepEqual(g, w["config"]) {
			t.Errorf("%s: config differs from Python\n--- go ---\n%s\n--- python ---\n%s",
				td.Name, indent(g), indent(w["config"]))
		}
	}
}

// An API tool whose headers name a credential it does not declare is refused,
// as an MCP tool is: the server would send the placeholder as literal text.
func TestAPIValidatesHeaderCredentials(t *testing.T) {
	td := API("stripe", "https://api.stripe.test/openapi.json", "",
		WithHeaders(map[string]string{"Authorization": "Bearer ${STRIPE_KEY}"}))
	if err := td.Validate(); err == nil || !strings.Contains(err.Error(), "STRIPE_KEY") {
		t.Fatalf("Validate = %v, want a complaint about STRIPE_KEY", err)
	}
	if err := API("", "", "").Validate(); err == nil {
		t.Fatal("an API tool with no url validated")
	}
}
