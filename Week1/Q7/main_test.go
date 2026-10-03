package main

import (
	"errors"
	"testing"
)

func TestErrorIs(t *testing.T) {
	err := DoSomething()

	if err == nil {
		t.Fatal("DoSomething() returned nil, expected an error")
	}

	if err == ErrPermissionDenied {
		t.Fatal("DoSomething() returned ErrPermissionDenied directly, but it should be wrapped with %w")
	}

	if !errors.Is(err, ErrPermissionDenied) {
		t.Errorf("errors.Is(err, ErrPermissionDenied) = false; want true")
	}
}
