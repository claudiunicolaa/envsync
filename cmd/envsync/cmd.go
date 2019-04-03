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

	r, err := envsync.EnvSync(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if r {
		log.Println("environment files are synced")
	} else {
		log.Println("environment files are out-of-sync")
	}
}
