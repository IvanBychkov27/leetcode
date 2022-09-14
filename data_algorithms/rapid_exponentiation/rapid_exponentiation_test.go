package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_rapidExponentiation(t *testing.T) {
	a, n := 2, 10
	expected := 1024
	res := rapidExponentiation(a, n)
	assert.Equal(t, expected, res)
}

func Test_rapidExponentiation2(t *testing.T) {
	a, n := 2, 10
	expected := 1024
	res := rapidExponentiation2(a, n)
	assert.Equal(t, expected, res)
}

func Test_pow(t *testing.T) {
	a, n := 2, 10
	expected := 1024
	res := int(math.Pow(float64(a), float64(n)))
	assert.Equal(t, expected, res)
}

func Test_naive(t *testing.T) {
	a, n := 2, 10
	expected := 1024
	res := naivePow(a, n)
	assert.Equal(t, expected, res)
}

func Benchmark_naive(b *testing.B) {
	a, n := 2, 36
	for i := 0; i < b.N; i++ {
		naivePow(a, n)
	}
}

func Benchmark_rapidExponentiation(b *testing.B) {
	a, n := 2, 36
	for i := 0; i < b.N; i++ {
		rapidExponentiation(a, n)
	}
}

func Benchmark_rapidExponentiation2(b *testing.B) {
	a, n := 2, 36
	for i := 0; i < b.N; i++ {
		rapidExponentiation2(a, n)
	}
}

func Benchmark_pow(b *testing.B) {
	a, n := 2, 36
	for i := 0; i < b.N; i++ {
		_ = int(math.Pow(float64(a), float64(n)))
	}
}

/*
goos: linux
goarch: amd64
pkg: leetcode/data_algorithms/rapid_exponentiation
cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
Benchmark_naive-8                  	62974935	        18.85 ns/op	       0 B/op	       0 allocs/op
Benchmark_rapidExponentiation-8    	100000000	        10.57 ns/op	       0 B/op	       0 allocs/op
Benchmark_rapidExponentiation2-8   	121819239	         9.753 ns/op	   0 B/op	       0 allocs/op
Benchmark_pow-8                    	55227450	        21.66 ns/op	       0 B/op	       0 allocs/op
*/
