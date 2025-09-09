#!/usr/bin/env bash
set -euo pipefail

# diff <(ls -1 <dir1>) <(ls -1 <dir2>)
diff <(ls -1 dir1) <(ls -1 dir2)
