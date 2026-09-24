// Package agents holds the Go ports of the Python SDK's examples/agents, one
// file per example. Each example has a Build function returning the agents it
// runs and a run function that plays the example the way the Python one does,
// reading approvals from in and printing to out. Catalog lists them for the
// command in ./cmd and for the playback test.
package agents

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

// Example is one runnable example, named after its Python source file.
type Example struct {
	Name  string
	Build func(model string) []*ai.Agent
	Run   func(ctx context.Context, rt *ai.Runtime, in io.Reader, out io.Writer) ([]*ai.AgentResult, error)
}

// Catalog lists every example in file order.
var Catalog = []Example{
	{Name: "01_basic_agent", Build: one(BasicAgent), Run: runBasicAgent},
	{Name: "02a_simple_tools", Build: one(SimpleTools), Run: runSimpleTools},
	{Name: "02c_tool_retry_config", Build: one(ToolRetryConfig), Run: runToolRetryConfig},
	{Name: "04_http_and_mcp_tools", Build: one(HTTPAndMCPTools), Run: runHTTPAndMCPTools},
	{Name: "05_handoffs", Build: one(Handoffs), Run: runHandoffs},
	{Name: "06_sequential_pipeline", Build: one(SequentialPipeline), Run: runSequentialPipeline},
	{Name: "07_parallel_agents", Build: one(ParallelAgents), Run: runParallelAgents},
	{Name: "09_human_in_the_loop", Build: one(HumanInTheLoop), Run: runHumanInTheLoop},
	{Name: "09c_hitl_streaming", Build: one(HITLStreaming), Run: runHITLStreaming},
	{Name: "103_plan_and_compile", Build: one(PlanAndCompile), Run: runPlanAndCompile},
	{Name: "10_guardrails", Build: one(Guardrails), Run: runGuardrails},
	{Name: "13_hierarchical_agents", Build: one(HierarchicalAgents), Run: runHierarchicalAgents},
	{Name: "16e_credentials_http_tool", Build: one(CredentialsHTTPTool), Run: runCredentialsHTTPTool},
	{Name: "17_swarm_orchestration", Build: one(SwarmOrchestration), Run: runSwarmOrchestration},
	{Name: "21_regex_guardrails", Build: RegexGuardrails, Run: runRegexGuardrails},
	{Name: "22_llm_guardrails", Build: one(LLMGuardrails), Run: runLLMGuardrails},
	{Name: "33_external_workers", Build: one(ExternalWorkers), Run: runExternalWorkers},
	{Name: "64_swarm_with_tools", Build: one(SwarmWithTools), Run: runSwarmWithTools},
	{Name: "66_handoff_to_parallel", Build: one(HandoffToParallel), Run: runHandoffToParallel},
}

// Lookup returns the example with that name.
func Lookup(name string) (Example, bool) {
	for _, ex := range Catalog {
		if ex.Name == name {
			return ex, true
		}
	}
	return Example{}, false
}

// Model is the model every example uses: CONDUCTOR_AGENT_LLM_MODEL, or the
// Python examples' default.
func Model() string {
	if m := os.Getenv("CONDUCTOR_AGENT_LLM_MODEL"); m != "" {
		return m
	}
	return "openai/gpt-4o-mini"
}

func one(build func(model string) *ai.Agent) func(model string) []*ai.Agent {
	return func(model string) []*ai.Agent { return []*ai.Agent{build(model)} }
}

// run starts one agent, waits for it and prints the outcome. A run the server
// finished as FAILED is an outcome, not an error: the example reports it the
// way the Python one does and carries on, and the playback check decides
// whether the failure was a legitimate guardrail rejection.
func run(ctx context.Context, rt *ai.Runtime, agent *ai.Agent, prompt string, out io.Writer) (*ai.AgentResult, error) {
	res, err := rt.Run(ctx, agent, prompt)
	return report(out, res, err)
}

func report(out io.Writer, res *ai.AgentResult, err error) (*ai.AgentResult, error) {
	if err != nil {
		if res == nil {
			return nil, err
		}
		fmt.Fprintf(out, "Execution ended: %s\n", res.Error)
		return res, nil
	}
	fmt.Fprintln(out, res.Output)
	fmt.Fprintf(out, "[%s] %s\n", res.Status, res.ExecutionID)
	return res, nil
}

// runWithApproval is run for an agent whose tools pause for a human. Events
// are printed as they stream. Each pause is answered the way the Python
// example answers it: one line from in per field of the pending tool's
// response schema, "y" or "yes" for a boolean field. The reviewer's answers
// become a message in the conversation, so they are part of what the
// recording has to match.
func runWithApproval(ctx context.Context, rt *ai.Runtime, agent *ai.Agent, prompt string, in io.Reader, out io.Writer) (*ai.AgentResult, error) {
	handle, err := rt.Start(ctx, agent, prompt)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(out, "Started: %s\n", handle.ExecutionID)

	streamCtx, stopStream := context.WithCancel(ctx)
	defer stopStream()
	events, err := handle.Events(streamCtx)
	if err != nil {
		return nil, err
	}
	go func() {
		for ev := range events {
			fmt.Fprintf(out, "[%s] %s\n", ev.Type, ev.Text)
		}
	}()

	answers := bufio.NewReader(in)
	for {
		status, statusErr := handle.Status(ctx)
		if statusErr != nil {
			return nil, statusErr
		}
		waiting, waitErr := handle.Waiting(ctx)
		if waitErr != nil {
			return nil, waitErr
		}
		if waiting {
			fmt.Fprintln(out, "\n--- Human input required ---")
			response, askErr := askReviewer(status, answers, out)
			if askErr != nil {
				return nil, askErr
			}
			if respondErr := handle.Respond(ctx, response); respondErr != nil {
				return nil, respondErr
			}
		} else if status.Status.Terminal() {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
		}
	}
	res, err := handle.Result(ctx)
	return report(out, res, err)
}

// askReviewer prompts for each field of the pending tool's response schema,
// in field-name order since the status document arrives as a Go map.
func askReviewer(status *ai.AgentResult, answers *bufio.Reader, out io.Writer) (map[string]any, error) {
	pending, _ := status.Raw["pendingTool"].(map[string]any)
	schema, _ := pending["response_schema"].(map[string]any)
	props, _ := schema["properties"].(map[string]any)
	fields := make([]string, 0, len(props))
	for field := range props {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	response := map[string]any{}
	for _, field := range fields {
		spec, _ := props[field].(map[string]any)
		desc, _ := spec["description"].(string)
		if desc == "" {
			if desc, _ = spec["title"].(string); desc == "" {
				desc = field
			}
		}
		boolean := spec["type"] == "boolean"
		if boolean {
			fmt.Fprintf(out, "  %s (y/n): ", desc)
		} else {
			fmt.Fprintf(out, "  %s: ", desc)
		}
		line, err := answers.ReadString('\n')
		if err != nil && line == "" {
			return nil, fmt.Errorf("no answer supplied for %s: %w", field, err)
		}
		answer := strings.TrimSpace(line)
		if boolean {
			lower := strings.ToLower(answer)
			response[field] = lower == "y" || lower == "yes"
		} else {
			response[field] = answer
		}
	}
	return response, nil
}
