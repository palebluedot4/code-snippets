#!/usr/bin/env bash
set -euo pipefail

# scp -r -i <private_key> <local_directory> <user>@<host>:<destination_path>
scp -r -i ~/.ssh/ecdsa-key.pem ./project ec2-user@10.0.1.100:/home/ec2-user/
