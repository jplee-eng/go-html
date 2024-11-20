---
title: Go Web Server
type: docs
---

# Welcome!

{{< columns >}}

## /api

The [api directory](https://github.com/jplee-eng/go-html/tree/main/api) contains the Gin API. By default it contains a health check handler that can be accessed under /api at [/active](http://localhost/api/active), [/health](http://localhost/api/health), [/isActive](http://localhost/api/isActive), and [/ping](http://localhost/api/ping).
<--->

## /web

The [web directory](https://github.com/jplee-eng/go-html/tree/main/web) contains this site you are currently exploring. It is statically rendered by [hugo](https://gohugo.io) usually under 50ms and is configured with the [book theme](https://themes.gohugo.io/themes/hugo-book/).
{{< /columns >}}

## Project goal

The end goal for this project is to have as much functionality out of the box as possible without giving up on performance or configuration time. It should be dead simple to set up, enough for non-technical businesses and individuals to easily setup a blog with a few steps.

- Lightweight and only using what's needed.
- Online in just a few clicks.
- Highly modular and extendable.

### Example ideas

- Sign up api route that allows newsletter subscriptions from a blog post.
- Automatic static site rebuilding from external data changes.

## Contributing

Please feel free to open PRs and contribute to make this project more useful and simpler to understand and use for everybody. Your contributions do not have to be anything related to the example ideas!
