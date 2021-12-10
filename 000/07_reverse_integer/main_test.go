package main

import (
	"testing"
)

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(-214748399)
	}
}

func Benchmark_reverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse2(-214748399)
	}
}

func Benchmark_reverse3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse3(-214748399)
	}
}
