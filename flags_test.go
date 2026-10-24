package main

import "testing"

func TestFlagsAfterStartPath(t *testing.T) {
	args := []string{"walk", "synthetic/start/path", "-h", "-v"}

	if !isFlagAt(args, 2, "--help", "-h") {
		t.Fatal("-h after a start path was not recognized")
	}
	if !isFlagAt(args, 3, "--version", "-v") {
		t.Fatal("-v after a start path was not recognized")
	}
	if isFlagAt(args, 1, "--help", "-h") {
		t.Fatal("start path was incorrectly recognized as a help flag")
	}
}

func TestFlagAtBounds(t *testing.T) {
	args := []string{"walk"}
	if isFlagAt(args, -1, "--help", "-h") || isFlagAt(args, len(args), "--help", "-h") {
		t.Fatal("out-of-range argument index matched a flag")
	}
}
