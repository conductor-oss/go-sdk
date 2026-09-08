"""Generate agent-config golden files from the Python SDK's serializer.

These files are the cross-SDK wire contract: the Go serializer must produce the
same agentConfig document that Python (and Java) produce for the same agent.
Comparison is semantic (parsed JSON deep-equal), not byte-for-byte, because
key ordering legitimately differs between languages.

Run from a checkout of conductor-oss/python-sdk:

    python3 <path-to-this-file> --out <path-to-this-directory>

No server or network is needed; AgentConfigSerializer is pure.
"""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional

from pydantic import BaseModel

from conductor.ai.agents import (
    Agent,
    CliConfig,
    CodeExecutionConfig,
    ConversationMemory,
    Guardrail,
    LLMGuardrail,
    MaxMessageTermination,
    OnCondition,
    OnFail,
    OnTextMention,
    OnToolResult,
    Position,
    PromptTemplate,
    RegexGuardrail,
    StopMessageTermination,
    Strategy,
    TextMentionTermination,
    TokenUsageTermination,
    agent_tool,
    guardrail,
    http_tool,
    human_tool,
    mcp_tool,
    tool,
)
from conductor.ai.agents.config_serializer import AgentConfigSerializer

MODEL = "openai/gpt-4o"


# ── tools ───────────────────────────────────────────────────────────────


@tool
def get_weather(city: str, days: int = 3) -> dict:
    """Get the current weather for a city."""
    return {}


@tool
def scalar_kinds(
    s: str, i: int, f: float, b: bool, opt: Optional[str] = None
) -> str:
    """Every scalar type the schema generator maps."""
    return ""


@tool
def container_kinds(names: List[str], meta: Dict[str, int]) -> dict:
    """List and dict parameters, to pin items/additionalProperties."""
    return {}


@tool(approval_required=True, timeout_seconds=45, max_calls=2)
def refund(order_id: str) -> str:
    """Issue a refund. Carries per-tool wire options."""
    return ""


@tool(credentials=["GITHUB_TOKEN"])
def push_branch(branch: str) -> str:
    """Declares a credential, which lands under config.credentials."""
    return ""


# ── callables referenced by name on the wire ────────────────────────────


def stop_when_done(output: str) -> bool:
    """Serializes to {"taskName": "<agent>_stop_when"}."""
    return "DONE" in output


def pick_specialist(prompt: str) -> str:
    """Serializes to {"taskName": "<agent>_router_fn"}."""
    return "billing"


def needs_escalation(state: Dict[str, Any]) -> bool:
    return True


@guardrail
def no_secrets(text: str) -> bool:
    """Custom guardrail, serializes as guardrailType=custom + taskName."""
    return "sk-" not in text


class Ticket(BaseModel):
    summary: str
    priority: int
    tags: List[str]


# ── fixtures ────────────────────────────────────────────────────────────


