# DB friendly Short UUID.V1 implementation for Go language

[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FNexoMichael%2Fshortuuid%2Fbadge&style=flat)](https://actions-badge.atrox.dev/NexoMichael/shortuuid/goto)
[![CodeCovWidget]][CodeCovResult]
[![GoReport Widget]][GoReport Status]
[![GoDoc Widget]][GoDoc Link]

[GoReport Status]: https://goreportcard.com/report/github.com/NexoMichael/shortuuid
[GoReport Widget]: https://goreportcard.com/badge/github.com/NexoMichael/shortuuid

[CodeCovResult]: https://coveralls.io/github/NexoMichael/shortuuid
[CodeCovWidget]: https://coveralls.io/repos/github/NexoMichael/shortuuid/badge.svg?branch=master

[GoDoc Link]: http://godoc.org/github.com/NexoMichael/shortuuid
[GoDoc Widget]: http://godoc.org/github.com/NexoMichael/shortuuid?status.svg

This package provides fast shortuuid implementation friendly for DB indexes (for UUID V1)

## Installation

Use the `go` command:

	$ go get github.com/NexoMichael/shortuuid

## Requirements

shortuuid package requires Go >= 1.5.

## Links
* [Store UUID in an optimized way](https://www.percona.com/blog/2014/12/19/store-uuid-optimized-way/)

## Copyright

Copyright (C) 2018 by Mikhail Kochegarov <mikhail@kochegarov.biz>.

UUID package released under MIT License.
See [LICENSE](https://github.com/NexoMichael/shortuuid/blob/master/LICENSE) for details.
