package agents

import (
	"testing"
)

func TestEveryExampleBuildsAValidAgent(t *testing.T) {
	seen := map[string]bool{}
	for _, ex := range Catalog {
		if seen[ex.Name] {
			t.Errorf("%s is listed twice", ex.Name)
		}
		seen[ex.Name] = true
		built := ex.Build("openai/gpt-4o-mini")
		if len(built) == 0 {
			t.Errorf("%s builds no agents", ex.Name)
		}
		for _, agent := range built {
			if err := agent.Validate(); err != nil {
				t.Errorf("%s: %s: %v", ex.Name, agent.Name, err)
			}
		}
	}
}
