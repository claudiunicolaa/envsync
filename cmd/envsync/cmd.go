package main

import (
	"flag"
	"fmt"
)

func main() {
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	help := `
	Check if environment example and environment file are synced.
	`

	args := flag.Args()
	if showHelp || len(args) == 0 {
		fmt.Println(help)
		return
	}

	//r, err := env()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if r {
	//	log.Println("Files are synced")
	//} else {
	//	log.Println("File are out-of-sync")
}
