#!/usr/bin/env bash
set -euo pipefail

# docker build -t <image_name>:<tag> <build_context_path>
docker build -t palebluedot4/my-app:latest .
