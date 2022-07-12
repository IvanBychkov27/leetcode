package main

import (
	"math"
	"testing"
)

const (
	d = 77.0
	s = 11
)

func Benchmark_Pow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myPow(d, s)
	}
}

func Benchmark_myPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(d, s)
	}
}
