#!/usr/bin/env bash
set -euo pipefail

# journalctl -u <service>
journalctl -u sshd
