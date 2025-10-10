package main

import (
	"testing"

	"github.com/aaccioly-coding-challenges/golang-tour-solutions/testutils"
	"golang.org/x/tour/wc"
)

func TestWordCount(t *testing.T) {
	wc.Test(WordCount)
}

func TestWordCountProgram(t *testing.T) {
	output := testutils.CaptureMainOutput(main)
	expected := "PASS\n f(\"I am learning Go!\") = \n  map[string]int{\"Go!\":1, \"I\":1, \"am\":1, \"learning\":1}\n" +
		"PASS\n f(\"The quick brown fox jumped over the lazy dog.\") = \n  map[string]int{\"The\":1, \"brown\":1, \"dog.\":1, \"fox\":1, \"jumped\":1, \"lazy\":1, \"over\":1, \"quick\":1, \"the\":1}\n" +
		"PASS\n f(\"I ate a donut. Then I ate another donut.\") = \n  map[string]int{\"I\":2, \"Then\":1, \"a\":1, \"another\":1, \"ate\":2, \"donut.\":2}\n" +
		"PASS\n f(\"A man a plan a canal panama.\") = \n  map[string]int{\"A\":1, \"a\":2, \"canal\":1, \"man\":1, \"panama.\":1, \"plan\":1}\n"

	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
