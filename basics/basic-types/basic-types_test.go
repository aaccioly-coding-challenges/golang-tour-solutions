package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestBasicTypesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Type: bool Value: false\nType: uint64 Value: 18446744073709551615\nType: complex128 Value: (2+3i)\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
