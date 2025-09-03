#!/usr/bin/env bash
set -euo pipefail

# find <path> -type f -name "<pattern>" -print0 | xargs -0 rm
find . -type f -name "*.tmp" -print0 | xargs -0 rm
