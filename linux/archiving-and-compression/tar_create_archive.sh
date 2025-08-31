#!/usr/bin/env bash
set -euo pipefail

# tar -czvf <archive_file> <directory_to_archive>
tar -czvf project_backup.tar.gz ./project
