package agents

import (
	"context"
	"io"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// HierarchicalAgents is the Python SDK's examples/agents/13_hierarchical_agents.py:
// a CEO routes to department leads, who route to specialists.
func HierarchicalAgents(model string) *ai.Agent {
	backendDev := &ai.Agent{
		Name:         "backend_dev",
		Model:        model,
		Instructions: "You are a backend developer. You design APIs, databases, and server architecture. Provide technical recommendations with code examples.",
	}
	frontendDev := &ai.Agent{
		Name:         "frontend_dev",
		Model:        model,
		Instructions: "You are a frontend developer. You design UI components, user flows, and client-side architecture. Provide recommendations with code examples.",
	}
	contentWriter := &ai.Agent{
		Name:         "content_writer",
		Model:        model,
		Instructions: "You are a content writer. You create blog posts, landing page copy, and marketing materials. Write engaging, clear content.",
	}
	seoSpecialist := &ai.Agent{
		Name:         "seo_specialist",
		Model:        model,
		Instructions: "You are an SEO specialist. You optimize content for search engines, suggest keywords, and improve page rankings.",
	}
	engineeringLead := &ai.Agent{
		Name:         "engineering_lead",
		Model:        model,
		Instructions: "You are the engineering lead. Route technical questions to the right specialist: backend_dev for APIs/databases/servers, frontend_dev for UI/UX/client-side.",
		Agents:       []*ai.Agent{backendDev, frontendDev},
		Strategy:     ai.StrategyHandoff,
	}
	marketingLead := &ai.Agent{
		Name:         "marketing_lead",
		Model:        model,
		Instructions: "You are the marketing lead. Route marketing questions to the right specialist: content_writer for blog posts/copy, seo_specialist for SEO/keywords/rankings.",
		Agents:       []*ai.Agent{contentWriter, seoSpecialist},
		Strategy:     ai.StrategyHandoff,
	}
	return &ai.Agent{
		Name:         "ceo",
		Model:        model,
		Instructions: "You are the CEO. Route requests to the right department: engineering_lead for technical/development questions, marketing_lead for marketing/content/SEO questions.",
		Agents:       []*ai.Agent{engineeringLead, marketingLead},
		Handoffs: []ai.HandoffCondition{
			&ai.OnTextMention{Text: "engineering_lead", Target: "engineering_lead"},
			&ai.OnTextMention{Text: "marketing_lead", Target: "marketing_lead"},
		},
		Strategy: ai.StrategySwarm,
	}
}

func runHierarchicalAgents(ctx context.Context, rt *ai.Runtime, _ io.Reader, out io.Writer) ([]*ai.AgentResult, error) {
	res, err := run(ctx, rt, HierarchicalAgents(Model()),
		"Design a REST API for a user management system with authentication, then ask the marketing team for a campaign to promote it.", out)
	if err != nil {
		return nil, err
	}
	return []*ai.AgentResult{res}, nil
}
