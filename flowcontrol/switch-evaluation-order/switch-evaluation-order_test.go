package main

import (
	"regexp"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSwithEvaluationOrderProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expectedPattern := `^When's Saturday\?\n(Today|Tomorrow|In two days|Too far away)\.\n$`
	matched, err := regexp.MatchString(expectedPattern, output)
	if err != nil {
		t.Fatalf("Failed to compile regex pattern: %v", err)
	}
	if !matched {
		t.Errorf("Expected output to match pattern %q, got %q", expectedPattern, output)
	}
}
