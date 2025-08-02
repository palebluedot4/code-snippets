#!/usr/bin/env bash
set -euo pipefail

# strace -p <pid>
strace -p $(pidof nginx)
