#!/usr/bin/env bash
set -euo pipefail

# scp -i <private_key> <user>@<host>:<remote_file> <local_path>
scp -i ~/.ssh/ecdsa-key.pem ec2-user@10.0.1.100:/var/log/app.log ./
