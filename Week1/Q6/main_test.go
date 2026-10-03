package main

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	err := ReadFile()
	
	if err == nil {
		t.Fatal("ReadFile() returned no error, expected an error")
	}

	if !strings.HasPrefix(err.Error(), "failed to open file:") {
		t.Errorf("Error message %q does not start with \"failed to open file:\"", err.Error())
	}

	unwrapped := errors.Unwrap(err)
	if unwrapped == nil {
		t.Fatal("errors.Unwrap(err) returned nil, error was not properly wrapped with %w")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("errors.Is(err, os.ErrNotExist) = false; want true. Ensure wrapped error is from os.Open")
	}
}
