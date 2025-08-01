package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestConstantsProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "Hello 世界\nHappy 3.14 Day\nGo rules? true"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}