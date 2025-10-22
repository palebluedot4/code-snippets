#!/usr/bin/env bash
set -euo pipefail

# docker run --rm -it --name <container_name> -p <host_port>:<container_port> <image_name>:<tag>
docker run --rm -it --name my-app -p 8080:8080 palebluedot4/my-app:latest
