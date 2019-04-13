package main

import "testing"

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
