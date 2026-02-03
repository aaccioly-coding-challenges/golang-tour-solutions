package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSwitchWithNoConditionProgram(t *testing.T) {
	output := strings.TrimSpace(testutils.CaptureMainOutput(main))
	switch output {
	case "Good morning!", "Good afternoon.", "Good evening.":
		// acceptable outputs
	default:
		t.Fatalf("unexpected output %q", output)
	}
}
