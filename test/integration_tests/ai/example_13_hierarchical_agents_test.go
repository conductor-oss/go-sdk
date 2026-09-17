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
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Hierarchical agents — the Python SDK's examples/agents/13_hierarchical_agents.py
// as a test.
//
// A CEO hands a two-part request to the engineering lead, who routes it to a
// backend developer, then on to the marketing lead, who routes to a content
// writer; twelve model calls in all. This is that flow, copied, with the
// print replaced by validation: the run completes and the CEO's final answer
// is the recorded one. Every one of the twelve recorded calls has to match
// for that to happen, which shows both SDKs drove the hierarchy the same way.
func TestExample13HierarchicalAgents(t *testing.T) {
	runtime := newRuntime(t)
	recorded := recordedAnswers(t, "13_hierarchical_agents")

	backendDev := &ai.Agent{
		Name:  "backend_dev",
		Model: mockModel,
		Instructions: "You are a backend developer. You design APIs, databases, and server " +
			"architecture. Provide technical recommendations with code examples.",
	}
	frontendDev := &ai.Agent{
		Name:  "frontend_dev",
		Model: mockModel,
		Instructions: "You are a frontend developer. You design UI components, user flows, " +
			"and client-side architecture. Provide recommendations with code examples.",
	}
	contentWriter := &ai.Agent{
		Name:  "content_writer",
		Model: mockModel,
		Instructions: "You are a content writer. You create blog posts, landing page copy, " +
			"and marketing materials. Write engaging, clear content.",
	}
	seoSpecialist := &ai.Agent{
		Name:  "seo_specialist",
		Model: mockModel,
		Instructions: "You are an SEO specialist. You optimize content for search engines, " +
			"suggest keywords, and improve page rankings.",
	}
	engineeringLead := &ai.Agent{
		Name:  "engineering_lead",
		Model: mockModel,
		Instructions: "You are the engineering lead. Route technical questions to the right " +
			"specialist: backend_dev for APIs/databases/servers, " +
			"frontend_dev for UI/UX/client-side.",
		Agents:   []*ai.Agent{backendDev, frontendDev},
		Strategy: ai.StrategyHandoff,
	}
	marketingLead := &ai.Agent{
		Name:  "marketing_lead",
		Model: mockModel,
		Instructions: "You are the marketing lead. Route marketing questions to the right " +
			"specialist: content_writer for blog posts/copy, " +
			"seo_specialist for SEO/keywords/rankings.",
		Agents:   []*ai.Agent{contentWriter, seoSpecialist},
		Strategy: ai.StrategyHandoff,
	}
	ceo := &ai.Agent{
		Name:  "ceo",
		Model: mockModel,
		Instructions: "You are the CEO. Route requests to the right department: " +
			"engineering_lead for technical/development questions, " +
			"marketing_lead for marketing/content/SEO questions.",
		Agents: []*ai.Agent{engineeringLead, marketingLead},
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "engineering_lead", Target: "engineering_lead"},
			&ai.OnTextMention{Text: "marketing_lead", Target: "marketing_lead"},
		},
		Strategy: ai.StrategySwarm,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	result, err := runtime.Run(ctx, ceo,
		"Design a REST API for a user management system with authentication, "+
			"then ask the marketing team for a campaign to promote it.")
	if err != nil {
		t.Fatalf("run failed: %v", err)
	}

	// Validation, in place of the example's result.print_result().
	if result.Status != ai.StatusCompleted {
		t.Fatalf("status = %q, want %q (error=%q)", result.Status, ai.StatusCompleted, result.Error)
	}
	if got, want := strings.TrimSpace(result.Output), strings.TrimSpace(recorded[len(recorded)-1]); got != want {
		t.Errorf("output is not the recorded answer\n--- got ---\n%.300s\n--- recorded ---\n%.300s", got, want)
	}
}
