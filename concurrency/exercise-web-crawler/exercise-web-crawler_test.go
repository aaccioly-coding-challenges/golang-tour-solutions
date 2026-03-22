package main

import (
	"strings"
	"sync"
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/internal/capture"
)

func TestWebCrawlerProgram(t *testing.T) {
	output := capture.CaptureOutput(main)
	expectedMessages := []string{
		"-> Fetching: https://golang.org/",
		"<- Found: https://golang.org/ \"The Go Programming Language\"",
		"-> Fetching: https://golang.org/cmd/",
		"<- Error fetching https://golang.org/cmd/: not found: https://golang.org/cmd/",
		"-> Fetching: https://golang.org/pkg/",
		"<- Found: https://golang.org/pkg/ \"Packages\"",
		"-> Fetching: https://golang.org/pkg/os/",
		"<- Found: https://golang.org/pkg/os/ \"Package os\"",
		"-> Fetching: https://golang.org/pkg/fmt/",
		"<- Found: https://golang.org/pkg/fmt/ \"Package fmt\"",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected output to contain message %q, but it didn't", msg)
		}
	}

	expectedStats := strings.Join([]string{
		"Fetching stats",
		"--------------",
		"https://golang.org/ was fetched",
		"https://golang.org/cmd/ failed: not found: https://golang.org/cmd/",
		"https://golang.org/pkg/ was fetched",
		"https://golang.org/pkg/fmt/ was fetched",
		"https://golang.org/pkg/os/ was fetched",
		"",
	}, "\n")

	if !strings.HasSuffix(output, expectedStats) {
		t.Errorf("Expected output to end with:\n%s\nGot:\n%s", expectedStats, output)
	}
}

func TestWebCrawlerDepthZero(t *testing.T) {
	fetched := NewURLCache()
	wg := &sync.WaitGroup{}
	output := capture.CaptureOutput(func() {
		Crawl("https://golang.org/", 0, fetcher, fetched, wg)
		wg.Wait()
	})

	expected := "<- Done with https://golang.org/, depth 0.\n"
	if output != expected {
		t.Errorf("Expected output %q, but got %q", expected, output)
	}
}
