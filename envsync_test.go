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
	r := canBeRun([]string{"one", "tow"})
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
			break;
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

	_, err = EnvSync([]string{"cmd_name", "random_file_name_789"})
	if err != nil && err.Error() == "open random_file_name_789: no such file or directory" {
		return
	}

	t.Error("Calling with one non-existing filename must return 'open random_file_name_789: no such file or directory'")
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
	if err != nil && err.Error() == "open 789_random_file_name: no such file or directory" {
		return
	}

	t.Error("Calling with one non-existing filename must return 'open 789_random_file_name: no such file or directory'")
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

	r, err := EnvSync([]string{"cmd_name", tmpfileEnvExample.Name(), tmpfileEnv.Name()})
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

	r, err := EnvSync([]string{"cmd_name", tmpfileEnvExample.Name(), tmpfileEnv.Name()})
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

	r, err := EnvSync([]string{"cmd_name", tmpfileEnvExample.Name(), tmpfileEnv.Name()})
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
