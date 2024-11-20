# go-html

go api with common functionality and hugo rendered docs.

## build and run locally

1. clone repo, change into the directory, change the name of `.env.example` to `.env` and update its values as desired.
2. for docker, start the docker daemon and run `make docker-run` otherwise use `make local-run`.
3. test the hugo static site at `http://localhost/` and the api health. check at `http://localhost/api/health` for output.

## dependencies to install

- [golang](https://go.dev/doc/install)
- [hugo](https://gohugo.io/installation/)
- [caddy](https://caddyserver.com/docs/install)
- [docker](https://docs.docker.com/get-docker/)
