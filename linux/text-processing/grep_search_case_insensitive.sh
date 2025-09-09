#!/usr/bin/env bash
set -euo pipefail

# grep -in "<pattern>" <file>
grep -in "denied" /var/log/auth.log
