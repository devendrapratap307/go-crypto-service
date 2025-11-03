#!/usr/bin/env bash
set -e
mkdir -p keys
# Generate 2048-bit RSA keypair for signing
openssl genrsa -out keys/private.pem 2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
# Generate random 32 byte AES key (base64)
head -c 32 /dev/urandom | base64 > keys/aes.key

echo "Keys generated in ./keys"