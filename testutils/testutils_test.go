package testutils

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestCaptureOutput_BasicOutput tests basic stdout capture functionality
func TestCaptureOutput_BasicOutput(t *testing.T) {
	expected := "Hello, World!"

	output := CaptureOutput(func() {
		fmt.Print(expected)
	})

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureOutput_MultilineOutput tests capture of multiple lines
func TestCaptureOutput_MultilineOutput(t *testing.T) {
	testFunc := func() {
		fmt.Println("Line 1")
		fmt.Println("Line 2")
		fmt.Print("Line 3")
	}

	output := CaptureOutput(testFunc)
	expected := "Line 1\nLine 2\nLine 3"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureOutput_UnicodeAndSpecialChars tests Unicode and special characters
func TestCaptureOutput_UnicodeAndSpecialChars(t *testing.T) {
	expected := "Hello, 世界! 🌍 Special chars: @#$%^&*()"

	output := CaptureOutput(func() {
		fmt.Print(expected)
	})

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// TestCaptureOutput_EmptyOutput tests handling of functions that produce no output
func TestCaptureOutput_EmptyOutput(t *testing.T) {
	output := CaptureOutput(func() {
		// Function that produces no output
	})

	if output != "" {
		t.Errorf("Expected empty output, got %q", output)
	}
}

// TestCaptureOutput_StdoutRestoration tests that stdout is properly restored
func TestCaptureOutput_StdoutRestoration(t *testing.T) {
	// Save original stdout
	originalStdout := os.Stdout

	// Capture some output
	output := CaptureOutput(func() {
		fmt.Print("test output")
	})

	// Verify output was captured
	if output != "test output" {
		t.Errorf("Expected 'test output', got %q", output)
	}

	// Verify stdout was restored to original
	if os.Stdout != originalStdout {
		t.Error("stdout was not properly restored after CaptureOutput")
	}
}

// TestCaptureOutput_MultipleCaptures tests multiple consecutive captures
func TestCaptureOutput_MultipleCaptures(t *testing.T) {
	// First capture
	output1 := CaptureOutput(func() {
		fmt.Print("First capture")
	})

	// Second capture
	output2 := CaptureOutput(func() {
		fmt.Print("Second capture")
	})

	// Third capture
	output3 := CaptureOutput(func() {
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

// TestCaptureOutput_LargeOutput tests handling of large amounts of output
func TestCaptureOutput_LargeOutput(t *testing.T) {
	// Generate a large string
	var expectedBuilder strings.Builder
	for i := 0; i < 1000; i++ {
		_, _ = fmt.Fprintf(&expectedBuilder, "Line %d\n", i)
	}
	expected := expectedBuilder.String()

	output := CaptureOutput(func() {
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

// TestCaptureOutput_FunctionExecution tests that the provided function is actually executed
func TestCaptureOutput_FunctionExecution(t *testing.T) {
	executed := false

	CaptureOutput(func() {
		executed = true
		fmt.Print("Function executed")
	})

	if !executed {
		t.Error("The provided function was not executed")
	}
}

// TestCaptureOutput_MixedOutputTypes tests different types of output functions
func TestCaptureOutput_MixedOutputTypes(t *testing.T) {
	output := CaptureOutput(func() {
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

// TestCaptureOutputWithPanic_NoPanic tests CaptureOutputWithPanic when no panic occurs
func TestCaptureOutputWithPanic_NoPanic(t *testing.T) {
	expectedOutput := "No panic here"
	output, recovered := CaptureOutputWithPanic(func() {
		fmt.Print(expectedOutput)
	})

	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
	if recovered != nil {
		t.Errorf("Expected nil recovered, got %v", recovered)
	}
}

// TestCaptureOutputWithPanic_WithPanic tests CaptureOutputWithPanic when a panic occurs
func TestCaptureOutputWithPanic_WithPanic(t *testing.T) {
	expectedOutput := "Before panic"
	panicValue := "something went wrong"

	output, recovered := CaptureOutputWithPanic(func() {
		fmt.Print(expectedOutput)
		panic(panicValue)
	})

	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
	if recovered != panicValue {
		t.Errorf("Expected recovered %v, got %v", panicValue, recovered)
	}
}

// TestCaptureOutputWithPanic_StdoutRestoration tests that stdout is properly restored even after a panic
func TestCaptureOutputWithPanic_StdoutRestoration(t *testing.T) {
	originalStdout := os.Stdout

	_, _ = CaptureOutputWithPanic(func() {
		panic("panic!")
	})

	if os.Stdout != originalStdout {
		t.Error("stdout was not properly restored after CaptureOutputWithPanic with panic")
	}
}
