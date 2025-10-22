#!/usr/bin/env bash
set -euo pipefail

docker run -d \
    --restart always \
    --name my-app \
    -p 8080:8080 \
    --volume my-app-data:/app/data \
    --memory="512m" \
    --cpus="0.5" \
    --log-driver json-file \
    --log-opt max-size=10m \
    --log-opt max-file=3 \
    palebluedot4/my-app:1.0.0
