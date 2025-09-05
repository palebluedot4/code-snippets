#!/usr/bin/env bash
set -euo pipefail

md5sum * | cut -c1-32 | sort | uniq -c | sort -nr | grep -v " 1 "
