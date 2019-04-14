# envsync

[![GoDoc](https://godoc.org/github.com/claudiunicolaa/envsync?status.svg)](https://godoc.org/github.com/claudiunicolaa/envsync)
[![Build Status](https://travis-ci.org/claudiunicolaa/envsync.svg?branch=master)](https://travis-ci.org/claudiunicolaa/envsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/claudiunicolaa/envsync)](https://goreportcard.com/report/github.com/claudiunicolaa/envsync)

A simple command for checking if the environment and the environment example files are synced. Under the hood, the command checks if the files have the same keys.
It can be used as a library or as a bin command.

### Operating system, programming language and framework agnostic
The command runs on all three main platforms (Linux, Mac, Windows), it is programming languages and frameworks agnostic.

Built on top of [gotdotenv](https://github.com/joho/godotenv).

## Installation

### Library

```shell
go get github.com/claudiunicolaa/envsync
```

### Bin Command

```shell
go get github.com/claudiunicolaa/envsync/cmd/envsync
```

## Usage


### [Library](examples/example.go)

```go
package main

import (
	"fmt"
	"github.com/claudiunicolaa/envsync"
)

func main() {
	_, err := envsync.EnvSync(".env", ".env.example")

	if err != nil {
		fmt.Println(err)
		return
	}
    
	// ...
	// all good, we'll supposed here is
	// a lot of magic written in Go
	// ...
}

```


### Bin command

Install as above and you can run it like as a bin command from your terminal.

```shell
envsync [-h] path/to/environment/example/file [path/to/environment/file]

// The above Go code can be translated into 
envsync .env.example .env
```

## Something wrong?
If you encounter some problems, please open an issue.
