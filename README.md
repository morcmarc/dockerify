dockerify
=========

[![Build Status](https://travis-ci.org/morcmarc/dockerify.svg?branch=master)](https://travis-ci.org/morcmarc/dockerify) [![GoDoc](https://godoc.org/github.com/morcmarc/dockerify?status.svg)](https://godoc.org/github.com/morcmarc/dockerify)[ ![Download](https://api.bintray.com/packages/morcmarc/dockerify/dockerify/images/download.svg) ](https://bintray.com/morcmarc/dockerify/dockerify/_latestVersion)

Small command-line utility for creating Dockerfiles easily. It will attempt to
discover the project type (e.g.: NodeJS, Go etc) and generate a Dockerfile
based on the curated list at [Dockerfile](http://dockerfile.github.io/).

## Building

For development:

```bash
$ make install
```

To compile the release version install [goxc](https://github.com/laher/goxc) first then run:

```bash
$ goxc
```

This will output all cross-compiled binaries into your $GOPATH/bin folder.

## Usage

```bash
$ dockerify /path/to/project
```

## Planned features (subject to change a lot)

- Full [fig](http://fig.sh) support

## Supported environments

- NodeJS (runtime)
- Go (runtime)