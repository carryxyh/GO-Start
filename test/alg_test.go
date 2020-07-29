package main

import (
	"bytes"
	"testing"
)

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	var output bytes.Buffer
	in := assembleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}
