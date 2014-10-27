dockerify
=========

Small command-line utility for creating Dockerfiles easily. It will attempt to
discover the project type (e.g.: NodeJS, Go etc) and generate a Dockerfile
using the curated list on [Dockerfile](http://dockerfile.github.io/).

Currently you have to compile it yourself, but I plan to add the cross-compiled
binaries in the future.

## Compiling

The project hasn't got any external dependencies so far, which means you can
just run `go build`, but you can also use `make all` (a bit more verbose).

## Usage

```bash
$ dfy ./
```

## Supported environments

- NodeJS (80%)