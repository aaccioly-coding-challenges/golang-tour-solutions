package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Table-driven test cases for successful program execution
	tests := []struct {
		name        string // Test case name
		module      string // Module to run
		program     string // Program to run
		wantContain string // String that should be in the output
	}{
		{
			name:        "hello_program_prints_greeting",
			module:      "welcome",
			program:     "hello",
			wantContain: "Hello, 世界",
		},
		{
			name:        "sandbox_program_prints_go_birthday",
			module:      "welcome",
			program:     "sandbox",
			wantContain: "2009-11-10 23:00:00 +0000 UTC",
		},
		{
			name:        "packages_program_prints_favorite_number",
			module:      "basics",
			program:     "packages",
			wantContain: "My favorite number is",
		},
	}

	for _, tt := range tests {
		// Using t.Run creates a subtest for each test case
		t.Run(tt.name, func(t *testing.T) {
			got, err := run(tt.module, tt.program)

			// Check that there is no error (we expect all tests to succeed)
			if err != nil {
				t.Fatalf("run(%q, %q) error = %v, expected no error",
					tt.module, tt.program, err)
			}

			// Check that the output contains the expected string
			if !strings.Contains(got, tt.wantContain) {
				t.Errorf("run(%q, %q) output = %q, want to contain %q",
					tt.module, tt.program, got, tt.wantContain)
			}
		})
	}
}

func TestRun_InvalidProgram(t *testing.T) {
	// Table-driven test for error cases
	tests := []struct {
		name          string // Test case name
		module        string // Module to run
		program       string // Program to run
		expectedError string // Expected error message
	}{
		{
			name:          "nonexistent_module_and_program",
			module:        "notamodule",
			program:       "nonexistent",
			expectedError: "program notamodule/nonexistent/nonexistent.go not found",
		},
		{
			name:          "valid_module_with_nonexistent_program",
			module:        "welcome",
			program:       "nonexistent",
			expectedError: "program welcome/nonexistent/nonexistent.go not found",
		},
		{
			name:          "empty_module_and_program",
			module:        "",
			program:       "",
			expectedError: "program .go not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := run(tt.module, tt.program)

			// Check that we got an error
			if err == nil {
				t.Fatalf("run(%q, %q) expected error but got nil",
					tt.module, tt.program)
			}

			// Check that the error message matches what we expect
			if err.Error() != tt.expectedError {
				t.Errorf("run(%q, %q) error = %q, want %q",
					tt.module, tt.program, err.Error(), tt.expectedError)
			}
		})
	}
}
