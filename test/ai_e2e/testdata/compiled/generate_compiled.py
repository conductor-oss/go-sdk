"""Capture what the Python SDK's plan() compiles for the suite-1 agents.

These are the agents of e2e/test_suite1_basic_validation.py in the Python SDK
that the Go SDK can express too. Each file is the plan() result the server
returned for the Python agentConfig; the Go e2e test compiles the same agent
and compares, so a difference here is a difference in what the two SDKs send.

Run from a checkout of conductor-oss/python-sdk with a server reachable at
CONDUCTOR_SERVER_URL:

    PYTHONPATH=src python3 <this file> --out <this directory>
"""

from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

from conductor.ai.agents import (
    Agent,
    AgentRuntime,
    Guardrail,
    GuardrailResult,
    OnTextMention,
    RegexGuardrail,
    Strategy,
    audio_tool,
    http_tool,
    image_tool,
    mcp_tool,
    pdf_tool,
    tool,
    video_tool,
)

MODEL = "anthropic/claude-sonnet-4-6"

# The MCP test server the suite's kitchen sink points at. The value is part of
# the compiled workflow, so the Go test uses the same literal.
MCP_URL = "http://localhost:3001"


@tool
def add(a: int, b: int) -> int:
    """Add two numbers."""
    return a + b


@tool
def multiply(x: int, y: int) -> int:
    """Multiply two numbers."""
    return x * y


@tool
def greet(name: str) -> str:
    """Greet someone."""
    return f"Hello {name}"


@tool(credentials=["API_KEY_1"])
def credentialed_tool(query: str) -> str:
    """A tool that needs credentials."""
    return query


@tool(credentials=["SECRET_A", "SECRET_B"])
def multi_cred_tool(data: str) -> str:
    """A tool needing multiple credentials."""
    return data


def no_pii(content: str) -> GuardrailResult:
    """Block PII patterns."""
    return GuardrailResult(passed=True)


def check_input(content: str) -> GuardrailResult:
    """Validate input."""
    return GuardrailResult(passed=True)


@tool
def local_tool(x: str) -> str:
    """A local worker tool."""
    return x


@tool(credentials=["KS_SECRET"])
def cred_local_tool(x: str) -> str:
    """Worker tool with credentials."""
    return x


def kitchen_sink() -> Agent:
    """The suite's _make_kitchen_sink_agent, verbatim."""
    ht = http_tool(name="ks_http", description="HTTP endpoint", url=f"{MCP_URL}/echo", method="POST")
    mt = mcp_tool(server_url=MCP_URL, name="ks_mcp", description="MCP tools")
    img = image_tool(name="ks_image", description="Generate image", llm_provider="openai", model="dall-e-3")
    aud = audio_tool(name="ks_audio", description="Generate audio", llm_provider="openai", model="tts-1")
    vid = video_tool(name="ks_video", description="Generate video", llm_provider="openai", model="sora")
    pdf = pdf_tool(name="ks_pdf", description="Generate PDF")
    router_lead = Agent(name="ks_router_lead", model=MODEL, instructions="Route to correct agent.")

    def team(name, strategy, kids, **kw):
        return Agent(name=name, model=MODEL, strategy=strategy,
                     agents=[Agent(name=k, model=MODEL, instructions=i) for k, i in kids], **kw)

    return Agent(
        name="e2e_kitchen_sink", model=MODEL, instructions="You are the kitchen sink agent.",
        tools=[local_tool, cred_local_tool, ht, mt, img, aud, vid, pdf],
        guardrails=[
            Guardrail(check_input, position="input", on_fail="retry"),
            Guardrail(no_pii, position="output", on_fail="retry"),
            RegexGuardrail(patterns=[r"password"], name="no_password",
                           message="No passwords in output.", on_fail="retry"),
        ],
        agents=[
            team("ks_handoff", Strategy.HANDOFF, [("ks_h1", "H1."), ("ks_h2", "H2.")], instructions="Route tasks."),
            team("ks_sequential", Strategy.SEQUENTIAL, [("ks_seq1", "Seq1."), ("ks_seq2", "Seq2.")]),
            team("ks_parallel", Strategy.PARALLEL, [("ks_p1", "P1."), ("ks_p2", "P2.")]),
            team("ks_router", Strategy.ROUTER, [("ks_r1", "R1."), ("ks_r2", "R2.")], router=router_lead),
            team("ks_round_robin", Strategy.ROUND_ROBIN, [("ks_rr1", "RR1."), ("ks_rr2", "RR2.")]),
            team("ks_random", Strategy.RANDOM, [("ks_rand1", "Rand1."), ("ks_rand2", "Rand2.")]),
            team("ks_swarm", Strategy.SWARM, [("ks_sw1", "SW1."), ("ks_sw2", "SW2.")],
                 handoffs=[OnTextMention(text="GOTO_SW2", target="ks_sw2"),
                           OnTextMention(text="GOTO_SW1", target="ks_sw1")]),
            team("ks_manual", Strategy.MANUAL, [("ks_m1", "M1."), ("ks_m2", "M2.")]),
        ],
        strategy=Strategy.HANDOFF,
    )


def agents() -> dict[str, Agent]:
    child = Agent(name="e2e_child", model=MODEL, instructions="You are a helper.")
    analyst = Agent(name="e2e_analyst", model=MODEL, instructions="You analyze data.")
    writer = Agent(name="e2e_writer", model=MODEL, instructions="You write reports.")
    return {
        "e2e_smoke": Agent(
            name="e2e_smoke", model=MODEL, instructions="You are a calculator.",
            tools=[add, multiply],
        ),
        "e2e_tools": Agent(
            name="e2e_tools", model=MODEL, instructions="Use tools.",
            tools=[add, multiply, greet],
        ),
        "e2e_creds": Agent(
            name="e2e_creds", model=MODEL, instructions="Use tools.",
            tools=[credentialed_tool, multi_cred_tool],
        ),
        "e2e_parent": Agent(
            name="e2e_parent", model=MODEL, instructions="Delegate to child.",
            agents=[child], strategy=Strategy.HANDOFF,
        ),
        "e2e_base_url": Agent(
            name="e2e_base_url", model=MODEL, instructions="Say hello.",
            base_url="https://my-custom-proxy.example.com/v1",
        ),
        "e2e_no_base_url": Agent(
            name="e2e_no_base_url", model=MODEL, instructions="Say hello.",
        ),
        "e2e_guardrails": Agent(
            name="e2e_guardrails", model=MODEL, instructions="Answer questions.",
            tools=[greet],
            guardrails=[
                Guardrail(check_input, position="input", on_fail="retry"),
                Guardrail(no_pii, position="output", on_fail="retry"),
                RegexGuardrail(patterns=[r"\b\d{3}-\d{2}-\d{4}\b"], name="no_ssn",
                               message="No SSNs allowed.", on_fail="retry"),
            ],
        ),
        "e2e_kitchen_sink": kitchen_sink(),
        "e2e_manager": Agent(
            name="e2e_manager", model=MODEL,
            instructions="Delegate analysis to analyst and writing to writer.",
            agents=[analyst, writer], strategy=Strategy.HANDOFF,
        ),
    }


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--out", required=True, type=Path)
    args = ap.parse_args()
    args.out.mkdir(parents=True, exist_ok=True)
    with AgentRuntime() as rt:
        for name, agent in agents().items():
            result = rt.plan(agent)
            (args.out / f"{name}.json").write_text(json.dumps(result, indent=2, sort_keys=True) + "\n")
            print("wrote", name)
    return 0


if __name__ == "__main__":
    sys.exit(main())
