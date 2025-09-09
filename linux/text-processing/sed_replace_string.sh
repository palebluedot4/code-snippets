#!/usr/bin/env bash
set -euo pipefail

# sed 's/<find>/<replace>/g' <file>
sed 's/localhost/127.0.0.1/g' /etc/hosts