def fixtures() -> Dict[str, Agent]:
    billing = Agent(name="billing", model=MODEL, instructions="Handle billing.")
    refunds = Agent(name="refunds", model=MODEL, instructions="Handle refunds.")
    tech = Agent(name="tech", model=MODEL, instructions="Handle tech support.")

    out: Dict[str, Agent] = {}

    out["01_minimal"] = Agent(
        name="minimal", model=MODEL, instructions="You are a helpful assistant."
    )

    out["02_tools_worker"] = Agent(
        name="tools_worker",
        model=MODEL,
        instructions="Use tools.",
        tools=[get_weather, refund, push_branch],
    )

    out["03_tools_schema_types"] = Agent(
        name="tools_schema_types",
        model=MODEL,
        instructions="Pin the JSON Schema mapping.",
        tools=[scalar_kinds, container_kinds],
    )

    out["04_strategy_sequential"] = Agent(
        name="pipeline",
        model=MODEL,
        instructions="Run in order.",
        agents=[billing, refunds],
        strategy=Strategy.SEQUENTIAL,
    )

    out["05_strategy_parallel"] = Agent(
        name="fanout",
        model=MODEL,
        agents=[billing, refunds, tech],
        strategy=Strategy.PARALLEL,
    )

    out["06_strategy_router_agent"] = Agent(
        name="router_agent",
        model=MODEL,
        agents=[billing, tech],
        strategy=Strategy.ROUTER,
        router=Agent(name="classifier", model=MODEL, instructions="Classify."),
    )

    out["07_strategy_router_fn"] = Agent(
        name="router_fn",
        model=MODEL,
        agents=[billing, tech],
        strategy=Strategy.ROUTER,
        router=pick_specialist,
    )

    out["08_handoffs"] = Agent(
        name="swarm",
        model=MODEL,
        agents=[billing, refunds, tech],
        strategy=Strategy.SWARM,
        handoffs=[
            OnToolResult(target="refunds", tool_name="refund", result_contains="ok"),
            OnTextMention(target="tech", text="broken"),
            OnCondition(target="billing", condition=needs_escalation),
        ],
        allowed_transitions={"billing": ["refunds"], "refunds": ["tech"]},
    )

    out["09_guardrails"] = Agent(
        name="guarded",
        model=MODEL,
        instructions="Be careful.",
        guardrails=[
            RegexGuardrail([r"\d{16}", r"sk-\w+"], mode="block", message="no cards"),
            LLMGuardrail(model=MODEL, policy="No medical advice.", max_tokens=256),
            Guardrail(no_secrets, position=Position.OUTPUT, on_fail=OnFail.RETRY),
        ],
    )

    out["10_termination"] = Agent(
        name="terminating",
        model=MODEL,
        termination=(
            TextMentionTermination("DONE", case_sensitive=True)
            | MaxMessageTermination(20)
        )
        & TokenUsageTermination(max_total_tokens=8000, max_completion_tokens=2000),
        stop_when=stop_when_done,
    )

    out["11_output_type"] = Agent(
        name="structured",
        model=MODEL,
        instructions="Return a ticket.",
        output_type=Ticket,
    )

    out["12_llm_knobs"] = Agent(
        name="knobs",
        model=MODEL,
        instructions=PromptTemplate(
            name="support_prompt", variables={"tier": "${workflow.input.tier}"}, version=2
        ),
        max_turns=7,
        max_tokens=4096,
        temperature=0.2,
        reasoning_effort="high",
        thinking_budget_tokens=2048,
        context_window_budget=100000,
        timeout_seconds=600,
        memory=ConversationMemory(max_messages=25),
        include_contents="none",
    )

    out["13_plan_execute"] = Agent(
        name="planner_root",
        model=MODEL,
        instructions="Plan then execute.",
        strategy=Strategy.PLAN_EXECUTE,
        tools=[get_weather, container_kinds],
        planner=Agent(name="the_planner", model=MODEL, instructions="Emit JSON plan."),
        fallback=Agent(name="the_fallback", model=MODEL, instructions="Best effort."),
        fallback_max_turns=4,
    )

    out["14_tools_nonworker"] = Agent(
        name="tools_nonworker",
        model=MODEL,
        instructions="Non-worker tool types.",
        tools=[
            http_tool(
                name="lookup",
                description="Look up a record.",
                url="https://example.test/api/{id}",
                method="GET",
                headers={"X-Api-Version": "2"},
            ),
            human_tool(name="ask_human", description="Ask a person to decide."),
            agent_tool(billing, name="delegate_billing", description="Delegate."),
        ],
    )

    out["17_tools_mcp"] = Agent(
        name="tools_mcp",
        model=MODEL,
        instructions="MCP tool types.",
        tools=[
            mcp_tool(server_url="http://localhost:3001/mcp"),
            mcp_tool(
                server_url="http://localhost:3002/mcp",
                name="secured_mcp",
                description="Authenticated MCP tools.",
                headers={"Authorization": "Bearer ${MCP_AUTH_KEY}"},
                tool_names=["get_weather", "math_add"],
                max_tools=16,
                credentials=["MCP_AUTH_KEY"],
            ),
        ],
    )

    out["15_execution_and_creds"] = Agent(
        name="executor",
        model=MODEL,
        instructions="Run code and commands.",
        code_execution=CodeExecutionConfig(
            allowed_languages=["python", "bash"], timeout=60
        ),
        cli_config=CliConfig(allowed_commands=["git", "ls"], timeout=45, allow_shell=True),
        credentials=["GITHUB_TOKEN", "OPENAI_API_KEY"],
        required_tools=["get_weather"],
        masked_fields=["ssn", "card"],
        introduction="Hi, I run code.",
        metadata={"team": "platform", "tier": 2},
        stateful=True,
    )

    out["16_nested_tree"] = Agent(
        name="root",
        model=MODEL,
        strategy=Strategy.SEQUENTIAL,
        agents=[
            Agent(
                name="mid",
                model=MODEL,
                strategy=Strategy.PARALLEL,
                agents=[billing, tech],
            ),
            refunds,
        ],
    )

    return out


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--out", required=True, type=Path)
    args = ap.parse_args()
    args.out.mkdir(parents=True, exist_ok=True)

    serializer = AgentConfigSerializer()
    written = 0
    for name, agent in sorted(fixtures().items()):
        config = serializer.serialize(agent)
        path = args.out / f"{name}.json"
        path.write_text(json.dumps(config, indent=2, sort_keys=True) + "\n")
        print(f"wrote {path.name}")
        written += 1
    print(f"\n{written} fixtures written to {args.out}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
