#!/usr/bin/env bash
set -euo pipefail

docker volume create --name=go-modules-cache

docker run --rm \
    -it \
    --name dev-app \
    -p 8080:8080 \
    -v "$(pwd):/app" \
    palebluedot4/development
