#!/usr/bin/env bash
set -euo pipefail

# sar -P <cpu> <interval> <count>
sar -P ALL 1 5
