package testutils

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestCaptureMainOutput_BasicOutput tests basic stdout capture functionality
func TestCaptureMainOutput_BasicOutput(t *testing.T) {
	expected := "Hello, World!"

	output := CaptureMainOutput(func() {
		fmt.Print(expected)
	})

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureMainOutput_MultilineOutput tests capture of multiple lines
func TestCaptureMainOutput_MultilineOutput(t *testing.T) {
	testFunc := func() {
		fmt.Println("Line 1")
		fmt.Println("Line 2")
		fmt.Print("Line 3")
	}

	output := CaptureMainOutput(testFunc)
	expected := "Line 1\nLine 2\nLine 3"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureMainOutput_UnicodeAndSpecialChars tests Unicode and special characters
func TestCaptureMainOutput_UnicodeAndSpecialChars(t *testing.T) {
	expected := "Hello, 世界! 🌍 Special chars: @#$%^&*()"

	output := CaptureMainOutput(func() {
		fmt.Print(expected)
	})

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureMainOutput_EmptyOutput tests handling of functions that produce no output
func TestCaptureMainOutput_EmptyOutput(t *testing.T) {
	output := CaptureMainOutput(func() {
		// Function that produces no output
	})

	if output != "" {
		t.Errorf("Expected empty output, got %q", output)
	}
}

// TestCaptureMainOutput_StdoutRestoration tests that stdout is properly restored
func TestCaptureMainOutput_StdoutRestoration(t *testing.T) {
	// Save original stdout
	originalStdout := os.Stdout

	// Capture some output
	output := CaptureMainOutput(func() {
		fmt.Print("test output")
	})

	// Verify output was captured
	if output != "test output" {
		t.Errorf("Expected 'test output', got %q", output)
	}

	// Verify stdout was restored to original
	if os.Stdout != originalStdout {
		t.Error("stdout was not properly restored after CaptureMainOutput")
	}
}

// TestCaptureMainOutput_MultipleCaptures tests multiple consecutive captures
func TestCaptureMainOutput_MultipleCaptures(t *testing.T) {
	// First capture
	output1 := CaptureMainOutput(func() {
		fmt.Print("First capture")
	})

	// Second capture
	output2 := CaptureMainOutput(func() {
		fmt.Print("Second capture")
	})

	// Third capture
	output3 := CaptureMainOutput(func() {
		fmt.Print("Third capture")
	})

	// Verify all captures worked independently
	if output1 != "First capture" {
		t.Errorf("First capture: expected 'First capture', got %q", output1)
	}
	if output2 != "Second capture" {
		t.Errorf("Second capture: expected 'Second capture', got %q", output2)
	}
	if output3 != "Third capture" {
		t.Errorf("Third capture: expected 'Third capture', got %q", output3)
	}
}

// TestCaptureMainOutput_LargeOutput tests handling of large amounts of output
func TestCaptureMainOutput_LargeOutput(t *testing.T) {
	// Generate a large string
	var expectedBuilder strings.Builder
	for i := 0; i < 1000; i++ {
		expectedBuilder.WriteString(fmt.Sprintf("Line %d\n", i))
	}
	expected := expectedBuilder.String()

	output := CaptureMainOutput(func() {
		for i := 0; i < 1000; i++ {
			fmt.Printf("Line %d\n", i)
		}
	})

	if output != expected {
		t.Errorf("Large output test failed. Expected length %d, got length %d", len(expected), len(output))
		// Don't print the full output as it would be too verbose
		if !strings.HasPrefix(output, "Line 0\n") {
			t.Errorf("Output doesn't start correctly. Got: %q", output[:20])
		}
		if !strings.HasSuffix(output, "Line 999\n") {
			t.Errorf("Output doesn't end correctly. Got: %q", output[len(output)-20:])
		}
	}
}

// TestCaptureMainOutput_FunctionExecution tests that the provided function is actually executed
func TestCaptureMainOutput_FunctionExecution(t *testing.T) {
	executed := false

	CaptureMainOutput(func() {
		executed = true
		fmt.Print("Function executed")
	})

	if !executed {
		t.Error("The provided function was not executed")
	}
}

// TestCaptureMainOutput_MixedOutputTypes tests different types of output functions
func TestCaptureMainOutput_MixedOutputTypes(t *testing.T) {
	output := CaptureMainOutput(func() {
		fmt.Print("Print: ")
		fmt.Printf("Printf %d ", 42)
		fmt.Println("Println")
		_, _ = fmt.Fprint(os.Stdout, "Fprint")
	})

	expected := "Print: Printf 42 Println\nFprint"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
