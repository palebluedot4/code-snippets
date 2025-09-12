#!/usr/bin/env bash
set -euo pipefail

# taskset -c <cpu> <command>
taskset -c 0 ./main &
