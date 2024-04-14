# Base
FROM golang:1.22 AS build-base
WORKDIR /app
COPY Makefile .

# API
FROM build-base AS api-builder
COPY api/ ./api/
RUN go mod init gowebserver
RUN make build-api

# Docs
FROM build-base AS docs-builder
COPY web/ ./web/
COPY web/config.toml .
RUN apt-get update && apt-get install -y hugo
RUN make build-docs

# Caddy
FROM caddy:latest AS caddy-builder

# Final
FROM golang:1.22
WORKDIR /app

# Get Hugo
RUN apt-get update && apt-get install -y hugo

COPY --from=caddy-builder /usr/bin/caddy /usr/bin/caddy
COPY --from=api-builder /app/api/bin/gowebserver .
COPY --from=docs-builder /app/web docs

# Copy Caddyfile
COPY Caddyfile /etc/caddy/Caddyfile

# Copy entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Start API, docs, caddy
CMD ["/app/entrypoint.sh"]
