#!/usr/bin/env bash
set -euo pipefail

# nice -n <nice_value> <command>
nice -n 19 ./main &
