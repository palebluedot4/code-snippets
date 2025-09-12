#!/usr/bin/env bash
set -euo pipefail

# nohup <command> &><output_file> &
nohup ./main &>app.log &
