package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestInterfacesAreSatisfiedImplicitlyProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "hello\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}

}
