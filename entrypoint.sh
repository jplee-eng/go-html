#!/bin/sh

./go-html &
caddy run --config /etc/caddy/Caddyfile
