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

from conductor.ai.agents import Agent, AgentRuntime, Strategy, tool

MODEL = "anthropic/claude-sonnet-4-6"


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
