package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSliceLiterals(t *testing.T) {
	output := testutils.CaptureOutput(main)
	expected := "[2 3 5 7 11 13]\n[true false true true false true]\n[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
