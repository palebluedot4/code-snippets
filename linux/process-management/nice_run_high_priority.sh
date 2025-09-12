#!/usr/bin/env bash
set -euo pipefail

# sudo nice -n <nice_value> <command>
sudo nice -n -20 ./main &
