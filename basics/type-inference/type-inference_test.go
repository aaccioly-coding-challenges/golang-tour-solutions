package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestTypeInferenceProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "v is of type int"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}