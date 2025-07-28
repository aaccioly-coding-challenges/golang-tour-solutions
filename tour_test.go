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
		{
			name:        "imports_program_prints_sqrt",
			module:      "basics",
			program:     "imports",
			wantContain: "Now you have 2.6457513110645907 problems.\n",
		},
		{
			name:        "exported_names_program_prints_pi",
			module:      "basics",
			program:     "exported-names",
			wantContain: "3.141592653589793",
		},
		{
			name:        "functions_program_prints_sum",
			module:      "basics",
			program:     "functions",
			wantContain: "55",
		},
		{
			name:        "functions_continued_program_also_prints_sum",
			module:      "basics",
			program:     "functions-continued",
			wantContain: "55",
		},
		{
			name:        "multiple_results_program_prints_swapped_values",
			module:      "basics",
			program:     "multiple-results",
			wantContain: "world hello",
		},
		{
			name:        "named_results_program_splits_numbers",
			module:      "basics",
			program:     "named-results",
			wantContain: "7 10",
		},
		{
			name:        "variables_program_prints_zero_values",
			module:      "basics",
			program:     "variables",
			wantContain: "0 false false false",
		},
		{
			name:        "variables_with_initializers_program_prints_values",
			module:      "basics",
			program:     "variables-with-initializers",
			wantContain: "1 2 true false no!",
		},
		{
			name:        "short_variable_declaration_program_prints_values",
			module:      "basics",
			program:     "short-variable-declaration",
			wantContain: "1 2 3 true false no!",
		},
		{
			name:        "basic_types_program_prints_types_and_values",
			module:      "basics",
			program:     "basic-types",
			wantContain: "Type: bool Value: false\nType: uint64 Value: 18446744073709551615\nType: complex128 Value: (2+3i)",
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
