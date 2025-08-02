#!/usr/bin/env bash
set -euo pipefail

# strace -o <output_file> <command>
strace -o strace_output.txt ./main
