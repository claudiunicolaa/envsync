package envsync

import (
	"errors"
	"github.com/joho/godotenv"
	"reflect"
	"sort"
)

func EnvSync(args []string) (bool, error) {
	if !canBeRun(args) {
		return false, errors.New("please provide the need it arguments")
	}

	envFileName, envFileNameExample := getEnvFileNames(args)

	envMap, err := godotenv.Read(envFileName)
	if err != nil {
		return false, err
	}

	envExampleMap, err := godotenv.Read(envFileNameExample)
	if err != nil {
		return false, err
	}

	envKeys := getMapKeysSorted(envMap)
	envExampleKeys := getMapKeysSorted(envExampleMap)

	eq := reflect.DeepEqual(envKeys, envExampleKeys)

	if !eq {
		return false, errors.New("environment files are not synced")
	}

	return true, nil
}

func canBeRun(args []string) bool {
	return len(args) > 0
}

func getMapKeysSorted(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getEnvFileNames(args []string) (string, string) {
	// default name of environment file
	var envFileName = ".env"
	envFileNameExample := args[0];
	// if second cli argument specified
	if len(args) > 1 {
		envFileName = args[1];
	}

	return envFileName, envFileNameExample
}
