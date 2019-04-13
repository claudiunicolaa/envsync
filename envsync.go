// Package envsync provide a simple checking tool if environment and environment example files are synced.
// The check take into consideration the keys, not the values.
package envsync

import (
	"errors"
	"github.com/joho/godotenv"
	"reflect"
	"sort"
	"strings"
)

// EnvSync get from args parameter the filenames of the environment and environment example.
// The default environment names is .env and the function needs at least one entry 
// into args array which will represent the name of environment example file. 
// The second entry is optional and will represent the environment file name and overrides the default .env.
func EnvSync(args []string) (bool, error) {
	if !canBeRun(args) {
		return false, errors.New("please provide the need it arguments")
	}

	envFileName, envFileNameExample := getEnvFileNames(args)

	envMap, err := godotenv.Read(envFileName)
	if err != nil {
		return false, errors.New(err.Error() + " for environment file (" + envFileName + ")")
	}

	envExampleMap, err := godotenv.Read(envFileNameExample)
	if err != nil {
		return false, errors.New(err.Error() + " for example environment file (" + envFileNameExample + ")")
	}

	envKeys := getMapKeysSorted(envMap)
	envExampleKeys := getMapKeysSorted(envExampleMap)

	eq := reflect.DeepEqual(envKeys, envExampleKeys)

	if !eq {
		diff := getKeysDiff(envExampleKeys, envKeys)
		return false, errors.New("environment files are not synced. Missing keys from " + envFileName + " are: " + strings.Join(diff, ", "))
	}

	return true, nil
}

func getKeysDiff(a, b []string) (diff []string) {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
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
	envFileNameExample := args[0]
	// if second cli argument specified
	if len(args) > 1 {
		envFileName = args[1]
	}

	return envFileName, envFileNameExample
}
