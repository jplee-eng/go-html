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

# Final
FROM golang:1.22
WORKDIR /app

# Get Hugo
RUN apt-get update && apt-get install -y hugo

COPY --from=api-builder /app/api/bin/gowebserver .
COPY --from=docs-builder /app/web docs

# Start API and docs
EXPOSE 1313
EXPOSE 8080
CMD ["sh", "-c", "hugo server --bind 0.0.0.0 --source docs --environment production & ./gowebserver"]
