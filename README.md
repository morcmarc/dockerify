dockerify
=========

[![Build Status](https://travis-ci.org/morcmarc/dockerify.svg?branch=master)](https://travis-ci.org/morcmarc/dockerify) [![GoDoc](https://godoc.org/github.com/morcmarc/dockerify?status.svg)](https://godoc.org/github.com/morcmarc/dockerify)

Small command-line utility for creating Dockerfiles easily. It will attempt to
discover the project type (e.g.: NodeJS, Go etc) and generate a Dockerfile
using the curated list on [Dockerfile](http://dockerfile.github.io/).

Currently you have to compile it yourself, but I plan to add the cross-compiled
binaries in the future.

## Compiling

Install dependencies with godep first, then run `go build`. You can also use `make install` (a bit more verbose).

## Usage

```bash
$ dfy /path/to/project/root
```

## Planned features (subject to change a lot)

- Full [fig](http://fig.sh) support

## Supported environments

- NodeJS (runtime)
- Go (runtime)