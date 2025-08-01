# A Tour of Go Solutions

[![tests](https://github.com/aaccioly-coding-challenges/golang-tour-solutions/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/aaccioly-coding-challenges/golang-tour-solutions/actions/workflows/test.yml)
[![lint](https://github.com/aaccioly-coding-challenges/golang-tour-solutions/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/aaccioly-coding-challenges/golang-tour-solutions/actions/workflows/lint.yml)
[![coverage](https://raw.githubusercontent.com/aaccioly-coding-challenges/golang-tour-solutions/badges/.badges/main/coverage.svg)](./.testcoverage.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aaccioly-coding-challenges/golang-tour-solutions?cache=v1)](https://goreportcard.com/report/github.com/aaccioly-coding-challenges/golang-tour-solutions)

This repository contains my solutions to the exercises in "[A Tour of Go](https://go.dev/tour/)".

## Running the Code

To run the code, you need to have Go installed on your machine. You can download it from the 
[official Go website](https://go.dev/dl/).

Once you have Go installed, you can use the `tour.go `file to run the exercises. Open a terminal, navigate to the
directory where the `tour.go` file is located, and run the following command:

```bash
go run tour.go <module> <program>
```

Replace `<module>` with the name of the module you want to run and `<program>` with the name of the program you want 
to run.  For example, to run the `hello` program in the `welcome` module, you would run:

```bash
go run tour.go welcome hello
```

Individual exercises are also runnable directly. For example, to run the `hello` exercise in the `welcome` module, 
you can run:

```bash
go run welcome/hello/hello.go
```

On Windows, you may need to use `go run .\welcome\hello\hello.go` instead.

## License

The Go Authors license the original code excerpts from the Go Tour, included or adapted from the official Go Tour 
repository under the BSD 3-Clause License.

My contributions and modifications are, in practice, not significant enough to warrant separate licensing, but I have
chosen to license this repository under the same BSD 3-Clause License for consistency. See the [LICENSE](LICENSE) file
for details.

Since intentions matter, let me state upfront that I'm using this as an exercise to refresh my knowledge of the Go
language and to practise source control and versioning with [Jujutsu](https://jj-vcs.github.io/jj/latest/). I couldn't
care less about how you use my code. Copy it, modify it, feed it to an AI, ignore it, whatever. Just don't do anything
that the original [Go Tour](https://go.dev/tour/) authors and the BSD-3-Clause License don't allow.
