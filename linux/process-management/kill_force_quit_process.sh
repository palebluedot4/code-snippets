#!/usr/bin/env bash
set -euo pipefail

# kill -9 <pid>
kill -9 $(pidof cron)
