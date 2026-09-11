#!/usr/bin/env python3
import sys

args = " ".join(sys.argv[1:]) if len(sys.argv) > 1 else "no-args"
print(f"ECHO_ARGS_RESULT:{args}")
