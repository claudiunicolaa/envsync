package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"sort"
)

func main() {
	if !canBeRun() {
		log.Fatal("Please provide the need it arguments")
	}

	// default name of environment file
	var envFileName = ".env"

	envFileNameExample := os.Args[1];
	// if second cli argument specified
	if len(os.Args) > 2 {
		envFileName = os.Args[2];
	}

	envMap, err := godotenv.Read(envFileName)
	if err != nil {
		log.Fatal(err);
	}

	envExampleMap, err := godotenv.Read(envFileNameExample)
	if err != nil {
		log.Fatal(err);
	}

	envKeys := getMapKeysSorted(envMap)
	envExampleKeys := getMapKeysSorted(envExampleMap)

	eq := reflect.DeepEqual(envKeys, envExampleKeys)

	if eq {
		fmt.Println("Environment files keys are the same")
	} else {
		fmt.Println("Environment files keys are not the same")
	}
}

func canBeRun() bool {
	return len(os.Args) > 1
}

func getMapKeysSorted(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
