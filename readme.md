# envsync

[![Build Status](https://travis-ci.org/claudiunicolaa/envsync.svg?branch=master)](https://travis-ci.org/claudiunicolaa/envsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/claudiunicolaa/envsync)](https://goreportcard.com/report/github.com/claudiunicolaa/envsync)

A simple tool for check if environment file has the same keys with the example environment file. Built on top of [gotdotenv](https://github.com/joho/godotenv).

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
