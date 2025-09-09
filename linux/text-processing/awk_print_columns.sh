#!/usr/bin/env bash
set -euo pipefail

ls -l | awk '{print $1, $9}'
