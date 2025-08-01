package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

// TestTourPositive tests successful execution scenarios using table-driven tests
func TestTourPositive(t *testing.T) {
	// Save original os.Args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Table of positive test cases
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		useRegex       bool
		regexPattern   string
	}{
		{
			name:           "successful_execution_welcome_hello",
			args:           []string{"tour", "welcome", "hello"},
			expectedOutput: "Hello, 世界\n",
			useRegex:       false,
		},
		{
			name:         "successful_execution_basics_packages",
			args:         []string{"tour", "basics", "packages"},
			useRegex:     true,
			regexPattern: `^My favorite number is [0-9]\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			output := testutils.CaptureMainOutput(main)
			
			if tt.useRegex {
				matched, err := regexp.MatchString(tt.regexPattern, output)
				if err != nil {
					t.Fatalf("Failed to compile regex pattern: %v", err)
				}
				if !matched {
					t.Errorf("Expected output to match pattern %q, got %q", tt.regexPattern, output)
				}
			} else {
				if output != tt.expectedOutput {
					t.Errorf("Expected output %q, got %q", tt.expectedOutput, output)
				}
			}
		})
	}
}

// mockLogger implements the Logger interface for testing
type mockLogger struct {
	fatalMsgs []string
}

func (m *mockLogger) Printf(format string, v ...interface{}) {
	// No-op: we don't need to track printf calls
}

func (m *mockLogger) Println(v ...interface{}) {
	// No-op: we don't need to track println calls
}

func (m *mockLogger) Fatal(v ...interface{}) {
	m.fatalMsgs = append(m.fatalMsgs, fmt.Sprint(v...))
	// Don't actually call os.Exit() in tests
}

func (m *mockLogger) Fatalf(format string, v ...interface{}) {
	m.fatalMsgs = append(m.fatalMsgs, fmt.Sprintf(format, v...))
	// Don't actually call os.Exit() in tests
}

// TestTourNegative tests error scenarios using table-driven tests with mock logger
func TestTourNegative(t *testing.T) {
	// Save original os.Args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Table of negative test cases
	tests := []struct {
		name             string
		args             []string
		expectedFatalMsg string
	}{
		{
			name:             "insufficient_arguments_no_args",
			args:             []string{"tour"},
			expectedFatalMsg: "insufficient arguments",
		},
		{
			name:             "insufficient_arguments_one_arg",
			args:             []string{"tour", "welcome"},
			expectedFatalMsg: "insufficient arguments",
		},
		{
			name:             "nonexistent_program_failure",
			args:             []string{"tour", "welcome", "nonexistent"},
			expectedFatalMsg: "Failed to run welcome/nonexistent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockLogger{}
			os.Args = tt.args

			runMain(mock)

			// Check fatal messages
			if tt.expectedFatalMsg != "" {
				if len(mock.fatalMsgs) == 0 {
					t.Error("Expected fatal call, got none")
				} else if !strings.Contains(mock.fatalMsgs[0], tt.expectedFatalMsg) {
					t.Errorf("Expected fatal message to contain %q, got %q", tt.expectedFatalMsg, mock.fatalMsgs[0])
				}
			}
		})
	}
}
