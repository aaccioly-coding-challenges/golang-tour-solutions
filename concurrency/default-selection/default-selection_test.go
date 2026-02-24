package main

import (
	"strings"
	"testing"
	"testing/synctest"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestDefaultSelectionProgram(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		output := testutils.CaptureOutput(main)

		if count := strings.Count(output, "tick"); count < 4 {
			t.Errorf("Expected 'tick' 4 or more times, got %d", count)
		}
		if count := strings.Count(output, "."); count < 14 {
			t.Errorf("Expected '.' 14 or more times, got %d", count)
		}
		if count := strings.Count(output, "BOOM!"); count != 1 {
			t.Errorf("Expected 'BOOM!' only once, got %d", count)
		}
	})
}
