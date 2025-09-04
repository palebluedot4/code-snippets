#!/usr/bin/env bash
set -euo pipefail

# scp -i <private_key> <local_file> <user>@<host>:<destination_path>
scp -i ~/.ssh/ecdsa-key.pem ./config.yaml ec2-user@10.0.1.100:/etc/app/
