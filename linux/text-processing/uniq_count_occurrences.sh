#!/usr/bin/env bash
set -euo pipefail

cat /var/log/auth.log | grep "Failed password" | grep -oE '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+' | sort | uniq -c | sort -nr
