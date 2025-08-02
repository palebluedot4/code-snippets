#!/usr/bin/env bash
set -euo pipefail

# taskset -c <cpu_list> <command>
taskset -c 0 ./main &
