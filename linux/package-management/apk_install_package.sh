#!/usr/bin/env bash
set -euo pipefail

# apk update && apk add --no-cache <package>
apk update && apk add --no-cache htop
