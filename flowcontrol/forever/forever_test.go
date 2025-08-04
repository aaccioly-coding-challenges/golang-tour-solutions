package main

import (
	"context"
	"testing"
	"time"
)

// TestForeverProgramDoesntEnd verifies that the main function runs indefinitely.
// This test is different from others in the project because it tests an infinite loop,
// so we can't use testutils.CaptureMainOutput which would hang forever.
func TestForeverProgramDoesntEnd(t *testing.T) {
	// Create a context with timeout for the test
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Channel to signal if main() unexpectedly returns
	done := make(chan struct{})

	// Start main() in a goroutine
	go func() {
		defer close(done)
		main()
	}()

	select {
	case <-done:
		t.Error("Expected the program to run forever, but it ended")
	case <-ctx.Done():
		// Test passes - the program ran for the expected duration without ending
		t.Log("Program correctly ran without ending within the timeout period")
	}
}
