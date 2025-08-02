#!/usr/bin/env bash
set -euo pipefail

# sar -P ALL <interval> <count>
sar -P ALL 1 5
