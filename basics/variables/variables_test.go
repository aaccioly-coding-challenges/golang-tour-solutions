package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestVariablesProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "0 false false false"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}