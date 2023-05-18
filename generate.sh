#!/bin/sh

# Generate the public and private keys for the command signatures

secrets="secrets"

# Generate keys
openssl genrsa -out $secrets/key.pem 2048
openssl rsa -in $secrets/key.pem -pubout -out $secrets/key.crt

# Embed publlic key to the bot source
pubkey=$(cat $secrets/key.crt | sed -e 's/$/\\n/' | tr -d '\n')
echo "package main\nvar certificate = \`$pubkey\`" > bot/certificate.go
