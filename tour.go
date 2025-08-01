// Package main provides a CLI tool for running Go tour examples.
// It executes example programs by specifying a module and program name,
// making it easier to run and explore the Go tour code samples.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Logger interface for dependency injection to enable testing of Fatal calls
type Logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

// main initializes and runs the tour example CLI, which executes
// specific examples from the Go tour by providing a module and program name.
func main() {
	logger := log.New(os.Stderr, "", 0)
	runMain(logger)
}

// runMain contains the main logic and accepts a logger for dependency injection
func runMain(logger Logger) {

	flag.Usage = func() {
		logger.Printf("Usage: %s <module> <program>", os.Args[0])
		logger.Println("\nArguments:")
		logger.Println("  module   The module name (e.g., welcome)")
		logger.Println("  program  The program name (e.g., hello)")
		logger.Println("\nFlags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		logger.Fatal("insufficient arguments")
	}

	module := flag.Arg(0)
	program := flag.Arg(1)

	output, err := run(module, program)
	if err != nil {
		logger.Fatalf("Failed to run %s/%s: %v", module, program, err)
	}
	fmt.Print(output)
}

// run executes the specified program in the given module and returns its output
func run(module string, program string) (string, error) {
	programPath := filepath.Join(module, program, program+".go")
	if _, err := os.Stat(programPath); os.IsNotExist(err) {
		return "", fmt.Errorf("program %s not found", programPath)
	}

	cmd := exec.Command("go", "run", programPath)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
