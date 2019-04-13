# envsync

[![GoDoc](https://godoc.org/github.com/claudiunicolaa/envsync?status.svg)](https://godoc.org/github.com/claudiunicolaa/envsync)
[![Build Status](https://travis-ci.org/claudiunicolaa/envsync.svg?branch=master)](https://travis-ci.org/claudiunicolaa/envsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/claudiunicolaa/envsync)](https://goreportcard.com/report/github.com/claudiunicolaa/envsync)

A simple command for check if the environment and environment example files are synced. Under the hood, the command checks if files have the same keys.
It can be used as a library or as a bin command.

### Operating system, programming language and framework agnostic
The command runs on all three main platforms (Linux, Mac, Windows), it is programming languages and frameworks agnostic. There is a single condition: key-value definition of environment variables into files.

Built on top of [gotdotenv](https://github.com/joho/godotenv).

## Installation

Library

```shell
go get github.com/claudiunicolaa/envsync
```

Bin command

```shell
go get github.com/claudiunicolaa/cmd/envsync
```

## Usage
