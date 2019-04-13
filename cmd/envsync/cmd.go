package main

import (
	"flag"
	"fmt"
	"github.com/claudiunicolaa/envsync"
	"log"
	"os"
)

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	help := `
Check if environment example and environment file are synced.

envsync [-h] path/to/environment/example/file [path/to/environment/file]

Second argument is optional. Default=.env

Usage example:
envsync .env.example .env
	`

	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Println(help)
		return
	}

	envFileName, envExampleFileName := getEnvFileNames(os.Args[1:])
	r, err := envsync.EnvSync(envFileName, envExampleFileName)
	if err != nil {
		log.Fatal(err)
	}

	if r {
		log.Println("environment files are synced")
	} else {
		log.Println("environment files are out-of-sync")
	}
}

func getEnvFileNames(args []string) (string, string) {
	// default name of environment file
	var envFileName = ".env"
	envFileNameExample := args[0]
	// if second cli argument specified
	if len(args) > 1 {
		envFileName = args[1]
	}

	return envFileName, envFileNameExample
}
