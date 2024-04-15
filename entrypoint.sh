#!/bin/sh

# Run Gin server binary and start Caddy
./gowebserver &
caddy run --config /etc/caddy/Caddyfile
