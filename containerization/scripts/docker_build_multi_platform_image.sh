#!/usr/bin/env bash
set -euo pipefail

# docker buildx build --platform <platform1,platform2,...> -t <image_name>:<tag> --push <build_context_path>
docker buildx build --platform linux/amd64,linux/arm64 -t palebluedot4/my-app:latest --push .
