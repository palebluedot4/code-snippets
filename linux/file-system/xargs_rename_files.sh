#!/usr/bin/env bash
set -euo pipefail

ls *.txt | cut -d'.' -f1 | xargs -I {} mv {}.txt {}.md
