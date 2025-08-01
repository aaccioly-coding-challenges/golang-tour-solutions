package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestPackagesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "My favorite number is"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}