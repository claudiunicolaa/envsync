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

// EnvSync has two parameter:
// 	- the environment filename
//  - the environment example filename
func EnvSync(envFileName, envFileNameExample string) (bool, error) {
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

func getMapKeysSorted(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
