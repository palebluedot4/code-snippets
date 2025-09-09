#!/usr/bin/env bash
set -euo pipefail

# grep -c "<pattern>" <file>
grep -c "session opened" /var/log/auth.log
