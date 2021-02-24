package envsync

import (
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

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

func TestCallWithOneNonExistingFilename(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filename := dir + "/.env"
	err = ioutil.WriteFile(filename, []byte{}, 0600)
	if err != nil {
		panic(err)
	}

	defer os.Remove(filename)

	_, err = EnvSync(".env", "random_file_name_789")
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
	err = ioutil.WriteFile(filename, []byte{}, 0600)
	if err != nil {
		panic(err)
	}

	defer os.Remove(filename)

	_, err = EnvSync("random_file_name_789", "789_random_file_name")
	// get just first 26 characters from error message because the rest differ linux&mac vs windows :|
	if err != nil && err.Error()[0:26] == "open random_file_name_789:" {
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
	err = ioutil.WriteFile(filename, []byte{}, 0600)
	if err != nil {
		panic(err)
	}

	tmpfile := createFile([]byte{}, ".env.example")

	defer os.Remove(filename)
	defer os.Remove(tmpfile.Name())

	_, err = EnvSync(filename, tmpfile.Name())
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

	r, err := EnvSync(tmpfileEnv.Name(), tmpfileEnvExample.Name())
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

	r, err := EnvSync(tmpfileEnv.Name(), tmpfileEnvExample.Name())
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

	r, err := EnvSync(tmpfileEnv.Name(), tmpfileEnvExample.Name())
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
