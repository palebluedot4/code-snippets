#!/usr/bin/env bash
set -euo pipefail

# kill -HUP <pid>
kill -HUP $(pidof nginx)
