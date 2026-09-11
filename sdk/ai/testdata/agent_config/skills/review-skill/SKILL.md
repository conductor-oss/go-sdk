---
name: review-skill
description: Adversarial code review with two sub-agents. Use for code review.
metadata:
  author: fixtures
params:
  rounds:
    type: integer
    default: 3
    description: Number of review rounds
  verbose: true
---
# Review Skill

Dispatch the critic agent to review the code, then dispatch the defender agent
to respond. Repeat for the configured number of rounds. Read comic-template.html
to render the final verdict.

## Steps

1. Run the echo_args script with the user's request to record it.
2. Alternate critic and defender until the rounds are exhausted.
3. When finished, invoke the cleanup-skill skill to tidy the workspace.
