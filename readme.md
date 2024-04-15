# gowebserver

go gin api with common functionality and docs side by side.

## build and run locally

1. edit `.env` as desired
2. run `make docker-run` or `make-run`
3. test the hugo static site at `http://localhost/` and the gin api health check at `http://localhost/api/ping` for output

## dependencies to install

- [golang](https://go.dev/doc/install)
- [hugo](https://gohugo.io/installation/)
- [caddy](https://caddyserver.com/docs/install)
- [docker](https://docs.docker.com/get-docker/)
