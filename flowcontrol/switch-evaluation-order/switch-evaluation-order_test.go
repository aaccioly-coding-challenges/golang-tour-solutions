package main

import (
	"regexp"
	"testing"
	"time"

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

func Test_whenIsSaturday(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{"Test_whenIsSaturday is today", time.Date(2025, time.August, 9, 0, 0, 0, 0, time.UTC), "Today."},
		{"Test_whenIsSaturday is tomorrow", time.Date(2025, time.August, 8, 0, 0, 0, 0, time.UTC), "Tomorrow."},
		{"Test_whenIsSaturday in two days", time.Date(2025, time.August, 7, 0, 0, 0, 0, time.UTC), "In two days."},
		{"Test_whenIsSaturday too far away", time.Date(2025, time.August, 10, 0, 0, 0, 0, time.UTC), "Too far away."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := whenIsSaturday(tt.t); got != tt.want {
				t.Errorf("whenIsSaturday() = %v, want %v", got, tt.want)
			}
		})
	}
}
