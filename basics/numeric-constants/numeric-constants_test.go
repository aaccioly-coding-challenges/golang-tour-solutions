package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestNumericConstantsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "21\n0.2\n1.2676506002282295e+29"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}