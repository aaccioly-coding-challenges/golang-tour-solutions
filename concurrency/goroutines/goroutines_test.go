package main

import (
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestGoroutinesProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	if count := strings.Count(output, "hello"); count != 5 {
		t.Errorf("Expected 'hello' 5 times, got %d", count)
	}
	if count := strings.Count(output, "world"); count < 4 {
		t.Errorf("Expected 'world' 4 or more times, got %d", count)
	}
}
