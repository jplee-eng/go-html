#!/bin/sh

# Run Gin server binary and start Caddy
./go-html &
caddy run --config /etc/caddy/Caddyfile
