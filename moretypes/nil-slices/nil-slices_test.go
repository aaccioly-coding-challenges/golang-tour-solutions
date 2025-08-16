package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestNilSlicesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "[] 0 0\nnil!\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
