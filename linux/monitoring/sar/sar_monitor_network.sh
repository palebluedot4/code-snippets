#!/usr/bin/env bash
set -euo pipefail

# sar -n DEV <interval> <count>
sar -n DEV 1 5
