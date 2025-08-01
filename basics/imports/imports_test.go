package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestImportsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Now you have 2.6457513110645907 problems.\n"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}