package main

import (
	"regexp"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSandboxProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)

	// Single multiline pattern that matches the entire expected output structure
	expectedPattern := "^Welcome to the playground!\n" +
		"The time is \\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d+ [+-]\\d{4} \\w+ m=\\+\\d+\\.\\d+\n" +
		"If this was the real playground is would be always 2009-11-10 23:00:00 \\+0000 UTC\n$"
	matched, err := regexp.MatchString(expectedPattern, output)
	if err != nil {
		t.Fatalf("Failed to compile regex pattern: %v", err)
	}
	if !matched {
		t.Errorf("Expected output to match pattern %q, got %q", expectedPattern, output)
	}
}
