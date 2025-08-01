package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestFunctionsContinuedProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "55"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}