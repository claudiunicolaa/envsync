package envsync

import (
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func TestCanBeRunWithEmptySlice(t *testing.T) {
	r := canBeRun([]string{})
	if r {
		t.Error("canBeRun with empty slice must return false")
	}
}

func TestCanBeRunWithOneElemSlice(t *testing.T) {
	r := canBeRun([]string{"one"})
	if !r {
		t.Error("canBeRun with one element in slice must return false")
	}
}

func TestCanBeRunWithTwoElemSlice(t *testing.T) {
	r := canBeRun([]string{"one", "two"})
	if !r {
		t.Error("canBeRun with two elements in slice must return true")
	}
}

func TestGetEnvFileNamesWithOneArgument(t *testing.T) {
	args := []string{".env.example"}
	s1, s2 := getEnvFileNames(args)

	if s1 != ".env" || s2 != args[0] {
		t.Error("getEnvFileName failed")
	}
}

func TestGetEnvFileNamesWithTwoArguments(t *testing.T) {
	args := []string{".env.example", ".env"}
	s1, s2 := getEnvFileNames(args)

	if s1 != args[1] || s2 != args[0] {
		t.Error("getEnvFileName failed")
	}
}

func TestGetMapKeysSorted(t *testing.T) {
	k := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

	kShuffled := make([]string, len(k))
	for {
		for i, v := range rand.Perm(len(k)) {
			kShuffled[v] = k[i]
		}
		if reflect.DeepEqual(kShuffled, k) {
			break
		}
	}

	m := make(map[string]string, len(k))
	for _, v := range kShuffled {
		m[v] = "a"
	}

	sorted := getMapKeysSorted(m)
	eq := reflect.DeepEqual(sorted, k)

	if !eq {
		t.Error("getMapKeysSorted failed")
	}
}

func TestGetKeysDiff(t *testing.T) {
	a := []string{"one", "two", "three", "four", "five"}
	b := []string{"two", "three", "five"}

	diff := getKeysDiff(a, b)
	eq := reflect.DeepEqual(diff, []string{"one", "four"})
	if !eq {
		t.Error("getKeysDiff must return true")
	}

	a = []string{"one", "two", "three", "four", "five"}
	b = []string{"one", "two", "three", "four", "five"}

	diff = getKeysDiff(a, b)
	eq = reflect.DeepEqual(diff, []string{})
	if eq {
		t.Error("getKeysDiff must return false")
	}
}

func TestCallWithoutArguments(t *testing.T) {
	_, err := EnvSync([]string{})
	if err != nil && err.Error() == "please provide the need it arguments" {
		return
	}

	t.Error("Calling without arguments must return 'please provide the need it arguments' error")
}

func TestCallWithOneNonExistingFilename(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filename := dir + "/.env"
	err = ioutil.WriteFile(filename, []byte{}, 0644)
	if err != nil {
		panic(err)
	}

	defer os.Remove(filename)

	_, err = EnvSync([]string{"random_file_name_789"})
	// get just first 26 characters from error message because the rest differ linux&mac vs windows :|
	if err != nil && err.Error()[0:26] == "open random_file_name_789:" {
		return
	}

	t.Error("Calling with one non-existing filename must return first 26 characters 'open random_file_name_789:'")
}

func TestCallWithTwoNonExistingFilenames(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filename := dir + "/.env"
	err = ioutil.WriteFile(filename, []byte{}, 0644)
	if err != nil {
		panic(err)
	}

	defer os.Remove(filename)

	_, err = EnvSync([]string{"random_file_name_789", "789_random_file_name"})
	// get just first 26 characters from error message because the rest differ linux&mac vs windows :|
	if err != nil && err.Error()[0:26] == "open 789_random_file_name:" {
		return
	}

	t.Error("Calling with one non-existing filename must return first 26 characters 'open 789_random_file_name:'")
}

func TestCallWithExampleFileExistingAndEnvFileDefault(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filename := dir + "/.env"
	err = ioutil.WriteFile(filename, []byte{}, 0644)
	if err != nil {
		panic(err)
	}

	tmpfile := createFile([]byte{}, ".env.example")

	defer os.Remove(filename)
	defer os.Remove(tmpfile.Name())

	_, err = EnvSync([]string{tmpfile.Name()})
	if err != nil {
		t.Error("Calling with example file existing must return nil")
	}
}

func TestCallWithExampleAndEnvFileSynced(t *testing.T) {
	tmpfileEnv := createFile(
		[]byte("APP_1=a\nAPP_2=a\nAPP_3=a"),
		".environment",
	)
	tmpfileEnvExample := createFile(
		[]byte("APP_1=a\nAPP_2=a\nAPP_3=a"),
		".environment.example",
	)
	// cleanup
	defer os.Remove(tmpfileEnv.Name())
	defer os.Remove(tmpfileEnvExample.Name())

	r, err := EnvSync([]string{tmpfileEnvExample.Name(), tmpfileEnv.Name()})
	if r && err != nil {
		t.Error("Calling with example and environment synced must return true, nil")
	}
}

func TestCallWithEnvFileNotSynced(t *testing.T) {
	tmpfileEnv := createFile(
		[]byte("APP_1=a\nAPP_2=a\n"),
		".environment",
	)
	tmpfileEnvExample := createFile(
		[]byte("APP_1=a\nAPP_2=a\nAPP_3=a"),
		".environment.example",
	)
	// cleanup
	defer os.Remove(tmpfileEnv.Name())
	defer os.Remove(tmpfileEnvExample.Name())

	r, err := EnvSync([]string{tmpfileEnvExample.Name(), tmpfileEnv.Name()})
	if r || err == nil {
		t.Error("Calling with example and environment not synced must return false, err")
	}
}

func TestCallWithEnvExampleFileNotSynced(t *testing.T) {
	tmpfileEnv := createFile(
		[]byte("APP_1=a\nAPP_2=a\nAPP_3=a"),
		".environment",
	)
	tmpfileEnvExample := createFile(
		[]byte("APP_1=a\nAPP_2=a\n"),
		".environment.example",
	)

	// cleanup
	defer os.Remove(tmpfileEnv.Name())
	defer os.Remove(tmpfileEnvExample.Name())

	r, err := EnvSync([]string{tmpfileEnvExample.Name(), tmpfileEnv.Name()})
	if r || err == nil {
		t.Error("Calling with example and environment not synced must return false, err")
	}
}

func createFile(content []byte, filename string) *os.File {
	tmpfile, err := ioutil.TempFile("", filename)
	if err != nil {
		panic(err)
	}
	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	return tmpfile
}
