#!/usr/bin/env bash
set -euo pipefail

# cut -d<delimiter> -f<fields> <file>
cut -d: -f1 /etc/passwd
