#!/usr/bin/env bash
set -euo pipefail

# nohup <command> ><output_file> 2>&1 &
nohup ./main >app.log 2>&1 &
