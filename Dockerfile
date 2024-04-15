### Base ###
FROM golang:1.22 AS build-base
WORKDIR /app
COPY Makefile .

### API and Docs Build ###
FROM build-base AS builder
WORKDIR /app
# Get Hugo
RUN apt-get update && apt-get install -y hugo
# API
COPY api/ ./api/
RUN go mod init gowebserver
RUN make build-api
# Docs
COPY web/ ./web/
COPY web/config.toml .
RUN make build-docs

### Final Image on Alpine ###
FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/api/bin/gowebserver .
RUN chmod +x gowebserver
COPY --from=builder /app/web/public /srv
# Copy Caddy binary
COPY --from=caddy:latest /usr/bin/caddy /usr/bin/caddy
COPY Caddyfile /etc/caddy/Caddyfile
# Copy entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Start and serve API, docs, caddy
ENTRYPOINT ["/app/entrypoint.sh"]
