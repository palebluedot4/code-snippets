#!/usr/bin/env bash
set -euo pipefail

# find <path> -type <f|d> -name "<pattern>"
find . -type f -name "*.log"
