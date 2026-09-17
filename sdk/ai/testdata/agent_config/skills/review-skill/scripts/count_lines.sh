#!/bin/bash
# Counts the lines in the text passed as arguments.
printf '%s\n' "$@" | wc -l | tr -d ' '
