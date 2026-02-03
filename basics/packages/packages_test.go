package main

import (
	"regexp"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestPackagesProgram(t *testing.T) {
	output := testutils.CaptureOutput(main)

	// The program outputs "My favorite number is X" where X is a random number from 0-9
	expectedPattern := `^My favorite number is [0-9]\n$`
	matched, err := regexp.MatchString(expectedPattern, output)
	if err != nil {
		t.Fatalf("Failed to compile regex pattern: %v", err)
	}
	if !matched {
		t.Errorf("Expected output to match pattern %q, got %q", expectedPattern, output)
	}
}
