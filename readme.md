# envsync

A simple tool for check if environment file has the same keys with the example environment file, built on top of [gotdotenv](https://github.com/joho/godotenv) library.

Arguments:
 - **required**: the example environment file name 
 - **optional**: the environment file name (default: `.env`)
 
 @todo:
 - [x] write tests 
 - [ ] write documentation
 - [ ] check if working on linux, ~~mac~~, windows
 - [ ] add a flag that let tool to write into environment file (https://github.com/joho/godotenv#writing-env-files)
 - [ ] check if it is ok to integrate with `godotenv` lib
 - [ ] benchmarking (optional)
 - [ ] add examples
