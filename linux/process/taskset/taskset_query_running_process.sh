#!/usr/bin/env bash
set -euo pipefail

# taskset -p <pid>
taskset -p $(pidof systemd)
