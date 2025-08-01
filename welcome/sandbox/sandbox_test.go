package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSandboxProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "2009-11-10 23:00:00 +0000 UTC"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}