//go:build ignore

// Hierarchical Agents — a CEO routes to department leads, who route to specialists.
//
// Run with:  go run agents/13_hierarchical_agents.go
//
// Three levels: the CEO hands off to the engineering or marketing lead by
// mentioning them (swarm strategy with text-mention handoffs); each lead
// routes to one of its specialists (handoff strategy).
//
// Requirements:
//   - Conductor server with LLM support
//   - CONDUCTOR_SERVER_URL=http://localhost:8080/api in the environment
//   - CONDUCTOR_AGENT_LLM_MODEL in the environment (optional)
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

func main() {
	model := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL")
	if model == "" {
		model = "openai/gpt-4o"
	}

	backendDev := &ai.Agent{
		Name:  "backend_dev",
		Model: model,
		Instructions: "You are a backend developer. You design APIs, databases, and server " +
			"architecture. Provide technical recommendations with code examples.",
	}
	frontendDev := &ai.Agent{
		Name:  "frontend_dev",
		Model: model,
		Instructions: "You are a frontend developer. You design UI components, user flows, " +
			"and client-side architecture. Provide recommendations with code examples.",
	}
	contentWriter := &ai.Agent{
		Name:  "content_writer",
		Model: model,
		Instructions: "You are a content writer. You create blog posts, landing page copy, " +
			"and marketing materials. Write engaging, clear content.",
	}
	seoSpecialist := &ai.Agent{
		Name:  "seo_specialist",
		Model: model,
		Instructions: "You are an SEO specialist. You optimize content for search engines, " +
			"suggest keywords, and improve page rankings.",
	}
	engineeringLead := &ai.Agent{
		Name:  "engineering_lead",
		Model: model,
		Instructions: "You are the engineering lead. Route technical questions to the right " +
			"specialist: backend_dev for APIs/databases/servers, " +
			"frontend_dev for UI/UX/client-side.",
		Agents:   []*ai.Agent{backendDev, frontendDev},
		Strategy: ai.StrategyHandoff,
	}
	marketingLead := &ai.Agent{
		Name:  "marketing_lead",
		Model: model,
		Instructions: "You are the marketing lead. Route marketing questions to the right " +
			"specialist: content_writer for blog posts/copy, " +
			"seo_specialist for SEO/keywords/rankings.",
		Agents:   []*ai.Agent{contentWriter, seoSpecialist},
		Strategy: ai.StrategyHandoff,
	}
	ceo := &ai.Agent{
		Name:  "ceo",
		Model: model,
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

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	fmt.Println("--- Technical question (CEO -> Engineering -> Backend) ---")
	result, err := runtime.Run(context.Background(), ceo,
		"Design a REST API for a user management system with authentication, "+
			"then ask the marketing team for a campaign to promote it.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
