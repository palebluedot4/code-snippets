#!/usr/bin/env bash
set -euo pipefail

# ln -s <target> <link_name>
ln -s /var/log/nginx/access.log ./nginx_access.log
