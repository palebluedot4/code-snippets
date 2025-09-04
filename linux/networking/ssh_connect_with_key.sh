#!/usr/bin/env bash
set -euo pipefail

# ssh -i <private_key> <user>@<host>
ssh -i ~/.ssh/ecdsa-key.pem ec2-user@10.0.1.100
