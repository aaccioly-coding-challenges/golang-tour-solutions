package main

import (
	"strings"
	"testing"
	"time"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
)

func TestSwitchWithNoConditionProgram(t *testing.T) {
	output := strings.TrimSpace(testutils.CaptureOutput(main))
	switch output {
	case "Good morning!", "Good afternoon.", "Good evening.":
		// acceptable outputs
	default:
		t.Fatalf("unexpected output %q", output)
	}
}

func Test_greet(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "morning",
			args: args{t: time.Date(2026, 1, 1, 11, 0, 0, 0, time.UTC)},
			want: "Good morning!",
		},
		{
			name: "afternoon",
			args: args{t: time.Date(2026, 1, 1, 16, 0, 0, 0, time.UTC)},
			want: "Good afternoon.",
		},
		{
			name: "evening",
			args: args{t: time.Date(2026, 1, 1, 17, 0, 0, 0, time.UTC)},
			want: "Good evening.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.TrimSpace(testutils.CaptureOutput(func() {
				greet(tt.args.t)
			}))
			if got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
