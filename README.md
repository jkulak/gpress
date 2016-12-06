# GPRESS

A command-line tool for compressing and decompressing files using GCOMPRESSOR library (https://github.com/jkulak/gcompressor).

![](https://raw.githubusercontent.com/jkulak/gpress/master/resources/compress.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkulak/gpress)](https://goreportcard.com/report/github.com/jkulak/gpress)
[![GoDoc](https://godoc.org/github.com/jkulak/gpress?status.svg)](https://godoc.org/github.com/jkulak/gpress)

## Usage

```
usage: gpress [--decompress] [--verbose] --sourcefile <file>

--decompress            decompress file
--help                  display help
--sourcefile            file to be compressed or decompressed
--verbose               add debug output
--version               display program version
```

To compress files
```
$ gpress --sourcefile file_to_compress.txt
```

To decompress files
```
$ gpress --decompress --sourcefile file_to_compress.txt
```

## Compilation

I like to use Docker 🐳 to work with Golang on my local dev machine.

Run the build

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 go build -v
```

Work inside the Docker container (you should set the GOPATH env variable)

```
$ docker run --rm -ti -e "GOPATH=/root/go" -v "$PWD":/root/go -w /root/go golang:1.6 bash
```

## Build and test status

[![Build Status](https://travis-ci.org/jkulak/gpress.svg?branch=master)](https://travis-ci.org/jkulak/gpress)
