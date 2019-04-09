# envsync

[![Build Status](https://travis-ci.org/claudiunicolaa/envsync.svg?branch=master)](https://travis-ci.org/claudiunicolaa/envsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/claudiunicolaa/envsync)](https://goreportcard.com/report/github.com/claudiunicolaa/envsync)

A simple command for check if the environment and environment example files are synced. Under the hood, the command checks if files have the same keys.

### Operating system, programming language and framework agnostic
The command runs on all three main platforms (Linux, Mac, Windows), it is programming languages and frameworks agnostic. There is a single condition: key-value definition of environment variables into files.

Built on top of [gotdotenv](https://github.com/joho/godotenv).

Arguments:
 - **required**: the example environment file name 
 - **optional**: the environment file name (default: `.env`)
 
 @todo:
 - [x] write tests 
 - [ ] write documentation
 - [x] check if working on linux, ~~mac~~, windows
 - [ ] add a flag that let tool to write into environment file (https://github.com/joho/godotenv#writing-env-files)
 - [ ] check if it is ok to integrate with `godotenv` lib
 - [ ] add examples
