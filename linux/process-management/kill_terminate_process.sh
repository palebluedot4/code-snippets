#!/usr/bin/env bash
set -euo pipefail

# kill <pid>
kill $(pidof nginx)
