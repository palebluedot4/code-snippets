#!/usr/bin/env bash
set -euo pipefail

# grep -r "<pattern>" <directory>
grep -r "ERROR" /var/log/
