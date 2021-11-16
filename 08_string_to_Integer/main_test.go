package main

import (
	"testing"
)

func Benchmark_myAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi(" -4193 with words")
	}
}

func Benchmark__myAtoi2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi_2(" -4193 with words")
	}
}
