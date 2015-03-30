dockerify
=========

[![Build Status](https://travis-ci.org/morcmarc/dockerify.svg?branch=master)](https://travis-ci.org/morcmarc/dockerify) [![GoDoc](https://godoc.org/github.com/morcmarc/dockerify?status.svg)](https://godoc.org/github.com/morcmarc/dockerify)[ ![Download](https://api.bintray.com/packages/morcmarc/dockerify/dockerify/images/download.svg) ](https://bintray.com/morcmarc/dockerify/dockerify/_latestVersion)

Small command-line utility for scaffolding Dockerfiles. It will attempt to
identify the project type (i.e.: node+express, golang) and generate a Docker and
Fig file -- using [Dockerfile](http://dockerfile.github.io/).

## Building

Development:

```bash
$ make install
```

To compile the release version install [goxc](https://github.com/laher/goxc) and run:

```bash
$ goxc
```

You can find the cross-compiled binaries at $GOPATH/bin

## Usage

```bash
$ dockerify /path/to/project
```

## Planned features (subject to change)

- Full [fig](http://fig.sh) support

## Supported environments

- Node (runtime)
- Go (runtime)