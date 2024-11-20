---
weight: 1
title: "Get started"
bookFlatSection: true
---

# Get started

## Configuring the project

### Downloads required

- The [Go programming language](https://go.dev/doc/install)
- [Hugo](https://gohugo.io/getting-started/quick-start/) for static site generation,
- [Caddy](https://caddyserver.com/docs/install) to reverse proxy,
- [Docker](https://docs.docker.com/engine/install/) to run the app in a container.

### Setup

After finishing the required downloads from the section above and testing that each of them works, follow these setup steps:

1. Open a terminal emulator and clone the repo and change into the new directory with `git clone https://github.com/jplee-eng/go-html.git && cd go-html` and leave this open.
2. Change the name of the `.env.example` file to `.env.` and edit the defaults as you wish.
3. In the same terminal as step 1, run `make local-run`.
4. From a web browser, check the url [http://localhost/](http://localhost/) for output. You should find the welcome page which describes the project. Then, navigate to [http://localhost/api/ping](http://localhost/api/ping) and check for the `{"message":"active"}` output.
